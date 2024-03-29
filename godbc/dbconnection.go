package godbc

import (
	"database/sql"
	"strings"
)

//DbConnection DbConnection controller
type DbConnection struct {
	driver         string
	datasourcename string
	db             *sql.DB
	dbmngr         *DbManager
}

//Execute - Execute (query) refer to golang sql connection Execute method
func (cn *DbConnection) Execute(query string, args ...interface{}) (lastInsertID int64, rowsAffected int64, err error) {
	stmnt := &DbStatement{cn: cn}
	lastInsertID, rowsAffected, err = stmnt.Execute(query, args...)
	stmnt = nil
	return
}

//ParseQuery parse query and return parsed query based on db connection type and list of paramnames
//query e.g 'SELECT :TEST-PARAM AS param' where :TEST-PARAM is name of parameter
func (cn *DbConnection) ParseQuery(query string) (parsedquery string, params []string) {
	startParam := false
	startText := false
	pname := ""
	for n := range query {
		c := string(query[n])
		if startParam {
			if strings.TrimSpace(c) == "" || strings.Index("[](),@$%&|!<>$*+-'", c) > -1 {
				if pname != "" {
					if params == nil {
						params = []string{}
					}
					params = append(params, pname)
					pname = ""
					parsedquery = parsedquery + "?" + c
				} else {
					parsedquery = parsedquery + ":" + c
				}
				startParam = false

			} else {
				pname = pname + c
			}
		} else {
			if c == "'" {
				if startText {
					startText = false
				} else {
					startText = true
				}
			}
			if !startParam {
				if !startText {
					if !startParam && c == ":" && n < len(query)-1 {
						if strings.TrimSpace(c) != "" && strings.Index("[](),@$%&|!<>$*+-'", c) == -1 {
							startParam = true
						} else {
							parsedquery = parsedquery + c
						}
					} else {
						parsedquery = parsedquery + c
					}
				} else {
					parsedquery = parsedquery + c
				}
			}
		}
	}

	if startParam {
		if pname != "" {
			if params == nil {
				params = []string{}
			}
			params = append(params, pname)
			pname = ""
			parsedquery = parsedquery + "?"
		} else {
			parsedquery = parsedquery + ":"
		}
		startParam = false
	}
	return parsedquery, params
}

//Query - Query (query) refer to golang sql connection Query method
//except that it returns and DbResultSet that extends standard resultset functionality
func (cn *DbConnection) Query(query string, args ...interface{}) (rset *DbResultSet, err error) {
	stmnt := &DbStatement{cn: cn}
	rset, err = stmnt.Query(query, args...)
	return rset, err
}

func openDB(drvr string, datasourcename string) (driver string, db *sql.DB, err error) {
	var drivers = sql.Drivers()
	var hasDriver = false
	for _, drv := range drivers {
		if drv == drvr {
			hasDriver = true
			driver = drvr
			break
		}
	}
	if hasDriver {
		db, err = sql.Open(driver, datasourcename)
		if err = db.Ping(); err != nil {
			db = nil
		}
	}
	return driver, db, err
}

func openConnection(dbMngr *DbManager, driver string, datasourcename string) (cn *DbConnection, err error) {
	if driver, db, dberr := openDB(driver, datasourcename); dberr == nil && db != nil {
		cn = &DbConnection{dbmngr: dbMngr, driver: driver, datasourcename: datasourcename, db: db}
	} else {
		err = dberr
	}
	return cn, err
}
