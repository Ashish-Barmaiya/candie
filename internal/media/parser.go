package media

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseMetadata(path string, ffprobeJSON []byte) (Metadata, error) {
	var out ffprobeOutput

	if err := json.Unmarshal(ffprobeJSON, &out); err != nil {
		return Metadata{}, err
	}

	var video *ffprobeStream

	for i := range out.Streams {
		if out.Streams[i].CodecType == "video" {
			video = &out.Streams[i]
			break
		}
	}

	if video == nil {
		return Metadata{}, fmt.Errorf("no video stream found")
	}

	durationSeconds, err := strconv.ParseFloat(out.Format.Duration, 64)
	if err != nil {
		return Metadata{}, err
	}

	fps, err := parseFPS(video.AvgFrameRate)
	if err != nil {
		return Metadata{}, err
	}

	return Metadata{
		Path: path,

		Format: out.Format.FormatName,
		Codec:  video.CodecName,

		Width:  video.Width,
		Height: video.Height,

		FPS: fps,

		Duration: time.Duration(durationSeconds * float64(time.Second)),
	}, nil
}

func parseFPS(rate string) (float64, error) {
	parts := strings.Split(rate, "/")

	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid frame rate: %q", rate)
	}

	num, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}

	den, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}

	if den == 0 {
		return 0, fmt.Errorf("frame rate denominator cannot be zero")
	}

	return num / den, nil
}
