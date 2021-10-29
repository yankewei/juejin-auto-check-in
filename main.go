package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"fmt"
)

func main() {
	sendAlert(checkIn())
	sendAlert(lottery())
}

func checkIn() string {
	url := "https://api.juejin.cn/growth_api/v1/check_in"
	response := sendRequest(url)
	return string(response)
}

func lottery() string {
	url := "https://api.juejin.cn/growth_api/v1/lottery/draw"
	response := sendRequest(url)
	return string(response)
}

func sendRequest(url string) []byte {
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString("{}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")

	for k, v := range getCookie() {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func getCookie() map[string]string {
	str := os.Getenv("COOKIE")
	slice := strings.Split(str, ";")
	cookieMap := make(map[string]string)
	for _, v := range slice {
		v = strings.Trim(v, " ")
		index := strings.Index(v, "=")
		cookieMap[v[0:index]] = v[index+1:]
	}
	return cookieMap
}

func sendAlert(send_text string) {
	url := os.Getenv("FEISHU_WEBHOOK")
	type content struct {
		Text string `json:"text"`
	}

	type parameter struct {
		MsgType string   `json:"msg_type"`
		Content *content `json:"content"`
	}

	str := &parameter{
		MsgType: "text",
		Content: &content{
			Text: send_text,
		},
	}

	byte, _ := json.Marshal(str)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
}
