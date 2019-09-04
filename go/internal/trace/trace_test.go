package trace

import (
	"bytes"
	"io"
	"testing"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
