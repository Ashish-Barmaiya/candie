package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ashish-barmaiya/candie/internal/collection"
	"github.com/ashish-barmaiya/candie/internal/extract"
	"github.com/ashish-barmaiya/candie/internal/media"
)

// RunExtract extracts wallpapers from a movie and saves them to a collection.
func RunExtract(movie string) error {
	metadata, err := media.Probe(movie)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	rootDir := filepath.Join(homeDir, "Pictures", "Candie")

	collectionName := strings.TrimSuffix(
		filepath.Base(movie),
		filepath.Ext(movie),
	)

	fmt.Printf("Extracting wallpapers from %q...\n\n", collectionName)

	coll, err := collection.Create(rootDir, collectionName)
	if err != nil {
		return err
	}

	err = extract.Extract(
		movie,
		extract.Options{
			Strategy: extract.Timer,
			Interval: 60 * time.Second,

			Collection: coll,
			Metadata:   metadata,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("✓ Extraction completed.")
	fmt.Printf("Location: %s\n", coll.Path)

	return nil
}

func RunPlay() {
	fmt.Println("Play command is not implemented yet")
}

func RunDoctor() {
	fmt.Println("Doctor command is not implemented yet")
}

func RunVersion() {
	fmt.Println("Candie 0.1.0-dev")
}
