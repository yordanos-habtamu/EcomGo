package cart

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

  