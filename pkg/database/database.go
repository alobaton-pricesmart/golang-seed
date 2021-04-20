package database

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database represents a reusable connection to a remote MySQL database.
type (
	Database struct {
		db    *gorm.DB
		debug bool
	}

	Conf struct {
		Debug           bool
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
	}
)

// Collection starts a query statement for a given Model
func (d *Database) Collection(model Model) *Collection {
	return newCollection(d, model)
}

// Migrate migrates a model to database
func (d *Database) Migrate(model Model) {
	d.db.AutoMigrate(model)
}

// SetupJoinTable migrates a JoinTable to database.
func (d *Database) SetupJoinTable(model Model, fieldName string, joinTable Model) {
	d.db.SetupJoinTable(model, fieldName, joinTable)
}

// Open open a database connection
func Open(credentials Credentials, conf Conf) (*Database, error) {
	database := new(Database)
	database.debug = conf.Debug

	if database.debug {
		log.WithField("credentials", credentials.String()).Debug("Open database connection")
	}

	l := logger.New(
		log.WithFields(log.Fields{}),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	dsn := credentials.String()
	var err error
	database.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		DisableAutomaticPing:   true,
		Logger:                 l,
	})
	if err != nil {
		return nil, fmt.Errorf("error opening db : %v", err)
	}

	if database.debug {
		database.db = database.db.Debug()
	}

	db, err := database.db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting inner db session : %v", err)
	}

	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime))

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping : %v", err)
	}

	return database, nil
}
