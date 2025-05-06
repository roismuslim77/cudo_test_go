package transaction

import (
	"context"
	"gorm.io/gorm"
	"log"
)

type repository struct {
	db *gorm.DB
}

func (r repository) TransactionFreqCheck(ctx context.Context) ([]FreqCheckResponse, error) {
	var data []FreqCheckResponse

	r.db.
		Table("transactions").
		Select("count(id), user_id, max(order_id) as order_id").
		Group("user_id, created_at, date_trunc('hour', created_at)").
		Order("created_at DESC").
		Find(&data)

	log.Println(data)
	return data, nil

}

func NewRepository(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}
