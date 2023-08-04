package service

import (
	"errors"
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *AuthService) Register(user entity.User) (int, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hashedPassword
	return s.repos.CreateUser(user)
}

func (s *AuthService) GetUserByID(id int) (entity.User, error) {
	return s.repos.GetUserByID(id)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repos.GetUser(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		fmt.Println("Password is correct!")
	} else {
		fmt.Println("Password is incorrect.")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) VerifyToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
