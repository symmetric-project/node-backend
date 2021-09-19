package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/symmetric-project/ngen"
	"github.com/symmetric-project/node-backend/graph/generated"
	"github.com/symmetric-project/node-backend/graph/model"
	"github.com/symmetric-project/node-backend/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	builder := SQ.Insert(`post`).Columns(`title`, `link`, `delta`).Values(newPost.Title, newPost.Link, newPost.Delta)
	var post model.Post
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &post, err
	}
	err = pgxscan.Get(context.Background(), DB, &post, query, args...)
	return &post, err
}

func (r *mutationResolver) CreateNode(ctx context.Context, newNode model.NewNode) (*model.Node, error) {
	builder := SQ.Insert(`node`).Columns(`name`, `access`, `nsfw`).Values(newNode.Name, newNode.Access, newNode.Nsfw).Suffix(`RETURNING *`)
	var node model.Node
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &node, err
	}
	err = pgxscan.Get(context.Background(), DB, &node, query, args...)
	return &node, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.User, error) {
	/* ctxUser, _ := ctx.Value("user").(model.User)
	log.Println((ctxUser)) */

	builder := SQ.Insert(`"user"`)
	if newUser.Name != nil {
		builder = builder.Columns(`"name"`).Values(newUser.Name)
	} else {
		builder = builder.Columns(`"name"`).Values(ngen.Generate())
	}

	builder = builder.Suffix(`RETURNING *`)

	var user model.User
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &user, err
	}
	err = pgxscan.Get(context.Background(), DB, &user, query, args...)
	return &user, err
}

func (r *queryResolver) Node(ctx context.Context, name string) (*model.Node, error) {
	var node model.Node
	builder := SQ.Select(`*`).From(`node`).Where(`name=$1`, name)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &node, err
	}
	err = pgxscan.Get(context.Background(), DB, &node, query, args...)
	return &node, err
}

func (r *queryResolver) Nodes(ctx context.Context) ([]*model.Node, error) {
	var nodes []*model.Node
	builder := SQ.Select(`*`).From(`node`)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return nodes, err
	}
	err = pgxscan.Select(context.Background(), DB, &nodes, query, args...)
	return nodes, err
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	builder := SQ.Select(`*`).From(`node`).Where(`id=$1`, id)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &post, err
	}
	err = pgxscan.Get(context.Background(), DB, &post, query, args...)
	return &post, err
}

func (r *queryResolver) Posts(ctx context.Context, nodeName *string) ([]*model.Post, error) {
	var posts []*model.Post
	builder := SQ.Select(`*`).From(`post`)
	if nodeName != nil {
		builder = builder.Where(`node_name=$1`, *nodeName)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return posts, err
	}
	err = pgxscan.Select(context.Background(), DB, &posts, query, args...)
	return posts, err
}

func (r *queryResolver) User(ctx context.Context, name *string) (*model.User, error) {
	var user model.User
	builder := SQ.Select(`*`).From(`"user"`)
	if name != nil {
		builder = builder.Where(`name=$1`, name)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &user, err
	}
	err = pgxscan.Get(context.Background(), DB, &user, query, args...)
	return &user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
