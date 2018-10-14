package linepay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Reservation struct {
	ProductName    string `json:"productName"`
	Amount         uint64 `json:"amount"`
	Currency       string `json:"currency"`
	ConfirmURL     string `json:"confirmUrl"`
	ConfirmURLType string `json:"confirmUrlType"`
	OrderID        string `json:"orderId"`
}

type PayMessage struct {
	ReturnMessage string      `json:"returnMessage"`
	Info          InfoMessage `json:"info"`
}

type InfoMessage struct {
	URL                PaymentURL `json:"paymentUrl"`
	TransactionId      int        `json:"transactionId"`
	PaymentAccessToken string     `json:"paymentAccessToken"`
}

type PaymentURL struct {
	Web string `json:"web"`
	App string `json:"app"`
}

func (pay *LinePay) Reserve(reservation Reservation) (*PayMessage, error) {
	url := fmt.Sprintf("https://%s/%s/payments/request", SandboxApiHostname, ApiVersion)
	data, err := json.Marshal(reservation)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal err")
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

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

	message := &PayMessage{}
	if err := json.Unmarshal(respBody, message); err != nil {
		return nil, errors.Wrap(err, "json unmarshal err")
	}

	return message, nil
}
