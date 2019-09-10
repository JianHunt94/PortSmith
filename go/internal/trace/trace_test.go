package trace

import (
	"bytes"
	"io"
	"testing"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
	recs := []Record{
		{TimestampNanos: 1000, Dir: Request, Session: "abc123", Proto: "redis", Payload: []byte("PING\r\n")},
		{TimestampNanos: 2000, Dir: Response, Session: "abc123", Proto: "redis", Payload: []byte("+PONG\r\n")},
