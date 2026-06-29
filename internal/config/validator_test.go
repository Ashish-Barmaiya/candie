package config

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     DefaultConfig(),
			wantErr: false,
		},
		{
			name: "invalid extraction strategy",
			cfg: func() Config {
				cfg := DefaultConfig()
				cfg.DefaultExtractionStrategy = "invalid"
				return cfg
			}(),
			wantErr: true,
		},
		{
			name: "invalid rotation strategy",
			cfg: func() Config {
				cfg := DefaultConfig()
				cfg.RotationStrategy = "invalid"
				return cfg
			}(),
			wantErr: true,
		},
		{
			name: "invalid default interval",
			cfg: func() Config {
				cfg := DefaultConfig()
				cfg.DefaultInterval = 0
				return cfg
			}(),
			wantErr: true,
		},
		{
			name: "invalid wallpaper interval",
			cfg: func() Config {
				cfg := DefaultConfig()
				cfg.WallpaperRotationInterval = -10
				return cfg
			}(),
			wantErr: true,
		},
		{
			name: "invalid log level",
			cfg: func() Config {
				cfg := DefaultConfig()
				cfg.LogLevel = "verbose"
				return cfg
			}(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.cfg)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error=%v, got=%v", tt.wantErr, err)
			}
		})
	}
}
