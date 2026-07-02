package terminal

import (
	"fmt"
	"sync"
	"time"
)

var frames = []rune{
	'⠋',
	'⠙',
	'⠹',
	'⠸',
	'⠼',
	'⠴',
	'⠦',
	'⠧',
	'⠇',
	'⠏',
}

type Spinner struct {
	message string

	done chan struct{}
	wg   sync.WaitGroup
}

func NewSpinner(message string) *Spinner {
	return &Spinner{
		message: message,
		done:    make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		i := 0

		for {
			select {
			case <-s.done:
				fmt.Print("\r\033[K")
				return

			case <-ticker.C:
				fmt.Printf("\r%c %s", frames[i], s.message)
				i = (i + 1) % len(frames)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	close(s.done)
	s.wg.Wait()
}
