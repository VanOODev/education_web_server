package storage

import "context"

type Storage interface {
	Add(ctx context.Context, data int64) (index int64, err error)
	Get(ctx context.Context, index int64) (data int64, err error)
	Delete(ctx context.Context, index int64, err error)
}
