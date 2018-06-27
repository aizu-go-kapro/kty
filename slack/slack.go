package slack

import (
  "fmt"
  "os"
  "net/http"
  "encoding/json"
  "bytes"
  "strings"
)

type Message struct {
  Channel string `json:"channel"`
  Text string `json:"text"`
  AsUser bool `json:"as_user"`
}

type Slack struct {
  token string
}

func New()*Slack{
  return &Slack{
    token: os.Getenv("SLACK_TOKEN"),
  }
}

func (s *Slack)Send(token, message string)error{

  message := &Message {
    token,
		message,
    true,
  }
  m,err:= json.Marshal(message)

  body := bytes.NewReader(m)
  req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", body)
  if err != nil {
		return error
	}

  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
  req.Header.Set("Content-Type", "application/json")

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
		return error
  }
  defer resp.Body.Close()

}

