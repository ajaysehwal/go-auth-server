package services

import (
	"database/sql"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	db *sql.DB
}

type Claims struct{
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) Register(email, password string)(int, error){
	hash, err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err !=nil{
		return 0 , err
	}

	var userID int
	if err =s.db.QueryRow("INSERT INTO users (email, passoword_hash) VALUES ($1, $2) RETURNING id",email,string(hash),).Scan(&userID); err !=nil {
		return 0 , err
	}

	return userID, nil
	
	
}

func (s *AuthService) Login(email string, password string) (string, error){
	var userId int 
	var passwordHash string
   if err:=s.db.QueryRow("SELECT id, password_hash FROM users WHERE email=$1",email).Scan(&userId,&passwordHash); err!=sql.ErrNoRows{
	return "",err;
   }else if err !=nil{
	return "",err;
   }

   if err:=bcrypt.CompareHashAndPassword([]byte(passwordHash),[]byte(password)); err !=nil{
	return "",err;
   }

   claims:=&Claims{
	UserID: userId,
	RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add((24 * time.Hour))),
	},
   }
   token:=jwt.NewWithClaims(jwt.SigningMethodES256,claims)
   return token.SignedString(middleware.JWTSecret)

}