package controller

import (
	"book-rentals/model"
	"gorm.io/gorm"
	"time"
)

type RequestUser struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	Age      int    `json: "age"`
	Sex      string `json: "sex"`
	ClientID int    `json: "client_id"`
}

func (req *RequestUser) toModel() model.User {
	return model.User{
		ID:       req.ID,
		Name:     req.Name,
		Age:      req.Age,
		Sex:      req.Sex,
		ClientID: req.ClientID,
	}
}

type ResponseUser struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ID        int            `json: "id"`
	Name      string         `json: "name"`
	Age       int            `json: "age"`
	Sex       string         `json: "sex"`
	ClientID  int            `json: "client_id"`
}

func NewResponse(modelUsers model.User) ResponseUser {
	return ResponseUser{
		CreatedAt: modelUsers.CreatedAt,
		UpdatedAt: modelUsers.UpdatedAt,
		DeletedAt: modelUsers.DeletedAt,
		ID:        modelUsers.ID,
		Name:      modelUsers.Name,
		Age:       modelUsers.Age,
		Sex:       modelUsers.Sex,
		ClientID:  modelUsers.ClientID,
	}
}

func newResponseArray(modelUsers []model.User) []ResponseUser {
	var response []ResponseUser
	for _, val := range modelUsers {
		response = append(response, NewResponse(val))
	}
	return response
}
