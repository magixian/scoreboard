package ops

import (
	"errors"
	"scoreboard/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxtFileReader(t *testing.T) {
	testCases := []struct {
		name       string
		file       string
		expIsValid bool
		exp        types.Results
		expErr     error
	}{
		{
			name:       "success test",
			file:       "./testfiles/txtFileReaderTest.txt",
			expIsValid: true,
			exp: types.Results{
				{TeamA: types.Team{Name: "Lions"}, ScoreA: 3, TeamB: types.Team{Name: "Snakes"}, ScoreB: 3},
				{TeamA: types.Team{Name: "Tarantulas"}, ScoreA: 1, TeamB: types.Team{Name: "FC Awesome"}, ScoreB: 0},
				{TeamA: types.Team{Name: "Lions"}, ScoreA: 1, TeamB: types.Team{Name: "FC Awesome"}, ScoreB: 1},
				{TeamA: types.Team{Name: "Tarantulas"}, ScoreA: 3, TeamB: types.Team{Name: "Snakes"}, ScoreB: 1},
				{TeamA: types.Team{Name: "Lions"}, ScoreA: 4, TeamB: types.Team{Name: "Grouches"}, ScoreB: 0},
			},
		},
		{
			name:       "invalid file",
			file:       "./testfiles/noop.txt",
			expIsValid: false,
			expErr:     errors.New(""),
		},
		{
			name:       "invalid format on line",
			file:       "./testfiles/txtFileReaderInvalidFormatTest.txt",
			expIsValid: true,
			expErr:     errors.New("invalid result line on line 1: Lions 3, Snakes 3, Tarantulas 1, FC Awesome 0,"),
		},
		{
			name:       "no scores on line",
			file:       "./testfiles/txtFileReaderNoScoresTest.txt",
			expIsValid: true,
			expErr:     errors.New("no score found on line 1: FC Awesome , Snakes 3"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := NewTxtFileReader(tc.file)
			err := reader.IsValidFile()
			assert.Equal(t, tc.expIsValid, err == nil)
			if err != nil {
				return
			}

			res, err := reader.ReadResults()
			assert.Equal(t, tc.expErr, err)
			if err != nil {
				return
			}

			for i, r := range res {
				assert.Equal(t, tc.exp[i].TeamA.Name, r.TeamA.Name)
				assert.Equal(t, tc.exp[i].ScoreA, r.ScoreA)
				assert.Equal(t, tc.exp[i].TeamB.Name, r.TeamB.Name)
				assert.Equal(t, tc.exp[i].ScoreB, r.ScoreB)
			}
		})
	}
}
