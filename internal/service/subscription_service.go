package service

import (
	"fmt"
	"subscriptions_service/internal/helper"
	"subscriptions_service/internal/model"
	"subscriptions_service/internal/repository"
)

type SubscriptionServiceInterface interface {
	Create(model.CreateSubscriptionRequest) (string, error)
	GetById(model.GetSubscriptionByIdRequest) (model.Subscription, error)
	Update(model.UpdateSubscriptionRequest) (int64, error)
	Delete(model.DeleteSubscriptionByIdRequest) (string, error)
	GetAll() ([]model.Subscription, error)
	GetTotalSummByFilter(model.GetTotalSummByFilterRequest) (model.GetTotalSummByFilterResponse, error)
}

type SubscriptionService struct {
	subscriptionRepository repository.SubscriptionRepositoryInterface
}

func NewSubscriptionService(subscriptionRepository repository.SubscriptionRepositoryInterface) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *SubscriptionService) Create(subscription model.CreateSubscriptionRequest) (string, error) {
	startDateTime, err := helper.ConvertTimeStringToBeginningOfMounth(subscription.StartDate)
	if err != nil {
		return "", err
	}
	subscription.StartDate = startDateTime

	if subscription.EndDate != nil {
		endDateTime, err := helper.ConvertTimeStringToTheEndOfMounth(*subscription.EndDate)
		if err != nil {
			return "", err
		}
		subscription.EndDate = &endDateTime
	}

	insertedId, err := s.subscriptionRepository.Create(subscription)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("Subscription with id = %v added srccessfully!", insertedId)
	return res, err
}

func (s *SubscriptionService) GetById(subscriptionRequest model.GetSubscriptionByIdRequest) (model.Subscription, error) {
	res, err := s.subscriptionRepository.GetById(subscriptionRequest)
	return res, err
}

func (s *SubscriptionService) Update(subscription model.UpdateSubscriptionRequest) (int64, error) {
	if subscription.StartDate != "" {
		startDateTime, err := helper.ConvertTimeStringToBeginningOfMounth(subscription.StartDate)
		if err != nil {
			return 0, err
		}
		subscription.StartDate = startDateTime
	}

	if subscription.EndDate != nil {
		endDateTime, err := helper.ConvertTimeStringToTheEndOfMounth(*subscription.EndDate)
		if err != nil {
			return 0, err
		}
		subscription.EndDate = &endDateTime
	}

	res, err := s.subscriptionRepository.Update(subscription)
	if res != 1 {
		return 0, err
	}
	return res, err
}

func (s *SubscriptionService) Delete(subscriptionRequest model.DeleteSubscriptionByIdRequest) (string, error) {
	rowAffected, err := s.subscriptionRepository.Delete(subscriptionRequest)
	res := fmt.Sprintf("Rows deleted %v successfully.", rowAffected)
	return res, err
}

func (s *SubscriptionService) GetAll() ([]model.Subscription, error) {
	res, err := s.subscriptionRepository.GetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *SubscriptionService) GetTotalSummByFilter(summRequest model.GetTotalSummByFilterRequest) (model.GetTotalSummByFilterResponse, error) {
	var totalSumm model.GetTotalSummByFilterResponse
	if summRequest.FromDate != "" {
		fromDate, err := helper.ConvertTimeStringToBeginningOfMounth(summRequest.FromDate)
		if err != nil {
			return totalSumm, err
		}
		summRequest.FromDate = fromDate
	}

	if summRequest.ToDate != "" {
		toDate, err := helper.ConvertTimeStringToTheEndOfMounth(summRequest.ToDate)
		if err != nil {
			return totalSumm, err
		}
		summRequest.ToDate = toDate
	}

	res, err := s.subscriptionRepository.GetTotalSummByFilter(summRequest)
	if err != nil {
		return totalSumm, err
	}

	totalSumm.TotalSumm = res
	return totalSumm, nil
}
