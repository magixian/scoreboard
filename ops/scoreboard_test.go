package ops

import (
	"errors"
	"scoreboard/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateScores(t *testing.T) {
	testCases := []struct {
		name   string
		res    types.Results
		exp    []types.ScoreBoardRow
		expErr error
	}{
		{
			name:   "nil",
			expErr: ErrNoResultsProvided,
		},
		{
			name: "draw scores",
			res: types.Results{
				{
					TeamA: types.Team{Name: "Tarantulas"}, ScoreA: 0,
					TeamB: types.Team{Name: "Lions"}, ScoreB: 0,
				},
				{
					TeamA: types.Team{Name: "Lions"}, ScoreA: 0,
					TeamB: types.Team{Name: "FC Awesome"}, ScoreB: 0,
				},
			},
			exp: []types.ScoreBoardRow{
				{
					Position: 1,
					Team:     types.Team{Name: "Lions"},
					Points:   2,
				},
				{
					Position: 2,
					Team:     types.Team{Name: "FC Awesome"},
					Points:   1,
				},
				{
					Position: 2,
					Team:     types.Team{Name: "Tarantulas"},
					Points:   1,
				},
			},
		},
		{
			name: "scored results",
			res: types.Results{
				{
					TeamA: types.Team{Name: "Lions"}, ScoreA: 3,
					TeamB: types.Team{Name: "Snakes"}, ScoreB: 3,
				},
				{
					TeamA: types.Team{Name: "Tarantulas"}, ScoreA: 1,
					TeamB: types.Team{Name: "FC Awesome"}, ScoreB: 0,
				},
				{
					TeamA: types.Team{Name: "Lions"}, ScoreA: 1,
					TeamB: types.Team{Name: "FC Awesome"}, ScoreB: 1,
				},
				{
					TeamA: types.Team{Name: "Tarantulas"}, ScoreA: 3,
					TeamB: types.Team{Name: "Snakes"}, ScoreB: 1,
				},
				{
					TeamA: types.Team{Name: "Lions"}, ScoreA: 4,
					TeamB: types.Team{Name: "Grouches"}, ScoreB: 0,
				},
			},
			exp: []types.ScoreBoardRow{
				{
					Position: 1,
					Team:     types.Team{Name: "Tarantulas"},
					Points:   6,
				},
				{
					Position: 2,
					Team:     types.Team{Name: "Lions"},
					Points:   5,
				},
				{
					Position: 3,
					Team:     types.Team{Name: "FC Awesome"},
					Points:   1,
				},
				{
					Position: 3,
					Team:     types.Team{Name: "Snakes"},
					Points:   1,
				},
				{
					Position: 5,
					Team:     types.Team{Name: "Grouches"},
					Points:   0,
				},
			},
		},
		{
			name: "invalid results same team name",
			res: types.Results{
				{
					TeamA: types.Team{Name: "Lions"}, ScoreA: 0,
					TeamB: types.Team{Name: "Lions"}, ScoreB: 0,
				},
			},
			expErr: errors.New("duplicate team name for result: {TeamA:{Name:Lions} TeamB:{Name:Lions} ScoreA:0 ScoreB:0}"),
		},
		{
			name: "invalid results negative score",
			res: types.Results{
				{
					TeamA: types.Team{Name: "Grouches"}, ScoreA: 0,
					TeamB: types.Team{Name: "Lions"}, ScoreB: -1,
				},
			},
			expErr: errors.New("invalid negative score for result: {TeamA:{Name:Grouches} TeamB:{Name:Lions} ScoreA:0 ScoreB:-1}"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb, err := CalculateScores(tc.res)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
				return
			}
			assert.Equal(t, len(tc.exp), len(sb.Rows))
			for i, r := range sb.Rows {
				assert.Equal(t, tc.exp[i].Position, r.Position)
				assert.Equal(t, tc.exp[i].Team.Name, r.Team.Name)
				assert.Equal(t, tc.exp[i].Points, r.Points)
			}
		})
	}
}

func Test_getSortedScores(t *testing.T) {
	testCases := []struct {
		name string
		res  map[string]types.ScoreBoardRow
		exp  []types.ScoreBoardRow
	}{
		{
			name: "nil",
			res:  map[string]types.ScoreBoardRow{},
		},
		{
			name: "zero scores",
			res: map[string]types.ScoreBoardRow{
				"Tarantulas": {Team: types.Team{Name: "Tarantulas"}, Points: 0},
				"Lions":      {Team: types.Team{Name: "Lions"}, Points: 0},
				"FC Awesome": {Team: types.Team{Name: "FC Awesome"}, Points: 0},
			},
			exp: []types.ScoreBoardRow{
				{
					Position: 1,
					Team:     types.Team{Name: "FC Awesome"},
					Points:   0,
				},
				{
					Position: 1,
					Team:     types.Team{Name: "Lions"},
					Points:   0,
				},
				{
					Position: 1,
					Team:     types.Team{Name: "Tarantulas"},
					Points:   0,
				},
			},
		},
		{
			name: "scored results",
			res: map[string]types.ScoreBoardRow{
				"Tarantulas": {Team: types.Team{Name: "Tarantulas"}, Points: 6},
				"Lions":      {Team: types.Team{Name: "Lions"}, Points: 5},
				"FC Awesome": {Team: types.Team{Name: "FC Awesome"}, Points: 1},
				"Snakes":     {Team: types.Team{Name: "Snakes"}, Points: 1},
				"Grouches":   {Team: types.Team{Name: "Grouches"}, Points: 0},
			},
			exp: []types.ScoreBoardRow{
				{
					Position: 1,
					Team:     types.Team{Name: "Tarantulas"},
					Points:   6,
				},
				{
					Position: 2,
					Team:     types.Team{Name: "Lions"},
					Points:   5,
				},
				{
					Position: 3,
					Team:     types.Team{Name: "FC Awesome"},
					Points:   1,
				},
				{
					Position: 3,
					Team:     types.Team{Name: "Snakes"},
					Points:   1,
				},
				{
					Position: 5,
					Team:     types.Team{Name: "Grouches"},
					Points:   0,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := getSortedScoreBoardRows(tc.res)
			for i, r := range res {
				assert.Equal(t, tc.exp[i].Position, r.Position)
				assert.Equal(t, tc.exp[i].Team.Name, r.Team.Name)
				assert.Equal(t, tc.exp[i].Points, r.Points)
			}
		})
	}
}
