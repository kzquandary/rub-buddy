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

func (s *PickupTransactionService) CreatePickupTransaction(newData pickuptransaction.PickupTransaction) (*pickuptransaction.PickupTransactionCreate, error) {
	return s.p.CreatePickupTransaction(newData)
}

func (s *PickupTransactionService) GetAllPickupTransaction(userId uint) ([]pickuptransaction.PickupTransactionInfo, error) {
	return s.p.GetAllPickupTransaction(userId)
}

func (s *PickupTransactionService) GetPickupTransactionByID(id int) (pickuptransaction.PickupTransactionInfo, error) {
	return s.p.GetPickupTransactionByID(id)
}
