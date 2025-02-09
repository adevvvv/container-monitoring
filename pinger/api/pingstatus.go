package api
import "time"

type PingStatus struct {
	IP          string    `json:"ip"`
	PingTime    float64   `json:"ping_time"`
	LastSuccess time.Time `json:"last_success"`
}
