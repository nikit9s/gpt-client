# GPT Client Library

## Description
This project is a library for interacting with the GPT API. The library aims to encapsulate the logic of working with the API key and interacting with the server, as well as supporting sessions for fine-tuning models and storing session IDs.

## Main functions
- Generate a response.
- Start a new session.
- Generate a response using a session and settings.
```bash
go get github.com/nikit9s/gpt-client
go mod tidy
```
## Usage Example
```golang
package main

import (
    "fmt"
    "github.com/nikit9s/gpt-client/pkg/gptclient"
)

func main() {
    apiKey := ""
    client := gptclient.NewClient(apiKey)

    // Setting the system prompt
    client.SetSystemPrompt("You are a helpful assistant that provides concise and accurate answers.")

    config := gptclient.Config{
        Model:       "gpt-3.5-turbo",
        MaxTokens:   100,
        Temperature: 0.7,
    }

    // Starting a new session
    sessionID, err := client.StartSession("Hello, world!")
    if err != nil {
        fmt.Println("Error starting session:", err)
        return
    }

    // Generating a response using the session
    response, err := client.GenerateResponseWithSession("Tell me a joke.", sessionID, config)
    if err != nil {
        fmt.Println("Error generating response:", err)
        return
    }

    fmt.Println("Response:", response)
}
```

## Contributing

# Feel free to submit issues or pull requests for any improvements or fixes.

