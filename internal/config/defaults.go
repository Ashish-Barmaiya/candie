package config

func DefaultConfig() Config {
	return Config{
		PictureDirectory:          "",
		DefaultExtractionStrategy: "scene",
		DefaultInterval:           60,

		CropBlackBars: true,

		RotationStrategy:          "sequential",
		WallpaperRotationInterval: 120,

		AutoResume: true,

		LogLevel: "info",
	}
}
