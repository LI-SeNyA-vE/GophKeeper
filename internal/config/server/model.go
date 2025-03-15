package serverconfig

import "github.com/sirupsen/logrus"

type ConfigServer struct {
	log *logrus.Entry
	Server
}

type Server struct {
	FlagAddressAndPort string `env:"ADDRESS" json:"address"`
	FlagLogLevel       string `env:"LOG_LEVEL"`
	FlagDatabaseDsn    string `env:"DATABASE_DSN" json:"database_dsn"`
	FlagAccessKey      string `env:"ACCESS_KEY"`
	FlagRefreshKey     string `env:"REFRESH_KEY"`
	FlagCryptoKey      string `env:"CRYPTO_KEY" json:"crypto_key"`
	FlagConfig         string `env:"CONFIG"`
}
