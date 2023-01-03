package sqldb

import (
	"context"
	"database/sql"
	"fmt"
)

type Config struct {
	Name   string
	Driver string
	DSN    string
	Select string
}

func NewSqlDbCheck(cfg *Config) func(ctx context.Context) error {
	return func(ctx context.Context) (chkErr error) {
		db, err := sql.Open(cfg.Driver, cfg.DSN)
		if err != nil {
			chkErr = fmt.Errorf("%s health check failed on connect: %w", cfg.Name, err)
			return
		}

		defer func() {
			err = db.Close()
			if err != nil && chkErr == nil {
				chkErr = fmt.Errorf("%s health check failed on close: %w", cfg.Name, err)
				return
			}
		}()

		err = db.PingContext(ctx)
		if err != nil {
			chkErr = fmt.Errorf("%s health check failed on ping: %w", cfg.Name, err)
			return
		}

		if cfg.Select != "" {
			rows, err := db.QueryContext(ctx, cfg.Select)
			if err != nil {
				chkErr = fmt.Errorf("%s health check failed on select: %w", cfg.Name, err)
				return
			}

			defer func() {
				err = rows.Close()
				if err != nil {
					chkErr = fmt.Errorf("sql server health check failed on rows closing: %w", err)
				}
			}()
		}

		return
	}
}
