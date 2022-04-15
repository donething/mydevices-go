package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
	"mydevices-go/models"
)

// go:embed "templates/router.html"
// var tplRouter string

// 入口函数
func hello(_ context.Context, r events.APIGatewayRequest) (*models.JResp, error) {
	return handle(r.QueryString, r.Body), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}

// 根据参数选择功能
func handle(query events.APIGatewayQueryString, body string) *models.JResp {
	if len(query["cmd"]) == 0 {
		return models.Gen(200, "未指定命令参数", "text/plain; charset=UTF-8")
	}

	cmd := query["cmd"][0]
	switch cmd {
	case "index":
		// GET ?cmd=index
		return models.Gen(200, "Index.", "text/plain; charset=UTF-8")
	default:
		return models.Gen(
			200,
			fmt.Sprintf("未知的命令参数 '%s'", cmd),
			"text/plain; charset=UTF-8",
		)
	}
}
