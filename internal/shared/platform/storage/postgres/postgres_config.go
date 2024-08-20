package postgres

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage"
)

func CreateConfig() (*storage.Dbconfig, error) {
	var cfg storage.Dbconfig
	err := envconfig.Process("DB", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
