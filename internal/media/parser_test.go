package media

import (
	"math"
	"testing"
	"time"
)

const sampleFFprobeOutput = `
{
  "streams": [
    {
      "codec_type": "video",
      "codec_name": "h264",
      "width": 1920,
      "height": 1080,
      "avg_frame_rate": "24000/1001"
    }
  ],
  "format": {
    "format_name": "matroska",
    "duration": "10123.456"
  }
}
`

func TestParseMetadata(t *testing.T) {
	metadata, err := parseMetadata(
		"/movies/interstellar.mkv",
		[]byte(sampleFFprobeOutput),
	)
	if err != nil {
		t.Fatalf("parseMetadata() returned error: %v", err)
	}

	if metadata.Path != "/movies/interstellar.mkv" {
		t.Errorf("expected path '/movies/interstellar.mkv', got %q", metadata.Path)
	}

	if metadata.Format != "matroska" {
		t.Errorf("expected format 'matroska', got %q", metadata.Format)
	}

	if metadata.Codec != "h264" {
		t.Errorf("expected codec 'h264', got %q", metadata.Codec)
	}

	if metadata.Width != 1920 {
		t.Errorf("expected width 1920, got %d", metadata.Width)
	}

	if metadata.Height != 1080 {
		t.Errorf("expected height 1080, got %d", metadata.Height)
	}

	if math.Abs(metadata.FPS-23.976) > 0.001 {
		t.Errorf("expected fps ≈ 23.976, got %f", metadata.FPS)
	}

	expectedDuration := time.Duration(10123.456 * float64(time.Second))

	if metadata.Duration != expectedDuration {
		t.Errorf(
			"expected duration %v, got %v",
			expectedDuration,
			metadata.Duration,
		)
	}
}

func TestParseFPS(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:  "23.976 fps",
			input: "24000/1001",
			want:  23.976023976,
		},
		{
			name:  "30 fps",
			input: "30/1",
			want:  30,
		},
		{
			name:  "60 fps",
			input: "60/1",
			want:  60,
		},
		{
			name:  "zero numerator",
			input: "0/1",
			want:  0,
		},
		{
			name:    "invalid format",
			input:   "abc",
			wantErr: true,
		},
		{
			name:    "division by zero",
			input:   "1/0",
			wantErr: true,
		},
		{
			name:    "invalid numerator",
			input:   "abc/1",
			wantErr: true,
		},
		{
			name:    "invalid denominator",
			input:   "1/abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFPS(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("expected %f, got %f", tt.want, got)
			}
		})
	}
}
