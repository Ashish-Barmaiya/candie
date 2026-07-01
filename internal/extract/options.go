package extract

import (
	"time"

	"github.com/ashish-barmaiya/candie/internal/collection"
	"github.com/ashish-barmaiya/candie/internal/media"
)

type Options struct {
	Strategy Strategy

	Interval time.Duration

	SkipStart time.Duration
	SkipEnd   time.Duration

	CropBlackBars bool

	Metadata media.Metadata

	Collection collection.Collection
}
