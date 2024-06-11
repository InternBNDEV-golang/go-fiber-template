package httpclient

import (
	"net/http"

	"github.com/goccy/go-json"
)

func GetIP() (string, error) {
	url := "https://api.ipify.org/?format=json"

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data["ip"].(string), nil

}
