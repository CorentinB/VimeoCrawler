package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/remeh/sizedwaitgroup"
)

var asciiArt = `
╔═══════════════════════════════════════════════════════════════════════╗
║                                                                       ║
║ /$$$$$$$$ /$$                         /$$$$$$$$                       ║
║ |__  $$__/| $$                        | $$_____/                      ║
║    | $$   | $$$$$$$   /$$$$$$         | $$       /$$   /$$  /$$$$$$   ║
║    | $$   | $$__  $$ /$$__  $$       | $$$$$   | $$  | $$ /$$__  $$   ║
║    | $$   | $$  \ $$| $$$$$$$$|       | $$__/   | $$  | $$| $$$$$$$$  ║
║    | $$   | $$  | $$| $$_____/        | $$      | $$  | $$| $$_____/  ║
║    | $$   | $$  | $$|  $$$$$$$        | $$$$$$$$|  $$$$$$$|  $$$$$$$  ║
║    |__/   |__/  |__/ \_______/        |________/ \____  $$ \_______/  ║
║                                                  /$$  | $$            ║
║                                                 |  $$$$$$/            ║
║                                                  \______/             ║
║                                                                       ║
║ vimeo.com Crawler v1.0.0 By The French Guy @ The-Eye.eu               ║
║                                                                       ║
╚═══════════════════════════════════════════════════════════════════════╝
`

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func testPage(URL string, f *os.File, worker *sizedwaitgroup.SizedWaitGroup) {
	defer worker.Done()

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(checkPre + color.Yellow(" [") + color.Green(URL) + color.Yellow("] ") + http.StatusText(resp.StatusCode))
		if _, err := f.WriteString(URL + "\n"); err != nil {
			panic(err)
		}
	} else {
		fmt.Println(crossPre + color.Yellow(" [") + color.Red(URL) + color.Yellow("] ") + http.StatusText(resp.StatusCode))
	}
}

func crawl() {
	var worker = sizedwaitgroup.New(arguments.Concurrency)
	var id string
	var index int

	f, err := os.OpenFile("links.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Loop through pages
	for ; ; index++ {
		worker.Add()
		id = strconv.Itoa(index)
		go testPage("https://vimeo.com/"+id, f, &worker)
	}
}

func main() {
	// Parse arguments
	parseArgs(os.Args)

	fmt.Println(asciiArt)
	time.Sleep(5 * time.Second)

	crawl()
}
