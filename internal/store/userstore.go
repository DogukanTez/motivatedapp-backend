package store

import (
	// "database/sql"
	// "errors"
	// "context"
	"motivatedapp-backend/internal/database"
	"motivatedapp-backend/internal/domain"
	// "github.com/matthewhartstonge/argon2"
	// "github.com/stephenafamo/bob/dialect/psql"
	// "github.com/stephenafamo/bob/dialect/psql/sm"
)

// UserStore defines the interface for user data operations
type UserStore interface {
	// GetUser(id int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
}

type userDBStore struct {
	db database.Service
}

// NewUserStore creates a new UserStore instance
func NewUserStore(db database.Service) UserStore {
	return &userDBStore{db: db}
}

func (s *userDBStore) GetAllUsers() ([]domain.User, error) {
	/*var users []domain.User

		ctx := context.Background()
		q := psql.Select(sm.Columns("id", "email", "password", "auth_provider", "created_at", "updated_at"),sm.From("users"))

		/*if !user.IsAdmin {
	    q.Apply(
	        sm.Where(psql.Quote("user_id").EQ(psql.Arg(user.ID))),
	    ) // SELECT * FROM projects WHERE "user_id" = $1
	}*/
	return nil, nil

}
