package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New(creds string) (*sql.DB, error) {

	db, err := sql.Open("mysql", creds)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE Table IF NOT EXISTS TASKS ( id int auto_increment primary key, task text, completed bool);")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE Table IF NOT EXISTS USERS ( uid int auto_increment primary key, name text);")
	if err != nil {
		return nil, err
	}

	return db, nil

}
