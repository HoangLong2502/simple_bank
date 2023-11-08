package api

import (
	"github.com/HoangLong2502/simple_bank/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// Kiểm tra xem giá trị truyền vào có thể chuyển đổi thành string ko, nếu được thì ok=true, ngược lại
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
