// Copyright © 2018 Patrick Nuckolls <nuckollsp+poorbox at gmail>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package poorboxdb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type sslType string

const (
	DB_NAME string = "poorbox"
	MOVIES_TABLE = "movies"
	SSL_DISABLE sslType = "disable"
	SSL_REQUIRE sslType = "require"
	//VERIFY_CA
	//VERIFY_FULL
)

var db *sql.DB

//Connect to the poorbox database
func Connect(username, password, endpoint string, ssl sslType) error {
	var err error
	connStr := "postgres://" + password + ":" + username + "@"
	connStr +=  endpoint + "/" + DB_NAME + "?sslmode=" + string(ssl)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func InitTables() error {
	_, err := db.Exec(`CREATE TABLE genres (
	id integer PRIMARY KEY,
	name text
);`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE movies (
	tmdb_id        int,
	imdb_id        text,
	background_url text,
	poster_url     text,
	full_file_path text NOT NULL,
	url            text NOT NULL,
	slug           text PRIMARY KEY,
	title          text NOT NULL,
	tagline        text NOT NULL,
	runtime        int  NOT NULL,
	release_year   int
);`)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

//func
