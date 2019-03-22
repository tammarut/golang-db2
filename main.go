package main

//!Practice to connect KTBDatabase(DB2)
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/asifjalil/cli"
)

func main() {
	con := "DATABASE=ODSS; HOSTNAME=10.2.154.16; PORT=50001; PROTOCOL=TCPIP; UID=odsadm; PWD=passw0rd;"

	//=>Connect
	db, err := sql.Open("cli", con)
	//=>Handle if fail
	if err != nil {
		log.Fatal()
	}
	//=>Close
	defer db.Close()

	//=>Test ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//=>Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//=>Close
	defer cancel()

	//=>Set Query SQL style
	rows, err := db.QueryContext(ctx, "SELECT * FROM ACTION_CODE_GO_URL")
	if err != nil {
		log.Fatal(err)
	}
	//=>Close
	defer rows.Close()

	//=>Wanna print all
	for rows.Next() {
		var actionCode string
		var url string

		//=>Scan data out
		if err := rows.Scan(&actionCode, &url); err != nil {
			log.Fatal(err)
		}

		fmt.Println(actionCode, url) //=>End
	}

}
