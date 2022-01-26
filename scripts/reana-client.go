// Testing `reana-client list` equivalent in Go.

package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	// read environment variables
	reanaServerURL := os.Getenv("REANA_SERVER_URL")
	if reanaServerURL == "" {
		fmt.Println("Please set REANA_SERVER_URL environment variable.")
		os.Exit(1)
	}

	reanaAccessToken := os.Getenv("REANA_ACCESS_TOKEN")
	if reanaAccessToken == "" {
		fmt.Println("Please set REANA_ACCESS_TOKEN environment variable.")
		os.Exit(1)
	}

	// read CLI arguments
	command := ""
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: reana-client <command>")
		fmt.Println("Example: reana-client list")
		os.Exit(2)
	} else {
		command = args[0]
	}

	// reana-client list
	if command == "list" {

		// disable certificate security checks
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}

		// make API query
		resp, err := http.Get(
			reanaServerURL + "/api/workflows?type=workflow&access_token=" + reanaAccessToken,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}

		// define response structure
		type rList struct {
			HasNext bool `json:"has_next"`
			HasPrev bool `json:"has_prev"`
			Items   []struct {
				Created  string `json:"created"`
				ID       string `json:"id"`
				Name     string `json:"name"`
				Progress struct {
					RunFinishedAt string `json:"run_finished_at"`
					RunStartedAt  string `json:"run_started_at"`
				} `json:"progress"`
				Size struct {
					HumanReadable string `json:"human_readable"`
					Raw           int    `json:"raw"`
				} `json:"size"`
				Status string `json:"status"`
				User   string `json:"user"`
			} `json:"items"`
			Page             int  `json:"page"`
			Total            int  `json:"total"`
			UserHasWorkflows bool `json:"user_has_workflows"`
		}

		// parse response
		p := rList{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			panic(err)
		}

		// format output
		fmt.Printf(
			"%-38s %-12s %-21s %-21s %-21s %-8s\n",
			"NAME",
			"RUN_NUMBER",
			"CREATED",
			"STARTED",
			"ENDED",
			"STATUS",
		)
		for _, workflow := range p.Items {
			workflowNameAndRunnumber := strings.SplitN(workflow.Name, ".", 2)
			fmt.Printf(
				"%-38s %-12s %-21s %-21s %-21s %-8s\n",
				workflowNameAndRunnumber[0],
				workflowNameAndRunnumber[1],
				workflow.Created,
				workflow.Progress.RunStartedAt,
				workflow.Progress.RunFinishedAt,
				workflow.Status,
			)
		}

	} else {
		fmt.Println("ERROR: Unknown command", command)
		os.Exit(2)
	}

}
