package dynect

/*
This struct represents the request body that would be sent to the DynECT API
for logging in and getting a session token for future requests.
*/
type LoginBlock struct {
	Username     string `json:"user_name"`
	Password     string `json:"password"`
	CustomerName string `json:"customer_name"`
}

// Type ResponseBlock holds the "header" information returned by any call to
// the DynECT API.
//
// All response-type structs should include this as an anonymous/embedded field.
type ResponseBlock struct {
	Status   string         `json:"status"`
	JobId    int            `json:"job_id,omitempty"`
	Messages []MessageBlock `json:"msgs,omitempty"`
}

// Type MessageBlock holds the message information from the server, and is
// nested within the ResponseBlock type.
type MessageBlock struct {
	Info      string `json:"INFO"`
	Source    string `json:"SOURCE"`
	ErrorCode string `json:"ERR_CD"`
	Level     string `json:"LVL"`
}

// Type LoginResponse holds the data returned by an HTTP POST call to
// https://api.dynect.net/REST/Session/.
type LoginResponse struct {
	ResponseBlock
	Data LoginDataBlock `json:"data"`
}

// Type LoginDataBlock holds the token and API version information from an HTTP
// POST call to https://api.dynect.net/REST/Session/.
//
// It is nested within the LoginResponse struct.
type LoginDataBlock struct {
	Token   string `json:"token"`
	Version string `json:"version"`
}

// RecordRequest holds the request body for a record create/update
type RecordRequest struct {
	RData DataBlock `json:"rdata"`
	TTL   string    `json:"ttl,omitempty"`
}

// CreateZoneBlock holds the request body for a zone create
type CreateZoneBlock struct {
	RName       string `json:"rname"`
	SerialStyle string `json:"serial_style,omitempty"`
	TTL         string `json:"ttl"`
}

// PublishZoneBlock holds the request body for a publish zone request
// https://help.dyn.com/update-zone-api/
type PublishZoneBlock struct {
	Publish bool `json:"publish"`
}

//PublishZoneResponseBlock holds the response for a zone publish request
type PublishZoneResponseBlock struct {
	Status string               `json:"status"`
	Data   PublishZoneDataBlock `json:"data"`
	JobID  int                  `json:"job_id"`
	Msgs   string               `json:"-"`
}

// PublishZoneDataBlock holds the response data filed for a zone publish request
type PublishZoneDataBlock struct {
	TaskID      string `json:"task_id"`
	Serial      int    `json:"serial"`
	SerialStyle string `json:"serial_style"`
	Zone        string `json:"zone"`
	ZoneType    string `json:"zone_type"`
}

// TaskStateResponse hold the response body of a get one DNS task request
// https://help.dyn.com/get-one-dns-task-api/
type TaskStateResponse struct {
	Status   string                `json:"status"`
	Blocking bool                  `json:"blocking"`
	Data     TaskStateResponseData `json:"data"`
}

// TaskStateResponseData hold the response data of a get one DNS task request
type TaskStateResponseData struct {
	TaskID       string `json:"task_id"`
	Name         string `json:"name"`
	CustomerName string `json:"customer_name"`
	ZoneName     string `json:"zone_name"`
	Status       string `json:"status"`
	StepCount    int    `json:"step_count"`
	TotalSteps   int    `json:"total_steps"`
	Blocking     bool   `json:"blocking"`
	Debug        string `json:"blocking,omitempty"`
	CreateTS     int    `json:"created_ts"`
	ModifiedTS   int    `json:"modified_ts"`
}
