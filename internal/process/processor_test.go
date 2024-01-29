package process

import (
	"io"
	"strings"
	"testing"

	"github.com/matnich89/osrs-track-search/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertStatsToHighScores(t *testing.T) {
	tests := []struct {
		name         string
		character    string
		input        string
		expected     *model.CharStats
		expectingErr bool
	}{
		{
			name:      "Valid Input",
			character: "TestCharacter",
			input:     "100,50,2000\n200,60,3000\n",
			expected: &model.CharStats{
				Character: "TestCharacter",
				Stats: []model.Stat{
					{Skill: "Overall", Rank: 100, Level: 50, Xp: 2000},
					{Skill: "Attack", Rank: 200, Level: 60, Xp: 3000},
				},
			},
			expectingErr: false,
		},
		{
			name:         "Invalid Input",
			character:    "TestCharacter",
			input:        "bad,data",
			expected:     &model.CharStats{},
			expectingErr: true,
		},
		{
			name:         "Empty Input",
			character:    "TestCharacter",
			input:        "",
			expected:     &model.CharStats{Character: "TestCharacter", Stats: nil},
			expectingErr: true,
		},
	}

	processor := NewProcessor()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := io.NopCloser(strings.NewReader(test.input))
			result, err := processor.convertStatsToHighScores(test.character, reader)
			if test.expectingErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
