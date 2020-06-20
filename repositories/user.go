package repositories

import (
	"context"

	"github.com/proxima78/x5-hackathon-backend/app"
	"github.com/proxima78/x5-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UserCollectionName = "users"
)

type UserRepository struct {
	app          *app.App
	dbCollection *mongo.Collection
	ctx          context.Context
}

func NewUserRepository(a *app.App) (*UserRepository, error) {
	ctx := context.Background()
	dbCol := a.DB.Collection(UserCollectionName)
	ur := &UserRepository{
		app:          a,
		dbCollection: dbCol,
		ctx:          ctx,
	}
	err := ur.initMongo()
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func (ur *UserRepository) initMongo() error {
	referalCardIDIndex := mongo.IndexModel{
		Keys: bson.M{
			"referalCardID": 1,
		}, Options: options.Index().SetSparse(true).SetUnique(true),
	}

	phoneNumberIndex := mongo.IndexModel{
		Keys: bson.M{
			"phoneNumber": 1,
		}, Options: options.Index().SetSparse(true).SetUnique(true),
	}

	_, err := ur.dbCollection.Indexes().CreateMany(ur.ctx, []mongo.IndexModel{referalCardIDIndex, phoneNumberIndex})
	return err
}

func (ur *UserRepository) Put(user *models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := ur.dbCollection.InsertOne(ur.ctx, user)
	return err
}

func (ur *UserRepository) GetByID(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := ur.dbCollection.FindOne(ur.ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (ur *UserRepository) GetByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	err := ur.dbCollection.FindOne(ur.ctx, bson.M{"phoneNumber": phoneNumber}).Decode(&user)
	return &user, err
}

func (ur *UserRepository) GetByReferalCardID(referalCardID string) (*models.User, error) {
	var user models.User
	err := ur.dbCollection.FindOne(ur.ctx, bson.M{"referalCardID": referalCardID}).Decode(&user)
	return &user, err
}
