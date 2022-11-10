package types

var (
	// MysqlSource mysql source, e.g: root:admin@(127.0.0.1:3306)/bench?charset=utf8&parseTime=True&loc=UTC.
	MysqlSource  = "root:lgh123456@(localhost:3306)/bench?charset=utf8&parseTime=True&loc=UTC"
	MysqlMaxIdle = 500
	MysqlMaxConn = 500
)
