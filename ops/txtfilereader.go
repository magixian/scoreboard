package ops

import (
	"fmt"
	"os"
	"scoreboard/types"
	"strconv"
	"strings"
)

type TxtFileReader struct {
	path string
}

func NewTxtFileReader(path string) FileReader {
	return &TxtFileReader{
		path: path,
	}
}

func (tfr TxtFileReader) IsValidFile() error {
	if tfr.path == "" {
		return ErrInvalidFilePath
	}
	if _, err := os.ReadFile(tfr.path); err != nil {
		return err
	}
	return nil
}

func (tfr TxtFileReader) ReadResults() (types.Results, error) {
	b, err := os.ReadFile(tfr.path)
	if err != nil {
		return types.Results{}, err
	}

	res := make(types.Results, 0)
	for i, l := range strings.Split(string(b), "\n") {
		ls := strings.Split(l, ",")
		if len(ls) != 2 {
			return types.Results{}, fmt.Errorf("invalid result line on line %d: %s", i+1, l)
		}

		// Read Team A
		ss := strings.SplitN(reverseString(ls[0]), " ", 2)
		if len(ss) != 2 {
			return types.Results{}, fmt.Errorf("invalid result line on line %d: %s", i+1, l)
		}

		// Validate score for team A
		if strings.TrimSpace(ss[0]) == "" {
			return types.Results{}, fmt.Errorf("no score found on line %d: %s", i+1, l)
		}
		scoreA, err := strconv.Atoi(reverseString(ss[0]))
		if err != nil {
			return types.Results{}, fmt.Errorf("no score found on line %d: %s", i+1, l)
		}

		// Get name for for team A
		teamA := strings.TrimSpace(reverseString(ss[1]))
		if teamA == "" {
			return types.Results{}, fmt.Errorf("no team name found on line %d: %s", i+1, l)
		}

		// Read Team B
		ss = strings.SplitN(reverseString(ls[1]), " ", 2)
		if len(ss) != 2 {
			return types.Results{}, fmt.Errorf("invalid result line on line %d: %s", i+1, ls[1])
		}

		// Validate score for team B
		if strings.TrimSpace(ss[0]) == "" {
			return types.Results{}, fmt.Errorf("no score found on line %d: %s", i+1, l)
		}
		scoreB, err := strconv.Atoi(reverseString(ss[0]))
		if err != nil {
			return types.Results{}, fmt.Errorf("no score found on line %d: %s", i+1, l)
		}

		// Get name for for team B
		teamB := strings.TrimSpace(reverseString(ss[1]))
		if teamB == "" {
			return types.Results{}, fmt.Errorf("no team name found on line %d: %s", i+1, l)
		}

		res = append(res, types.Result{
			TeamA:  types.Team{Name: teamA},
			ScoreA: scoreA,
			TeamB:  types.Team{Name: teamB},
			ScoreB: scoreB,
		})
	}
	return res, nil
}

func reverseString(s string) string {
	var ns string
	for _, c := range s {
		ns = string(c) + ns
	}
	return ns
}
