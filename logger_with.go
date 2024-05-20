package main

import (
	"log/slog"
	"math/rand"
	"os"
)

// START OMIT
func main() {
	logCorrelationID := rand.Int()
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler).With("logCorrelationID", logCorrelationID)

	logger.Info("Dispositivo criado")
	logger.Debug("Dispositivo integrado com Alexa")
	logger.Warn("Limite de dispositivos atingido")
	logger.Error("Falha ao acessar serviço externo")
}

// END OMIT
