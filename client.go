package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SurveillanceStationClient struct {
	BaseURL  string
	Session  string
	Username string
	Password string
	Client   *http.Client
}

func NewClient(baseURL, username, password string) *SurveillanceStationClient {
	// Create a custom HTTP client that skips TLS verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &SurveillanceStationClient{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
		Client:   client,
	}
}

// Login to the Surveillance Station
func (c *SurveillanceStationClient) Login() error {
	endpoint := fmt.Sprintf("%s/webapi/SurveillanceStation/ThirdParty/Auth/Login/v1", c.BaseURL)
	params := url.Values{}
	params.Set("account", c.Username)
	params.Set("passwd", c.Password)

	resp, err := c.Client.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Sid string `json:"sid"`
		} `json:"data"`
		Success bool `json:"success"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if !result.Success {
		return fmt.Errorf("login failed")
	}

	c.Session = result.Data.Sid
	fmt.Println("Login successful, session ID:", c.Session)
	return nil
}

type Camera struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	NewName string `json:"newName"`
}

// List available cameras and return them
func (c *SurveillanceStationClient) ListCameras() ([]Camera, error) {
	endpoint := fmt.Sprintf("%s/webapi/entry.cgi", c.BaseURL)
	params := url.Values{}
	params.Set("api", "SYNO.SurveillanceStation.Camera")
	params.Set("method", "List")
	params.Set("version", "9")
	params.Set("_sid", c.Session)

	resp, err := c.Client.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
		Data    struct {
			Cameras []Camera `json:"cameras"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf("failed to list cameras")
	}

	return result.Data.Cameras, nil
}

// TakeSnapshot returns the camera snapshot as bytes
func (c *SurveillanceStationClient) TakeSnapshot(camera Camera) ([]byte, error) {
	fmt.Printf("Taking snapshot for camera ID: %d, Name: %s\n", camera.ID, camera.NewName)

	endpoint := fmt.Sprintf("%s/webapi/entry.cgi", c.BaseURL)
	params := url.Values{}
	params.Set("api", "SYNO.SurveillanceStation.Camera")
	params.Set("method", "GetSnapshot")
	params.Set("version", "9")
	params.Set("id", fmt.Sprintf("%d", camera.ID))
	params.Set("_sid", c.Session)

	resp, err := c.Client.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("failed to take snapshot for camera ID %d: %v", camera.ID, err)
	}
	defer resp.Body.Close()

	// Read the image data into a byte slice
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read snapshot data for camera ID %d: %v", camera.ID, err)
	}

	return imgData, nil
}
