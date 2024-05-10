package data

import (
	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
	"rub_buddy/features/midtranspayment"
	"time"

	"gorm.io/gorm"
)

type MidtransData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) midtranspayment.MidtransDataInterface {
	return &MidtransData{
		DB: db,
	}
}

func (data *MidtransData) CreatePayment(newPayment midtranspayment.Midtrans) error {
	newPayment.CreatedAt = time.Now()
	newPayment.UpdatedAt = time.Now()
	err := data.DB.Create(&newPayment).Error
	if err != nil {
		return constant.ErrPaymentTransactionCreate
	}
	return nil
}

func (data *MidtransData) VerifyPayment(orderId string) error {
	var payment midtranspayment.Midtrans
	if err := data.DB.Where("id = ?", orderId).First(&payment).Error; err != nil {
		return constant.ErrGetPaymentTransaction
	}

	if payment.Status == 0 {
		if err := data.DB.Model(&payment).Update("status", 1).Error; err != nil {
			return constant.ErrPaymentTransactionUpdate
		}
		if err := data.DB.Table(tablename.UserTableName).Where("id = ?", payment.UserID).Update("balance", gorm.Expr("balance + ?", payment.Amount)).Error; err != nil {
			return constant.ErrPaymentTransactionUpdate
		}
		return nil
	} else if payment.Status == 1 {
		return constant.ErrAlreadyVerified
	}
	return constant.ErrAlreadyVerified
}
