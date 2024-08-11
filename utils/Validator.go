package utils

import (
	"Glyphrz-go/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HandleValidatorError(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		response.Fail(c, http.StatusInternalServerError, 500, "字段校验错误", err.Error())
	}
	response.Fail(c, http.StatusBadRequest, 400, "字段校验错误", errs.Error())
	return
}
