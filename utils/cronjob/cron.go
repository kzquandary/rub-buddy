package cronjob

import (
	"rub_buddy/constant/tablename"
	PickupRequestData "rub_buddy/features/pickup_request/data"

	"gorm.io/gorm"
)

type CronInterface interface {
	HandleDeletePickupRequest()
}

type Cron struct {
	DB *gorm.DB
}

func New(db *gorm.DB) CronInterface {
	return &Cron{
		DB: db,
	}
}

func (c *Cron) HandleDeletePickupRequest() {
	c.DB.Where("status = ?", tablename.TablePickupRequestStatusPending).Delete(&PickupRequestData.PickupRequest{})
}
