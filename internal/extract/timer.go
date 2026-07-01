package extract

import (
	"fmt"
	"path/filepath"

	"github.com/ashish-barmaiya/candie/internal/media"
)

// extractTimer extracts frames from a movie at regular intervals using the Timer strategy.
func extractTimer(
	movie string,
	metadata media.Metadata,
	options Options,
) error {
	end := metadata.Duration - options.SkipEnd

	if end <= options.SkipStart {
		return fmt.Errorf("invalid extraction window")
	}

	cmd := commandOptions{
		Input: movie,

		Start: options.SkipStart,
		End:   end,

		Interval: options.Interval,

		Output: filepath.Join(
			options.Collection.Path,
			FramePattern,
		),

		CropBlackBars: options.CropBlackBars,
	}

	return runFFmpeg(cmd)
}
