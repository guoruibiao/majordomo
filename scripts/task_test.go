package scripts

import (
	"testing"
	"time"
)

func TestStartClipboardListening(t *testing.T) {
	StartClipboardListening()
	time.Sleep(time.Second * 10)
}
