package main

type Post struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type User struct {
	ID    string `json:"id" db:"id"`
	Karma int    `json:"karma" db:"karma"`
}
