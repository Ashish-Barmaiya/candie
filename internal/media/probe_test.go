package media

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProbe(t *testing.T) {
	if _, err := exec.LookPath("ffprobe"); err != nil {
		t.Skip("ffprobe is not installed")
	}

	video := filepath.Join("testdata", "sample.mp4")

	metadata, err := Probe(video)
	if err != nil {
		t.Fatalf("Probe() returned error: %v", err)
	}

	if metadata.Path != video {
		t.Errorf("expected path %q, got %q", video, metadata.Path)
	}

	if metadata.Width != 1920 {
		t.Errorf("expected width 1920, got %d", metadata.Width)
	}

	if metadata.Height != 1080 {
		t.Errorf("expected height 1080, got %d", metadata.Height)
	}

	if metadata.Codec == "" {
		t.Error("expected codec to be populated")
	}

	if metadata.Format == "" {
		t.Error("expected format to be populated")
	}

	if metadata.Duration <= 0 {
		t.Error("expected positive duration")
	}

	if metadata.FPS <= 0 {
		t.Error("expected positive FPS")
	}
}

func TestProbeNonExistentFile(t *testing.T) {
	if _, err := exec.LookPath("ffprobe"); err != nil {
		t.Skip("ffprobe is not installed")
	}

	_, err := Probe("does-not-exist.mp4")
	if err == nil {
		t.Fatal("expected error")
	}
}
