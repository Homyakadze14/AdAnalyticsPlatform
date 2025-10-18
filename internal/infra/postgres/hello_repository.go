package postgres

import (
	"context"
	"errors"
	"fmt"

	"test/internal/entity"
	"test/pkg/postgres"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	*postgres.Postgres
}

func NewHelloRepository(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

func (r *Repo) Get(ctx context.Context) (hello entity.HelloResponse, err error) {
	const op = "repositories.postgresql.HelloRepository.Get"

	row := r.Pool.QueryRow(ctx, "SELECT text FROM hello WHERE id=1")

	err = row.Scan(&hello.Text)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return hello, nil
		}
		return hello, fmt.Errorf("%s: %w", op, err)
	}

	return hello, nil
}
