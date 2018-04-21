# tmdb
This package handles api requests to themoviedb.org, and the baggage that goes along with those
api requests.

## cmd-add.go
cmd-add provides the command line interface for having users select the movie that
matches the given file best

### Why is cmd-add here and not in the cmd package?
Due to the amount of api requests involved, keeping it in here allows
this package to expose less to the outside world, and allow for a more consistent
interface. Also there's a lot of code involved and I don't want it all in cmd/add.go
