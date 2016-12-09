package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	repo := repository{path: "."}
	mainWithArgs(&repo, os.Args[1:]...)
}

func mainWithArgs(repo *repository, args ...string) {

	octopusConfig := getOctopusConfig(repo, args)

	if octopusConfig.printVersion {
		fmt.Println("2.0")
		return
	}

	if len(octopusConfig.patterns) == 0 {
		fmt.Println("Nothing to merge. No pattern given")
		return
	}

	branchList := resolveBranchList(repo, octopusConfig.patterns, octopusConfig.excludedPatterns)

	if len(branchList) == 0 {
		fmt.Printf("No branch matching \"%v\" were found\n", strings.Join(octopusConfig.patterns, " "))
		return
	}
}

func resolveBranchList(repo *repository, patterns []string, excludedPatterns []string) map[string]string {
	result := parseLsRemote(repo.git(append([]string{"ls-remote", "."}, patterns...)...))

	if len(excludedPatterns) == 0 {
		return result
	}

	excludedRefs := parseLsRemote(repo.git(append([]string{"ls-remote", "."}, excludedPatterns...)...))
	for excludedRef, _ := range excludedRefs {
		delete(result, excludedRef)
	}

	return result
}
