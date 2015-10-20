package database

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"github.com/marcopennekamp/gosrs/core"
)

func Open(conf *core.Config) (*sqlx.DB, error) {
	type DBConfig struct {
		Development struct {
			Driver string
			Open   string
		}
	}

	configText, err := ioutil.ReadFile(conf.GetFilePath(filepath.Join("database", "dbconf.yml")))
	if err != nil {
		return nil, err
	}

	dbConfig := DBConfig{}
	err = yaml.Unmarshal(configText, &dbConfig)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect(dbConfig.Development.Driver, dbConfig.Development.Open)
	if err != nil {
		return nil, err
	}

	return db, nil
}



