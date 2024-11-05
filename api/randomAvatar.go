package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Avatar struct {
	Counts int `json:"counts"`
}

type Response struct {
	Avatar Avatar `json:"avatar"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// 发送 HTTP GET 请求
	resp, err := http.Get("https://static.mggo.xyz/config.json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 解析 JSON 响应
	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// 生成范围 [a, b] 内的随机数
	randomNumber := rand.Intn(data.Avatar.Counts) + 1
	redirect := fmt.Sprintf("https://static.mggo.xyz/avatar/%d.gif", randomNumber)
	w.Header().Set("Location", redirect)
	w.WriteHeader(http.StatusFound)
}
