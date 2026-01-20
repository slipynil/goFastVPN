package service

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

func NewHttpClient(endpoint string) client {
	httpClient := &http.Client{}
	return client{httpClient, endpoint}
}

// AddPeer adds a new peer, use method post, and returns the response body with publicKey
func (c *client) AddPeer(virtualEndpoint string, id int64) (dto.AddPeerResponse, error) {

	// parse request body
	req := dto.Request{
		VirtualEndpoint: virtualEndpoint,
		ID:              id,
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

func (c *client) DeletePeer(id string) error {
	url := fmt.Sprintf("%s/peers/%s", c.url, id)

	req, err := http.NewRequest("DELETE", url, nil)
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete peer, status %d", resp.StatusCode)
	}

	return nil
}

func (c *client) DownloadConfFile(publicKey string) ([]byte, error) {
	url := fmt.Sprintf("%s/peers/%s/config", c.url, publicKey)

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

	return data, nil
}
