// Code written manually, THIS MAKES SENSE TO BE EDITED.

package model

type Node struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Tags   *[]*string `json:"tags"`
	Access NodeAccess `json:"access"`
	Nsfw   bool       `json:"nsfw"`
}
