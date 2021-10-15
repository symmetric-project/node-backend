package model

type Post struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Link              *string `json:"link"`
	RawState          *string `json:"rawState"`
	NodeName          string  `json:"nodeName"`
	Slug              string  `json:"slug"`
	CreationTimestamp int     `json:"creationTimestamp"`
	AuthorID          string  `json:"authorId"`
	Author            *User   `json:"author"`
	Bases             int     `json:"bases"`
	ThumbnaillURL     *string `json:"thumbnaillUrl" db:"thumbnail_url"`
	ImageURL          *string `json:"imageUrl" db:"image_url"`
}
