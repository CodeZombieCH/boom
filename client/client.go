package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiClient struct {
	baseUrl string
}

func NewApiClient(baseUrl string) ApiClient {
	return ApiClient{baseUrl: baseUrl}
}

func (c *ApiClient) ListBooks() ([]Book, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("%s/books", c.baseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("client: unexpected HTTP status code %v", resp.StatusCode)
	}

	bodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: failed to read response: %w", err)
	}
	var body []Book
	if err := json.Unmarshal(bodyRaw, &body); err != nil {
		return nil, fmt.Errorf("client: failed to parse response: %w", err)
	}

	return body, nil
}

func (c *ApiClient) CreateBook(book *Book) (*Book, error) {
	jsonPayload, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/books", c.baseUrl), "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("client: unexpected HTTP status code %v", resp.StatusCode)
	}

	bodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("client: failed to read response: %w", err)
	}
	var body Book
	if err := json.Unmarshal(bodyRaw, &body); err != nil {
		return nil, fmt.Errorf("client: failed to parse response: %w", err)
	}

	return &body, nil
}
