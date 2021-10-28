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

	var jsonStr = []byte(`{"msg_type":"text","content":{"text":"Test"}}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
	str := "_ga=GA1.2.939679032.1605796620; MONITOR_WEB_ID=88cc506b-ae22-413a-b58e-998fad01ec4c; passport_csrf_token_default=95c0488b5ba4e9aecae43903b17bf26f; passport_csrf_token=95c0488b5ba4e9aecae43903b17bf26f; sid_guard=d4cf4401e695fe219dc5aaa21593c672|1632753166|5184000|Fri,+26-Nov-2021+14:32:46+GMT; uid_tt=ab8cdaae710128c1db4a932c982c3348; uid_tt_ss=ab8cdaae710128c1db4a932c982c3348; sid_tt=d4cf4401e695fe219dc5aaa21593c672; sessionid=d4cf4401e695fe219dc5aaa21593c672; sessionid_ss=d4cf4401e695fe219dc5aaa21593c672; sid_ucp_v1=1.0.0-KDEzM2Q1YzhmNjliMDNhOGU3NWYzODFkMGJiM2NhZDZlZDdiYmVlZGQKFwjIxpC__fXBBhCOrMeKBhiwFDgCQO8HGgJsZiIgZDRjZjQ0MDFlNjk1ZmUyMTlkYzVhYWEyMTU5M2M2NzI; ssid_ucp_v1=1.0.0-KDEzM2Q1YzhmNjliMDNhOGU3NWYzODFkMGJiM2NhZDZlZDdiYmVlZGQKFwjIxpC__fXBBhCOrMeKBhiwFDgCQO8HGgJsZiIgZDRjZjQ0MDFlNjk1ZmUyMTlkYzVhYWEyMTU5M2M2NzI; n_mh=Fl1946g9fh2hxMpM9pj9wspuMcohK7EC_G5C6eorvuk; _gid=GA1.2.1646189504.1635334047"
	slice := strings.Split(str, ";")
	cookieMap := make(map[string]string)
	for _, v := range slice {
		index := strings.Index(strings.Trim(v, " "), "=")
		cookieMap[v[0:index]] = v[index+1:]
	}
	return cookieMap
}
