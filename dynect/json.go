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

type ResponseBlock struct {
	Status   string         `json:"string"`
	JobId    int            `json:"job_id,omitempty"`
	Messages []MessageBlock `json:"msgs,omitempty"`
}

type MessageBlock struct {
	Info      string `json:"INFO"`
	Source    string `json:"SOURCE"`
	ErrorCode string `json:"ERR_CD"`
	Level     string `json:"LVL"`
}

type LoginResponse struct {
	ResponseBlock
	Data LoginDataBlock `json:"data"`
}

type LoginDataBlock struct {
	Token   string `json:"token"`
	Version string `json:"version"`
}
