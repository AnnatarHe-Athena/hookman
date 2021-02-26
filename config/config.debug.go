// +build !release

package config

const (
	IsDev     = true
	RedisHost = "localhost:6379"
	RavenDSN  = ""
	NRKey     = ""
	DB_DSN    = "host=localhost user=postgres password=admin dbname=postgres sslmode=disable"
	// DB_DSN  = "host=annatarhe.com user=postgres password=d8fdd2eb46a84a7c13656f012d118760ff43aef1e580e8577c7bcd157abac033 dbname=athena sslmode=disable"
	MaxPage = 1

	// IsDev     = false
	// RedisHost = "redis:6379"
	// RavenDSN  = "https://d340b4089d794feb8c892d820ce121df@sentry.io/1530743"
	// NRKey     = "e7d539a08f3e1f7b1b9291c28d2ca6f275ea4d2a"
	// DB_DSN    = "host=annatarhe.com user=postgres password=d8fdd2eb46a84a7c13656f012d118760ff43aef1e580e8577c7bcd157abac033 dbname=athena sslmode=disable"
	// MaxPage   = 2
)
