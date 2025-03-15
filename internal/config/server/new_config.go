package serverconfig

import (
	"flag"
	"github.com/sirupsen/logrus"
)

func NewConfigServer(log *logrus.Entry) *ConfigServer {
	return &ConfigServer{
		log:    log,
		Server: Server{},
	}
}

func NewConfigServerMock(log *logrus.Entry) *ConfigServer {
	return &ConfigServer{
		log: log,
		Server: Server{
			FlagAddressAndPort: "localhost:8901",
			FlagLogLevel:       "debug",
			FlagDatabaseDsn:    "postgresql://Senya:1q2w3e4r5t@localhost:5432/GopheKeeper?sslmode=disable",
			FlagAccessKey:      "f4pq3792h3dy4g82o63R84P265o3874wgfiy2p947gf7qo5hcnbvtbo8y2c9upnox3q9E3",
			FlagRefreshKey:     "x53416ucertiyvuybiunb5yp6no78iu65cr34exqyto839p28u320kjfnubviry3294bdf",
			FlagCryptoKey:      "",
			FlagConfig:         "",
		},
	}
}

func (c *ConfigServer) InitializeServerConfig() {
	//Парсит флаги
	flag.Parse()
}
