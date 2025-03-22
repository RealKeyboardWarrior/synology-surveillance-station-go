package sssg

import "encoding/json"

type ActionRule struct {
	RuleType    int    `json:"ruleType"`
	ExtUrl      string `json:"extUrl"`
	ActDevName  string `json:"actDevName"`
	EvtDevName  string `json:"evtDevName"`
	ActType     int    `json:"actType"`
	ID          int    `json:"id"`
	ActID       int    `json:"actId"`
	ActSchedule string `json:"actSchedule"`
	ActDevID    int    `json:"actDevId"`
	ActTimes    int    `json:"actTimes"`
	EvtID       int    `json:"evtId"`
	ActRetItem  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"actRetItem"`
	Status      int    `json:"status"`
	UserName    string `json:"userName"`
	ActTimeDur  int    `json:"actTimeDur"`
	ActDsID     int    `json:"actDsId"`
	EvtSrc      int    `json:"evtSrc"`
	Password    string `json:"password"`
	ActTimeUnit int    `json:"actTimeUnit"`
	EvtDsID     int    `json:"evtDsId"`
	Name        string `json:"name"`
	ActItem     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"actItem"`
	ActSrc   int `json:"actSrc"`
	EvtDevID int `json:"evtDevId"`
	EvtItem  int `json:"evtItem"`
}

type IOModule struct {
	IP     string `json:"ip"`
	Mac    string `json:"mac"`
	Model  string `json:"model"`
	Port   int    `json:"port"`
	Vendor string `json:"vendor"`
}

type HomeModeInfo struct {
	ActRuleOn            bool         `json:"actrule_on"`
	ActRules             []ActionRule `json:"actrules"`
	Cameras              []Camera     `json:"cameras"`
	Custom1Det           int          `json:"custom1_det"`
	Custom1Di            int          `json:"custom1_di"`
	Custom2Det           int          `json:"custom2_det"`
	Custom2Di            int          `json:"custom2_di"`
	DualRecOff           bool         `json:"dual_rec_off"`
	GeoDelayTime         int          `json:"geo_delay_time"`
	GeoLat               float64      `json:"geo_lat"`
	GeoLng               float64      `json:"geo_lng"`
	GeoMobiles           []string     `json:"geo_mobiles"` // Assuming geo_mobiles is a list of strings
	GeoRadius            int          `json:"geo_radius"`
	IoModules            []IOModule   `json:"io_modules"`
	LastUpdateTime       int64        `json:"last_update_time"` // As per the large number in the JSON
	ModeSchedule         string       `json:"mode_schedule"`
	ModeScheduleNextTime int          `json:"mode_schedule_next_time"`
	ModeScheduleOn       bool         `json:"mode_schedule_on"`
	NotifyEventList      []struct {
		EventGroupType int `json:"eventGroupType"`
		EventType      int `json:"eventType"`
		Filter         int `json:"filter"`
	} `json:"notify_event_list"`
	NotifyOn               bool `json:"notify_on"`
	On                     bool `json:"on"`
	OneTimeDisableOn       bool `json:"onetime_disable_on"`
	OneTimeDisableTime     int  `json:"onetime_disable_time"`
	OneTimeEnableOn        bool `json:"onetime_enable_on"`
	OneTimeEnableTime      int  `json:"onetime_enable_time"`
	Reason                 int  `json:"reason"`
	RecSchCustomDetAppList []struct {
		Custom1AppDet int `json:"custom1_app_det"`
		Custom2AppDet int `json:"custom2_app_det"`
	} `json:"rec_sch_custom_det_app_list"`
	RecSchedule   string `json:"rec_schedule"`
	RecScheduleOn bool   `json:"rec_schedule_on"`
	StreamProfile string `json:"stream_profile"`
	StreamingOn   bool   `json:"streaming_on"`
	WifiSsid      string `json:"wifi_ssid"`
}

// HomeModeInfoJSON struct for parsing the intermediate representation
// It uses `interface{}` for cameras field to handle different possible formats
type HomeModeInfoJSON struct {
	ActRuleOn            bool        `json:"actrule_on"`
	ActRules             interface{} `json:"actrules"`
	Cameras              interface{} `json:"cameras"`
	Custom1Det           int         `json:"custom1_det"`
	Custom1Di            int         `json:"custom1_di"`
	Custom2Det           int         `json:"custom2_det"`
	Custom2Di            int         `json:"custom2_di"`
	DualRecOff           bool        `json:"dual_rec_off"`
	GeoDelayTime         int         `json:"geo_delay_time"`
	GeoLat               float64     `json:"geo_lat"`
	GeoLng               float64     `json:"geo_lng"`
	GeoMobiles           []string    `json:"geo_mobiles"`
	GeoRadius            int         `json:"geo_radius"`
	IoModules            interface{} `json:"io_modules"`
	LastUpdateTime       int64       `json:"last_update_time"`
	ModeSchedule         string      `json:"mode_schedule"`
	ModeScheduleNextTime int         `json:"mode_schedule_next_time"`
	ModeScheduleOn       bool        `json:"mode_schedule_on"`
	NotifyEventList      []struct {
		EventGroupType int `json:"eventGroupType"`
		EventType      int `json:"eventType"`
		Filter         int `json:"filter"`
	} `json:"notify_event_list"`
	NotifyOn               bool `json:"notify_on"`
	On                     bool `json:"on"`
	OneTimeDisableOn       bool `json:"onetime_disable_on"`
	OneTimeDisableTime     int  `json:"onetime_disable_time"`
	OneTimeEnableOn        bool `json:"onetime_enable_on"`
	OneTimeEnableTime      int  `json:"onetime_enable_time"`
	Reason                 int  `json:"reason"`
	RecSchCustomDetAppList []struct {
		Custom1AppDet int `json:"custom1_app_det"`
		Custom2AppDet int `json:"custom2_app_det"`
	} `json:"rec_sch_custom_det_app_list"`
	RecSchedule   string `json:"rec_schedule"`
	RecScheduleOn bool   `json:"rec_schedule_on"`
	StreamProfile string `json:"stream_profile"`
	StreamingOn   bool   `json:"streaming_on"`
	WifiSsid      string `json:"wifi_ssid"`
}

// UnmarshalJSON custom deserializer for HomeModeInfo that converts HomeModeInfoJSON
func (h *HomeModeInfo) UnmarshalJSON(data []byte) error {
	// Step 1: First, unmarshal into the intermediate HomeModeInfoJSON struct
	var temp HomeModeInfoJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Step 2: Handle the "cameras" field
	switch cameras := temp.Cameras.(type) {
	case string:
		if cameras == "-1" {
			h.Cameras = []Camera{} // If cameras is "-1", set it to an empty array
		}
	case []interface{}:
		// If cameras is an array, unmarshal it into the Cameras field
		for _, cam := range cameras {
			camData, _ := json.Marshal(cam)
			var camera Camera
			if err := json.Unmarshal(camData, &camera); err != nil {
				return err
			}
			h.Cameras = append(h.Cameras, camera)
		}
	}

	// Step 2b: Handle the "actionRules" field
	switch actionRules := temp.ActRules.(type) {
	case string:
		if actionRules == "-1" {
			h.ActRules = []ActionRule{} // If actionRules is "-1", set it to an empty array
		}
	case []interface{}:
		// If actionRules is an array, unmarshal it into the ActionRules field
		for _, cam := range actionRules {
			camData, _ := json.Marshal(cam)
			var actionRule ActionRule
			if err := json.Unmarshal(camData, &actionRule); err != nil {
				return err
			}
			h.ActRules = append(h.ActRules, actionRule)
		}
	}

	// Step 2c: Handle the "ioModules" field
	switch ioModules := temp.IoModules.(type) {
	case string:
		if ioModules == "" { // For some reason "" instead of "-1" here
			h.IoModules = []IOModule{}
		}
	case []interface{}:
		for _, ioModule := range ioModules {
			ioModuleData, _ := json.Marshal(ioModule)
			var ioModule IOModule
			if err := json.Unmarshal(ioModuleData, &ioModule); err != nil {
				return err
			}
			h.IoModules = append(h.IoModules, ioModule)
		}
	}

	// Step 3: Copy other fields from temp to h
	h.ActRuleOn = temp.ActRuleOn
	h.Custom1Det = temp.Custom1Det
	h.Custom1Di = temp.Custom1Di
	h.Custom2Det = temp.Custom2Det
	h.Custom2Di = temp.Custom2Di
	h.DualRecOff = temp.DualRecOff
	h.GeoDelayTime = temp.GeoDelayTime
	h.GeoLat = temp.GeoLat
	h.GeoLng = temp.GeoLng
	h.GeoMobiles = temp.GeoMobiles
	h.GeoRadius = temp.GeoRadius
	h.LastUpdateTime = temp.LastUpdateTime
	h.ModeSchedule = temp.ModeSchedule
	h.ModeScheduleNextTime = temp.ModeScheduleNextTime
	h.ModeScheduleOn = temp.ModeScheduleOn
	h.NotifyEventList = temp.NotifyEventList
	h.NotifyOn = temp.NotifyOn
	h.On = temp.On
	h.OneTimeDisableOn = temp.OneTimeDisableOn
	h.OneTimeDisableTime = temp.OneTimeDisableTime
	h.OneTimeEnableOn = temp.OneTimeEnableOn
	h.OneTimeEnableTime = temp.OneTimeEnableTime
	h.Reason = temp.Reason
	h.RecSchCustomDetAppList = temp.RecSchCustomDetAppList
	h.RecSchedule = temp.RecSchedule
	h.RecScheduleOn = temp.RecScheduleOn
	h.StreamProfile = temp.StreamProfile
	h.StreamingOn = temp.StreamingOn
	h.WifiSsid = temp.WifiSsid

	return nil
}

// MarshalJSON custom serializer for HomeModeInfo using HomeModeInfoJSON
func (h HomeModeInfo) MarshalJSON() ([]byte, error) {
	// Step 1: Prepare the HomeModeInfoJSON struct as an intermediate representation
	temp := HomeModeInfoJSON{
		ActRuleOn:              h.ActRuleOn,
		Custom1Det:             h.Custom1Det,
		Custom1Di:              h.Custom1Di,
		Custom2Det:             h.Custom2Det,
		Custom2Di:              h.Custom2Di,
		DualRecOff:             h.DualRecOff,
		GeoDelayTime:           h.GeoDelayTime,
		GeoLat:                 h.GeoLat,
		GeoLng:                 h.GeoLng,
		GeoMobiles:             h.GeoMobiles,
		GeoRadius:              h.GeoRadius,
		LastUpdateTime:         h.LastUpdateTime,
		ModeSchedule:           h.ModeSchedule,
		ModeScheduleNextTime:   h.ModeScheduleNextTime,
		ModeScheduleOn:         h.ModeScheduleOn,
		NotifyEventList:        h.NotifyEventList,
		NotifyOn:               h.NotifyOn,
		On:                     h.On,
		OneTimeDisableOn:       h.OneTimeDisableOn,
		OneTimeDisableTime:     h.OneTimeDisableTime,
		OneTimeEnableOn:        h.OneTimeEnableOn,
		OneTimeEnableTime:      h.OneTimeEnableTime,
		Reason:                 h.Reason,
		RecSchCustomDetAppList: h.RecSchCustomDetAppList,
		RecSchedule:            h.RecSchedule,
		RecScheduleOn:          h.RecScheduleOn,
		StreamProfile:          h.StreamProfile,
		StreamingOn:            h.StreamingOn,
		WifiSsid:               h.WifiSsid,
	}

	// Step 2: Handle special serialization of "cameras"
	if len(h.Cameras) == 0 {
		// If cameras is empty, set it to "-1"
		temp.Cameras = "-1"
	} else {
		// Otherwise, set cameras to the list of camera objects
		temp.Cameras = h.Cameras
	}

	// Step 2b: Handle special serialization of "cameras"
	if len(h.ActRules) == 0 {
		temp.ActRules = "-1"
	} else {
		temp.ActRules = h.ActRules
	}

	// Step 2c: Handle special serialization of "ioModules"
	if len(h.IoModules) == 0 {
		temp.IoModules = ""
	} else {
		temp.IoModules = h.IoModules
	}

	// Step 3: Marshal the intermediate HomeModeInfoJSON struct into JSON
	return json.Marshal(temp)
}
