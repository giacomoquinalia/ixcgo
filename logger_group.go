package main

import (
	"log/slog"
	"os"
)

// START OMIT
func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Info("Dispositivo criado",
		slog.String("name", "Luz da cozinha"),
		slog.String("room", "6° andar"),
		slog.String("local", "Prédio IXC Soft"),
		slog.Group("metadata", "foo", "bar"),
	)
}

// END OMIT
