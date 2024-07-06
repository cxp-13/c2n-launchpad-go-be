package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	DBHost          string
	DBPort          string
	DBName          string
	DBUsername      string
	DBPassword      string
	DBSSLMode       string
	OwnerPrivateKey string
)

func Init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
		os.Exit(1)
	}

	OwnerPrivateKey = cfg.Section("private_key").Key("OWNER_PRIVATE_KEY").String()

	DBHost = cfg.Section("postgres").Key("DB_HOST").String()
	DBPort = cfg.Section("postgres").Key("DB_PORT").String()
	DBName = cfg.Section("postgres").Key("DB_NAME").String()
	DBUsername = cfg.Section("postgres").Key("DB_USERNAME").String()
	DBPassword = cfg.Section("postgres").Key("DB_PWD").String()
	DBSSLMode = cfg.Section("postgres").Key("DB_SSLMODE").String()
}
