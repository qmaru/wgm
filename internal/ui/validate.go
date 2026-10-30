package ui

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	maxPort    = 65535
	minMTU     = 64
	maxMTU     = 1518
	maxHostLen = 253
)

var (
	ipv4Pattern = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
	hostPattern = regexp.MustCompile(`(?i)^[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$`)
)

func isIPv4(v string) bool {
	if !ipv4Pattern.MatchString(v) {
		return false
	}
	for _, part := range strings.Split(v, ".") {
		n, _ := strconv.Atoi(part)
		if n > 255 {
			return false
		}
	}
	return true
}

func isHost(v string) bool {
	if len(v) > maxHostLen {
		return false
	}
	for _, label := range strings.Split(v, ".") {
		if !hostPattern.MatchString(label) {
			return false
		}
	}
	return true
}

func parseOptionalInt(v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	return strconv.Atoi(v)
}

// displayKey shortens a WireGuard key the way the cards show it.
func displayKey(key string) string {
	if len(key) > 10 {
		return key[:4] + "..." + key[len(key)-6:]
	}
	return key
}
