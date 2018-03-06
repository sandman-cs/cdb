package cdb

import (
	"database/sql"
	"time"

	"github.com/sandman-cs/core"

	//Load database drivers
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

// ConnectToDB -- Try to connect to the DB server as
// long as it takes to establish a connection
func ConnectToDB(dbServer string, dbUsr string, dbPwd string, dbName string, dbType string) *sql.DB {

	var err error
	var dB *sql.DB
	for {
		if dbType == "mysql" {
			dB, err = sql.Open("mysql", dbUsr+":"+dbPwd+"@tcp("+dbServer+")/"+dbName)
		} else if dbType == "mssql" {
			dB, err = sql.Open("mssql", "server="+dbServer+";user id="+dbUsr+"; password="+dbPwd+";database="+dbName)
		}
		if err == nil {
			err = dB.Ping()
			if err != nil {
				core.SyslogCheckError(err)
				time.Sleep(500 * time.Millisecond)
			} else {
				return dB
			}
		}
	}
}
