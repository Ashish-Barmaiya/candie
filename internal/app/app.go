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
	"github.com/ashish-barmaiya/candie/internal/terminal"
)

// RunExtract extracts wallpapers from a movie and saves them to a collection.
func RunExtract(movie string) error {
	start := time.Now()

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

	fmt.Printf("Movie      : %s\n", collectionName)
	fmt.Printf("Strategy   : Timer\n")
	fmt.Printf("Interval   : 60s\n")
	fmt.Printf("Output     : %s\n\n", filepath.Join(rootDir, collectionName))

	coll, err := collection.Create(rootDir, collectionName)
	if err != nil {
		return err
	}

	spin := terminal.NewSpinner("Extracting wallpapers...")
	spin.Start()

	err = extract.Extract(
		movie,
		extract.Options{
			Strategy: extract.Timer,
			Interval: 60 * time.Second,

			Collection: coll,
			Metadata:   metadata,
		},
	)

	spin.Stop()

	if err != nil {
		return err
	}

	fmt.Println("✓ Extraction completed.")
	fmt.Printf("Location   : %s\n", coll.Path)
	fmt.Printf("Elapsed    : %s\n", time.Since(start).Round(time.Millisecond))

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
