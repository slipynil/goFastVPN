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
	http *http.Client
	url  string
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

	// get response
	resp, err := c.http.Post(c.url, "application/json", data)
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

func (c *client) DeletePeer(publicKey string) error {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return err
	}

	// get url params
	q := req.URL.Query()
	q.Add("publicKey", publicKey)

	// set url params
	req.URL.RawQuery = q.Encode()

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status:", resp.Status)
		return errors.New("failed to delete peer")
	}

	fmt.Println("Peer deleted successfully")
	return nil
}
