package templates

import (
	"fmt"
	"strconv"
	"time"
)

// Format duration string, such as "03:55", from second number
func formatDuration(t int64) string {
	return pad(t/60) + ":" + pad(t%60)
}

// Pad with zero, if needed
func pad(n int64) (s string) {
	s = strconv.FormatInt(n, 10)
	if n < 10 {
		s = "0" + s
	}
	return
}

// Renders readable elapsed/remaining time
func renderTime(t int64) string {
	t = time.Now().Unix() - t
	isFuture := false
	if t < 1 {
		isFuture = true
		t = -t
	}

	if t < 60 {
		if isFuture {
			return "soon™"
		}
		return "just now"
	}
	t = t / 60
	if t < 60 {
		return ago(t, "minute", isFuture)
	}
	return ago(t, "hour", isFuture)
}

// Formats "56 minutes ago" or "in 56 minutes" like relative time text
func ago(t int64, word string, isFuture bool) string {
	count := pluralize(int(t), word)
	if isFuture {
		return "in " + count
	}
	return count + " ago"
}

// Return either the singular or plural form of a word, depending on n
func pluralize(n int, word string) string {
	s := fmt.Sprintf("%d %s", n, word)
	if n != 1 && n != -1 {
		s += "s"
	}
	return s
}
