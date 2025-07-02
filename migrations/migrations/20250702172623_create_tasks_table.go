package migrations

import "gofr.dev/pkg/gofr/migration"

const createTask = `CREATE Table IF NOT EXISTS TASKS ( 
    id int auto_increment primary key, 
    task text, 
    completed bool, 
    uid int
    );`

func create_tasks_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTask)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
