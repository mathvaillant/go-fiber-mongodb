package user

import (
	"context"
	"example/go-fiber-mongodb/database"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection string = "users"

type mongoDBRepository struct {
	mongodb *database.MongoInstance
}

// Create a new User Repository
func NewUserRepository(mongoInstance *database.MongoInstance) UserRepository {
	return &mongoDBRepository{
		mongodb: mongoInstance,
	}
}

// GetUsers returns all users from the database
func (r *mongoDBRepository) GetUsers(ctx context.Context) ([]*User, error) {
	var items []*User

	collection := r.mongodb.Client.Database(r.mongodb.DatabaseName).Collection(userCollection)
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var item User
		if err := cursor.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

// GetUser returns a user from the database by ID
func (r *mongoDBRepository) GetUser(ctx context.Context, userID string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user := &User{}

	collection := r.mongodb.Client.Database(r.mongodb.DatabaseName).Collection(userCollection)
	filter := bson.M{"id": userID} // ?? What is bson???

	err := collection.FindOne(ctx, filter).Decode(user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, errors.Wrap(err, "repository research")
	}
	return user, nil
}

func (r *mongoDBRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

	defer cancel()

	collection := r.mongodb.Client.Database(r.mongodb.DatabaseName).Collection(userCollection)

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			// TODO: Missing settings and bank accounts
		},
	)

	if err != nil {
		return nil, errors.Wrap(err, "Error writing to repository")
	}

	return user, nil
}

// UpdateUser updates a user in the database
func (r *mongoDBRepository) UpdateUser(ctx context.Context, userID string, user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

	defer cancel()

	collection := r.mongodb.Client.Database(r.mongodb.DatabaseName).Collection(userCollection)

	updates := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id": userID},
		updates,
	)

	if err != nil {
		return nil, err
	}

	return user, err
}

// DeleteUser deletes a user from the database
func (r *mongoDBRepository) DeleteUser(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	filter := bson.M{"id": userID}

	collection := r.mongodb.Client.Database(r.mongodb.DatabaseName).Collection(userCollection)
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
