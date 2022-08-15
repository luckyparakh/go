package main

import (
	"bufio"
	"example/app"
	"example/domain/repo"
	optiona "example/logging/option_a"
	optionb "example/logging/option_b"
	"example/services"
	"example/utils/errors"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	optiona.Log.Info("Starting app")
	optiona.Log.Debug("Debug app") // will not print as log level is info
	optiona.Log.Warn("Warn app")
	optiona.Info("Msg with tag", "step:1", "status:ok")
	optionb.Info("optionb", optionb.Feild("step", "2"))
	app.StartApp()
	conCreateRepos()
}

type repoOutput struct {
	Resp *repo.ServiceRepoResponse
	Err  errors.ApiError
	Name string
}

func getNames() []repo.ServiceRepoRequest {
	requests := make([]repo.ServiceRepoRequest, 0)
	file, err := os.Open("./utils/reponame.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("Error while reading file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		requests = append(requests, repo.ServiceRepoRequest{
			Name: scanner.Text(),
		})

	}
	return requests
}

func conCreateRepos() {
	input := make(chan repoOutput, 10)
	var wg sync.WaitGroup
	requests := getNames()
	// Can run many handler by putting it in loop
	go handleOutput(&wg, input)

	for _, request := range requests {
		wg.Add(1)
		go createRepo(input, request)
	}
	wg.Wait()
	close(input)
}

func handleOutput(wg *sync.WaitGroup, input chan repoOutput) {
	for ip := range input {
		// do something
		fmt.Println(ip.Name)
		wg.Done()
	}
}

func createRepo(output chan<- repoOutput, request repo.ServiceRepoRequest) {
	resp, err := services.RepoService.CreateRepo(request)
	output <- repoOutput{
		Resp: resp,
		Err:  err,
		Name: request.Name,
	}
}
