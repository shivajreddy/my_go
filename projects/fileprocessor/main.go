package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func foo() {
	start := time.Now()
	totalCount := 0
	for _, path := range getPaths("testdata") {
		count := getTotalWordsFromFile(path)
		totalCount += count
		fmt.Printf("%s : %d\n", filepath.Base(path), count)
	}
	fmt.Printf("total words = %d\n", totalCount)
	fmt.Println(time.Since(start))
	fmt.Println()
}

func bar() {
	start := time.Now()
	totalCount := 0
	var wg sync.WaitGroup
	for _, path := range getPaths("testdata") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := getTotalWordsFromFile(path)
			totalCount += count
			fmt.Printf("%s : %d\n", filepath.Base(path), count)
		}()
	}
	wg.Wait()
	fmt.Printf("total words = %d\n", totalCount)
	fmt.Println(time.Since(start))
	fmt.Println()
}

func foobar() {
	startTime := time.Now()
	totalCount := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, path := range getPaths("testdata") {
		wg.Go(func() {
			count := getTotalWordsFromFile(path)
			mu.Lock()
			totalCount += count
			mu.Unlock()
			fmt.Printf("%s : %d\n", filepath.Base(path), count)
		})
	}
	wg.Wait()
	fmt.Printf("total words = %d\n", totalCount)
	fmt.Println(time.Since(startTime))
	fmt.Println()
}

func main() {
	fmt.Println("REAL TIME?")
	// foo()
	// bar()
	foobar()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPaths(dir string) []string {
	var paths []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		check(err)
		if !info.IsDir() && strings.HasSuffix(path, ".txt") {
			paths = append(paths, path)
		}
		return nil
	})
	return paths
}

func getTotalWordsFromFile(fileName string) int {
	dat, e := os.ReadFile(fileName)
	check(e)
	return len(string(dat))
}
