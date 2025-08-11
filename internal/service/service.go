package service

import "subscriptions_service/internal/model"

type SubscriptionServiceInterface interface {
	Create(model.CreateSubscriptionRequest) (string, error)
	GetById(model.GetSubscriptionByIdRequest) (model.Subscription, error)
	Update(model.UpdateSubscriptionRequest) (int64, error)
	Delete(model.DeleteSubscriptionByIdRequest) (string, error)
	GetAll() ([]model.Subscription, error)
	GetTotalSummByFilter(model.GetTotalSummByFilterRequest) (model.GetTotalSummByFilterResponse, error)
}
