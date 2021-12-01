package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
	"encoding/csv"
	"github.com/gocolly/colly"
)

func main() {
	// Create a file to store the scraped data.
	fName := "data.csv"
	file, err := os.Create(fName)
	if (err != nil) {
		log.Fatal("Could not create file. Error: ", err)
		return
	}
	defer file.Close()

	// Create a writer to write the scraped data to the file.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Initialize the collector.
	collector := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	// Point to web structure and specify what is needed from the page.
	collector.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		// Write data to csv file.
		writer.Write([]string {
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	for i := 0; i < 312; i++ {
		fmt.Printf("Scraping page: %d\n", i)

		collector.Visit("https://www.internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping complete")
	log.Print(collector)
}
