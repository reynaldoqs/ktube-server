package main

import "github.com/reynaldoqs/ka/internal/infrastructure/server"

func main() {
	server.RegisterRouter(":8080")
}

/*
var (
	query      = flag.String("query", "audi", "Search term")
	maxResults = flag.Int64("max-results", 4, "Max Youtube results")
)

const developerKey = "AIzaSyCuHwjTi26PHYP_v2qrBh9jRdgAeVzmXeg"

func main() {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new Youtube client: %v", err)
	}
	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).Q(*query).MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	printIDs("Channels", channels)
	printIDs("Playlists", playlists)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
f("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
*/
