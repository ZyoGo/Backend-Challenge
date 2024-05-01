package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name       string   `toml:"name"`
		Port       int      `toml:"port"`
		CorsOrigin []string `toml:"corsOrigin"`
		LogLevel   string   `toml:"logLevel"`
	} `toml:"app"`
	Database struct {
		Name                   string `toml:"name"`
		Host                   string `toml:"host"`
		Port                   int    `toml:"port"`
		User                   string `toml:"user"`
		Password               string `toml:"password"`
		TimeZone               string `toml:"timeZone"`
		MaxConns               int32  `toml:"maxConns"`
		MinConns               int32  `toml:"minConns"`
		MaxConnLife            int    `toml:"maxConnLife"`
		MaxConnLifeUnit        string `toml:"maxConnLifeUnit"`
		MaxConnIdle            int    `toml:"maxConnIdle"`
		MaxConnIdleUnit        string `toml:"maxConnIdleUnit"`
		HealthCheckPeriod      int    `toml:"healthCheckPeriod"`
		HealthCheckPeriodUnit  string `toml:"healthCheckPeriodUnit"`
		ConnConnectTimeout     int    `toml:"connConnectTimeout"`
		ConnConnectTimeoutUnit string `toml:"connConnectTimeoutUnit"`
	} `toml:"database"`
	// JWTPublicKey  string   `toml:"jwt_public_key"`
	// JWTPrivateKey string   `toml:"jwt_private_key"`
}

var (
	lock      = &sync.Mutex{}
	appConfig *AppConfig
)

func LoadConfig(path string) *AppConfig {
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")

	env := os.Getenv("APP_ENV")
	if env == "PRODUCTION" || env == "DEVELOPMENT" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("test")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config not found, use env variable, current path = ", path)
		}
	}

	var finalConfig AppConfig
	if err := viper.Unmarshal(&finalConfig); err != nil {
		panic(err)
	}
	return &finalConfig
}

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)

		appConfig = LoadConfig(basePath)
	}

	return appConfig
}

// func GetJWTPublicKey() *rsa.PublicKey {
// 	jwtPublicKey := []byte(fmt.Sprintf(`
// -----BEGIN PUBLIC KEY-----
// %s
// -----END PUBLIC KEY-----`, GetConfig().JWTPublicKey))

// 	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPublicKey)
// 	if err != nil {
// 		log.Fatal("Cannot parse public key: ", publicKey)
// 	}

// 	return publicKey
// }

// func GetJWTPrivateKey() *rsa.PrivateKey {
// 	jwtPrivateKey := []byte(fmt.Sprintf(`
// ----BEGIN PRIVATE KEY----
// %s
// ----END PRIVATE KEY----`, GetConfig().JWTPrivateKey))

// 	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtPrivateKey)
// 	if err != nil {
// 		log.Fatal("Cannot parse private key: ", privateKey)
// 	}

// 	return privateKey
// }
