package sssg

import (
	"encoding/json"
	"testing"
)

func TestParseAndReencodeCameraJSON(t *testing.T) {
	jsonData := `
	{
		"DINum": 0,
		"DONum": 0,
		"addedTime": 0,
		"audioCodec": 4,
		"channel": "1",
		"connectionOverSSL": true,
		"dsId": 0,
		"dsName": "Local host",
		"enableLowProfile": true,
		"enableRecordingKeepDays": true,
		"enableRecordingKeepSize": true,
		"enableSRTP": false,
		"fov": "",
		"highProfileStreamNo": 1,
		"id": 61,
		"idOnRecServer": 0,
		"ip": "192.168.1.1",
		"lowProfileStreamNo": 1,
		"mac": "-",
		"mediumProfileStreamNo": 1,
		"model": "RLC-811A",
		"newName": "Camera1",
		"port": 443,
		"postRecordTime": 5,
		"preRecordTime": 5,
		"recordPrefix": "",
		"recordSchedule": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
		"recordTime": 30,
		"recordingKeepDays": 30,
		"recordingKeepSize": "100",
		"status": 1,
		"stream1": {
			"bitrateCtrl": 2,
			"constantBitrate": "1024",
			"fps": 15,
			"quality": "5",
			"resolution": "2560x1440"
		},
		"tvStandard": 0,
		"userName": "User503",
		"vendor": "Reolink",
		"videoCodec": 3,
		"videoMode": ""
	}`

	// Decode JSON into Camera struct
	var camera Camera
	err := json.Unmarshal([]byte(jsonData), &camera)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Re-encode Camera struct back to JSON
	encodedData, err := json.Marshal(camera)
	if err != nil {
		t.Fatalf("Failed to encode JSON: %v", err)
	}

	// Decode re-encoded JSON into a map for normalized comparison
	var originalData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &originalData); err != nil {
		t.Fatalf("Failed to parse original JSON into map: %v", err)
	}

	var reEncodedData map[string]interface{}
	if err := json.Unmarshal(encodedData, &reEncodedData); err != nil {
		t.Fatalf("Failed to parse re-encoded JSON into map: %v", err)
	}

	// Compare the maps directly
	if !deepEqual(originalData, reEncodedData) {
		t.Errorf("Mismatch between original and re-encoded JSON:\nOriginal: %s\nRe-encoded: %s", jsonData, string(encodedData))
	}
}

// Helper function for deep comparison of maps
func deepEqual(a, b map[string]interface{}) bool {
	original, _ := json.Marshal(a)
	reEncoded, _ := json.Marshal(b)
	return string(original) == string(reEncoded)
}
