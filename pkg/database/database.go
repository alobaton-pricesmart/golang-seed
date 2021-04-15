package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database represents a reusable connection to a remote MySQL database.
type Database struct {
	db    *gorm.DB
	debug bool
}

func (d *Database) Collection(model Model) *Collection {
	return newCollection(d, model)
}

func (d *Database) Migrate(model Model) {
	d.db.AutoMigrate(model)
}

func (d *Database) SetupJoinTable(model Model, fieldName string, joinTable Model) {
	d.db.SetupJoinTable(model, fieldName, joinTable)
}

func Open(credentials Credentials, debug bool) (*Database, error) {
	database := new(Database)
	database.debug = debug

	if database.debug {
		log.WithField("credentials", credentials.String()).Debug("Open database connection")
	}

	dsn := credentials.String()
	var err error
	database.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db : %v", err)
	}

	db, err := database.db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting inner db session : %v", err)
	}

	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(0)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping : %v", err)
	}

	return database, nil
}
