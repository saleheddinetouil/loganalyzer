package main

import (
    "encoding/json"
    "os"
)

// Config struct to define structure of config.json
type Config struct {
    LogFile  string            `json:"logfile"`
    Parser   ParserConfig     `json:"parser"`
    Rules    []RuleConfig      `json:"rules"`
    Alerters map[string]AlertConfig `json:"alerters"`
}

type ParserConfig struct {
    Regex string `json:"regex"`
    TimeFormat string `json:"timeFormat"`
}

type AlertConfig struct {
    To string `json:"to"`
    From string `json:"from"`
    // Add other alerter specific config here.

}

type RuleConfig struct {
    Level    string `json:"level"`
    Message  string `json:"message"`
    Alerter string `json:"alerter"`
}


func LoadConfig(filepath string) (*Config, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()


    decoder := json.NewDecoder(file)
    var config Config
    err = decoder.Decode(&config)
    if err != nil {
        return nil, err
    }


    return &config, nil
}
