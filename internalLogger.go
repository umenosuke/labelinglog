package labelinglog

import (
	"fmt"
	"os"
)

func internalLog(timestamp string, fileName string, msg string) {
	fmt.Fprintln(os.Stderr, timestamp+"[logger][FATAL] "+fileName+": "+msg)
}
