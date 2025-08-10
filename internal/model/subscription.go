package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Id          uuid.UUID `json:"id"`
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
	EndDate     *string   `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" validate:"required,min=3,max=150"`
	Price       int       `json:"price" validate:"required,numeric"`
	UserID      uuid.UUID `json:"user_id" validate:"required,uuid"`
	StartDate   string    `json:"start_date" validate:"required"`
	EndDate     *string   `json:"end_date"`
}

type UpdateSubscriptionRequest struct {
	Id          *uuid.UUID `json:"id" validate:"required,uuid"`
	ServiceName string     `json:"service_name" validate:"omitempty,min=3,max=150"`
	Price       int        `json:"price" validate:"omitempty,numeric"`
	UserID      uuid.UUID  `json:"user_id" validate:"omitempty,uuid"`
	StartDate   string     `json:"start_date" validate:"omitempty"`
	EndDate     *string    `json:"end_date"`
}

type GetSubscriptionByIdRequest struct {
	Id string `json:"id" validate:"required,uuid"`
}

type DeleteSubscriptionByIdRequest struct {
	Id string `json:"id" validate:"required,uuid"`
}

type GetTotalSummByFilterRequest struct {
	ServiceName string `json:"service_name" validate:"omitempty,min=3,max=150"`
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	FromDate    string `json:"from_date" validate:"omitempty"`
	ToDate      string `json:"to_date"`
}

type GetTotalSummByFilterResponse struct {
	TotalSumm int `json:"total_sum"`
}
