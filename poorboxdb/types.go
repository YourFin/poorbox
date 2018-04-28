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
	"fmt"
	"crypto/md5"
	"strings"
	"unicode"
	"errors"
	"encoding/base64"
)

type Genre struct {
	TmdbId int `json: "id"`
	Name   string
}

//WARNING: Fiddling with the order of the fields
// will cause slugs to be calculated differently, and
// possibly break stuff
type Movie struct {
	TmdbId             int     `json: "id",sql: "tmdb_id"`
	ImdbId             string  `json: "imdb_id,sql`
	Genres             []Genre `json: "genres"`
	TmdbPosterPath     string  `json: "poster_path"`
	TmdbBackgroundPath string  `json: "backdrop_path"`
	Slug               string  //Short name, not from tmdb
	Title              string
	Tagline            string  `json: "tagline"`
	Runtime            int
	ReleaseYear        int //not from tmdb

	// This is close to the same thing as the overview provided by tmdb,
	// however we actually want this to come from the initial query, so we
	// give it a bettor (read: different) name.
	Description        string
	Path               string
}


// Create a slug
func (m *Movie) GiveSlug(allowOverwrite bool) error {
	if !allowOverwrite && m.Slug != "" {
		return errors.New("Attempted to overwrite existing slug")
	}
	pathProxy := m.Path
	m.Path = ""
	defer func() { m.Path = pathProxy }()
	hash := md5.Sum([]byte(fmt.Sprintf("%v", *m)))
	shortTitle := strings.Map(func(rr rune) rune {
		if unicode.IsSpace(rr) {
			return -1
		}
		return unicode.ToLower(rr)
	}, m.Title)
	strHash := base64.URLEncoding.EncodeToString(hash[:16])
	m.Slug = fmt.Sprintf("%s%v", shortTitle[:10], strHash)
	return nil
}
