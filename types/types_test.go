package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintScoreBoard(t *testing.T) {
	testCases := []struct {
		name       string
		scoreBoard ScoreBoard
		exp        string
	}{
		{
			name: "nil",
			exp:  "",
		},
		{
			name: "simple board single line",
			scoreBoard: ScoreBoard{
				Rows: []ScoreBoardRow{
					{
						Position: 1,
						Team: Team{
							Name: "Tarantulas",
						},
						Points: 6,
					},
				},
			},
			exp: "1. Tarantulas, 6 pts\n",
		},
		{
			name: "bigger board multi line",
			scoreBoard: ScoreBoard{
				Rows: []ScoreBoardRow{
					{
						Position: 1,
						Team: Team{
							Name: "Tarantulas",
						},
						Points: 6,
					},
					{
						Position: 2,
						Team: Team{
							Name: "Lions",
						},
						Points: 5,
					},
					{
						Position: 3,
						Team: Team{
							Name: "FC Awesome",
						},
						Points: 1,
					},
					{
						Position: 4,
						Team: Team{
							Name: "Snakes",
						},
						Points: 1,
					},
					{
						Position: 5,
						Team: Team{
							Name: "Grouches",
						},
						Points: 0,
					},
				},
			},
			exp: `1. Tarantulas, 6 pts
2. Lions, 5 pts
3. FC Awesome, 1 pts
4. Snakes, 1 pts
5. Grouches, 0 pts
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tc.scoreBoard.PrintScoreBoard())
		})
	}
}
