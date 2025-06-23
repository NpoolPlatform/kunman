package webhook

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	logger "github.com/NpoolPlatform/kunman/framework/logger"

	"github.com/NpoolPlatform/kunman/mal/payment/paypal"

	"github.com/go-chi/chi/v5"
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
		logger.Sugar().Errorw("Paypal Webhook", "Error", err)
		http.Error(w, "Invalid payload", http.StatusInternalServerError)
		return
	}

	cli, err := paypal.NewPaymentClient(context.Background())
	if err != nil {
		logger.Sugar().Errorw("Paypal Webhook", "Error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := cli.OnWebhook(context.Background(), &event); err != nil {
		logger.Sugar().Errorw("Paypal Webhook", "Error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func Initialize(r *chi.Mux) {
	r.Post("/api/webhook/v1/paypak/callback", onPaypalCallback)
}
