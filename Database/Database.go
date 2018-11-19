package Database

import (
	"database/sql"
	"fmt"
	"google.golang.org/appengine"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "google.golang.org/appengine"
)

func Database() (*sql.DB, error) {
	var (
		//local_proxy= os.Getenv("CLOUDSQL_PROXY")
		//local_proxy= os.Getenv("CLOUDSQL_PROXY")
		connectionName= os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user= os.Getenv("CLOUDSQL_USER")
		password= os.Getenv("CLOUDSQL_PASSWORD") // NOTE: password may be empty
		dsn string
	)

	cloudengine :=  appengine.IsDevAppServer()
	appserver := appengine.IsAppEngine()

	if ! appserver || cloudengine  {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s",
				user,
				password,
				connectionName,
				"mntbudget")
	} else {
		dsn = fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName)
	}
	db, err := sql.Open( "mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db,err
}
