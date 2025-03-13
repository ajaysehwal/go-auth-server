package services

import (
	"database/sql"

	"github.com/ajaysehwal/go-rest-server/internal/models"
)



type UserService struct{
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}
func (s *UserService) GetUserServices(userId string) (*models.User,error){
   var user models.User
   err := s.db.QueryRow(
	"SELECT id, email, created_at FROM users WHERE id = $1",
	userId,
 ).Scan(&user.ID, &user.Email, &user.CreatedAt)

 if err == sql.ErrNoRows {
	return nil, err
 } else if err != nil {
	return nil, err
}
return &user, nil
}


