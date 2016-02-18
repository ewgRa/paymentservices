package aliasdirectlink

// Response for Alias Direct Link request
type Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}
