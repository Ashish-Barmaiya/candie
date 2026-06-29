package config

import "fmt"

func Validate(cfg Config) error {
	switch cfg.DefaultExtractionStrategy {
	case SceneExtractionStrategy, TimerExtractionStrategy:
	default:
		return fmt.Errorf(
			"invalid default_extraction_strategy: %q",
			cfg.DefaultExtractionStrategy,
		)
	}

	switch cfg.RotationStrategy {
	case SequentialRotationStrategy, ShuffleRotationStrategy:
	default:
		return fmt.Errorf(
			"invalid rotation_strategy: %q",
			cfg.RotationStrategy,
		)
	}

	if cfg.DefaultInterval <= 0 {
		return fmt.Errorf(
			"default_interval must be greater than zero",
		)
	}

	if cfg.WallpaperRotationInterval <= 0 {
		return fmt.Errorf(
			"wallpaper_rotation_interval must be greater than zero",
		)
	}

	switch cfg.LogLevel {
	case DebugLogLevel,
		InfoLogLevel,
		WarnLogLevel,
		ErrorLogLevel:
	default:
		return fmt.Errorf(
			"invalid log_level: %q",
			cfg.LogLevel,
		)
	}

	return nil
}
