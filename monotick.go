// Package monotick provides monotonic timestamps as int64 nanoseconds.
//
// # Why int64?
//
// For measuring durations, you don't need wall clock time — just a counter.
// int64 nanoseconds can represent ~292 years, while int32 overflows in ~2 seconds.
//
// # How it works
//
// A reference time.Time is captured at package init. Now() returns
// time.Since(start) as int64 nanoseconds. This approach:
//
//   - Uses only public Go APIs (no go:linkname or runtime internals)
//   - Requires no CGO
//   - Works on all platforms
//
// Timestamps are relative to process start, not system boot. This is
// sufficient for measuring durations within a process.
//
// # Example
//
//	start := monotick.Now()
//	// ... do work ...
//	elapsed := monotick.Since(start)
package monotick

import "time"

var start = time.Now()

// Now returns nanoseconds elapsed since process start (monotonic).
func Now() int64 {
	return int64(time.Since(start))
}

// Since returns nanoseconds elapsed since t, where t is a value
// previously returned by Now.
func Since(t int64) int64 {
	return Now() - t
}
