package linepay

import (
	"encoding/json"
	"fmt"

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
	TransactionID      int        `json:"transactionId"`
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

	respBody, err := pay.Post(url, data)
	if err != nil {
		return nil, errors.Wrap(err, "linepay Pout err")
	}

	message := &PayMessage{}
	if err := json.Unmarshal(respBody, message); err != nil {
		return nil, errors.Wrap(err, "json unmarshal err")
	}

	return message, nil
}
