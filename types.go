package main


type InputData struct {
	Name      string `json:"name"`
	InputFrom []uint `json:"inputFrom"`
}

var vertices map[string]int