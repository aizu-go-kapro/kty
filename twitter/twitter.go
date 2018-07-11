package twitter

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
	client *twitter.Client
}

func NewClient() *Twitter {

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials

	return &Twitter{
		client: client,
	}
}

func (t *Twitter) Send(user map[string]string, message string) error {

	params := &twitter.DirectMessageNewParams{
		ScreenName: user["recipient"],
		Text:       message,
	}
	_, _, err := t.client.DirectMessages.New(params)

	return err
}
