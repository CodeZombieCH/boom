package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publicationDate"`
	Edition         string    `json:"edition"`
	Description     string    `json:"description"`
	Genre           string    `json:"genre"`
}

type ApiClient struct {
	baseUrl string
}

func NewApiClient(baseUrl string) ApiClient {
	return ApiClient{baseUrl: baseUrl}
}

func (c *ApiClient) ListBooks() ([]Book, error) {
	// Wait for new mail to arrive
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("%s/books", c.baseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//t.Fatalf("Receiving test mail failed with HTTP status code %v", resp.StatusCode)
		return nil, err
	}

	bodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		//t.Fatalf("Failed to read response: %v", err)
		return nil, err
	}
	var body []Book
	if err := json.Unmarshal(bodyRaw, &body); err != nil {
		//t.Fatalf("Failed to parse response: %v", err)
		return nil, err
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
		//t.Fatalf("Failed to read response: %v", err)
		return nil, err
	}
	var body Book
	if err := json.Unmarshal(bodyRaw, &body); err != nil {
		//t.Fatalf("Failed to parse response: %v", err)
		return nil, err
	}

	return &body, nil
}
