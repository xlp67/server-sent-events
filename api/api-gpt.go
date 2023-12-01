// api.go
package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sse/config"
	"sse/file"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type OpenAIRequest struct {
	Model     string        `json:"model"`
	Messages  []interface{} `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func OpenAiResponse(tokens int, pathFile string) string {
	apiKey := config.LoadConfig()
	request := OpenAIRequest{
		Model:     "gpt-3.5-turbo",
		Messages:  []interface{}{map[string]interface{}{"role": "system", "content": file.ReadFile(pathFile)}},
		MaxTokens: tokens,
	}
	request_body, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Erro no corpo da requisição: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(request_body))
	if err != nil {
		log.Println("Erro na requisição")
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	log.Println("Processando a resposta")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {log.Println("Erro na resposta do servidor")}
	defer resp.Body.Close()
	var openai_response OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&openai_response)
	if err != nil {log.Println("Erro ao processar o corpo da resposta")}

	return openai_response.Choices[0].Message.Content
}
