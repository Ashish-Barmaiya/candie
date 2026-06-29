package config

func DefaultConfig() Config {
	return Config{
		PictureDirectory:          "",
		DefaultExtractionStrategy: SceneExtractionStrategy,
		DefaultInterval:           60,

		CropBlackBars: true,

		RotationStrategy:          SequentialRotationStrategy,
		WallpaperRotationInterval: 120,

		AutoResume: true,

		LogLevel: InfoLogLevel,
	}
}
