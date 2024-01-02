package model

type CharStats struct {
	Character string `json:"character"`
	Stats     []Stat `json:"stats"`
}

type Stat struct {
	Skill string `json:"skill"`
	Rank  int    `json:"rank"`
	Level int    `json:"level"`
	Xp    int    `json:"xp"`
}
