package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"telegram-service/internal/dto"
)

type client struct {
	http     *http.Client
	endpoint string
}

func NewHttpClient(endpoint string) client {
	httpClient := &http.Client{}
	return client{httpClient, endpoint}
}

func (c *client) AddPeer(virtualEndpoint, fileName string) (*dto.AddPeerResponse, error) {

	// parse request body
	req := dto.Request{
		VirtualEndpoint: virtualEndpoint,
		FileName:        fileName,
	}
	reqBytes, _ := json.Marshal(req)
	data := bytes.NewReader(reqBytes)

	url := fmt.Sprintf("http://%s/peer", c.endpoint)

	// get response
	resp, err := c.http.Post(url, "application/json", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Status:", resp.Status)
		return nil, errors.New("failed to add peer")
	}

	// decode response body
	respBody := dto.AddPeerResponse{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	fmt.Println("Peer added successfully")
	return &respBody, nil
}
