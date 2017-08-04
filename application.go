package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type Repository struct {
	Name   string `json:"name"`
	SshUrl string `json:"ssh_url,omitempty"`
}

func main() {
	url := os.Args[1]
	token := os.Args[2]
	folder := os.Args[3]

	url = strings.TrimSuffix(url, "/")
	if !strings.HasSuffix(folder, "/") {
		folder = folder + "/"
	}

	for {
		index := 0
		index++

		req, err := http.NewRequest("GET", url+"?per_page=100&page="+strconv.Itoa(index), nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
		}
		req.Header.Set("Authorization", "token "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		var repos []Repository

		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			return
		}

		var wg sync.WaitGroup
		wg.Add(len(repos))

		for _, repo := range repos {
			go getrepo(folder, repo, wg)
		}

		wg.Wait()
		return
	}
	log.Println("Done :D")
}

func getrepo(folder string, repo Repository, wg sync.WaitGroup) {
	defer wg.Done()
	_, err := exec.Command("git", "clone", repo.SshUrl, folder+repo.Name).Output()
	if err != nil {
		log.Println("Unable to get " + repo.Name)
		log.Fatal(err)
	}
	log.Println("Done " + repo.Name + "saved on " + folder + repo.Name)
}
