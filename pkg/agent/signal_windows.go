package agent

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
