package utility

import (
	"context"
	"os"

	"github.com/aws/aws-xray-sdk-go/awsplugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
)

func init() {
	// conditionally load plugin
	if os.Getenv("ENVIRONMENT") == "production" {
		ecs.Init()
	}
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
func DbXrayMiddle(level1ctx context.Context, opName string, f func() error) error {
	addAdvanceInfo(level1ctx, "keyName", "data in level_1_ctx") //
	level2ctx, subSeg := xray.BeginSubsegment(level1ctx, "DbXrayMiddle")
	defer subSeg.Close(nil)
	addAdvanceInfo(level2ctx, "keyName", "data in level_2_ctx") //

	return xray.Capture(level2ctx, opName, func(level3ctx context.Context) error {
		var err error
		err = f()
		if err != nil {
			return err
		}
		addAdvanceInfo(level3ctx, "keyName", "data in level_3_ctx") //
		return err
	})
}

func addAdvanceInfo(ctx context.Context, key, value string) {
	xray.AddMetadata(ctx, key, value)
	xray.AddAnnotation(ctx, key, value)
}
