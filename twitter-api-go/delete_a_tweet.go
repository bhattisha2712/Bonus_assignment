package main

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
)

func deleteTweet(tweetID string) {
	config := oauth1.NewConfig("TMZyTPKKdwghO9AMoZD9eattN", "MXOqwEl6BmMgPGDzN6IIHSHCiqnfXiNzjTDQ1yv3vr6hGiEosI")
	token := oauth1.NewToken("1844782891340857346-PCt0haIXwaF3x3t9rQu92gttDl6Ftt", "5RvviQwf09Dp3snAc3QZW4q453ZpGPLIQNpRMODsfBXNB")
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/1.1/statuses/destroy/" + tweetID + ".json"
	req, _ := http.NewRequest("POST", url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting tweet:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Tweet deleted successfully")
}

func main() {
	deleteTweet("1844813741524693239") // Replace with actual Tweet ID
}
