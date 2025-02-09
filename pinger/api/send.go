package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendToAPI(apiURL string, status PingStatus) error {
	data, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("error marshaling: %v", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error sending data to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unsuccessful response status from API: %s", resp.Status)
	}

	return nil
}