package data

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// withTransaction создание и управление транзакцией
func (d *ServiceData) withTransaction(ctx context.Context, txFunc func(pgx.Tx) error) error {
	tx, err := d.dbPool.Begin(ctx)
	if err != nil {
		logrus.Error("Error starting transaction: ", err)
		return err
	}
	//управление транзакцией при помощи замыкания
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = txFunc(tx)
	return err
}
