package github_api

import (
	"io"
	"net/http"
	"os"
)

func ReadFile(fileUrl *string) (string, error) {
	req, err := http.NewRequest("GET", *fileUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/text")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()

	defer resp.Body.Close()

	configFile, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(configFile), nil
}
