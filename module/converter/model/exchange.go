package converter

import (
	common "github.com/currency-converter/common"
)

//Exchange : database model for currency exchange data
type Exchange struct {
	common.BaseModel
	Source string `gorm:"type:char(5)"`
	Target string `gorm:"type:char(5)"`
}
