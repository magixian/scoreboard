package ops

import (
	"errors"
	"fmt"
	"scoreboard/types"
	"sort"
)

var (
	ErrNoResultsProvided = errors.New("no results provided for calculations")
)

func CalculateScores(results types.Results) (types.ScoreBoard, error) {
	if results == nil {
		return types.ScoreBoard{}, ErrNoResultsProvided
	}

	// Validate results
	if err := validateResults(results); err != nil {
		return types.ScoreBoard{}, err
	}

	// Calculate team points
	res := make(map[string]types.ScoreBoardRow)
	for _, r := range results {
		teamA, ok := res[r.TeamA.Name]
		if !ok {
			teamA = types.ScoreBoardRow{Team: r.TeamA}
		}
		teamB, ok := res[r.TeamB.Name]
		if !ok {
			teamB = types.ScoreBoardRow{Team: r.TeamB}
		}

		if r.ScoreA > r.ScoreB {
			teamA.Points += 3
			teamB.Points += 0
		} else if r.ScoreA == r.ScoreB {
			teamA.Points += 1
			teamB.Points += 1
		}

		res[r.TeamA.Name] = teamA
		res[r.TeamB.Name] = teamB
	}

	// Return sorted results
	return types.ScoreBoard{
		Rows: getSortedScoreBoardRows(res),
	}, nil
}

func validateResults(results types.Results) error {
	for _, r := range results {
		// Validate same team for result
		if r.TeamA.Name == r.TeamB.Name {
			return fmt.Errorf("duplicate team name for result: %+v", r)
		}
		// Validate negative results
		if r.ScoreA < 0 || r.ScoreB < 0 {
			return fmt.Errorf("invalid negative score for result: %+v", r)
		}
	}

	return nil
}

func getSortedScoreBoardRows(res map[string]types.ScoreBoardRow) []types.ScoreBoardRow {
	// Sort teams by points
	keys := make([]string, 0)
	for k := range res {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		// Sort by points and to maintain predictability for same points, sort by team name alphabetical order
		if res[keys[i]].Points == res[keys[j]].Points {
			return res[keys[i]].Team.Name < res[keys[j]].Team.Name
		}
		return res[keys[i]].Points >= res[keys[j]].Points
	})

	var sbr []types.ScoreBoardRow
	for i, k := range keys {
		// When points are same with previous row, maintain position. By default index 0 is position 1
		p := i + 1
		if len(sbr) > 0 && sbr[i-1].Points == res[k].Points {
			p = sbr[i-1].Position
		}
		sbr = append(sbr, types.ScoreBoardRow{
			Position: p,
			Team:     res[k].Team,
			Points:   res[k].Points,
		})
	}
	return sbr
}
