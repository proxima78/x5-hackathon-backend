package services

import (
	"fmt"
	"github.com/proxima78/x5-hackathon-backend/app"
	"github.com/proxima78/x5-hackathon-backend/models"
	httpModels "github.com/proxima78/x5-hackathon-backend/models/http"
	"github.com/proxima78/x5-hackathon-backend/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	StatusCodeAccountExists       = 100
	StatusCodeAccountDoesntExists = 101
	StatusCodeAccountCreated      = 102
)

type GeneralService struct {
	app            *app.App
	userRepository *repositories.UserRepository
}

func NewGeneralService(a *app.App, ur *repositories.UserRepository) *GeneralService {
	return &GeneralService{
		app:            a,
		userRepository: ur,
	}
}

func (gs *GeneralService) CheckAccountExists(checkAccountExistsRequest *httpModels.CheckAccountExistsRequest) (*httpModels.HTTPResponse, error) {
	_, err := gs.userRepository.GetByPhoneNumber(checkAccountExistsRequest.PhoneNumber)
	if err == mongo.ErrNoDocuments {
		_, err = gs.userRepository.GetByReferalCardID(checkAccountExistsRequest.ReferalCardID)
	} else if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return &httpModels.HTTPResponse{
			Status: StatusCodeAccountDoesntExists,
		}, nil
	} else if err != nil {
		return nil, err
	}
	return &httpModels.HTTPResponse{
		Status: StatusCodeAccountExists,
	}, nil
}

func (gs *GeneralService) CreateAccount(createAccountRequest httpModels.CreateAccountRequest) (*httpModels.HTTPResponse, error) {
	user := &models.User{}
	if createAccountRequest.PhoneNumber == "" {
		if createAccountRequest.ReferalCardID == "" {
			return nil, fmt.Errorf("all ids are empty")
		}
		user.ReferalCardID = createAccountRequest.ReferalCardID
	} else {
		user.PhoneNumber = createAccountRequest.PhoneNumber
	}
	user.Name = createAccountRequest.Name
	user.Age = createAccountRequest.Age
	user.Gender = createAccountRequest.Gender
	err := gs.userRepository.Put(user)
	if err != nil {
		return nil, err
	}
	return &httpModels.HTTPResponse{
		Status: StatusCodeAccountCreated,
	}, nil
}
