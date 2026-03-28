package main

import (
	"context"
	"log/slog"
	"match/internal/service"
)

func main() {

	err := service.PerformMatch(context.Background())

	if err != nil {
		slog.Error("Error matching chargebacks and sales", slog.Any("Error", err))
	}

}
