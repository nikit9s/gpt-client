package gpt

// Config представляет собой конфигурацию клиента
type Config struct {
	Model       string
	MaxTokens   int
	Temperature float32
}

// NewConfig создает новую конфигурацию с заданным API ключом
func NewConfig(model string, maxTokens int, temperature float32) *Config {
	return &Config{
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}
}
