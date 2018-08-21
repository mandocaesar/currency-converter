package converter

import (
	common "github.com/currency-converter/common"
)

//Exchange : database model for currency exchange data
type Exchange struct {
	common.BaseModel
	source string `gorm:"type:char(5)"`
	target string `gorm:"type:char(5)"`
}
