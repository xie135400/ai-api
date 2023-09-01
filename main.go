package main

import (
	"ai-api/utils"
	"log"
)

func main() {
	gpt := utils.NewChatGptTool("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")
	message := []utils.Gpt3Dot5Message{
		{
			Role:    "system",
			Content: "你是一个精通开发的资深工程师，熟悉全栈技术，任何问题都难不倒你",
		},
		{
			Role:    "user",
			Content: "帮我使用golang开发一个在线客服系统",
		},
	}
	res, err := gpt.ChatGPT3Dot5Turbo(message)
	log.Println(res, err)
}
