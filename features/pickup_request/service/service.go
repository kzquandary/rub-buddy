package service

import (
	"errors"
	"rub_buddy/constant"
	pickuprequest "rub_buddy/features/pickup_request"
)

type PickupRequestService struct {
	d pickuprequest.PickupRequestDataInterface
}

func New(d pickuprequest.PickupRequestDataInterface) pickuprequest.PickupRequestServiceInterface {
	return &PickupRequestService{
		d: d,
	}
}

func (s *PickupRequestService) CreatePickupRequest(newData pickuprequest.PickupRequest) (*pickuprequest.PickupRequest, error) {
	return s.d.CreatePickupRequest(&newData)
}

func (s *PickupRequestService) GetAllPickupRequest() ([]pickuprequest.PickupRequest, error) {
	return s.d.GetAllPickupRequest()
}

func (s *PickupRequestService) GetPickupRequestByID(id int) (pickuprequest.PickupRequest, error) {
	return s.d.GetPickupRequestByID(id)
}

func (s *PickupRequestService) DeletePickupRequestByID(id int, userID uint) error {
	PickupRequestAvailable, _ := s.d.GetPickupRequestByID(id)
	if PickupRequestAvailable.ID == 0 {
		return errors.New(constant.NotFound)
	}
	return s.d.DeletePickupRequestByID(id, userID)
}
