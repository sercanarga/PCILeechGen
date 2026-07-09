package mmio

import (
	"io"
	"strings"
	"sync"
	"testing"
	"time"
)

type quietTraceReader struct {
	closed chan struct{}
	once   sync.Once
}

func newQuietTraceReader() *quietTraceReader {
	return &quietTraceReader{closed: make(chan struct{})}
}

func (r *quietTraceReader) Read([]byte) (int, error) {
	<-r.closed
	return 0, io.EOF
}

func (r *quietTraceReader) Close() error {
	r.once.Do(func() { close(r.closed) })
	return nil
}

func TestCollectLiveTracePreservesKernelTimestampDeltas(t *testing.T) {
	input := strings.Join([]string{
		"R 4 99.000 0x9000 0x0",
		"R 4 100.125 0x1000 0x1",
		"W 4 100.375 0x1004 0x2",
	}, "\n")
	target := TraceTarget{BDF: "0000:03:00.0", BARIndex: 1, BARBase: 0x1000, BARSize: 0x1000}
	records, err := collectLiveTrace(io.NopCloser(strings.NewReader(input)), target, time.Second)
	if err != nil {
		t.Fatalf("collectLiveTrace: %v", err)
	}
	if len(records) != 2 {
		t.Fatalf("records = %+v, want two target records", records)
	}
	if records[0].Timestamp != 0 || records[1].Timestamp != 250*time.Millisecond {
		t.Fatalf("normalized timestamps = %v, %v, want 0 and 250ms", records[0].Timestamp, records[1].Timestamp)
	}
}

func TestCollectLiveTraceQuietReaderReturnsAtDeadline(t *testing.T) {
	reader := newQuietTraceReader()
	target := TraceTarget{BDF: "0000:03:00.0", BARIndex: 1, BARBase: 0x1000, BARSize: 0x1000}
	started := time.Now()
	records, err := collectLiveTrace(reader, target, 20*time.Millisecond)
	elapsed := time.Since(started)
	if err != nil {
		t.Fatalf("collectLiveTrace: %v", err)
	}
	if len(records) != 0 {
		t.Fatalf("quiet capture records = %+v, want none", records)
	}
	if elapsed < 10*time.Millisecond {
		t.Fatalf("quiet capture returned before deadline after %v", elapsed)
	}
	if elapsed > 500*time.Millisecond {
		t.Fatalf("quiet capture took %v, expected bounded deadline return", elapsed)
	}
	select {
	case <-reader.closed:
	default:
		t.Fatal("quiet reader was not closed at deadline")
	}
}
