package inmem

import (
	"testing"
	"time"
)

const fmtTimestamp = "2006-01-02T15:04:05.00Z"

func TestCache_TrimOlderThan(t *testing.T) {
	events := []Event[any]{
		{
			Timestamp: parseTime(t, "2024-01-01T12:00:00.04Z"),
		}, {
			Timestamp: parseTime(t, "2024-01-01T12:00:00.05Z"),
		},
		{
			Timestamp: parseTime(t, "2024-01-01T12:00:00.06Z"),
		},
	}

	cache := &Cache{now: func() time.Time {
		return parseTime(t, "2024-01-01T12:00:00.06Z")
	}}

	cache.AddAll(events)
	cache.TimeOlderThan(parseTime(t, "2024-01-01T12:00:00.06Z").Add(-15 * time.Millisecond))

	got := cache.GetAll()
	expected := 2
	if len(got) != expected {
		t.Fatalf("expected %d, got %d", expected, len(got))
	}
}

func parseTime(t *testing.T, timestamp string) time.Time {
	parsedTime, err := time.Parse(fmtTimestamp, timestamp)
	if err != nil {
		t.Fatalf("invalid timestamp %s, err: %v", timestamp, err)
	}
	return parsedTime
}
