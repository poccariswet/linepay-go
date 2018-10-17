package linepay

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Confirmation struct {
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

func (pay *LinePay) Confirm(transactionID int, confirmation Confirmation) error {
	url := fmt.Sprintf("https://%s/%s/payments/%d/confirm", SandboxAPIHostname, APIVersion, transactionID)

	data, err := json.Marshal(confirmation)
	if err != nil {
		return errors.Wrap(err, "json marshal err")
	}

	respBody, err := pay.Post(url, data)
	if err != nil {
		return errors.Wrap(err, "linepay Post err")
	}
	log.Println(string(respBody))

	return nil
}
