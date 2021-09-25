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
	"github.com/symmetric-project/node-backend/slug"
	"github.com/symmetric-project/node-backend/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	id := utils.NewOctid()
	slug := slug.Slugify(newPost.Title)
	creationTimestamp := utils.CurrentTimestamp()
	builder := SQ.Insert(`post`).Columns(`id`, `title`, `link`, `delta_ops`, `node_name`, `slug`, `creation_timestamp`).Values(id, newPost.Title, newPost.Link, newPost.DeltaOps, newPost.NodeName, slug, creationTimestamp).Suffix("RETURNING *")
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
		jwt, _ := middleware.NewUserJWT(user.ID)
		cookie := middleware.NewCookie(jwt)
		middleware.SetCookie(*resolverContext.Writer, cookie)
	}
	return &user, err
}

func (r *mutationResolver) CreateComment(ctx context.Context, newComment model.NewComment) (*model.Comment, error) {
	id := utils.NewOctid()
	creationTimestamp := utils.CurrentTimestamp()
	builder := SQ.Insert(`comment`).Columns(`id`, `post_id`, `creation_timestamp`, `delta_ops`, `user_id`, `post_slug`).Values(id, newComment.PostID, creationTimestamp, newComment.DeltaOps, newComment.UserID, newComment.PostSlug).Suffix(`RETURNING *`)
	var comment model.Comment
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &comment, err
	}
	err = pgxscan.Get(context.Background(), DB, &comment, query, args...)
	return &comment, err
}

func (r *queryResolver) Node(ctx context.Context, name string) (*model.Node, error) {
	var node model.Node
	builder := SQ.Select(`*`).From(`node`).Where(`name = $1`, name)
	query, args, err := builder.ToSql()
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

func (r *queryResolver) Post(ctx context.Context, id string, slug string) (*model.Post, error) {
	var post model.Post
	builder := SQ.Select(`*`).From(`post`)
	builder = builder.Where(`id = $1 AND slug=$2`, id, slug)
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
		builder = builder.Where(`node_name = $1`, *nodeName)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return posts, err
	}
	err = pgxscan.Select(context.Background(), DB, &posts, query, args...)
	return posts, err
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	var user model.User

	resolverContext := middleware.GetResolverContext(ctx)
	if resolverContext.JWT == nil {
		return &user, errors.New("no jwt in context")
	}
	_, claims, err := middleware.VerifyJWT(*resolverContext.JWT)
	if err != nil {
		return &user, errors.New("invalid jwt")
	}

	builder := SQ.Select(`*`).From(`"user"`)

	// If a name is not provided, it is replaced with the name from the JWT claims
	if id == nil {
		id = &claims.Id
	}
	builder = builder.Where(`id = $1`, *id)

	query, args, err := builder.ToSql()
	log.Println(query)
	log.Println(args[0])
	if err != nil {
		utils.Stacktrace(err)
		return &user, err
	}
	err = pgxscan.Get(context.Background(), DB, &user, query, args...)
	return &user, err
}

func (r *queryResolver) Users(ctx context.Context, nameSubstring string) ([]*model.User, error) {
	var users []*model.User

	resolverContext := middleware.GetResolverContext(ctx)
	if resolverContext.JWT == nil {
		return users, errors.New("no jwt in context")
	}
	_, _, err := middleware.VerifyJWT(*resolverContext.JWT)
	if err != nil {
		return users, errors.New("invalid jwt")
	}

	builder := SQ.Select(`*`).From(`comment`)
	builder = builder.Where(`name ~ $1`, nameSubstring)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return users, err
	}
	err = pgxscan.Select(context.Background(), DB, &users, query, args...)
	return users, err
}

func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment
	builder := SQ.Select(`*`).From(`comment`).Where(`id = $1`, id)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return &comment, err
	}
	err = pgxscan.Get(context.Background(), DB, &comment, query, args...)
	user, err := r.Resolver.Query().User(ctx, &comment.UserID)
	comment.Author = user
	return &comment, err
}

func (r *queryResolver) Comments(ctx context.Context, postID string, postSlug string) ([]*model.Comment, error) {
	var comments []*model.Comment
	builder := SQ.Select(`*`).From(`comment`).Where(`post_id = $1 AND post_slug = $2`, postID, postSlug)
	query, args, err := builder.ToSql()
	if err != nil {
		utils.Stacktrace(err)
		return comments, err
	}
	err = pgxscan.Select(context.Background(), DB, &comments, query, args...)
	for i, comment := range comments {
		ctx := context.WithValue(ctx, "resolverContext", middleware.SystemResolverContext)
		author, err := r.Resolver.Query().User(ctx, &comment.UserID)
		log.Println(comment.UserID)
		log.Println(err)
		/* if err != nil {
			return comments, err
		} */
		comments[i].Author = author
	}
	return comments, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
