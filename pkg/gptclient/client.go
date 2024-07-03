package gptclient

import (
	"gpt-client/internal/gpt"
)

// Client представляет собой публичный интерфейс клиента для работы с GPT API
type Client struct {
	internalClient *gpt.Client
}

// NewClient создает новый клиент с заданным API ключом
func NewClient(apiKey string) *Client {
	return &Client{
		internalClient: gpt.NewClient(apiKey),
	}
}

// SetSystemPrompt устанавливает системный промт, который будет использоваться в каждом запросе
func (c *Client) SetSystemPrompt(prompt string) {
	c.internalClient.SetSystemPrompt(prompt)
}

// StartSession начинает новую сессию и возвращает session_id
func (c *Client) StartSession(prompt string) (string, error) {
	return c.internalClient.StartSession(prompt)
}

// GenerateResponse генерирует ответ на заданный запрос с указанной конфигурацией
func (c *Client) GenerateResponse(prompt string, cfg Config) (string, error) {
	return c.internalClient.GenerateResponse(prompt, gpt.Config(cfg))
}

// GenerateResponseWithSession генерирует ответ с использованием session_id и настроек
func (c *Client) GenerateResponseWithSession(prompt string, sessionID string, cfg Config) (string, error) {
	return c.internalClient.GenerateResponseWithSession(prompt, sessionID, gpt.Config(cfg))
}
