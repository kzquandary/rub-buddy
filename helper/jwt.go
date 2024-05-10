package helper

import (
	"fmt"
	"net/http"
	"rub_buddy/constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type JWTInterface interface {
	GenerateJWT(userID uint, role string, email string, address string) (string, error)
	GenerateToken(id uint, role string, email string, address string) string
	ExtractToken(token *jwt.Token) map[string]interface{}
	ValidateToken(token string) (*jwt.Token, error)
	GetID(c echo.Context) (uint, error)
	CheckID(c echo.Context) interface{}
}

type JWT struct {
	signKey string
}

func New(signKey string) JWTInterface {
	return &JWT{
		signKey: signKey,
	}
}

func (j *JWT) GenerateJWT(userID uint, role string, email string, address string) (string, error) {
	var accessToken = j.GenerateToken(userID, role, email, address)
	if accessToken == "" {
		return "", constant.ErrLoginJWT
	}

	return accessToken, nil
}

func (j *JWT) GenerateToken(id uint, role string, email string, address string) string {
	var claims = jwt.MapClaims{}
	claims[constant.JWT_ID] = id
	claims[constant.JWT_ROLE] = role
	claims[constant.JWT_EMAIL] = email
	claims[constant.JWT_ADDRESS] = address
	claims[constant.JWT_IAT] = time.Now().Unix()
	claims[constant.JWT_EXP] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) ExtractToken(token *jwt.Token) map[string]interface{} {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			var result = map[string]interface{}{}
			result[constant.JWT_ID] = uint(mapClaim[constant.JWT_ID].(float64))
			result[constant.JWT_ROLE] = mapClaim[constant.JWT_ROLE]
			result[constant.JWT_ADDRESS] = mapClaim[constant.JWT_ADDRESS]
			result[constant.JWT_EMAIL] = mapClaim[constant.JWT_EMAIL]
			return result
		}
		return nil
	}
	return nil
}

func (j *JWT) ValidateToken(token string) (*jwt.Token, error) {
	var authHeader = token[7:]
	parsedToken, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.signKey), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func (j *JWT) GetID(c echo.Context) (uint, error) {
	authHeader := c.Request().Header.Get(constant.HeaderAuthorization)

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return 0, err
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	idFloat, ok := mapClaim["id"].(uint)
	if !ok {
		return 0, fmt.Errorf("ID not found or not a valid number")
	}

	idUint := uint(idFloat)
	return idUint, nil
}

func (j *JWT) CheckID(c echo.Context) any {
	authHeader := c.Request().Header.Get(constant.HeaderAuthorization)

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusUnauthorized, FormatResponse(false, "Token is not valid", nil))
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	id := mapClaim["id"]

	return id
}
