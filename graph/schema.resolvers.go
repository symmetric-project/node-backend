package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/symmetric-project/ngen"
	"github.com/symmetric-project/node-backend/graph/generated"
	"github.com/symmetric-project/node-backend/graph/model"
	"github.com/symmetric-project/node-backend/middleware"
	"github.com/symmetric-project/node-backend/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	builder := SQ.Insert(`post`).Columns(`title`, `link`, `delta`, `node_name`).Values(newPost.Title, newPost.Link, newPost.Delta, newPost.NodeName).Suffix("RETURNING *")
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
	if err == nil {
		resolverContext := middleware.GetResolverContext(ctx)
		jwt, _ := middleware.GenerateUserJWT(user.Name)
		cookie := middleware.NewCookie(jwt)
		middleware.SetCookie(*resolverContext.Writer, cookie)
	}
	return &user, err
}

func (r *queryResolver) Node(ctx context.Context, name string) (*model.Node, error) {
	var node model.Node
	builder := SQ.Select(`*`).From(`node`).Where(`name=$1`, name)
	query, args, err := builder.ToSql()
	log.Println(query)
	if err != nil {
		utils.Stacktrace(err)
		return &node, err
	}
	err = pgxscan.Get(context.Background(), DB, &node, query, args...)
	return &node, err
}

func (r *queryResolver) Nodes(ctx context.Context, substring *string) ([]*model.Node, error) {
	var nodes []*model.Node
	builder := SQ.Select(`*`).From(`node`)

	if substring != nil {
		builder = builder.Where(`name ~ $1`, *substring)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return nodes, err
	}
	err = pgxscan.Select(context.Background(), DB, &nodes, query, args...)
	return nodes, err
}

func (r *queryResolver) Post(ctx context.Context, id string, title string) (*model.Post, error) {
	var post model.Post
	builder := SQ.Select(`*`).From(`post`)
	builder = builder.Where(`id=$1 AND title=$2`, id, title)
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
	resolverContext := middleware.GetResolverContext(ctx)
	if resolverContext.JWT == nil {
		return &user, errors.New("unauthorized")
	}
	_, claims, err := middleware.VerifyJWT(*resolverContext.JWT)
	if err != nil {
		return &user, errors.New("unauthorized")
	}

	// If a name is not provided, it is replaced with the name from the JWT claims
	if name == nil {
		name = &claims.Id
	}

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
