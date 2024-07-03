package gptclient

import (
	"testing"
)

func TestClient(t *testing.T) { 
	apiKey := "" 
	client := NewClient(apiKey)

	// Установка системного промта
	client.SetSystemPrompt("You are a helpful assistant that provides concise and accurate answers.")

	config := Config{
		Model:       "gpt-3.5-turbo",
		MaxTokens:   100,
		Temperature: 0.7,
	}

	// Начало новой сессии
	sessionID, err := client.StartSession("Hello, world!")
	if err != nil {
		t.Fatalf("Error starting session: %v", err)
	}

	// Генерация ответа с использованием сессии
	response, err := client.GenerateResponseWithSession("Tell me a joke.", sessionID, config)
	if err != nil {
		t.Fatalf("Error generating response: %v", err)
	}

	if response == "" {
		t.Error("Expected a non-empty response")
	}
}

func TestGenerateResponse(t *testing.T) {
	apiKey := ""
	client := NewClient(apiKey)

	// Установка системного промта
	client.SetSystemPrompt("You are a helpful assistant that provides concise and accurate answers.")

	config := Config{
		Model:       "gpt-3.5-turbo",
		MaxTokens:   100,
		Temperature: 0.7,
	}

	response, err := client.GenerateResponse("Hello, world!", config)
	if err != nil {
		t.Fatalf("Error generating response: %v", err)
	}

	if response == "" {
		t.Error("Expected a non-empty response")
	}
}
