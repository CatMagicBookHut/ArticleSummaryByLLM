package main

import (
	moepackage "ASBLLM/Moepackage"
	"fmt"
)

// func main() {
// 	// Input Tongyi Qianwen's API url.
// 	url := ""
// 	// Input your api key.
// 	apiKey := ""
// 	// Use model(qwen-turbo,qwen-plus and so on)
// 	model := "qwen-turbo"
// 	// The article content.
// 	question := ""
// 	// Which language.
// 	language := "Chinese"

// 	// Call function QianWen and import parameter.
// 	data := moepackage.QianWen(url, apiKey, model, question, language)
// 	// You can get the Status,Text,TotleToken,OutputTokens,InputTokens
// 	fmt.Println(data.Text)
// }

func main() {
	// Input GPT's API url.
	url := ""
	// Input your api key.
	apiKey := ""
	// Use model(qwen-turbo,qwen-plus and so on)
	model := "gpt-3.5-turbo"
	// The article content.
	question := "这是一段测试文字"
	// Which language.
	language := "Chinese"

	// Call function QianWen and import parameter.
	data := moepackage.GPT(url, apiKey, model, question, language)
	// You can get the Status,Text,TotleToken,OutputTokens,InputTokens
	fmt.Println(data)
}
