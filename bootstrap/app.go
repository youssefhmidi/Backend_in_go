package bootstrap

import "github.com/youssefhmidi/Backend_in_go/database"

type Application struct {
	Database database.SqliteDatabase
	Env      Env
}

func App() *Application {
	db := InitDB("./database/db/testdb.db")
	env := NewEnv(".env")

	return &Application{
		Database: db,
		Env:      env,
	}
}
