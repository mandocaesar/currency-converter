package converter

import (
	"time"

	common "github.com/currency-converter/common"
	"github.com/satori/go.uuid"
)

//DailyRate : struct for Daily rate exchange data
type DailyRate struct {
	common.BaseModel
	ExchangeID   uuid.UUID `gorm:"type:char(36); not null"`
	ExchangeDate *time.Time
	Rate         float64 `gorm:"decimal"`
}
