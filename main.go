package main

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/hellerox/playlistSync/config"
)

func main() {

	configuration, err := config.New("./conf")
	if err != nil {
		log.Fatalln(err)
	}

	confi := &clientcredentials.Config{
		ClientID:     configuration.Spotify.ClientID,
		ClientSecret: configuration.Spotify.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := confi.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	user, err := client.GetUsersPublicProfile(spotify.ID(configuration.Spotify.Users[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	fmt.Println("+++++ User ID:", user.ID)

	playlists, err := client.GetPlaylistsForUser(configuration.Spotify.Users[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	if playlists.Playlists != nil {
		for _, item := range playlists.Playlists {
			log.Println("-", item.Name, item.ID)
			items, _ := client.GetPlaylist(item.ID)

			for _, item := range items.Tracks.Tracks {
				log.Println("+ ", item.Track.Name)

				for _, artist := range item.Track.Artists {
					log.Println("*--", artist.Name)
				}
			}
		}
	}
}
