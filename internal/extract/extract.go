package extract

import "fmt"

func Extract(movie string, options Options) error {
	switch options.Strategy {
	case Timer:
		return fmt.Errorf("timer extraction not implemented")

	case Scene:
		return fmt.Errorf("scene extraction not implemented")

	default:
		return fmt.Errorf("unknown extraction strategy")
	}
}
