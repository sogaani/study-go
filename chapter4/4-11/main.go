// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const IssuesURLFmt string = "https://api.github.com/repos/%s/%s/issues"
const IssueURLFmt string = "https://api.github.com/repos/%s/%s/issues/%s"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type CreIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type EdiIssue struct {
	//Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

func readIssue(user string, repo string) error {
	url := fmt.Sprintf(IssuesURLFmt, user, repo)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("read issue failed: %s", resp.Status)
	}
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Printf("issues:\n")
	for _, item := range result {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.State, item.Title)
	}
	return nil
}

func editIssue(user string, repo string, number string, password string, item EdiIssue) error {
	url := fmt.Sprintf(IssueURLFmt, user, repo, number)
	input, err := json.Marshal(item)

	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, password)
	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("edit issue failed: %s", resp.Status)
	}

	return nil
}

func createIssue(user string, repo string, password string, item CreIssue) error {
	url := fmt.Sprintf(IssuesURLFmt, user, repo)
	input, err := json.Marshal(item)

	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, password)
	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}

	return nil
}

func getPassword(user string, password *string) {
	if *password == "" {
		fmt.Fprintf(os.Stderr, "enter %s Password:", user)
		fmt.Scan(password)
	}
}

func getComment(comment *string) error {
	const text = "temporary_comment.txt"

	if *comment != "" {
		return nil
	}
	f, err := os.Create(text)
	if err != nil {
		return err
	}
	defer os.Remove(text)
	f.Close()

	switch runtime.GOOS {
	case "linux":
		{
			exec.Command(os.Getenv("EDITOR"), f.Name()).Run()
		}
	case "windows":
		{
			exec.Command("cmd", "/c", f.Name()).Run()
		}
	}
	f, err = os.Open(text)

	if err != nil {
		return err
	}
	defer f.Close()
	var m = bytes.NewBuffer(make([]byte, 0, 256))
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m.WriteString(sc.Text())
		m.WriteString("\n")
	}
	*comment = m.String()
	return nil
}

//!+
func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [COMMANDS] owner repo ARGS... [OPTIONS]
COMMANDS
	create owner repo title
	read owner repo
	update owner repo number
	close owner repo number
OPTIONS
`, os.Args[0], os.Args[0])
		fs.PrintDefaults()
	}
	var (
		password = fs.String("p", "", "oners password")
		comment  = fs.String("m", "", "issue comment")
	)

	cmd := os.Args[1]
	user := os.Args[2]
	repo := os.Args[3]

	if len(os.Args) > 4 {
		fs.Parse(os.Args[5:])
	}

	switch cmd {
	case "create":
		{
			getPassword(user, password)
			getComment(comment)
			item := CreIssue{Title: os.Args[4], Body: *comment}
			err := createIssue(user, repo, *password, item)
			if err != nil {
				fmt.Print(err)
			}
		}
	case "read":
		{
			readIssue(user, repo)
		}
	case "update":
		{
			number := os.Args[4]
			getPassword(user, password)
			getComment(comment)
			item := EdiIssue{Body: *comment}
			err := editIssue(user, repo, number, *password, item)
			if err != nil {
				fmt.Print(err)
			}
		}
	case "close":
		{
			number := os.Args[4]
			getPassword(user, password)
			getComment(comment)
			item := EdiIssue{Body: *comment, State: "closed"}
			err := editIssue(user, repo, number, *password, item)
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}

//!-
