package gpt

import (
	"context"
	"sync"

	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
)

// Client представляет собой клиента для работы с GPT API
type Client struct {
	apiKey       string
	client       *openai.Client
	systemPrompt string
	sessions     map[string][]openai.ChatCompletionMessage
	mu           sync.Mutex
}

// NewClient создает новый клиент с заданным API ключом
func NewClient(apiKey string) *Client {
	c := openai.NewClient(apiKey)
	return &Client{
		apiKey:   apiKey,
		client:   c,
		sessions: make(map[string][]openai.ChatCompletionMessage),
	}
}

// SetSystemPrompt устанавливает системный промт, который будет использоваться в каждом запросе
func (c *Client) SetSystemPrompt(prompt string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.systemPrompt = prompt
}

// StartSession начинает новую сессию и возвращает session_id
func (c *Client) StartSession(prompt string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	sessionID := generateSessionID()
	c.sessions[sessionID] = []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: c.systemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	return sessionID, nil
}

// GenerateResponse генерирует ответ на заданный запрос с указанной конфигурацией
func (c *Client) GenerateResponse(prompt string, cfg Config) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: c.systemPrompt,
		},
		{
			Role:    "user",
			Content: prompt,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:       cfg.Model,
		Messages:    messages,
		MaxTokens:   cfg.MaxTokens,
		Temperature: cfg.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateResponseWithSession генерирует ответ с использованием session_id и настроек
func (c *Client) GenerateResponseWithSession(prompt string, sessionID string, cfg Config) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	messages, ok := c.sessions[sessionID]
	if !ok {
		return "", ErrSessionNotFound
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: prompt,
	})

	req := openai.ChatCompletionRequest{
		Model:       cfg.Model,
		Messages:    messages,
		MaxTokens:   cfg.MaxTokens,
		Temperature: cfg.Temperature,
	}

	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	c.sessions[sessionID] = append(messages, resp.Choices[0].Message)

	return resp.Choices[0].Message.Content, nil
}

func generateSessionID() string {
	return uuid.New().String()
}
