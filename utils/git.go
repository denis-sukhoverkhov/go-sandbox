package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Branch struct {
	datetime time.Time
	author   string
	ref      string
}

func main() {
	pathPtr := flag.String("path", "", "path to git-repo")
	flag.Parse()

	if *pathPtr == "" {
		color.Red("ERROR: empty path to git repo!!\n")
		return
	}

	color.Green("Path to repo: %s\n", *pathPtr)
	err := os.Chdir(*pathPtr)
	if err != nil {
		panic(err)
	}

	out, err := exec.Command("bash", "-c", "git branch -r | wc -l").Output()
	if err != nil {
		log.Fatal(err)
	}
	color.Green(fmt.Sprintf("Count remote branches: %s\n", out))

	cmdGetAllRemoteBranches :=
		"git for-each-ref --format='%(committerdate) %09 %(authorname) %09 %(refname)' --sort=committerdate --merged"

	out, err = exec.Command("bash", "-c", cmdGetAllRemoteBranches).Output()
	if err != nil {
		log.Fatal(err)
	}

	var output []Branch
	var buffer bytes.Buffer
	for _, s := range out {
		buffer.WriteByte(s)

		if s == uint8('\n') {
			branch := parseRow(buffer.String())

			if branch.author == "Суховерхов Денис Владиславович" {
				now := time.Now()
				delta := now.Sub(branch.datetime)

				// 2 месяца
				treshold := time.Duration(1000) * time.Millisecond * 60 * 60 * 24 * 60

				if delta > treshold {
					output = append(output, branch)
				}
			}

			buffer.Reset()
		}
	}

	err = writeLines(output, "output2.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range output {
		out, err = exec.Command("bash", "-c", "git push -d origin "+i.ref).Output()
		if err != nil {
			log.Fatal(err)
		}

		branch_parts := strings.Split(i.ref, "/")
		branch_name := branch_parts[len(branch_parts)-1]
		out, err = exec.Command("bash", "-c", "git branch -d "+branch_name).Output()
		if err != nil {
			log.Fatal(err)
		}

	}

}

func parseRow(row string) Branch {
	parts := strings.Split(row, "\t")
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])
	parts[2] = strings.TrimSpace(parts[2])

	const rfc2822 = "Mon Jan 2 15:04:05 2006 -0700"
	datetime, err := time.Parse(rfc2822, parts[0])
	if err != nil {
		log.Fatal(err)
	}
	return Branch{
		datetime: datetime,
		author:   parts[1],
		ref:      parts[2],
	}
}

// writeLines writes the lines to the given file.
func writeLines(lines []Branch, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line.ref)
	}
	return w.Flush()
}
