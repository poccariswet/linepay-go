package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	linepay "github.com/soeyusuke/linepay-go"
)

const (
	userID = "U56470f69e2877678040ce7df336dc2eb"
)

var (
	hostname = os.Getenv("PAY_HOSTNAME")
)

func main() {
	pay := linepay.New(
		os.Getenv("PAY_CHANNEL_ID"),
		os.Getenv("PAY_CHANNEL_SECRET_KEY"),
	)

	reservation := linepay.Reservation{
		ProductName:    "Sample Product",
		Amount:         1,
		Currency:       "JPY",
		ConfirmURL:     fmt.Sprintf("https://%s/pay/confirm", hostname),
		ConfirmURLType: "SERVER",
		OrderID:        fmt.Sprintf("%s-%s", userID, time.Now().Unix()),
	}

	message, err := pay.Reserve(reservation)
	if err != nil {
		log.Fatalln(err)
	}
	// access to message.Info.Web
	log.Println(message)

	http.HandleFunc("/pay/confirm", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println(r.Method)
		}

		confirmation := linepay.Confirmation{
			Amount:   reservation.Amount,
			Currency: reservation.Currency,
		}

		if err := pay.Confirm(message.Info.TransactionID, confirmation); err != nil {
			log.Fatalln(err)
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalln(err)
	}
}
