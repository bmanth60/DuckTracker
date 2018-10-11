package data

import (
	"database/sql"

	dterr "github.com/bmanth60/DuckTracker/errors"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	// blank import to support the use of file:// connector
	_ "github.com/golang-migrate/migrate/source/file"
	// blank import to integrate the postgres driver to sql api
	_ "github.com/lib/pq"
)

var (
	//gDatabase global variable for database singleton
	gDatabase *Database
)

//Database object to handle database interactions
type Database struct {
	Db       *sql.DB
	migrator *migrate.Migrate
	*Helper
}

//Connect singleton to handle the databse connection pool
func Connect() (*Database, error) {
	if gDatabase != nil {
		return gDatabase, nil
	}

	db, err := sql.Open("postgres", getConfig().DatabaseDSN)
	if err != nil {
		db.Close()
		return nil, errors.Wrap(err, "Couldn't open connection to postgres database")
	}

	// Although the database connection pool has been created, golang
	// does not actually attempt to connect to the database until we
	// ping
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

//initMigrator initialize the migration library
func (d *Database) initMigrator() error {
	if d.Db == nil {
		return dterr.ErrDbNotConnected
	}

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

//Migrate runs all up migration scripts that have not yet been executed
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

//Reset drops all tables in the database. Ideally, this should be abstracted
//into a consolidated test tool to prevent accidental data loss
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

//Rollback run all down migration scripts
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

//Close database and prevent new queries from starting
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
