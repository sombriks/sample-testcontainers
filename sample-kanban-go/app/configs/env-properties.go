package configs

import (
	"errors"
	"os"
)

type DbProps struct {
	Database   string
	Username   string
	Password   string
	Hostname   string
	SslMode    string
	InitScript string
}

func NewDbProps() (*DbProps, error) {
	var props DbProps
	var ok bool

	// collect configuration from environment
	props.Database, ok = os.LookupEnv("PG_DATABASE")
	if !ok {
		return nil, errors.New("PG_DATABASE environment variable not set")
	}
	props.Username, ok = os.LookupEnv("PG_USERNAME")
	if !ok {
		return nil, errors.New("PG_USERNAME environment variable not set")
	}
	props.Password, ok = os.LookupEnv("PG_PASSWORD")
	if !ok {
		return nil, errors.New("PG_PASSWORD environment variable not set")
	}
	props.Hostname, ok = os.LookupEnv("PG_HOSTNAME")
	if !ok {
		return nil, errors.New("PG_HOSTNAME environment variable not set")
	}
	props.SslMode, ok = os.LookupEnv("PG_SSL_MODE")
	if !ok {
		return nil, errors.New("PG_SSL_MODE environment variable not set")
	}
	props.InitScript, ok = os.LookupEnv("PG_INIT_SCRIPT")
	if !ok {
		return nil, errors.New("PG_INIT_SCRIPT environment variable not set")
	}

	return &props, nil
}
