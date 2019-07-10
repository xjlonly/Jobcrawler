package model

type DBConfig struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	DbDriver string
	DbCharset string
	// if false, local default UTC +0
	IsLocalTime bool
}

