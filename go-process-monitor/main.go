package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type process struct {
	id   int
	tty  string
	time string
	cmd  string
}

func (p process) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", p.id, p.tty, p.time, p.cmd)
}

func parseProcess(fiends []string) process {
	if len(fiends) < 4 {
		return process{}
	}
	id, err := strconv.Atoi(fiends[0])
	if err != nil {
		id = -1 // Because 0 is the root process
	}
	return process{
		id:   id,
		tty:  fiends[1],
		time: fiends[2],
		cmd:  fiends[3],
	}
}

func main() {
	// Parse flags
	pid := flag.Int("p", -1, "filter by process ID")
	pname := flag.String("n", "", "filter by process name")
	flag.Parse()

	for {
		// Run a command to capture the output
		output, err := exec.Command("ps").Output()
		if err != nil {
			log.Fatal(err)
		}

		// Parse output into a struct
		processes := []process{}
		scanner := bufio.NewScanner(bytes.NewReader(output))
		scanner.Scan() // Skips header
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			process := parseProcess(fields)
			processes = append(processes, process)
		}

		// Filter processes
		filtered := make([]process, 0, len(processes))
		if *pid > 0 || len(*pname) > 0 {
			for _, p := range processes {
				if *pid == p.id || len(*pname) > 0 && strings.Contains(p.cmd, *pname) {
					filtered = append(filtered, p)
				}
			}
		} else {
			filtered = append(filtered, processes...)
		}

		// Print filtered processes
		for _, p := range filtered {
			fmt.Println(p)
		}

		// Sleep for one second
		time.Sleep(1 * time.Second)
	}
}
