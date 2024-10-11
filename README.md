# Bonus_assignment
# Twitter API Interaction with Go

## Introduction
This project demonstrates how to interact with the Twitter API using the Go programming language. The program can post and delete tweets using Twitter’s REST API and OAuth authentication.

## Setup Instructions

1. **Sign up for Twitter Developer Account**: 
   - Follow the steps in the Twitter Developer documentation to create an app and generate API keys.

2. **Generate API Keys**: 
   - Generate `API Key`, `API Secret Key`, `Access Token`, and `Access Token Secret`.

3. **Run the Program**:
   - Clone the repository.
   - Install dependencies: `go get github.com/dghubble/oauth1`
   - Run the program:
     ```bash
     go run main.go
     ```

## Program Details

### Post a New Tweet
The `postTweet()` function uses the `POST statuses/update` endpoint to create a new tweet.

### Delete a Tweet
The `deleteTweet()` function uses the `POST statuses/destroy/:id` endpoint to delete a tweet by its ID.

## Error Handling
The program handles invalid tokens, missing parameters, and invalid tweet IDs by checking API responses and outputting meaningful error messages.

