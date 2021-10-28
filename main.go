package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main() {
	fmt.Println(os.Getenv("FEISHU_WEBHOOK"))

	url := "https://open.feishu.cn/open-apis/bot/v2/hook/b858da68-00db-4629-aca7-767cc37b22fc"

	var str = "{'msg_type':'text', 'content': {'text': '新更新提醒'}}"
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(str))
	for k, v := range getCookie() {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getCookie() map[string]string {
	str := os.Getenv("COOKIE")
	slice := strings.Split(str, ";")
	cookieMap := make(map[string]string)
	for _, v := range slice {
		index := strings.Index(strings.Trim(v, " "), "=")
		cookieMap[v[0:index]] = v[index+1:]
	}
	return cookieMap
}
