package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/v1/chat/completions", HttpPost)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe("0.0.0.0:8089", nil)
	log.Println(err)
}

func HttpPost(w http.ResponseWriter, r *http.Request) {
	api := "https://api.openai.com/v1/chat/completions"
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "无法读取请求体", http.StatusInternalServerError)
		return
	}

	fmt.Println(w, "请求体内容：%s", string(body))
	//	str := `{
	//    "model": "gpt-3.5-turbo",
	//    "messages": [
	//
	//        {
	//            "role": "user",
	//            "content": "你知道劲舞团这个游戏吗？"
	//        },{
	//            "role": "user",
	//            "content": "那么怎么玩好这个游戏呢？"
	//        }
	//    ],
	//    "top_p":0.8,
	//    "stream":true
	//}`
	data := make(map[string]interface{})
	_ = json.Unmarshal(body, &data)
	dataByte, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postData := bytes.NewBuffer(dataByte)
	var resp *http.Response
	req, err := http.NewRequest(http.MethodPost, api, postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "服务器返回错误状态码", http.StatusInternalServerError)
		return
	}
	for {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取HTTP响应失败：", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(body) == 0 {
			break
		}
		fmt.Print(string(body))
		fmt.Fprintf(w, string(body))
		w.(http.Flusher).Flush()
	}
	return
}
