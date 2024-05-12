package service

import (
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
	if newData.Weight == 0 || newData.Description == "" || newData.Image == "" {
		return nil, constant.ErrPickupRequestEmptyInput
	}
	return s.d.CreatePickupRequest(&newData)
}

func (s *PickupRequestService) GetAllPickupRequest(role string, id uint) ([]pickuprequest.PickupRequestInfo, error) {
	return s.d.GetAllPickupRequest(role, id)
}

func (s *PickupRequestService) GetPickupRequestByID(id int) (pickuprequest.PickupRequestInfo, error) {
	return s.d.GetPickupRequestByID(id)
}

func (s *PickupRequestService) DeletePickupRequestByID(id int, userID uint) error {
	PickupRequestAvailable, _ := s.d.GetPickupRequestByID(id)
	if PickupRequestAvailable.ID == 0 {
		return constant.ErrPickupRequestNotFound
	}
	err := s.d.DeletePickupRequestByID(id, userID)
	if err != nil {
		return err
	}
	return nil
}
