package gpt

import "errors"

var (
	ErrInvalidAPIKey            = errors.New("invalid API key")
	ErrInvalidRequest           = errors.New("invalid request")
	ErrResponseFailed           = errors.New("failed to get response from GPT API")
	ErrUnexpectedResponseFormat = errors.New("unexpected response format")
	ErrSessionNotFound          = errors.New("session not found")
)
