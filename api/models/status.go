package models

import "time"

type Status struct {
	Version string        `json:"version"`
	Uptime  time.Duration `json:"uptime" swaggertype:"integer"`
}