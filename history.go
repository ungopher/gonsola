package gonsola

import "github.com/google/uuid"

type history struct {
	ID    string
	Kind  string // info, warn, error, start, box, success, fail
	Lines int
}

func createHistory(kind string, lines int) history {
	return history{
		ID:    uuid.New().String(),
		Kind:  kind,
		Lines: lines,
	}
}
