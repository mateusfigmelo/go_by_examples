package main

import (
	"strings"
	"testing"
	"time"

	"github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestSelectWithTimeout(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "timeout occurs",
			want: "Timeout: operation took too long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(selectWithTimeout)
			if !strings.Contains(output, tt.want) {
				t.Errorf("selectWithTimeout() output = %q, want %q", output, tt.want)
			}
		})
	}
}

func TestNonBlockingChannelOps(t *testing.T) {
	tests := []struct {
		name      string
		wantLines []string
	}{
		{
			name: "non-blocking operations",
			wantLines: []string{
				"No message received",
				"No message sent",
				"No activity",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(nonBlockingChannelOps)
			for _, want := range tt.wantLines {
				if !strings.Contains(output, want) {
					t.Errorf("nonBlockingChannelOps() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestClosingChannels(t *testing.T) {
	tests := []struct {
		name      string
		wantLines []string
	}{
		{
			name: "channel closing sequence",
			wantLines: []string{
				"Sent job 1",
				"Sent job 2",
				"Sent job 3",
				"Sent all jobs",
				"Received all jobs",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(closingChannels)
			for _, want := range tt.wantLines {
				if !strings.Contains(output, want) {
					t.Errorf("closingChannels() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestRangeOverChannels(t *testing.T) {
	tests := []struct {
		name      string
		wantLines []string
	}{
		{
			name: "channel iteration",
			wantLines: []string{
				"Queue contents:",
				"- one",
				"- two",
				"- three",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(rangeOverChannels)
			for _, want := range tt.wantLines {
				if !strings.Contains(output, want) {
					t.Errorf("rangeOverChannels() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestTimerExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "timer operations",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(timerExample)
			// Since timer operations are asynchronous and timing-dependent,
			// we only check that either Timer 1 expired or Timer 2 stopped
			if !strings.Contains(output, "Timer 1 expired") && !strings.Contains(output, "Timer 2 stopped") {
				t.Errorf("timerExample() output = %q, want either Timer 1 expired or Timer 2 stopped", output)
			}
		})
	}
}

func TestTickerExample(t *testing.T) {
	tests := []struct {
		name     string
		minTicks int
	}{
		{
			name:     "ticker operations",
			minTicks: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(tickerExample)
			tickCount := strings.Count(output, "Tick at")
			if tickCount < tt.minTicks {
				t.Errorf("tickerExample() got %d ticks, want at least %d", tickCount, tt.minTicks)
			}
		})
	}
}

func TestRateLimiter(t *testing.T) {
	tests := []struct {
		name      string
		wantLines []string
	}{
		{
			name: "rate limited requests",
			wantLines: []string{
				"Rate limited requests:",
				"Processing request 1",
				"Processing request 2",
				"Processing request 3",
				"Processing request 4",
				"Processing request 5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(rateLimiter)
			for _, want := range tt.wantLines {
				if !strings.Contains(output, want) {
					t.Errorf("rateLimiter() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestMain(t *testing.T) {
	start := time.Now()
	output := testutil.CaptureOutput(main)
	duration := time.Since(start)

	// Check for section headers
	wantSections := []string{
		"=== Go Advanced Channel Operations ===",
		"1. Select with Timeout:",
		"2. Non-Blocking Channel Operations:",
		"3. Closing Channels:",
		"4. Range over Channels:",
		"5. Timer Example:",
		"6. Ticker Example:",
		"7. Rate Limiter Example:",
	}

	for _, section := range wantSections {
		if !strings.Contains(output, section) {
			t.Errorf("main() missing section %q in output", section)
		}
	}

	// Check execution time is reasonable
	if duration > 5*time.Second {
		t.Errorf("main() took %v, want less than 5s", duration)
	}
}

// Example outputs for documentation
func Example() {
	main()
	// Output is non-deterministic due to concurrent operations
	// We only verify the section headers
	// Output:
	// === Go Advanced Channel Operations ===
	// 1. Select with Timeout:
	// 2. Non-Blocking Channel Operations:
	// 3. Closing Channels:
	// 4. Range over Channels:
	// 5. Timer Example:
	// 6. Ticker Example:
	// 7. Rate Limiter Example:
}
