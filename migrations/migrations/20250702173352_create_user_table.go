package migrations

import "gofr.dev/pkg/gofr/migration"

const createUser = `CREATE Table IF NOT EXISTS USERS (
    uid int auto_increment primary key, 
    name text
	);`

func create_user_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createUser)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
