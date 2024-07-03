package gptclient

// Config представляет собой конфигурацию для запроса
type Config struct {
	Model       string
	MaxTokens   int
	Temperature float32
}
