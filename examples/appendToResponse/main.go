package main

import (
	"fmt"
	"os"

	tmdb "github.com/lunarr-app/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.Init(os.Getenv("APIKey"))

	if err != nil {
		fmt.Println(err)
	}

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "watch/providers"

	movie, err := tmdbClient.GetMovieDetails(299536, options)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range movie.MovieWatchProvidersAppend.WatchProviders.Results {
		fmt.Println(v.Flatrate)
	}
}
