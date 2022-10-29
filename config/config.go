package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4/middleware"
)

// Configuration :
type Configuration struct {
	Host            string
	Port            string
	CORSConfig      middleware.CORSConfig
	MongoHost       string
	MongoPort       string
	MongoUser       string
	MongoPassword   string
	DatabaseName    string
	CollectionName  string
	MongoReplicaSet string
	S3URL           string
	MaxConnections,
	MaxIdleConnections,
	ConnectionLifetime,
	DBTimeout int
	MongoConnectionTimeout int
	MongoMaxConnections    int64
}

var config Configuration

func init() {
	configureEnvironment()
	configureCors()
}

// GetConfig :
func GetConfig() Configuration {
	return config
}

func configureEnvironment() {
	if os.Getenv("GO_ENV") == "" {
		godotenv.Load()
	}
	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
	config.MongoHost = os.Getenv("MONGODB_HOSTS")
	config.MongoPort = os.Getenv("MONGODB_PORT")
	config.MongoUser = os.Getenv("MONGODB_USER")
	config.MongoPassword = os.Getenv("MONGODB_PASSWORD")
	config.DatabaseName = os.Getenv("MONGODB_DATABASE")
	config.CollectionName = os.Getenv("MONGODB_COLLECTION")
	config.MongoReplicaSet = os.Getenv("MONGODB_REPLICA_SET")
	config.S3URL = os.Getenv("S3_URL")

	config.MaxConnections, _ = strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
	config.MaxIdleConnections, _ = strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	config.DBTimeout, _ = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	config.ConnectionLifetime, _ = strconv.Atoi(os.Getenv("DB_CONNECTION_LIFETIME"))
	config.MongoConnectionTimeout, _ = strconv.Atoi(os.Getenv("MONGODB_CONNECTION_TIMEOUT"))
	config.MongoMaxConnections, _ = strconv.ParseInt(os.Getenv("MONGODB_MAX_CONNECTIONS"), 10, 64)

}

func configureCors() {

	corsAllowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	corsAllowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	corsAllowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	config.CORSConfig = middleware.CORSConfig{
		AllowHeaders: corsAllowedHeaders,
		AllowMethods: corsAllowedMethods,
		AllowOrigins: corsAllowedOrigins,
		MaxAge:       7200,
	}
}
