package data

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// withTransaction creates and manages a database transaction.
// If txFunc completes successfully, the transaction is committed.
// If txFunc returns an error or a panic occurs, the transaction is rolled back.
//
// Parameters:
// - ctx: The context for the transaction.
// - txFunc: A function that takes a transaction (pgx.Tx) as an argument.
//
// Returns:
// - An error if any issues occurred during the transaction or within txFunc.
func (d *ServiceData) withTransaction(ctx context.Context, txFunc func(pgx.Tx) error) error {
	tx, err := d.dbPool.Begin(ctx)
	if err != nil {
		logrus.WithError(err).Error("Error starting transaction")
		return err
	}
	//управление транзакцией при помощи замыкания
	defer func() {
		if p := recover(); p != nil {
			if err = tx.Rollback(ctx); err != nil {
				logrus.WithError(err).Error("Error rolling back transaction.")
			}
			panic(p)
		} else if err != nil {
			if err = tx.Rollback(ctx); err != nil {
				logrus.WithError(err).Error("Error rolling back transaction.")
			}
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = txFunc(tx)
	return err
}
