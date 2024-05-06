package helper

import (
	"net/http"
	"rub_buddy/constant"

	"github.com/labstack/echo/v4"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constant.ErrBadRequest:
		return http.StatusBadRequest
	// User Error
	case constant.UserNotFound:
		return http.StatusNotFound
	case constant.ErrLoginEmptyInput:
		return http.StatusBadRequest
	case constant.ErrLoginNotFound:
		return http.StatusUnauthorized
	case constant.ErrLoginIncorrectPassword:
		return http.StatusUnauthorized
	case constant.ErrLoginJWT:
		return http.StatusInternalServerError
	case constant.ErrHashPassword:
		return http.StatusInternalServerError
	case constant.ErrRegisterUserExists:
		return http.StatusConflict
	case constant.ErrUpdateUserEmailExists:
		return http.StatusConflict
	case constant.ErrUpdateUser:
		return http.StatusInternalServerError

	// Collector Error
	case constant.ErrorCollectorRegister:
		return http.StatusInternalServerError
	case constant.ErrCollectorUserEmailExists:
		return http.StatusConflict
	case constant.ErrCollectorUserNotFound:
		return http.StatusNotFound
	case constant.ErrCollectorIncorrectPassword:
		return http.StatusUnauthorized
	case constant.ErrCollectorLoginJWT:
		return http.StatusInternalServerError
	case constant.ErrUpdateCollectorEmailExists:
		return http.StatusConflict
	case constant.ErrorUpdateCollector:
		return http.StatusInternalServerError

	default:
		return http.StatusInternalServerError
	}
}

func HandleEchoError(err error) (int, string) {
	switch e := err.(type) {
	case *echo.HTTPError:
		return ConvertResponseCode(e), e.Message.(string)
	default:
		return ConvertResponseCode(err), err.Error()
	}
}

func UnauthorizedError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, FormatResponse(false, constant.Unauthorized, []interface{}{}))
}
func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, FormatResponse(false, constant.InternalServerError, []interface{}{}))
}
