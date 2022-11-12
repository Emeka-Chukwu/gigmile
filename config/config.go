package config

import (
	"database/sql"
	"gigmile-task/data"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}
