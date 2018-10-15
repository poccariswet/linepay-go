package linepay

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type LinePay struct {
	ChannelID     string
	ChannelSecret string
	Hostname      string
	IsSandBox     bool
}

const (
	ApiVersion         = "v2"
	SandboxApiHostname = "sandbox-api-pay.line.me"
)

func New(channelID, channelSecret string) *LinePay {
	return &LinePay{
		ChannelID:     channelID,
		ChannelSecret: channelSecret,
		IsSandBox:     true,
	}
}

func (pay *LinePay) Post(url string, data []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("X-LINE-ChannelId", pay.ChannelID)
	req.Header.Set("X-LINE-ChannelSecret", pay.ChannelSecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client do err")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "readall err")
	}

	return respBody, nil
}
