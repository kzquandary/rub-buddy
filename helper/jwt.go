package helper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type JWTInterface interface {
	GenerateJWT(userID uint) map[string]any
	GenerateToken(id uint) string
	ExtractToken(token *jwt.Token) map[string]interface{}
	ValidateToken(token string) (*jwt.Token, error)
	GetID(c echo.Context) (uint, error)
	CheckID(c echo.Context) interface{}
}

type JWT struct {
	signKey    string
}

func New(signKey string) JWTInterface {
	return &JWT{
		signKey:    signKey,
	}
}

func (j *JWT) GenerateJWT(userID uint) map[string]any {
	var result = map[string]any{}
	var accessToken = j.GenerateToken(userID)
	if accessToken == "" {
		return nil
	}
	result["access_token"] = accessToken

	return result
}

func (j *JWT) GenerateToken(id uint) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

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
			result["id"] = mapClaim["id"]
			result["role"] = mapClaim["role"]
			result["status"] = mapClaim["status"]
			return result
		}
		logrus.Error("Token Expired")
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
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return 0, err
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	idFloat, ok := mapClaim["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("ID not found or not a valid number")
	}

	idUint := uint(idFloat)
	return idUint, nil
}

func (j *JWT) CheckID(c echo.Context) any {
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.ValidateToken(authHeader)
	if err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusUnauthorized, FormatResponse("Token is not valid", nil))
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	id := mapClaim["id"]

	return id
}
