package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/cindy-cyber/simpleBank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}

// after coding this up, we need to register this custom validator with Gin
// with `binding.Validator.Engine()` in server.go