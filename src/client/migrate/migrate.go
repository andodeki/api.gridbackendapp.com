package migrate

import (
	"database/sql"
	"fmt"

	"github.com/andodeki/api.gridbackendapp.com/src/config/config"

	"github.com/golang-migrate/migrate/v4"

	// "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func MigrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "Connecting to Database")
	}

	migrationSourse := fmt.Sprintf("file://%ssrc/client/migrate/migrations", *config.DataDirectory)

	migrator, err := migrate.NewWithDatabaseInstance(migrationSourse, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "Creating Migrator")
	}
	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "Executing Migration")
	}
	version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "Getting Migration Version")
	}

	logrus.WithFields(logrus.Fields{
		"version": version,
		"dirty":   dirty,
	}).Debug("Database Migrated")

	return nil

}
