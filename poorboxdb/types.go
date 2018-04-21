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
	TmdbId             int     `json: "id"`
	ImdbId             string  `json: "imdb_id`
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
}


// Create a slug
func (m *Movie) GiveSlug(allowOverwrite bool) error {
	if !allowOverwrite && m.Slug != "" {
		return errors.New("Attempted to overwrite existing slug")
	}
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
