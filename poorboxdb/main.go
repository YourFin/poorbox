package poorboxdb

import (
	"fmt"

	"github.com/go-pg/pg"
)

var db *pg.DB

func Connect(username, password, endpoint string) {
	db = pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Addr:     endpoint,
		Database: "poorbox",
	})
	fmt.Println(db)
}

func Close() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
