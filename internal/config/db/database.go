package db

type DatabaseList struct {
	Product struct {
		Mysql Database
	}
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
