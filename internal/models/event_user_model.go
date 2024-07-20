package models

import (
	"database/sql"
	"time"
)

var TYPE_EVENT_USER = "eventUser"

type EventUser struct {
	ID            int
	AccountNumber string
	Name          string
	PhoneNumber   string
	Email         string
	Password      string
	Status        string
	State         string
	Role          string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     sql.NullTime
}

func (eu *EventUser) ToResponse() *GetEventUserResponse {
	return &GetEventUserResponse{
		Type:          TYPE_EVENT_USER,
		Name:          eu.Name,
		AccountNumber: eu.AccountNumber,
		Email:         eu.Email,
		Status:        eu.Status,
		Role:          eu.Role,
	}
}

type GetEventUserResponse struct {
	Type          string `json:"type" example:"coolCategory"`
	Name          string `json:"name"`
	Email         string `json:"email,omitempty"`
	PhoneNumber   string `json:"phoneNumber,omitempty"`
	AccountNumber string `json:"accountNumber"`
	Role          string `json:"role"`
	Status        string `json:"status" example:"active"`
}

func (eu *CreateEventUserResponse) ToCreateEventUser() *CreateEventUserResponse {
	return &CreateEventUserResponse{
		Type:          TYPE_EVENT_USER,
		Name:          eu.Name,
		Email:         eu.Email,
		AccountNumber: eu.AccountNumber,
		Role:          eu.Role,
		Token:         "token",
		Status:        eu.Status,
	}
}

type (
	CreateEventUserRequest struct {
		Name  string `json:"name" validate:"required,min=1,max=50,nospecial,noStartEndSpaces" example:"Professionals"`
		Email string `json:"email" validate:"required,noStartEndSpaces,emailFormat" example:"jeremy@gmail.com"`
	}

	CreateEventUserResponse struct {
		Type          string `json:"type" example:"coolCategory"`
		Name          string `json:"name" example:"Profesionals"`
		Email         string `json:"email"`
		AccountNumber string `json:"accountNumber"`
		Role          string `json:"role"`
		Token         string `json:"token"`
		Status        string `json:"status" example:"active"`
	}
)

func (eu *CreateEventUserManualResponse) ToCreateEventUserManual() CreateEventUserManualResponse {
	return CreateEventUserManualResponse{
		Type:          TYPE_EVENT_USER,
		Name:          eu.Name,
		Email:         eu.Email,
		PhoneNumber:   eu.PhoneNumber,
		AccountNumber: eu.AccountNumber,
		Role:          eu.Role,
		Token:         eu.Token,
		Status:        eu.Status,
	}
}

type (
	CreateEventUserManualRequest struct {
		Name        string `json:"name" validate:"required,min=1,max=50,nospecial,noStartEndSpaces" example:"Professionals"`
		PhoneNumber string `json:"phoneNumber" validate:"omitempty,noStartEndSpaces,phoneFormat"`
		Email       string `json:"email" validate:"omitempty,noStartEndSpaces,emailFormat" example:"jeremy@gmail.com"`
		Password    string `json:"password" validate:"required,min=6,max=50,noStartEndSpaces" example:"Professionals"`
	}
	CreateEventUserManualResponse struct {
		Type          string `json:"type" example:"coolCategory"`
		Name          string `json:"name" example:"Profesionals"`
		Email         string `json:"email,omitempty"`
		PhoneNumber   string `json:"phoneNumber,omitempty"`
		AccountNumber string `json:"accountNumber"`
		Role          string `json:"role"`
		Token         string `json:"token"`
		Status        string `json:"status" example:"active"`
	}
)

func (eu *LoginEventUserManualResponse) ToLoginEventUserManual() LoginEventUserManualResponse {
	return LoginEventUserManualResponse{
		Type:          TYPE_EVENT_USER,
		Name:          eu.Name,
		Email:         eu.Email,
		PhoneNumber:   eu.PhoneNumber,
		AccountNumber: eu.AccountNumber,
		Role:          eu.Role,
		Token:         eu.Token,
		Status:        eu.Status,
	}
}

type (
	LoginEventUserManualRequest struct {
		Identifier string `json:"identifier" validate:"required,noStartEndSpaces"`
		Password   string `json:"password" validate:"required,noStartEndSpaces"`
	}
	LoginEventUserManualResponse struct {
		Type          string `json:"type" example:"coolCategory"`
		Name          string `json:"name"`
		Email         string `json:"email,omitempty"`
		PhoneNumber   string `json:"phoneNumber,omitempty"`
		AccountNumber string `json:"accountNumber"`
		Token         string `json:"token"`
		Role          string `json:"role"`
		Status        string `json:"status" example:"active"`
	}
)

// type (
// 	InquiryUserRequest struct {
// 		AccountNumber string         `json:"accountNumber" binding:"required"`
// 		Additional    AdditionalInfo `json:"additionalInfo"`
// 	}

// 	InquiryUserResponse struct {
// 		ResponseCode    string `json:"responseCode"`
// 		ResponseMessage string `json:"responseMessage"`
// 		AccountNumber   string `json:"accountNumber"`
// 		Name            string `json:"name"`
// 		State           string `json:"state"`
// 		Role            string `json:"role"`
// 		Email           string `json:"email"`
// 	}
// )
