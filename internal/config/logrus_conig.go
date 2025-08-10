package config

import (
	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Log messages
	// logrus.Info("Application started successfully.")
	// logrus.Warn("Configuration file not found, using default settings.")
	// logrus.Error("Failed to connect to the database.")
}
