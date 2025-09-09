package main

import (
	"fmt"
	"strings"
	"time"
)

func parseTimestampPrefix(s string) (time.Time, string) {
	sp := strings.IndexByte(s, ' ')
	if sp > 0 {
		if ts, err := time.Parse(time.RFC3339Nano, s[:sp]); err == nil {
			return ts, strings.TrimSpace(s[sp+1:])
		}
	}
	return time.Now(), s
}

func enviarLog(sistema, linha string, timestamp time.Time, isErr bool) error {
	hour := timestamp.Local().Format("15:04:05")
	desc := fmt.Sprintf("[%s] %s", hour, linha)

	logType := "info"
	if isErr {
		logType = "error"
	}
	if err := sendDiscordMessage(sistema, desc, logType); err != nil {
		return err
	}

	if isErr {
		if err := sendTelegramMessage(sistema, linha); err != nil {
			return err
		}
	}

	return nil
}
