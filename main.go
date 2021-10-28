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
	url := os.Getenv("FEISHU_WEBHOOK")



	var str = []byte(`{"msg_type":"text", "content":"{"text":"你中奖了"}"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(str))
	// for k, v := range getCookie() {
	// 	req.AddCookie(&http.Cookie{Name: k, Value: v})
	// }
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

func checkIn() {
	url := 
}