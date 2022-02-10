package utility

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-xray-sdk-go/awsplugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	"gorm.io/gorm"
)

func init() {
	// conditionally load plugin
	if os.Getenv("ENVIRONMENT") == "production" {
		ecs.Init()
	}
	xray.Configure(xray.Config{
		DaemonAddr:     GetXRAYAddr(), // default
		ServiceVersion: "0.0.1",
		// ContextMissingStrategy: ctxmissing.NewDefaultLogErrorStrategy(),
	})
	xray.SetLogger(xraylog.NewDefaultLogger(os.Stderr, xraylog.LogLevelInfo))
}

/*
XRaySetAWSServices - make AWS Service to enable XRay
ex : XRaySetAWSServices(s3.Client)
*/
func XRaySetAWSServices(client *client.Client) {
	xray.AWS(client)
}

/*
GetXRAYAppName ...
*/
func GetXRAYAppName() string {
	appName := os.Getenv("XRAY_APP_NAME")
	if appName != "" {
		return appName
	}
	return "xray_tutorial"
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
XrayMiddle ... 抓取資料送到 xray
*/
func XrayMiddle(level1ctx context.Context, opName string, myfunc func() error) error {
	addAdvanceInfo(level1ctx, "keyName", "data in level_1_ctx") //
	//
	level2ctx, subSeg := xray.BeginSubsegment(level1ctx, "XrayMiddle")
	defer subSeg.Close(nil)
	addAdvanceInfo(level2ctx, "keyName", "data in level_2_ctx") //

	return xray.Capture(level2ctx, opName, func(level3ctx context.Context) error {
		err := myfunc()
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

/*
- 因 gorm 與 xray, 互相無法支援, 所以用 wrap 的方式包裝起來
*/
func XrayGormWrap(ctx context.Context, gorm_exec func() *gorm.DB, options ...string) *gorm.DB {
	//
	ctx, sub_segment := xray.BeginSubsegment(ctx, "gorm")
	//
	tx := gorm_exec()
	//
	xray.AddMetadata(ctx, "sql_string", options)
	sub_segment.Close(tx.Error)
	return tx
}
