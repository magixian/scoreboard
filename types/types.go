package types

import "fmt"

type Team struct {
	Name string
}

type Results []Result

type Result struct {
	TeamA, TeamB   Team
	ScoreA, ScoreB int
}

type ScoreBoard struct {
	Rows []ScoreBoardRow
}

func (sc ScoreBoard) PrintScoreBoard() string {
	var res string
	for _, r := range sc.Rows {
		res += fmt.Sprintf("%s\n", r)
	}
	return res
}

type ScoreBoardRow struct {
	Position int
	Team     Team
	Points   int
}

func (sbr ScoreBoardRow) String() string {
	return fmt.Sprintf("%d. %s, %d pts", sbr.Position, sbr.Team.Name, sbr.Points)
}
