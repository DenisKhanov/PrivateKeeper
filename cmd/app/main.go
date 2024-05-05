package main

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	privateKeeper, err := app.NewApp(ctx)
	if err != nil {
		logrus.Fatalf("failed to init app: %s", err.Error())
	}

	privateKeeper.Run()
}
