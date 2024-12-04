package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yordanos-habtamu/EcomGo.git/types"
)

type Store struct {
  db *sql.DB
}

func NewStore (db *sql.DB) *Store {
  return &Store{db:db}
}

func (s *Store) GetUserByEmail (email string) (*types.User,error){
   rows,err := s.db.Query("SELECT * FROM users WHERE email = ?",email)
   if err != nil{
    return nil,err
   }
   u := new(types.User)
   for rows.Next(){
    u,err = scanRowsIntoUsers(rows)
    if err!=nil{
      return nil,err
    }
   }
   if u.ID == 0 {
    return nil , fmt.Errorf("User Not Found")
   }
   return u,nil
}

func scanRowsIntoUsers(rows *sql.Rows) (*types.User,error){
  user := new(types.User)
  err := rows.Scan(
     &user.ID,
     &user.FirstName,
     &user.LastName,
     &user.Email,
     &user.Age,
     &user.DoB,
     &user.Password,
     &user.CreatedAt,
  )
  if err != nil{
    log.Fatal(err)
  }
    return user , nil
}

func (s *Store) GetUserById(id int) (*types.User,error){
  return nil,nil
}

func (s *Store) CreateUser(user types.User) error {
  return nil
}