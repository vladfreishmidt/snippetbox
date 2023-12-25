package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// INSERT  a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Return by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return the 10 most recently created
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
