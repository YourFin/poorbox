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

package tmdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"strconv"

	pdb "github.com/yourfin/poorbox/poorboxdb"
)

type searchResponseMovie struct {
	Title       string
	ReleaseDate string `json:"release_date"`
	Overview    string
	Id          int
}
type mSearchResponse struct {
	Results      []searchResponseMovie
	TotalResults int `json:"total_results"`
}

// Doing it this way instead of having a struct
// as I don't really see any reason that we'll
// need to have multiple instances of the api
// up
var (
	apiKey       string
	client       *http.Client
	regexParens  *regexp.Regexp
	regexYearEOL *regexp.Regexp
)

func ApiInit(ApiKey string) {
	apiKey = ApiKey
	client = &http.Client{
		Timeout: time.Second * 60,
	}
	regexParens = regexp.MustCompile("[()]")
	regexYearEOL = regexp.MustCompile("\\d{4}$")
}
func maybeInit() {
	if client == nil {
		panic("uninitialized api")
	}
}

func searchResponseMovieToMovie(in searchResponseMovie) pdb.Movie {
	maybeInit()
	url := "https://api.themoviedb.org/3/movie/"
	url += strconv.Itoa(in.Id) + "?api_key=" + apiKey + "&language=en-US"
	body, err := tmdbAPIHit(url)
	if err != nil {
		panic("could not find movie with id" + strconv.Itoa(in.Id))
	}
	var out pdb.Movie
	err = json.Unmarshal(body, &out)
	out.Description = in.Overview
	if len(in.ReleaseDate) >= 4 {
		out.ReleaseYear, _ = strconv.Atoi(in.ReleaseDate[:4])
	}
	err = out.GiveSlug(false)
	if err != nil {
		panic("Slug somehow got parsed")
	}
	return out
}

func tmdbAPIHit(url string) ([]byte, error) {
	maybeInit()
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
func movieSearch(rawMovieName string) (results []searchResponseMovie, err error) {
	maybeInit()
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
	body, err := tmdbAPIHit(url)
	if err != nil {
		return
	}
	var m mSearchResponse
	err = json.Unmarshal(body, &m)
	results = m.Results
	return
}
