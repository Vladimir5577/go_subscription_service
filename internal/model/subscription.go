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
	ServiceName string    `json:"service_name" validate:"required,min=3,max=150" default:"Yandex Plus"`
	Price       int       `json:"price" validate:"required,numeric" default:"400"`
	UserID      uuid.UUID `json:"user_id" validate:"required,uuid" default:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string    `json:"start_date" validate:"required" default:"07-2025"`
	EndDate     *string   `json:"end_date" default:"07-2026"`
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
