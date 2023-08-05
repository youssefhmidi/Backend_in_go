package bootstrap

import "github.com/youssefhmidi/Backend_in_go/database"

type Application struct {
	Database database.SqliteDatabase
	Env      Env
}

func App() *Application {
	var db database.SqliteDatabase
	env := NewEnv(".env")
	if env.ReleaseMode {
		db = InitDB("./database/db/officialdb.db")
		return &Application{
			Database: db,
			Env:      env,
		}
	}
	db = InitDB("./database/db/testdb.db")
	return &Application{
		Database: db,
		Env:      env,
	}

}
