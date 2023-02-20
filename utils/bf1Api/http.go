package bf1Api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// 基础请求
func sendCommand(session string, data io.Reader) (string, error) {
	client := &http.Client{}
	json.Marshal(data)
	req, err := http.NewRequest("POST", baseUrl, data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Host", host)
	req.Header.Set("X-Gatewaysession", session)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
