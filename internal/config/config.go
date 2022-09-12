package config

const DatastoreType string = "sqlite3"

type Config struct {
	Datastore struct {
		File string `env:"DB_FILE_PATH" env-default:"employees.db?_fk=1"`
	}
}
