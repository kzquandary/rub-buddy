package seed

import (
	"database/sql"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type TPS struct {
	*gorm.Model
	ID        uint
	Name      string
	Address   string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *sql.NullTime
}

func SeedTPS(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		tps := TPS{
			ID:        uint(111111 * (i + 1)),
			Name:      "TPS " + strconv.Itoa(i+1),
			Address:   "Alamat TPS " + strconv.Itoa(i+1),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}

		// Menambahkan TPS ke database
		if err := db.Create(&tps).Error; err != nil {
			return err
		}
	}

	return nil
}
