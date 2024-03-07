package bootstrap

import (
	"github.com/go_fundaments/internal/domain"
	"github.com/go_fundaments/internal/user"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Adrian",
			LastName:  "Pucheta",
			Email:     "jorge@hotmail.com",
		}, {
			ID:        2,
			FirstName: "Adri",
			LastName:  "Puche",
			Email:     "jor@hotmail.com",
		}, {
			ID:        3,
			FirstName: "Jorge",
			LastName:  "Diaz",
			Email:     "jorgeDiaz@hotmail.com",
		}},
		MaxUserID: 3,
	}
}
