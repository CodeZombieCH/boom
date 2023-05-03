package utils

import (
	"canonical/assessment/client"
	"encoding/json"
	"fmt"
	"os"
)

func PrintErrorJSON(s string) {
	PrintJSON(struct {
		Error string `json:"error"`
	}{
		Error: s,
	})
}

func PrintJSON(s interface{}) {
	json, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(json))
}

func CreateApiClient() client.ApiClient {
	return client.NewApiClient("http://localhost:8080/api")
}
