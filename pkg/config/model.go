package config

// M is config model
type M struct {
	Port       string // Heroku default port env.
	HerokuSlug string // HEROKUSLUG => HEROKU_APP_NAME

	// https://devcenter.heroku.com/articles/heroku-postgresql
	// example: postgres://xx:yyy@host:5432/d8slm9t7b5mjnd
	DatabaseHost string // DATABASEHOST => DATABASE_URL
}
