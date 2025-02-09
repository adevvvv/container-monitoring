package model

type PingStatus struct {
	ID          int       `json:"id"`
	IP          string    `json:"ip"`
	PingTime    float64   `json:"ping_time"`
	LastSuccess string    `json:"last_success"`
	Links       []Link    `json:"_links"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
