package model

type Config struct {
	Server *Server
	Mysql  *Mysql
	Redis  *Redis
}

type Server struct {
	Domain string
	Addr   string
	Port   int
}

type Mysql struct {
	Name     string
	Password string
	Addr     string
	Database string
	Port     int
}

type Redis struct {
	Addr     string
	Password string
}
