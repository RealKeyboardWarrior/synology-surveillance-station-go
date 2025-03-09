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

type Stream struct {
	BitrateCtrl     int    `json:"bitrateCtrl"`
	ConstantBitrate string `json:"constantBitrate"`
	FPS             int    `json:"fps"`
	Quality         string `json:"quality"`
	Resolution      string `json:"resolution"`
}

type Camera struct {
	DINum                   int    `json:"DINum"`
	DONum                   int    `json:"DONum"`
	AddedTime               int    `json:"addedTime"`
	AudioCodec              int    `json:"audioCodec"`
	Channel                 string `json:"channel"`
	ConnectionOverSSL       bool   `json:"connectionOverSSL"`
	DsID                    int    `json:"dsId"`
	DsName                  string `json:"dsName"`
	EnableLowProfile        bool   `json:"enableLowProfile"`
	EnableRecordingKeepDays bool   `json:"enableRecordingKeepDays"`
	EnableRecordingKeepSize bool   `json:"enableRecordingKeepSize"`
	EnableSRTP              bool   `json:"enableSRTP"`
	FOV                     string `json:"fov"`
	HighProfileStreamNo     int    `json:"highProfileStreamNo"`
	ID                      int    `json:"id"`
	IDOnRecServer           int    `json:"idOnRecServer"`
	IP                      string `json:"ip"`
	LowProfileStreamNo      int    `json:"lowProfileStreamNo"`
	MAC                     string `json:"mac"`
	MediumProfileStreamNo   int    `json:"mediumProfileStreamNo"`
	Model                   string `json:"model"`
	NewName                 string `json:"newName"`
	Port                    int    `json:"port"`
	PostRecordTime          int    `json:"postRecordTime"`
	PreRecordTime           int    `json:"preRecordTime"`
	RecordPrefix            string `json:"recordPrefix"`
	RecordSchedule          string `json:"recordSchedule"`
	RecordTime              int    `json:"recordTime"`
	RecordingKeepDays       int    `json:"recordingKeepDays"`
	RecordingKeepSize       string `json:"recordingKeepSize"`
	Status                  int    `json:"status"`
	Stream1                 Stream `json:"stream1"`
	TVStandard              int    `json:"tvStandard"`
	UserName                string `json:"userName"`
	Vendor                  string `json:"vendor"`
	VideoCodec              int    `json:"videoCodec"`
	VideoMode               string `json:"videoMode"`
}

func NewClient(baseURL, username, password string, insecureSkipVerify bool) *SurveillanceStationClient {
	// Create a custom HTTP client that skips TLS verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
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
