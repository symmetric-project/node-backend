package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/symmetric-project/node-backend/errors"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/symmetric-project/node-backend/graph/generated"
	"github.com/symmetric-project/node-backend/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	builder := SQ.Insert("post").Columns("title", "link", "delta").Values(newPost.Title, newPost.Link, newPost.Delta)
	var post model.Post
	query, args, err := builder.ToSql()
	if err != nil {
		errors.Stacktrace(err)
		return &post, err
	}
	err = DB.QueryRow(query, args...).Scan(post)
	return &post, err
}

func (r *mutationResolver) CreateNode(ctx context.Context, newNode model.NewNode) (*model.Node, error) {
	builder := SQ.Insert("node").Columns("name", "access", "nsfw").Values(newNode.Name, newNode.Access, newNode.Nsfw).Suffix(`RETURNING *`)
	var node model.Node
	query, args, err := builder.ToSql()
	if err != nil {
		errors.Stacktrace(err)
		return &node, err
	}
	err = DB.Get(&node, query, args...)
	return &node, err
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	builder := SQ.Select("*").From("post")
	query, args, err := builder.ToSql()
	if err != nil {
		errors.Stacktrace(err)
		return posts, err
	}
	err = DB.Select(&posts, query, args...)
	return posts, err
}

func (r *queryResolver) Nodes(ctx context.Context) ([]*model.Node, error) {
	var nodes []*model.Node
	builder := SQ.Select("*").From("node")
	query, args, err := builder.ToSql()
	if err != nil {
		errors.Stacktrace(err)
		return nodes, err
	}
	err = DB.Select(&nodes, query, args...)
	return nodes, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
