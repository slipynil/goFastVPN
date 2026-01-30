package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"telegram-service/internal/dto"
)

type client struct {
	http *http.Client
	url  string
}

func New(endpoint string) *client {
	return &client{&http.Client{}, endpoint}
}

// AddPeer adds a new peer, use method post, and returns the response body with publicKey
func (c *client) AddPeer(hostID int, telegramID int64) (dto.AddPeerResponse, error) {

	virtualEndpoint := fmt.Sprintf("10.66.66.%d/32", hostID)
	// parse request body
	req := dto.Request{
		VirtualEndpoint: virtualEndpoint,
		ID:              telegramID,
	}
	reqBytes, _ := json.Marshal(req)
	data := bytes.NewReader(reqBytes)

	url := fmt.Sprintf("%s/peers", c.url)

	// get response
	resp, err := c.http.Post(url, "application/json", data)
	if err != nil {
		return dto.AddPeerResponse{}, err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusCreated {
		return dto.AddPeerResponse{}, fmt.Errorf("failed to add peer, status %d", resp.StatusCode)
	}

	// decode response body
	respBody := dto.AddPeerResponse{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	return respBody, nil
}

func (c *client) DeletePeer(publicKey string) error {
	url := fmt.Sprintf("%s/peers", c.url)
	reqDto := dto.DelRequest{PublicKey: publicKey}
	byte, err := json.Marshal(reqDto)
	if err != nil {
		return err
	}
	buf := bytes.NewReader(byte)
	req, err := http.NewRequest("DELETE", url, buf)
	if err != nil {
		return err
	}

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check status code
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return fmt.Errorf("config already deleted, status %d", resp.StatusCode)
	default:
		return fmt.Errorf("failed to delete peer, status %d", resp.StatusCode)
	}

}

func (c *client) DownloadConfFile(telegramID int64) ([]byte, error) {
	url := fmt.Sprintf("%s/peers/%d/config", c.url, telegramID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download config file, status %d", resp.StatusCode)
	}

	// read body to buffer
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	fmt.Printf("загружен конфиг %s\n", string(data))

	return data, nil
}
