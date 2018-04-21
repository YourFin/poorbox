package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"log"
	"net/url"
	"time"
	"regexp"
	"strings"
)

//type Genre struct {
//	TmdbId int
//	Name   string
//}

// This is meant as a temporary data structure,
// and should really just be an intermediary to
// postgress
//type Movie struct {
//TmdbId int
//ImdbId string
//Genres []Genre
//TmdbPosterPath string
//TmdbBackgroundPath string
//Slug string //Short name
//Title string
//Runtime int
//ReleaseYear int
//Description string
//}

type searchResponseMovie struct {
	Title       string
	ReleaseDate string `json:"release_date"`
	Overview    string
	Id          int
}
type mSearchResponse struct {
	Results      []SearchResponseMovie
	TotalResults int `json:"total_results"`
}

// Doing it this way instead of having a struct
// as I don't really see any reason that we'll
// need to have multiple instances of the api
// up
var (
	apiKey string
	client *http.Client
	regexParens  regexp.Regexp
	regexYearEOL regexp.Regexp
)


func Init(ApiKey string) {
	apiKey = ApiKey
	client = &http.Client {
		Timeout: time.Second * 60
	}
	regexParens  = regexp.MustCompile("[()]")
	regexYearEOL = regexp.MustCompile("\\d{4}$")
}

func tmdbAPIHit(url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body, err
}

// Search for a movie on tmdb.
// Does a little bit of cleanup on movie
// names before running them into the api.
func movieSearch(rawMovieName, apiKey string, client *http.Client) (results []searchResponseMovie, err error) {
	// In its current iteration, this doesn't specifically attack the (2001)
	// at the end of file names, instead stripping parens and then dumping
	// anything that looks like a movie at the end of a string.
	// This could be changed to something more reasonable in the future.
	cleanMoviename := regexParens.ReplaceAllString(rawMovieName, "")
	cleanMoviename = strings.TrimSpace(cleanMoviename)
	proxy := regexYearEOL.ReplaceAllString(cleanMoviename, "")
	// Make sure that we haven't just removed the entire movie name
	// i.e. we were searching for a movie called 1984 or something
	if len(proxy) != 0 {
		cleanMoviename = proxy
	}
	cleanMoviename = url.QueryEscape(cleanMoviename)
	url := "https://api.themoviedb.org/3/search/movie?include_adult=false&page=1&query="
	url += cleanMoviename + "&language=en-US&api_key=" + apiKey
	body, err := tmdbAPIHit(url, client)
	if err != nil {
		return
	}
	var m mSearchResponse
	err = json.Unmarshal(body, &m)
	results = m.Results
	return
}
