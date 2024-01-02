package model

type CharStats struct {
	Character string `json:"character"`
	Stats     []Stat `json:"stats"`
}

type Stat struct {
	Skill string `json:"skill"`
	Rank  int64  `json:"rank"`
	Level int64  `json:"level"`
	Xp    int64  `json:"xp"`
}
