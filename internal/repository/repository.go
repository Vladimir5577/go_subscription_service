package repository

import "subscriptions_service/internal/model"

type SubscriptionRepositoryInterface interface {
	Create(model.CreateSubscriptionRequest) (string, error)
	GetById(model.GetSubscriptionByIdRequest) (model.Subscription, error)
	Update(model.UpdateSubscriptionRequest) (int64, error)
	Delete(model.DeleteSubscriptionByIdRequest) (int64, error)
	GetAll() ([]model.Subscription, error)
	GetTotalSummByFilter(model.GetTotalSummByFilterRequest) (int, error)
}
