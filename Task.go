package main

import "time"

type Task struct {
	Id int `json:"id"`
	Navn string `json:"navn"`
	Type string `json:"type"`
	LastDone time.Time `json:"last_done"`
	MidTime time.Time `json:"mid_time"`
	CritTime time.Time `json:"crit_time"`
}