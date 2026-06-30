package media

import "time"

type Metadata struct {
	Path string

	Format string
	Codec  string

	Width  int
	Height int

	FPS float64

	Duration time.Duration
}
