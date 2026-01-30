package dto

import (
	"encoding/base64"
	"encoding/json"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type DelRequest struct {
	PublicKey string `json:"public_key"`
}

type Request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	ID              int64  `json:"id"`
}

type AddPeerResponse struct {
	Message   Message `json:"message"`
	PublicKey string  `json:"public_key"`
}

type CallbackData struct {
	Action string `json:"action"`
}

type PaymentHandler struct {
	InvoicePayload string
	TotalAmount    int
	Currency       string
}

func DecodeCallbackData(raw string) (*CallbackData, error) {
	bs, err := base64.RawURLEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var data CallbackData
	if err := json.Unmarshal(bs, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// base64 кодирование для безопасной передачи действий в Inline кнопках
func EncodeCallbackData(action string) string {
	data := CallbackData{Action: action}
	bs, _ := json.Marshal(data)
	// можно добавить base64 encoding, если бояться спецсимволов
	return base64.RawURLEncoding.EncodeToString(bs)
}
