package media

import (
	"fmt"
	"os/exec"
)

func Probe(path string) (Metadata, error) {
	ffprobeJSON, err := runFFprobe(path)
	if err != nil {
		return Metadata{}, err
	}

	return parseMetadata(path, ffprobeJSON)
}

func runFFprobe(path string) ([]byte, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		path,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ffprobe failed: %w", err)
	}

	return output, nil
}
