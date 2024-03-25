package moepackage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// LLM return data:status, content, total token, output token, input token
// AI返回的数据：状态，回复文本，总Token，输出Token，输入Token
type QianWenData struct {
	Status       string
	Text         string
	TotalTokens  int64
	OutputTokens int64
	InputTokens  int64
}

func QianWen(url string, apiKey string, model string, question string, language string) (data QianWenData) {
	// the struct about return data
	// 构建返回数据
	var qianWenData QianWenData

	// return body data
	// 构建请求的body数据
	requestBody := map[string]interface{}{
		"model": model,
		"input": map[string]interface{}{
			"prompt": "这是一篇文章:" + question + "。请你使用" + language + "写一则简短的文章摘要。",
		},
		"parameters": map[string]interface{}{
			"enable_search": true,
		},
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		panic(err)
	}

	// Create POST query
	// 创建POST请求
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBodyBytes)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	// Set header
	// 设置请求的header
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Sent query
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	// Get query status code
	// 获取响应的状态码
	qianWenData.Status = resp.Status

	// Get query content
	// 读取响应的内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		panic(err)
	}

	// Get LLM reply
	// 读取AI的回复
	var respContent map[string]interface{}
	err = json.Unmarshal([]byte(string(responseBody)), &respContent)
	if err != nil {
		panic(err)
	}
	respText := respContent["output"].(map[string]interface{})["text"].(string)
	qianWenData.Text = respText

	// Get token status
	// 读取token状态
	respTotalTokens := respContent["usage"].(map[string]interface{})["total_tokens"].(float64)
	respOutputTokens := respContent["usage"].(map[string]interface{})["output_tokens"].(float64)
	respInputTokens := respContent["usage"].(map[string]interface{})["input_tokens"].(float64)
	qianWenData.TotalTokens = int64(respTotalTokens)
	qianWenData.OutputTokens = int64(respOutputTokens)
	qianWenData.InputTokens = int64(respInputTokens)
	return qianWenData
}
