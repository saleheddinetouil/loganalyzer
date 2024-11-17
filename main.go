package main



import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
    "sync"
)



func main() {

	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}



	file, err := os.Open(config.LogFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


    var wg sync.WaitGroup // WaitGroup to manage goroutines


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

        wg.Add(1)

        // process line in go routine to improve performance for large number of alerts
        go func(line string){
            defer wg.Done()

		    entry, err := ParseLog(line, config.Parser)
		    if err != nil {
			    log.Println("Error parsing log:", err)
			    return  // Or handle error differently
		    }


		    if entry != nil { // Ensure entry is not nil if parsing fails


                err := SendAlert(config, entry)
                if err != nil{
                    log.Println("Error sending alert", err)
                }



		    }

        }(line)

	}
    wg.Wait() // Wait for all goroutines to finish



	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
