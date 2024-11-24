package main

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/app/server"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	privateKeeper, err := server.NewApp(ctx)
	if err != nil {
		logrus.Fatalf("failed to init app: %s", err.Error())
	}

	privateKeeper.Run()
}
