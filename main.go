package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"net/http"
)

type Job struct {
	ID          int    `json:"id"`
}

func main () {
	var project_id int
	var per_page int
	var start_page int
	var current_page int
	var token string
	var server string

	flag.IntVar(&project_id, "project_id", 0, "Project ID")
	flag.IntVar(&per_page, "per_page", 100, "Number of jobs per page")
	flag.IntVar(&start_page, "start-page", 1, "Start pages")
	flag.StringVar(&token, "token", "", "Private token")
	flag.StringVar(&server, "server", "", "Gitlab server")
	flag.Parse()

	current_page = start_page

	for true {
		url := fmt.Sprintf("%v/api/v4/projects/%v/jobs?pagination=keyset&page=%d&per_page=%d&order_by=id&sort=asc&artifact_expired=false", server, project_id, current_page, per_page)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("PRIVATE-TOKEN", token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		var jobs []Job
		err = json.NewDecoder(resp.Body).Decode(&jobs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("page number: %s of %s\n", resp.Header.Get("x-page"), resp.Header.Get("x-total-pages"))

		for _, job := range jobs {
			fmt.Println("deleting artifacts of job:", job.ID)
			url := fmt.Sprintf("%v/api/v4/projects/%v/jobs/%d/artifacts", server, project_id, job.ID)
			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("PRIVATE-TOKEN", token)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
		}

		next := resp.Header.Get("x-next-page")
		next_page, err := strconv.Atoi(next)
		if err != nil || next == "" || current_page == next_page {
			println("no more pages. exiting")
			return
		}
		current_page = next_page
	}
}
