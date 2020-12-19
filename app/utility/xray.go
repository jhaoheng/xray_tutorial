package utility

import (
	"context"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
)

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     GetXRAYAddr(), // default
		ServiceVersion: "0.0.1",
	})
	xray.SetLogger(xraylog.NewDefaultLogger(os.Stderr, xraylog.LogLevelInfo))
}

/*
GetXRAYAppName ...
*/
func GetXRAYAppName() string {
	appName := os.Getenv("XRAY_APP_NAME")
	if appName != "" {
		return appName
	}
	return "myApp"
}

/*
GetXRAYAddr ...
*/
func GetXRAYAddr() string {
	XrayDaemonAddr := os.Getenv("XRAY_DAEMON_ADDR")
	if XrayDaemonAddr != "" {
		return XrayDaemonAddr
	}
	return "localhost:2000"
}

/*
DbXrayMiddle ...
*/
func DbXrayMiddle(ctx context.Context, opName string, f func() error) error {
	subCtx, subSeg := xray.BeginSubsegment(ctx, "DbXrayMiddle")
	defer subSeg.Close(nil)

	return xray.Capture(subCtx, opName, func(ctx1 context.Context) error {
		var err error
		err = f()
		if err != nil {
			return err
		}
		err = xray.AddMetadata(ctx1, "tag", "my private metadata")
		return err
	})
}
