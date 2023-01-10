package ops

import (
	"errors"
	"scoreboard/types"
)

var (
	ErrInvalidFilePath = errors.New("invalid file path")
)

type FileReader interface {
	IsValidFile() error
	ReadResults() (types.Results, error)
}

func ReadResults(path string) (types.Results, error) {
	// Currently using TxtFileReader otherwise create other FileReaders as needed
	reader := NewTxtFileReader(path)

	// Validate file
	if err := reader.IsValidFile(); err != nil {
		return nil, err
	}

	// Read results
	return reader.ReadResults()
}
