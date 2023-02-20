package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/Garfield247/gonf.git/pkg/e"
	"github.com/Garfield247/gonf.git/pkg/response"
	"github.com/Garfield247/gonf.git/utils/rediscli"
	"github.com/gin-gonic/gin"
)

type heartBeat struct {
	OnlineState int       `json:"onlineState"`
	SessionID   string    `json:"sessionId"`
	User        string    `json:"user"`
	GameID      string    `json:"gameId"`
	Name        string    `json:"name"`
	State       string    `json:"state"`
	Path        string    `json:"path"`
	Time        time.Time `json:"time"`
}

func InpectHeartBeat(ctx *gin.Context) {
	var hb heartBeat
	rctx := context.Background()
	err := ctx.ShouldBindJSON(&hb)
	if err != nil {
		response.FailWithResponse(ctx, e.UnknowError, fmt.Sprintf("Json解析错误%s", err))
		return
	}
	fmt.Println(hb)
	// 将心跳信息缓存到Redis
	if err := rediscli.RDB.Set(rctx, "bot01", hb, 0).Err(); err != nil {
		response.FailWithResponse(ctx, e.UnknowError, fmt.Sprintf("Json解析错误%s", err))
		return
	}
	response.SuccessWithResponse(ctx, map[string]string{})
}
