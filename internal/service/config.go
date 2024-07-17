package service

type Config struct {
	Port int
	Env  string
	DB   struct {
		dsn         string
		maxConns    int
		minConns    int
		maxIdleTime string
	}
	Limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	Smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}
