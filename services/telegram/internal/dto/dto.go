package dto

type Message struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type Request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	ID              int64  `json:"id"`
}

type AddPeerResponse struct {
	Message   Message `json:"message"`
	PublicKey string  `json:"public_key"`
}
