package process

import (
	"bufio"
	"io"
	"osrs-track-search/model"
	"strconv"
	"strings"
)

type Processor interface {
}

type SkillsProcessor struct {
	skillsPlacement map[int]string
}

func NewProcessor() *SkillsProcessor {
	return &SkillsProcessor{skillsPlacement: map[int]string{
		0:  "Overall",
		1:  "Attack",
		2:  "Defence",
		3:  "Strength",
		4:  "Hitpoints",
		5:  "Ranged",
		6:  "Prayer",
		7:  "Magic",
		8:  "Cooking",
		9:  "Woodcutting",
		10: "Fletching",
		11: "Fishing",
		12: "Firemaking",
		13: "Crafting",
		14: "Smithing",
		15: "Mining",
		16: "Herblore",
		17: "Agility",
		18: "Thieving",
		19: "Slayer",
		20: "Farming",
		21: "Runecrafting",
		22: "Hunter",
		23: "Construction",
	},
	}
}

func (p *SkillsProcessor) Process(character string, body io.ReadCloser) (*model.CharStats, error) {
	stats, err := p.convertStatsToHighScores(character, body)

	if err != nil {
		return nil, err
	}

	return &stats, err
}

func (p *SkillsProcessor) convertStatsToHighScores(character string, body io.ReadCloser) (model.CharStats, error) {
	scanner := bufio.NewScanner(body)
	var i = 0
	var stats []model.Stat
	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n")
		tokens := strings.Split(str, ",")
		if len(tokens) == 3 {
			skill := p.skillsPlacement[i]
			i++
			if skill != "" {
				rank, _ := strconv.ParseInt(tokens[0], 10, 64)
				level, _ := strconv.ParseInt(tokens[1], 10, 64)
				xp, _ := strconv.ParseInt(tokens[2], 10, 64)
				stats = append(stats, model.Stat{
					Skill: skill,
					Rank:  rank,
					Level: level,
					Xp:    xp,
				})
			}
		}
	}

	char := model.CharStats{
		Character: character,
		Stats:     stats,
	}

	return char, nil
}
