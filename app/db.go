package app

import (
	
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DbConnector struct {
	Dbname string
	Db *sql.DB
}



func (con *DbConnector) Connect() error {
	var err error 
	con.Db, err = sql.Open("sqlite3", con.Dbname)
	if err != nil {
		return err
	}
	return nil
}

func (con *DbConnector) Initialise() error {
	if con.Db != nil {
		_, err := con.Db.Exec(`Create Table kfkconfig (
					kfkid  INTEGER PRIMARY KEY AUTOINCREMENT,
					title VARCHAR(75) NOT NULL,
					description TEXT NULL,
					last_updated_on DATETIME 
				)`)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Connection not established to kfkconfig db.")
}

func (con *DbConnector) AddDefaultConfig(title, description string) error {
	if con.Db!= nil {
		stmt, err := con.Db.Prepare(`insert into kfkconfig(title,description,last_updated_on) values (?,?,datetime('now'))`)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(title, description)
		return err
	}
	return fmt.Errorf("Connection not established to kfkconfig db.") 
}