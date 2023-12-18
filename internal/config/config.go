package config

const DatastoreType string = "sqlite3"

type Config struct {
	Port      uint16 `json:"port" env:"PORT" env-default:"50051"`
	Datastore struct {
		File string `env:"DB_FILE_PATH" env-default:"employees.db?_fk=1"`
	}
}
