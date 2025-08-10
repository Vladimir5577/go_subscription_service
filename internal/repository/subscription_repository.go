package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"subscriptions_service/internal/model"
	"time"

	"github.com/Masterminds/squirrel"
)

type SubscriptionRepositoryInterface interface {
	Create(model.CreateSubscriptionRequest) (string, error)
	GetById(model.GetSubscriptionByIdRequest) (model.Subscription, error)
	Update(model.UpdateSubscriptionRequest) (int64, error)
	Delete(model.DeleteSubscriptionByIdRequest) (int64, error)
	GetAll() ([]model.Subscription, error)
	GetTotalSummByFilter(model.GetTotalSummByFilterRequest) (int, error)
}

type SubscriptionRepository struct {
	Db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		Db: db,
	}
}

func (s *SubscriptionRepository) Create(subscription model.CreateSubscriptionRequest) (string, error) {
	query, args, err := squirrel.Insert("subscription").
		PlaceholderFormat(squirrel.Dollar).
		Columns("service_name", "price", "user_id", "start_date", "end_date").
		Values(
			subscription.ServiceName,
			subscription.Price,
			subscription.UserID,
			subscription.StartDate,
			subscription.EndDate,
		).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", err
	}

	var lastInsertedId string
	err = s.Db.QueryRow(query, args...).Scan(&lastInsertedId)
	if err != nil {
		return "", err
	}

	return lastInsertedId, nil
}

func (s *SubscriptionRepository) GetById(subscriptionRequest model.GetSubscriptionByIdRequest) (model.Subscription, error) {
	id := subscriptionRequest.Id
	var subscription model.Subscription
	query, args, err := squirrel.Select(
		"id",
		"service_name",
		"price",
		"user_id",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From("subscription").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "id")), id).
		Limit(1).
		ToSql()
	if err != nil {
		return subscription, err
	}

	row := s.Db.QueryRow(query, args...)
	err = row.Scan(&subscription.Id, &subscription.ServiceName, &subscription.Price, &subscription.UserID, &subscription.StartDate, &subscription.EndDate, &subscription.CreatedAt, &subscription.UpdatedAt)

	if err != nil {
		return subscription, fmt.Errorf("subscription with id = %v does not found", id)
	}

	return subscription, nil
}

func (s *SubscriptionRepository) Update(subscription model.UpdateSubscriptionRequest) (int64, error) {
	var subscriptionRequest model.GetSubscriptionByIdRequest
	subscriptionRequest.Id = subscription.Id.String()
	_, err := s.GetById(subscriptionRequest)
	if err != nil {
		return 0, err
	}

	builder := squirrel.
		Update("subscription").
		PlaceholderFormat(squirrel.Dollar)
	if subscription.ServiceName != "" {
		builder = builder.Set("service_name", subscription.ServiceName)
	}
	if subscription.Price != 0 {
		builder = builder.Set("price", subscription.Price)
	}
	if subscription.UserID.String() != "" {
		builder = builder.Set("user_id", subscription.UserID)
	}
	if subscription.StartDate != "" {
		builder = builder.Set("start_date", subscription.StartDate)
	}
	if subscription.EndDate != nil {
		builder = builder.Set("end_date", subscription.EndDate)
	}
	builder = builder.Set("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	query, args, err := builder.
		Where((fmt.Sprintf("%s = ?", "id")), subscription.Id).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	res, err := s.Db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (s *SubscriptionRepository) Delete(subscriptionDeleteRequest model.DeleteSubscriptionByIdRequest) (int64, error) {
	var subscriptionRequest model.GetSubscriptionByIdRequest
	subscriptionRequest.Id = subscriptionDeleteRequest.Id
	_, err := s.GetById(subscriptionRequest)
	if err != nil {
		return 0, err
	}

	query, args, err := squirrel.
		Delete("subscription").
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", "id")), subscriptionRequest.Id).
		ToSql()
	if err != nil {
		return 0, err
	}

	res, err := s.Db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowAffected == 0 {
		return 0, fmt.Errorf("record not deleted.")
	}
	return rowAffected, nil
}

func (s *SubscriptionRepository) GetAll() ([]model.Subscription, error) {
	var subscription model.Subscription
	var subscriptions []model.Subscription

	queryBuilder := squirrel.Select(
		"id",
		"service_name",
		"price",
		"user_id",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From("subscription").
		PlaceholderFormat(squirrel.Dollar).
		OrderBy(fmt.Sprintf("%s %s", "created_at", "DESC"))

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return subscriptions, err
	}

	rows, err := s.Db.Query(query, args...)
	if err != nil {
		return subscriptions, err
	}

	for rows.Next() {
		err = rows.Scan(&subscription.Id, &subscription.ServiceName, &subscription.Price, &subscription.UserID, &subscription.StartDate, &subscription.EndDate, &subscription.CreatedAt, &subscription.UpdatedAt)
		if err != nil {
			return subscriptions, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (s *SubscriptionRepository) GetTotalSummByFilter(summRequest model.GetTotalSummByFilterRequest) (int, error) {

	builder := squirrel.
		Select(
			"SUM(price) AS total",
			// "id",
			// "service_name",
			// "price",
			// "user_id",
			// "start_date",
			// "end_date",
			// "created_at",
			// "updated_at",
		).
		From("subscription").
		PlaceholderFormat(squirrel.Dollar)
	if summRequest.UserID != "" {
		builder = builder.Where((fmt.Sprintf("%s = ?", "user_id")), summRequest.UserID)
	}
	if summRequest.ServiceName != "" {
		builder = builder.Where((fmt.Sprintf("%s = ?", "service_name")), summRequest.ServiceName)
	}
	if summRequest.FromDate != "" {
		builder = builder.Where((fmt.Sprintf("%s >= ?", "start_date")), summRequest.FromDate)
	}
	if summRequest.ToDate != "" {
		builder = builder.Where((fmt.Sprintf("%s <= ?", "end_date")), summRequest.ToDate)
	}

	query, args, err := builder.ToSql()

	if err != nil {
		return 0, err
	}

	var totalSumm int
	res := s.Db.QueryRow(query, args...)
	res.Scan(&totalSumm)

	return totalSumm, nil
}
