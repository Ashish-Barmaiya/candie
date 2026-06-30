package collection

import "time"

const ManifestVersion = 1

type Manifest struct {
	Version          int       `json:"version"`
	Movie            string    `json:"movie"`
	Strategy         string    `json:"strategy"`
	FrameCount       int       `json:"frame_count"`
	CreatedAt        time.Time `json:"created_at"`
	SkipStartSeconds int       `json:"skip_start_seconds"`
	SkipEndSeconds   int       `json:"skip_end_seconds"`
	CropBlackBars    bool      `json:"crop_black_bars"`
}
