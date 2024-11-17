package main


import (
    "fmt"
    "net/smtp"
    "strings"
)


func SendEmailAlert(config AlertConfig, entry LogEntry) error {

    // Build the email message
    subject := fmt.Sprintf("Log Alert: %s", entry.Level)
    body := fmt.Sprintf("Timestamp: %s\nLevel: %s\nMessage: %s\n", entry.Timestamp, entry.Level, entry.Message)
    msg := []byte("To: " + config.To + "\r\n" +
        "From: " + config.From + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + body)


    // Connect to the SMTP server


    err := smtp.SendMail("your_smtp_server:587",
        smtp.PlainAuth("", "your_smtp_username", "your_smtp_password", "your_smtp_server"),
        config.From, []string{config.To}, msg)

    if err != nil {
        return err
    }


    return nil
}


func SendAlert(config *Config, entry *LogEntry) error{
    for _, rule := range config.Rules {

        matchedLevel := rule.Level == "" || rule.Level == entry.Level
        matchedMessage, _ := regexp.MatchString(rule.Message, entry.Message) // error handling omitted for brevity



        if matchedLevel && matchedMessage{


             if alerterConfig, ok := config.Alerters[rule.Alerter]; ok {

                // choose and call correct alerter function
                if rule.Alerter == "email"{
                   return SendEmailAlert(alerterConfig, *entry)
                }  // ...Handle other alerter types

             }
        }

    }
    return nil

}
