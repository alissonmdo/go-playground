package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const TIMES_TO_MONITOR_WEBwebsites = 3
const DELAY_IN_MILLISECONDS = 5

func main() {
	showWelcome()
	for {
		showOptions()
		selectOption()
	}
}

func selectOption() {
	switch readCommand() {
	case 1:
		monitor()
	case 2:
		printLogs()
	case 0:
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid command!")
		os.Exit(-1)
	}
}

func monitor() {
	fmt.Println("Monitoring...")
	websites := readFile()

	for i := 0; i < TIMES_TO_MONITOR_WEBwebsites; i++ {
		for _, site := range websites {
			testSite(site)
		}
		time.Sleep(DELAY_IN_MILLISECONDS * time.Second)
		fmt.Println()
	}
}

func testSite(site string) {
	fmt.Println("🚀 Request sent to", site)
	response, err := http.Get(site)
	if err != nil {
		os.Exit(-1)
	}
	if response.StatusCode == 200 {
		fmt.Println("🟢 Successfully loaded")
		registerLog(site, true)
	} else {
		fmt.Println("🔴 Page failed to laod with message", response.Status)
		registerLog(site, false)
	}
}

func readCommand() int {
	command := -1
	fmt.Scan(&command)
	return command
}

func showWelcome() {
	clearTerminal()
	person := "Alisson"
	fmt.Println("Hello", person)
}

func showOptions() {
	fmt.Println("1. Start Monitoring")
	fmt.Println("2. Show Logs")
	fmt.Println("0. Exit")
}

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func readFile() []string {
	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("An error ocurred:", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	websites := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		websites = append(websites, string(line))
	}

	file.Close()
	return websites
}

func registerLog(site string, isSuccess bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error ocurred:", err)
	}

	file.WriteString(
		time.Now().Format("02/01/2006 15:04:05") +
			" | " + site + " | online: " +
			strconv.FormatBool(isSuccess) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println(string(file))
}
