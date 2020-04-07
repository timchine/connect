package connect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbs        = make(map[string]*gorm.DB)
	configs    = make(map[string]*mysqlConfig)
	messageFmt = "============%s\n"
	DefaultDb  = ""
)

type mysqlConfig struct {
	host     string
	port     string
	user     string
	password string
	database string
	db       *gorm.DB
}

func NewMysqlConfig(host, port, user, password, database string) *mysqlConfig {
	conf := &mysqlConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
		db:       nil,
	}
	if len(configs) == 0 {
		DefaultDb = database
	}
	configs[database] = conf
	return conf
}

func (m *mysqlConfig) Connect() error {
	var (
		err error
	)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		m.user, m.password, m.host, m.port, m.database, "Local")
	fmt.Printf(messageFmt, "mysql 开始连接")
	m.db, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Printf(messageFmt, "mysql 连接失败")
		return err
	}
	fmt.Printf(messageFmt, "mysql 连接成功")
	//将连接放入缓存
	dbs[m.database] = m.db
	return nil
}

func GDB(databases ...string) *gorm.DB {
	return getDb(databases...)
}

func getDb(databases ...string) *gorm.DB {
	var (
		database  string
		connector *mysqlConfig
		gdb       *gorm.DB
	)

	if len(databases) == 0 {
		database = DefaultDb
	} else {
		database = databases[0]
	}
	if db, ok := dbs[database]; !ok || db == nil {
		c := configs[database]
		connector = NewMysqlConfig(c.host, c.port, c.user, c.password, c.database)
		err := connector.Connect()
		if err != nil {
			fmt.Printf(messageFmt, "mysql 连接失败")
		}
		dbs[database] = connector.db
		gdb = connector.db
	} else {
		gdb = dbs[database]
	}
	return gdb
}

func (m *mysqlConfig) Close() error {
	return m.db.Close()
}
