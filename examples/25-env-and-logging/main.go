package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// Configuration holds application settings from environment variables
type Configuration struct {
	LogLevel      string
	LogFormat     string
	DatabaseURL   string
	APIKey        string
	EnableFeature bool
	Port          string
	Environment   string
}

// loadConfig loads configuration from environment variables
func loadConfig() (*Configuration, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	// Get environment variables with defaults
	config := &Configuration{
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		LogFormat:     getEnv("LOG_FORMAT", "json"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://localhost:5432/mydb"),
		APIKey:        getEnv("API_KEY", "default-key"),
		EnableFeature: getEnvBool("ENABLE_FEATURE", false),
		Port:          getEnv("PORT", "8080"),
		Environment:   getEnv("ENV", "development"),
	}

	// Validate configuration
	if config.APIKey == "default-key" && config.Environment == "production" {
		return nil, fmt.Errorf("API_KEY must be set in production environment")
	}

	return config, nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvBool gets a boolean environment variable with a default value
func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value == "true" || value == "1" || value == "yes"
}

// setupStandardLogging demonstrates the standard library logger
func setupStandardLogging(config *Configuration) *log.Logger {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal(err)
	}

	// Open log file
	logFile, err := os.OpenFile(
		filepath.Join("logs", "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create multi-writer for both file and stdout
	multiWriter := zerolog.MultiLevelWriter(logFile, os.Stdout)

	// Configure standard logger
	return log.New(multiWriter,
		"[STANDARD] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

// setupZerologging demonstrates structured logging with zerolog
func setupZerologging(config *Configuration) zerolog.Logger {
	// Set global log level
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Configure output format
	var output zerolog.ConsoleWriter
	if config.LogFormat == "pretty" {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		}
	} else {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    true,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %-50s|", i)
			},
		}
	}

	// Create logger
	return zerolog.New(output).With().
		Timestamp().
		Str("environment", config.Environment).
		Logger()
}

// demoLogging demonstrates different logging scenarios
func demoLogging(stdLog *log.Logger, logger zerolog.Logger) {
	// Standard library logging
	fmt.Println("\n=== Standard Library Logging ===")
	stdLog.Println("This is a standard log message")
	stdLog.Printf("Hello, %s!", "World")

	// Structured logging with zerolog
	fmt.Println("\n=== Structured Logging (zerolog) ===")

	// Different log levels
	logger.Trace().Msg("Trace level message")
	logger.Debug().Msg("Debug level message")
	logger.Info().Msg("Info level message")
	logger.Warn().Msg("Warning level message")
	logger.Error().Msg("Error level message")

	// Logging with context
	logger.Info().
		Str("user", "john").
		Int("age", 30).
		Bool("verified", true).
		Msg("User logged in")

	// Logging with error
	err := fmt.Errorf("database connection failed")
	logger.Error().
		Err(err).
		Str("database", "postgres").
		Int("retry_attempt", 3).
		Msg("Failed to connect to database")

	// Logging with timing
	logger.Info().
		Dur("duration", time.Second*2).
		Msg("Operation completed")

	// Sub-logger with additional context
	subLogger := logger.With().
		Str("component", "auth").
		Str("request_id", "123").
		Logger()

	subLogger.Info().Msg("Processing authentication")
}

func main() {
	fmt.Println("=== Environment Variables and Logging Example ===")

	// Load configuration from environment variables
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Print current configuration
	fmt.Println("\n=== Current Configuration ===")
	fmt.Printf("Log Level: %s\n", config.LogLevel)
	fmt.Printf("Log Format: %s\n", config.LogFormat)
	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("API Key: %s\n", maskSecret(config.APIKey))
	fmt.Printf("Enable Feature: %v\n", config.EnableFeature)
	fmt.Printf("Port: %s\n", config.Port)
	fmt.Printf("Environment: %s\n", config.Environment)

	// Setup loggers
	stdLog := setupStandardLogging(config)
	logger := setupZerologging(config)

	// Demonstrate logging
	demoLogging(stdLog, logger)
}

// maskSecret masks a secret string for display
func maskSecret(secret string) string {
	if len(secret) <= 4 {
		return "****"
	}
	return secret[:2] + "****" + secret[len(secret)-2:]
}
