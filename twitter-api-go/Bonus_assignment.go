package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
)

func main() {
	tweetText := "Hello, All!"

	config := oauth1.NewConfig("TMZyTPKKdwghO9AMoZD9eattN", "MXOqwEl6BmMgPGDzN6IIHSHCiqnfXiNzjTDQ1yv3vr6hGiEosI")
	token := oauth1.NewToken("1844782891340857346-PCt0haIXwaF3x3t9rQu92gttDl6Ftt", "5RvviQwf09Dp3snAc3QZW4q453ZpGPLIQNpRMODsfBXNB")

	// Create an HTTP client that authenticates with OAuth1
	httpClient := config.Client(oauth1.NoContext, token)

	// Prepare the API request to post a tweet using Twitter API v2
	endpoint := "https://api.x.com/2/tweets"
	payload := map[string]string{"text": tweetText} // Prepare the JSON payload
	jsonPayload, err := json.Marshal(payload)       // Marshal it to JSON
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Make the POST request with the JSON payload
	resp, err := httpClient.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusCreated {
		var responseBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseBody)
		fmt.Printf("Tweet posted successfully! Response: %v\n", responseBody)
	} else {
		var responseBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseBody)
		fmt.Printf("Failed to post tweet: Status %d %s, Response: %v\n", resp.StatusCode, http.StatusText(resp.StatusCode), responseBody)
	}
}
