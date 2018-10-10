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

type Database struct {
	Db *sql.DB
	*Helper
}

func Connect() (*Database, error) {
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

	return &Database{
		Db: db,
		Helper: &Helper{
			db: db,
		},
	}, nil
}

func (d *Database) Migrate() error {
	if d.Db == nil {
		return dterr.ErrDbNotConnected
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

	err = m.Up()
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
