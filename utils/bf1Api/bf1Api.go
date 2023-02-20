package bf1Api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

const (
	baseUrl     = "https://sparta-gw-bf1.battlelog.com/jsonrpc/pc/api"
	contentType = "application/json"
	host        = "sparta-gw.battlelog.com"
)

type requestData struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      string      `json:"id"`
}

func (d *requestData) toIoReader() (io.Reader, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	} else {
		return strings.NewReader(string(b)), err
	}
}

// 退出游戏
func LeaveGame(session, gameId string) {
	currentUUID := uuid.New().String()
	param := map[string]string{
		"game":   "tunguska",
		"gameId": gameId,
	}
	rd := requestData{
		Jsonrpc: "2.0",
		Method:  "Game.leaveGame",
		Params:  param,
		Id:      currentUUID,
	}
	data, err := rd.toIoReader()
	if err != nil {
		fmt.Println(err)
		// 后续插入到暖服日志
	}
	res, err := sendCommand(session, data)
	if err != nil {
		fmt.Println(err)
		// 后续插入到暖服日志
	}
	fmt.Println(res)
	// 后续插入到暖服日志
}

// 获取账号正在游玩并更新状态
func ServersByPersonaIds(session string, personaIds []string) {
	currentUUID := uuid.New().String()
	param := map[string][]string{
		"personaIds": personaIds,
	}
	rd := requestData{
		Jsonrpc: "2.0",
		Method:  "GameServer.getServersByPersonaIds",
		Params:  param,
		Id:      currentUUID,
	}
	data, err := rd.toIoReader()
	if err != nil {
		fmt.Println(err)
	}
	res, err := sendCommand(session, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 获取游戏信息
func ServerDetails(session, gameId string) (string, error) {
	currentUUID := uuid.New().String()
	param := map[string]string{
		"game":   "tunguska",
		"gameId": gameId,
	}
	rd := requestData{
		Jsonrpc: "2.0",
		Method:  "GameServer.getServerDetails",
		Params:  param,
		Id:      currentUUID,
	}
	data, err := rd.toIoReader()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := sendCommand(session, data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(res)
	return res, nil
}

// 搜索服务器
func SearchServers(session, serverName string) (string, error) {
	currentUUID := uuid.New().String()
	fliterJson := fmt.Sprintf("{\"name\":\"%s\"}", serverName)
	param := map[string]string{
		"game":       "tunguska",
		"fliterJson": fliterJson,
	}
	rd := requestData{
		Jsonrpc: "2.0",
		Method:  "GameServer.getServerDetails",
		Params:  param,
		Id:      currentUUID,
	}
	data, err := rd.toIoReader()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := sendCommand(session, data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(res)
	return res, nil
}
