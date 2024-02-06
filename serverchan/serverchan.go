package serverchan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	ApiKey string
}

type ServerChan struct {
	config *Config
	client *http.Client
}

type ChannelType int

const (
	Test       ChannelType = 0
	Bark       ChannelType = 8
	FangTang   ChannelType = 9
	WeComApp   ChannelType = 66
	WeComGroup ChannelType = 1
	DingTalk   ChannelType = 2
	FeiShu     ChannelType = 3
	PushDeer   ChannelType = 18
)

func (c ChannelType) String() string {
	return fmt.Sprintf("%d", c)
}

type Message struct {
	Title    string
	Desp     string
	Short    string
	NoIP     bool
	Channels []ChannelType
	OpenID   string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PushID    string `json:"pushid"`
		ReadKey   string `json:"readkey"`
		Error     string `json:"error"`
		ErrorCode int    `json:"errno"`
	}
}

func New(config *Config) *ServerChan {
	return &ServerChan{
		config: config,
		client: http.DefaultClient,
	}
}

// https://sct.ftqq.com/sendkey#发起推送
func (s *ServerChan) Send(message *Message) (resp *Response, err error) {
	url := fmt.Sprintf("https://sctapi.ftqq.com/%s.send?title=%s", s.config.ApiKey, message.Title)
	if message.Desp != "" {
		url = fmt.Sprintf("%s&desp=%s", url, message.Desp)
	}
	if message.Short != "" {
		url = fmt.Sprintf("%s&short=%s", url, message.Short)
	}
	if message.NoIP {
		url = fmt.Sprintf("%s&noip=1", url)
	}
	if message.OpenID != "" {
		url = fmt.Sprintf("%s&openid=%s", url, message.OpenID)
	}
	var channels []string
	for _, c := range message.Channels {
		channels = append(channels, c.String())
	}
	if len(channels) > 0 {
		url = fmt.Sprintf("%s&channel=%s", url, strings.Join(channels, "|"))
	}
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return
	}
	res, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resp)
	return
}

// https://sct.ftqq.com/sendkey#查询推送状态
func (s *ServerChan) Query() {
	// https://sctapi.ftqq.com/push?id={pushid}&readkey={readkey}
}
