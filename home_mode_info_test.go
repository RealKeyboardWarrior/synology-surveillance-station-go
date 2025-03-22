package sssg

import (
	"encoding/json"
	"testing"
)

func TestParseAndReencodeHomeModeInfoJSON(t *testing.T) {
	testCases := []struct {
		name     string
		jsonData string
	}{
		{
			name: "Example with all -1 or empty string arrays",
			jsonData: `{
      "actrule_on":false,
      "actrules":"-1",
      "cameras":"-1",
      "custom1_det":1,
      "custom1_di":1,
      "custom2_det":1,
      "custom2_di":1,
      "dual_rec_off":false,
      "geo_delay_time":60,
      "geo_lat":0,
      "geo_lng":0,
      "geo_mobiles":[
         
      ],
      "geo_radius":100,
      "io_modules":"",
      "last_update_time":1741624494295081,
      "mode_schedule":"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      "mode_schedule_next_time":-1,
      "mode_schedule_on":true,
      "notify_event_list":[
         {
            "eventGroupType":2,
            "eventType":3,
            "filter":0
         },
				 {
            "eventGroupType":2,
            "eventType":3,
            "filter":0
         }
      ],
      "notify_on":true,
      "on":false,
      "onetime_disable_on":false,
      "onetime_disable_time":0,
      "onetime_enable_on":false,
      "onetime_enable_time":0,
      "reason":1,
      "rec_sch_custom_det_app_list":[
         {
            "custom1_app_det":0,
            "custom2_app_det":0
         },
         {
            "custom1_app_det":0,
            "custom2_app_det":0
         }
      ],
      "rec_schedule":"111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
      "rec_schedule_on":false,
      "stream_profile":"1,1,1,1,1,1,0",
      "streaming_on":false,
      "wifi_ssid":""
   }`,
		},
		{
			name: "Example With All Arrays Filled In",
			jsonData: `{
      "actrule_on":false,
      "actrules": [
				{
					"ruleType": 0,
					"extUrl": "",
					"actDevName": "axis p3384",
					"evtDevName": "axis p3384",
					"actType": 0,
					"id": 82,
					"actId": 9,
					"actSchedule": "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
					"actDevId": 63,
					"actTimes": 1,
					"evtId": 8,
					"actRetItem": {
						"id": -1,
						"name": ""
					},
					"status": 2,
					"userName": "",
					"actTimeDur": 1,
					"actDsId": 0,
					"evtSrc": 0,
					"password": "",
					"actTimeUnit": 1,
					"evtDsId": 0,
					"name": "p3384 audio detected p3384 audio output",
					"actItem": {
						"id": 20,
						"name": "syno1"
					},
					"actSrc": 0,
					"evtDevId": 63,
					"evtItem": -1
				}
			],
      "cameras": [
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
				}
			],
      "custom1_det":1,
      "custom1_di":1,
      "custom2_det":1,
      "custom2_di":1,
      "dual_rec_off":false,
      "geo_delay_time":60,
      "geo_lat":0,
      "geo_lng":0,
      "geo_mobiles":[
         
      ],
      "geo_radius":100,
      "io_modules": [
				{
					"ip": "192.168.1.1",
					"mac": "MAC",
					"model": "fakemodel",
					"port": 0,
					"vendor": "Synology"
			  }
			],
      "last_update_time":1741624494295081,
      "mode_schedule":"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      "mode_schedule_next_time":-1,
      "mode_schedule_on":true,
      "notify_event_list":[
         {
            "eventGroupType":2,
            "eventType":3,
            "filter":0
         },
				 {
            "eventGroupType":2,
            "eventType":3,
            "filter":0
         }
      ],
      "notify_on":true,
      "on":false,
      "onetime_disable_on":false,
      "onetime_disable_time":0,
      "onetime_enable_on":false,
      "onetime_enable_time":0,
      "reason":1,
      "rec_sch_custom_det_app_list":[
         {
            "custom1_app_det":0,
            "custom2_app_det":0
         },
         {
            "custom1_app_det":0,
            "custom2_app_det":0
         }
      ],
      "rec_schedule":"111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
      "rec_schedule_on":false,
      "stream_profile":"1,1,1,1,1,1,0",
      "streaming_on":false,
      "wifi_ssid":""
   }`,
		},
	}

	// Iterate through all test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Decode JSON into HomeModeInfo struct
			var homeModeInfo HomeModeInfo
			err := json.Unmarshal([]byte(testCase.jsonData), &homeModeInfo)
			if err != nil {
				t.Fatalf("Failed to parse JSON for %s: %v", testCase.name, err)
			}

			// Re-encode HomeModeInfo struct back to JSON
			encodedData, err := json.Marshal(homeModeInfo)
			if err != nil {
				t.Fatalf("Failed to encode JSON for %s: %v", testCase.name, err)
			}

			// Decode re-encoded JSON into a map for normalized comparison
			var originalData map[string]interface{}
			if err := json.Unmarshal([]byte(testCase.jsonData), &originalData); err != nil {
				t.Fatalf("Failed to parse original JSON into map for %s: %v", testCase.name, err)
			}

			var reEncodedData map[string]interface{}
			if err := json.Unmarshal(encodedData, &reEncodedData); err != nil {
				t.Fatalf("Failed to parse re-encoded JSON into map for %s: %v", testCase.name, err)
			}

			// Compare the maps directly
			if !deepEqual(originalData, reEncodedData) {
				t.Errorf("Mismatch between original and re-encoded JSON for %s:\nOriginal: %s\nRe-encoded: %s", testCase.name, testCase.jsonData, string(encodedData))
			}
		})
	}
}
