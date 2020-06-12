package models

import "fmt"

type PracticeItem struct {
	tableName   struct{} `pg:"practice_item"`
	Id          int
	Text        string
	Description string
	Url         string
	IsCompleted bool
}

type Item struct {
	Id          int
	Text        string
	Description string
	Url         string
	IsCompleted bool
}

func (p Item) String() string {
	return fmt.Sprintf("Item<%d %s %v>", p.Id, p.Text, p.IsCompleted)
}