package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	defaultHeaders = map[string]string{
		"Accept":                       "application/json",
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
		"Access-Control-Allow-Headers": "Content-Type, Authorization, ApiKey",
	}
)

func GetRequest(url string, defaultHeaders map[string]string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	// Buat request GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error membuat request: %v", err)
	}

	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error mengirim request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error membaca response: %v", err)
	}

	return body, nil
}

func PostRequest(url string, payload []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error membuat request: %v", err)
	}

	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error mengirim request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error membaca response: %v", err)
	}

	return body, nil
}

func PostRefreshToken(link string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	basicAuth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("CONSUMER_KEY") + ":" + os.Getenv("CONSUMER_SECRET")))

	req, err := http.NewRequest("POST", link, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error membuat request: %v", err)
	}

	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+basicAuth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error mengirim request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error membaca response: %v", err)
	}

	return body, nil
}
