package webhook

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	logger "github.com/NpoolPlatform/kunman/framework/logger"

	"github.com/NpoolPlatform/kunman/mal/payment/paypal"
)

func onPaypalCallback(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var event paypal.WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		logger.Sugar().Errorw(
			"Paypal Webhook",
			"Error", err,
		)
		return
	}

	cli, err := paypal.NewPaymentClient(context.Background())
	if err != nil {
		logger.Sugar().Errorw(
			"Paypal Webhook",
			"Error", err,
		)
		return
	}

	if err := cli.OnWebhook(context.Background(), &event); err != nil {
		logger.Sugar().Errorw(
			"Paypal Webhook",
			"Error", err,
		)
	}
}

func Initialize() {
	r := mux.NewRouter()

	r.HandleFunc("/api/webhook/v1/paypak/callback", onPaypalCallback).Methods("POST")
}
