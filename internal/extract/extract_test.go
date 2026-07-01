package extract

import "testing"

func TestUnknownStrategy(t *testing.T) {
	err := Extract("movie.mkv", Options{
		Strategy: Strategy(100),
	})

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTimerNotImplemented(t *testing.T) {
	err := Extract("movie.mkv", Options{
		Strategy: Timer,
	})

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSceneNotImplemented(t *testing.T) {
	err := Extract("movie.mkv", Options{
		Strategy: Scene,
	})

	if err == nil {
		t.Fatal("expected error")
	}
}
