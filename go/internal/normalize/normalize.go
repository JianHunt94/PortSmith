// Package normalize converts loosely-structured session logs into canonical
// trace records and computes summary statistics over a trace stream.
package normalize

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/portsmith/portcap/internal/trace"
)

// FromRawLog parses a simple line-oriented session log into trace records.
//
// Each line is "<dir> <text>" where <dir> is ">" or "<". Consecutive lines
