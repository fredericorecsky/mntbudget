package Database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func Database() (*sql.DB, error){
	db, err := sql.Open( "mysql", "user:password@tcp(ipaddress:port)/mntbudget")

	if ( err != nil ){
		log.Fatal(err)
	}

	return db,err
}
