package postgres

import (
	"context"
	"github.com/VanOODev/education_web_server/adapters/gates/postgres/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Client struct {
	db *pgxpool.Pool
}

func NewClient(config config.Config) *Client {
	return &Client{}
}

func (c *Client) Close() {
}

func (c *Client) Add(ctx context.Context, data int64) (index int64, err error) {
	defer func() {
		err = errors.Wrap(err, "postgres (c *Client) Add()")
	}()

	row := c.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "mytable")
	if err != nil {
		return 0, errors.Wrap(err, "c.db.QueryRow()")
	}

	// TODO обработать данные из бд
	_ = row
	return
}

func (c *Client) Get(ctx context.Context, index int64) (data int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Delete(ctx context.Context, index int64, err error) {
	//TODO implement me
	panic("implement me")
}
