package service

import (
	"rub_buddy/configs"
	"rub_buddy/features/midtranspayment"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService struct {
	d      midtranspayment.MidtransDataInterface
	config configs.MidtransConfig
}

func New(d midtranspayment.MidtransDataInterface, config configs.MidtransConfig) midtranspayment.MidtransServiceInterface {
	return &MidtransService{
		d:      d,
		config: config,
	}
}

func (s *MidtransService) GenerateSnapURL(payment midtranspayment.Midtrans) (midtranspayment.Midtrans, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.ID,
			GrossAmt: int64(payment.Amount),
		},
	}
	var client snap.Client
	client.New(s.config.ServerKey, midtrans.Sandbox)
	snapResponse, err := client.CreateTransaction(req)
	if err != nil {
		return midtranspayment.Midtrans{}, err
	}

	paymentData := midtranspayment.Midtrans{
		ID:      payment.ID,
		UserID:  payment.UserID,
		Amount:  payment.Amount,
		Status:  0,
		SnapURL: snapResponse.RedirectURL,
	}

	if err := s.d.CreatePayment(paymentData); err != nil {
		return midtranspayment.Midtrans{}, err
	}
	return paymentData, nil
}

func (s *MidtransService) VerifyPayment(orderId string) error {
	err := s.d.VerifyPayment(orderId)
	if err != nil {
		return err
	}
	return nil
}