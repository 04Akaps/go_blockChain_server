package mysql

import (
	"database/sql"
	"fmt"
	"log"

	sqlc "go_blockChain_server/mysql/sqlc"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"

	_ "github.com/go-sql-driver/mysql"
)

func MigratMysql(db *sql.DB) *sqlc.Queries {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("launchpad instance Error : ", err)
	}
	// migrate --path db/migration -database "postgresql://root:hojin@localhost:5432/simple_bank?sslmode=disable" -verbose up

	m, err := migrate.NewWithDatabaseInstance(
		"file://mysql/migrate",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("db instance not found : ", err)
	}

	version, dirty, err := m.Version()

	fmt.Printf("%v\n", version)
	fmt.Printf("%v\n", dirty)

	if dirty {
		// dirty version이 껴있는 경우
		m.Drop()
		// drop 시켜버리고, migrate를 재 실행
	}

	// 이런식으로 작성하면 내가 migrate안에 있는 값을 수정할 떄마다 version을 갱신 시켜주지 않아도 된다.
	// 다른 방법으로는 version을 up하는 경우도 있는데, 그것보다는 개인적으로 개발할떄는 이게 훨씬 빠르게 가능

	if err := m.Migrate(version); err != nil {
		log.Fatal("-----: ", err)
	}
	return sqlc.New(db)
}
