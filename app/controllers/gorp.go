package controllers
import (
	"github.com/coopernurse/gorp"
	"database/sql"
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/memikequinn/lfs-server-go/app/models"
	"fmt"
)

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

type DbConfig struct {
	Host string
	Port string
	User string
	Pass string
	Dbname string
	Protocol string
	Dbargs string
	Driver string
	ConnectionString string
}

func (dbConfig *DbConfig) getConnectionString() (*DbConfig) {
	dbConfig.Host = getParamString("db.host", "")
	dbConfig.Port = getParamString("db.port", "3306")
	dbConfig.User = getParamString("db.user", "")
	dbConfig.Pass = getParamString("db.password", "")
	dbConfig.Dbname = getParamString("db.name", "lfs_server_go")
	dbConfig.Protocol = getParamString("db.protocol", "tcp")
	dbConfig.Dbargs = getParamString("db.dbargs", " ")
	dbConfig.Driver = getParamString("db.driver", " ")
	if strings.Trim(dbConfig.Dbargs, " ") != "" {
		dbConfig.Dbargs = "?" + dbConfig.Dbargs
	} else {
		dbConfig.Dbargs = ""
	}
	dbConfig.ConnectionString = fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		dbConfig.User, dbConfig.Pass, dbConfig.Protocol, dbConfig.Host, dbConfig.Port, dbConfig.Dbname, dbConfig.Dbargs)
	fmt.Printf("Set connection string %s\n", dbConfig.ConnectionString)
	fmt.Printf("Set Driver to %s\n", dbConfig.Driver)
	return dbConfig
}

var (
	Dbm *gorp.DbMap
)

var InitDB func() = func(){
	dbConfig := new(DbConfig)
	config := dbConfig.getConnectionString()
	if db, err := sql.Open(config.Driver, config.ConnectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db: db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}
	// Defines the table for use by GORP
	// This is a function we will create soon.
	defineObjectTable(Dbm)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

// Set up tables here - a bit janky but it looks like you need to
// define tables manually
func defineObjectTable(dbm *gorp.DbMap){
	// set "id" as primary key and autoincrement
	t := dbm.AddTable(models.Object{}).SetKeys(true, "id")
	// e.g. VARCHAR(25)
	t.ColMap("oid").SetMaxSize(65)
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			fmt.Printf("Set config for %s to %s\n", param, p)
			return defaultValue
		}
	}
	return p
}


