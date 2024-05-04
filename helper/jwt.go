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
	GenerateJWT(userID uint, role string, email string) (string, error)
	GenerateToken(id uint, role string, email string) string
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

func (j *JWT) GenerateJWT(userID uint, role string, email string) (string, error) {
	var accessToken = j.GenerateToken(userID, role, email)
	if accessToken == "" {
		return "", fmt.Errorf("failed to generate access token")
	}

	return accessToken, nil
}

func (j *JWT) GenerateToken(id uint, role string, email string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["email"] = email
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
			result["id"] = uint(mapClaim["id"].(float64))
			result["role"] = mapClaim["role"]
			result["email"] = mapClaim["email"]
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
	idFloat, ok := mapClaim["id"].(uint)
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
		return c.JSON(http.StatusUnauthorized, FormatResponse(false, "Token is not valid", nil))
	}

	mapClaim := token.Claims.(jwt.MapClaims)
	id := mapClaim["id"]

	return id
}
