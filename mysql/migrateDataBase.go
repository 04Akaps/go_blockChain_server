package mysql

import (
	"database/sql"
	"log"

	sqlc "go_blockChain_server/mysql/sqlc"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func MigratMysql() *sqlc.Queries {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/launchPad?multiStatements=true&parseTime=true")
	if err != nil {
		log.Fatal("launchpad sql Open Error : ", err)
	}

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
