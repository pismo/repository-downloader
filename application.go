package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

type Repository struct {
	Id     int64  `json:id`
	Name   string `json:name`
	Url    string `json:html_url`
	SshUrl string `json:"ssh_url,omitempty"`
	Size   int64  `json:size`
}

func main() {
	url := fmt.Sprintf("https://api.github.com/orgs/pismo/repos?per_page=100&page=1")
	//url := fmt.Sprintf("https://api.github.com/orgs/pismo/repos?per_page=100&page=2")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	req.Header.Set("Authorization", "token poasdpoasdkpaskdopsakdpokas")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Server return non-200 status: %v\n", resp.Status)
	}

	defer resp.Body.Close()
	var repos []Repository

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(repos))

	for _, repo := range repos {
		go getrepo(repo, wg)
	}
	wg.Wait()
	log.Println("Done :D")
}

func getrepo(repo Repository, wg sync.WaitGroup) {
	defer wg.Done()
	_, err := exec.Command("git", "clone", repo.SshUrl, "~/pismo/"+repo.Name).Output()
	if err != nil {
		log.Println("Unable to get " + repo.Name)
		log.Fatal(err)
	}
	log.Println("Done " + repo.Name)
}
