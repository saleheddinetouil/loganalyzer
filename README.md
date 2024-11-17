# Log Analyzer and Alerting System

This project implements a log analysis and alerting system written in Go. It reads log files, parses them according to a configurable format, and triggers alerts based on predefined rules.

## Features

* Customizable log parsing using regular expressions.
* Flexible rule-based alerting.
* Supports multiple alerting methods (e.g., email, Slack).
* Concurrent processing for efficient log analysis.

## Getting Started

1. **Configuration:**  Modify `config.json` to specify the log file path, parsing rules, and alerting settings.
2. **Build:** `go build`
3. **Run:** `./loganalyzer`

## Configuration (config.json)

```json
{
  "logfile": "/path/to/logfile.txt",
  "parser": {
    "regex": "^(?P<timestamp>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}) (?P<level>\\w+) (?P<message>.*)$",
    "timeFormat": "2006-01-02 15:04:05"
  },
  "rules": [
    {
      "level": "ERROR",
      "message": ".*database connection error.*",
      "alerter": "email"
    }
  ],
  "alerters": {
    "email": {
      "to": "your_email@example.com",
      "from": "alerting_system@example.com"
      // ... other email settings
    },
    "slack": {
      "webhook_url": "your_slack_webhook_url"
    }
  }
}
