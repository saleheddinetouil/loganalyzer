package main

import (
	"regexp"
	"time"
)


type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func ParseLog(line string, parserConfig ParserConfig) (*LogEntry, error) {

	r, err := regexp.Compile(parserConfig.Regex)
	if err != nil {
		return nil, err  // Handle regex compilation error
	}


	match := r.FindStringSubmatch(line)
	if match == nil {
		return nil, nil // No match, potentially skip line
	}

    timestamp, err := time.Parse(parserConfig.TimeFormat, match[1]) // Assuming timestamp is the first capture group

	if err != nil {
        return nil, err // Handle timestamp parsing error

    }

	entry := &LogEntry{
		Timestamp: timestamp,
		Level:     match[2],  // Assuming level is second capture group
		Message:   match[3],  // Assuming message is third capture group
	}


	return entry, nil
}
