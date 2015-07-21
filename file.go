package sshmux

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// read a list of hosts to run command on
func GetHostsByFile(file string) []string {
	var hosts []string
	fr, _ := os.Open(file)
	defer fr.Close()
	r := bufio.NewReader(fr)
	for {
		readline, _ := r.ReadString('\n')
		// ignore 'comment' lines
		if !strings.Contains(readline, "#") {
			// add host to list
			if readline != "" {
				hosts = append(hosts, strings.Trim(readline, "\n"))
			}
		}
		// check for end of file
		var err error
		if err == io.EOF {
			break
		}
	}
	return hosts
}
