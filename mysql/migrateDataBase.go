package mysql

import (
	"database/sql"
	"log"

	sqlc "go_blockChain_server/mysql/sqlc"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func MigratMysql(db *sql.DB) *sqlc.Queries {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("launchpad instance Error : ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://mysql/migrate",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("db instance not found : ", err)
	}

	m.Steps(2)
	return sqlc.New(db)
}
