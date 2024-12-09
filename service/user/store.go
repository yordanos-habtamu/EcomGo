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
    return nil , fmt.Errorf("user not found")
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
     &user.Password,
     &user.CreatedAt,
     &user.DoB,
     &user.Sex,
  )
  if err != nil{
    log.Fatal(err)
  }
    return user , nil
}

func (s *Store) GetUserById(id int) (*types.User,error){
  rows,err := s.db.Query("SELECT * FROM users WHERE id = ?",id)
  if err != nil {
    return nil,err
  }
  u := new(types.User)
  for rows.Next(){
      u,err = scanRowsIntoUsers(rows)
      if err != nil{
         return nil,err
      }
  }
  if u.ID == 0{
    return nil, fmt.Errorf("user not found")
  }
  return u,nil
}

func (s *Store) CreateUser(user types.User) error {
  _,err := s.db.Exec("INSERT INTO users (firstName,lastName,email,password,DoB,sex) VALUES (?,?,?,?,?,?)",user.FirstName,user.LastName,user.Email,user.Password,user.DoB,user.Sex)
  if err != nil {
    return err
  }
  return nil
}