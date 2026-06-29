package config

type Config struct {
	PictureDirectory          string `yaml:"picture_directory"`
	DefaultExtractionStrategy string `yaml:"default_extraction"`
	DefaultInterval           int    `yaml:"default_interval"`

	CropBlackBars bool `yaml:"crop_black_bars"`

	RotationStrategy          string `yaml:"rotation_mode"`
	WallpaperRotationInterval int    `yaml:"wallpaper_interval"`

	AutoResume bool `yaml:"auto_resume"`

	LogLevel string `yaml:"log_level"`
}
