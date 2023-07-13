package clockface

import (
	"hello/mathclockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	mathclockface.SVGWriter(os.Stdout, t)
}
