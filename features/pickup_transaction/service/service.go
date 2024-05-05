package service

import (
	pickuptransaction "rub_buddy/features/pickup_transaction"
)

type PickupTransactionService struct {
	p pickuptransaction.PickupTransactionDataInterface
}

func New(p pickuptransaction.PickupTransactionDataInterface) pickuptransaction.PickupTransactionServiceInterface {
	return &PickupTransactionService{
		p: p,
	}
}

func (s *PickupTransactionService) CreatePickupTransaction(newData pickuptransaction.PickupTransaction) (*pickuptransaction.PickupTransaction, error) {
	return s.p.CreatePickupTransaction(newData)
}

func (s *PickupTransactionService) GetAllPickupTransaction(userId uint) ([]pickuptransaction.PickupTransaction, error) {
	query, err := s.p.GetAllPickupTransaction(userId)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *PickupTransactionService) GetPickupTransactionByID(id int) (pickuptransaction.PickupTransaction, error) {
	return s.p.GetPickupTransactionByID(id)
}