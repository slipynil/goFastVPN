package transport

import "net/http"

type message struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
type response struct {
	Message message `json:"message"`
}

type request struct {
	VirtualEndpoint string `json:"virtual_endpoint"`
	FileName        string `json:"file_name"`
}

type createPeer struct {
	Message   message `json:"message"`
	PublicKey string  `json:"public_key"`
	FilePath  string  `json:"file_path"`
}

func newCreatePeer(publicKey, filePath string) createPeer {
	return createPeer{
		Message: message{
			StatusCode: http.StatusCreated,
		},
		PublicKey: publicKey,
		FilePath:  filePath,
	}
}

func newResp(statusCode int, err error) response {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return response{
		Message: message{
			StatusCode: statusCode,
			Error:      errMsg,
		},
	}
}
