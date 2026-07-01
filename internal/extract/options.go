package extract

import (
	"time"

	"github.com/ashish-barmaiya/candie/internal/collection"
)

type Options struct {
	Strategy Strategy

	Interval time.Duration

	SkipStart time.Duration
	SkipEnd   time.Duration

	CropBlackBars bool

	Collection collection.Collection
}
