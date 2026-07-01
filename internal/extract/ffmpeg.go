package extract

import (
	"os/exec"
	"strconv"
	"time"
)

type commandOptions struct {
	Input string

	Start time.Duration
	End   time.Duration

	Interval time.Duration

	Output string

	CropBlackBars bool
}

// runFFmpeg runs the ffmpeg command with the specified options to extract frames from a video.
func runFFmpeg(opts commandOptions) error {
	filter := "fps=1/" + strconv.FormatFloat(
		opts.Interval.Seconds(),
		'f',
		-1,
		64,
	)

	_ = opts.CropBlackBars

	args := []string{
		"-y",

		"-ss",
		strconv.FormatFloat(opts.Start.Seconds(), 'f', -1, 64),

		"-to",
		strconv.FormatFloat(opts.End.Seconds(), 'f', -1, 64),

		"-i",
		opts.Input,

		"-vf",
		filter,

		"-q:v",
		JPEGQuality,

		opts.Output,
	}

	cmd := exec.Command("ffmpeg", args...)

	return cmd.Run()
}
