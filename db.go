package main

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

func GetPostsFromDB(email string) ([]Post, error) {
	var posts []Post
	err := pgxscan.Get(context.Background(), DBPool, &posts, `SELECT * FROM "post"`)
	return posts, err
}
