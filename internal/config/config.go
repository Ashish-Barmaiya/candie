package config

type Config struct {
	PictureDirectory          string             `yaml:"picture_directory"`
	DefaultExtractionStrategy ExtractionStrategy `yaml:"default_extraction_strategy"`
	DefaultInterval           int                `yaml:"default_interval"`

	CropBlackBars bool `yaml:"crop_black_bars"`

	RotationStrategy          RotationStrategy `yaml:"rotation_strategy"`
	WallpaperRotationInterval int              `yaml:"wallpaper_rotation_interval"`

	AutoResume bool `yaml:"auto_resume"`

	LogLevel LogLevel `yaml:"log_level"`
}
