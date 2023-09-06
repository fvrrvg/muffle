package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	for {
		startServer()
		time.Sleep(2 * time.Second)
	}
}

func startServer() {
	print("Starting server...")
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	authConfig := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)
	client.AutoRetry = true

	app := fiber.New(
		fiber.Config{
			Prefork: true,
			AppName: "Muffle",
			GETOnly: true,
		},
	)

	// Middleware
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET",
		}),
		limiter.New(limiter.Config{
			Max:        30,
			Expiration: 30 * time.Second,

			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).SendString("Too Many Requests")
			},
		}),
	)

	app.Get("/random", func(c *fiber.Ctx) error {
		var searchResults *spotify.SearchResult

		for {
			searchResults, _ = client.Search(Querygen(), spotify.SearchTypeTrack)
			if searchResults.Tracks.Total > 0 && searchResults.Tracks.Tracks[0].PreviewURL != "" {
				break
			}
		}

		return c.JSON(fiber.Map{
			"track":      searchResults.Tracks.Tracks[0].Name,
			"artist":     searchResults.Tracks.Tracks[0].Artists[0].Name,
			"album":      searchResults.Tracks.Tracks[0].Album.Name,
			"image":      searchResults.Tracks.Tracks[0].Album.Images[0].URL,
			"preview":    searchResults.Tracks.Tracks[0].PreviewURL,
			"year":       searchResults.Tracks.Tracks[0].Album.ReleaseDate,
			"spotifyurl": "https://open.spotify.com/track/" + searchResults.Tracks.Tracks[0].ID,
		})
	})

	err = app.Listen(":5000")
	if err != nil {
		log.Printf("Server error occurred: %v\n", err)
		startServer()
	}
}
