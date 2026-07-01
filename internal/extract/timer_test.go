package extract

import (
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/ashish-barmaiya/candie/internal/collection"
	"github.com/ashish-barmaiya/candie/internal/media"
)

func TestInvalidExtractionWindow(t *testing.T) {
	err := Extract(
		"movie.mkv",
		Options{
			Metadata: media.Metadata{
				Duration: 100 * time.Second,
			},

			Strategy: Timer,

			Interval: 60 * time.Second,

			SkipStart: 80 * time.Second,
			SkipEnd:   30 * time.Second,

			Collection: collection.Collection{},
		},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTimerExtraction(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg is not installed")
	}

	if _, err := exec.LookPath("ffprobe"); err != nil {
		t.Skip("ffprobe is not installed")
	}

	video := filepath.Join(
		"..",
		"media",
		"testdata",
		"sample.mp4",
	)

	metadata, err := media.Probe(video)
	if err != nil {
		t.Fatalf("Probe() failed: %v", err)
	}

	root := t.TempDir()

	coll, err := collection.Create(root, "sample")
	if err != nil {
		t.Fatalf("Create() failed: %v", err)
	}

	err = Extract(
		video,
		Options{
			Metadata:   metadata,
			Strategy:   Timer,
			Interval:   time.Second,
			Collection: coll,
		},
	)
	if err != nil {
		t.Fatalf("Extract() failed: %v", err)
	}

	frames, err := filepath.Glob(
		filepath.Join(coll.Path, "*.jpg"),
	)
	if err != nil {
		t.Fatalf("Glob() failed: %v", err)
	}

	if len(frames) == 0 {
		t.Fatal("expected at least one extracted frame")
	}
}
