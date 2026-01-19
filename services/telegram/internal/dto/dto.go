package dto

type Message struct {
	StatusCode int   `json:"status_code"`
	Error      error `json:"error"`
}

type Request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	FileName        string `json:"file_name"`
}

type AddPeerResponse struct {
	Message   Message `json:"message"`
	PublicKey string  `json:"public_key"`
	FilePath  string  `json:"file_path"`
}
