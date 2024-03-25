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
type GPTData struct {
	Status       string
	Text         string
	TotalTokens  int64
	OutputTokens int64
	InputTokens  int64
}

func GPT(url string, apiKey string, model string, question string, language string) (data GPTData) {
	// the struct about return data
	// 构建返回数据
	var GPTData GPTData
	requestBody := strings.NewReader(`{
		    "model": "gpt-3.5-turbo",
		   "messages": [{"role": "user", "content": "这是一篇文章:` + question + `。请你使用` + language + `写一则简短的文章摘要。"}]
		  }`)
	// Create POST query
	// 创建POST请求
	req, err := http.NewRequest("POST", url, requestBody)
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
	GPTData.Status = resp.Status

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
	respText := ((respContent["choices"]).([]interface{}))[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	GPTData.Text = respText

	// Get token status
	// 读取token状态
	respTotalTokens := (respContent["usage"]).(map[string]interface{})["total_tokens"].(float64)
	respOutputTokens := (respContent["usage"]).(map[string]interface{})["completion_tokens"].(float64)
	respInputTokens := (respContent["usage"]).(map[string]interface{})["prompt_tokens"].(float64)
	GPTData.TotalTokens = int64(respTotalTokens)
	GPTData.OutputTokens = int64(respOutputTokens)
	GPTData.InputTokens = int64(respInputTokens)
	return GPTData
}
