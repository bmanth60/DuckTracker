package data

import (
	"database/sql"

	dterr "github.com/bmanth60/DuckTracker/errors"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

var gDatabase *Database

type Database struct {
	Db       *sql.DB
	migrator *migrate.Migrate
	*Helper
}

func Connect() (*Database, error) {
	if gDatabase != nil {
		return gDatabase, nil
	}

	db, err := sql.Open("postgres", getConfig().DatabaseDSN)
	if err != nil {
		db.Close()
		return nil, errors.Wrap(err, "Couldn't open connection to postgres database")
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, errors.Wrap(err, "Couldn't ping postgres database")
	}

	gDatabase = &Database{
		Db: db,
		Helper: &Helper{
			db: db,
		},
	}

	return gDatabase, nil
}

func (d *Database) initMigrator() error {
	if d.Db == nil {
		return dterr.ErrDbNotConnected
	}

	// Migrator has already been instantiated
	if d.migrator != nil {
		return nil
	}

	driver, err := postgres.WithInstance(d.Db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://data/migrations/",
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	d.migrator = m

	return nil
}

func (d *Database) Migrate() error {
	err := d.initMigrator()
	if err != nil {
		return err
	}

	err = d.migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d *Database) Reset() error {
	err := d.initMigrator()
	if err != nil {
		return err
	}

	err = d.migrator.Drop()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) Rollback() error {
	err := d.initMigrator()
	if err != nil {
		return err
	}

	err = d.migrator.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d *Database) Close() error {
	if d.Db == nil {
		return dterr.ErrDbNotConnected
	}

	// TODO - Brian - 20181009 - What happens if
	// db is closed and not nil?
	if err := d.Db.Close(); err != nil {
		return errors.Wrap(err, "Errored closing database connection")
	}

	return nil
}
