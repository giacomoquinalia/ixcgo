package main

import (
	"log/slog"
	"os"
	"time"
)

// START OMIT
func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Info("Dispositivo criado",
		"name", "Luz da cozinha",
		"room", "6° andar",
		"local", "Prédio IXC Soft",
		"now", time.Now())
}

// END OMIT
