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
		"name", "Luz da cozinha",
		"room", "6° andar",
		"local", "Prédio IXC Soft")
}

// END OMIT
