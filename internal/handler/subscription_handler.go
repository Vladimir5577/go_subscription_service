package handler

import (
	"fmt"
	"net/http"
	"subscriptions_service/internal/helper"
	"subscriptions_service/internal/model"
	"subscriptions_service/internal/service"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// type SubscriptionHandlerInterface interface {
// 	Create()
// 	GetById()
// 	Update()
// 	Delete()

// 	GetByFilter()
// 	GetAll()
// }

type SubscriptionHandler struct {
	subscriptionService service.SubscriptionServiceInterface
}

func NewSubscriptionHandler(subscriptionService service.SubscriptionServiceInterface) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

// Create godoc
// @Summary Create subscription
// @Description Create subscription
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param input body model.CreateSubscriptionRequest  true "create subscription"
// @Success 200 strig json
// @Router /subscription [post]
func (s *SubscriptionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subscription, err := helper.HandleBody[model.CreateSubscriptionRequest](&w, r)
		if err != nil {
			return
		}

		logrus.Infof("Create: %#v", subscription)

		res, err := s.subscriptionService.Create(*subscription)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

// GetById godoc
// @Summary Get subscription by id
// @Description Get subscriptions by id
// @Tags subscriptions
// @Produce  json
// @Param id path string true "Subscription Id"
// @Success 200 {strig} json
// @Router /subscription/{id} [get]
func (s *SubscriptionHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var subscriptionRequest model.GetSubscriptionByIdRequest
		subscriptionRequest.Id = chi.URLParam(r, "id")
		err := helper.IsValid(subscriptionRequest)
		if err != nil {
			errorResponse := helper.ParseValidationErrors(err)
			helper.JsonResponse(w, errorResponse, http.StatusBadRequest)
			return
		}

		logrus.Infof("Get by id: %#v", subscriptionRequest)

		res, err := s.subscriptionService.GetById(subscriptionRequest)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

// Create godoc
// @Summary Update subscription
// @Description Update subscription
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param input body model.UpdateSubscriptionRequest  true "update subscription"
// @Success 200 strig json
// @Router /subscription [put]
func (s *SubscriptionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subscription, err := helper.HandleBody[model.UpdateSubscriptionRequest](&w, r)
		if err != nil {
			return
		}

		logrus.Infof("Update: %#v", subscription)

		res, err := s.subscriptionService.Update(*subscription)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, fmt.Sprintf("Updated successfully %v row", res), http.StatusOK)
	}
}

// Delete godoc
// @Summary Delete subscription by id
// @Description Delete subscriptions by id
// @Tags subscriptions
// @Produce  json
// @Param id path string true "Subscription Id"
// @Success 200 {strig} json
// @Router /subscription/{id} [delete]
func (s *SubscriptionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var subscriptionRequest model.DeleteSubscriptionByIdRequest
		subscriptionRequest.Id = chi.URLParam(r, "id")
		err := helper.IsValid(subscriptionRequest)
		if err != nil {
			errorResponse := helper.ParseValidationErrors(err)
			helper.JsonResponse(w, errorResponse, http.StatusBadRequest)
			return
		}

		logrus.Infof("Deelte: %#v", subscriptionRequest)

		res, err := s.subscriptionService.Delete(subscriptionRequest)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

// GetTotalSummByFilter godoc
// @Summary Get summ subscriptions
// @Description Get summ subscriptions
// @Tags subscriptions
// @Produce  json
// @Param id path string true "Subscription Id"
// @Success 200 {strig} json
// @Router /subscriptions/total_summ [get]
func (s *SubscriptionHandler) GetTotalSummByFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var totalSummRequest model.GetTotalSummByFilterRequest

		totalSummRequest.UserID = r.URL.Query().Get("user_id")
		totalSummRequest.ServiceName = r.URL.Query().Get("service_name")
		totalSummRequest.FromDate = r.URL.Query().Get("from_date")
		totalSummRequest.ToDate = r.URL.Query().Get("to_date")

		logrus.Info("Get total summ: %#v", totalSummRequest)

		err := helper.IsValid(totalSummRequest)
		if err != nil {
			errorResponse := helper.ParseValidationErrors(err)
			helper.JsonResponse(w, errorResponse, http.StatusBadRequest)
			return
		}

		res, err := s.subscriptionService.GetTotalSummByFilter(totalSummRequest)
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}

// GetAll godoc
// @Summary Get add subscriptions
// @Description Get add subscriptions
// @Tags subscriptions
// // @Accept  json
// @Produce  json
// @Success 200 {array} model.Subscription
// @Router /subscriptions [get]
func (s *SubscriptionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.subscriptionService.GetAll()
		if err != nil {
			helper.JsonResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		helper.JsonResponse(w, res, http.StatusOK)
	}
}
