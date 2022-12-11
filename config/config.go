package config

import "os"

var (
	DSN     = os.Getenv("DSN")
	SignKey = os.Getenv("SIGNING-KEY")
)
