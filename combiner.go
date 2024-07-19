package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	exists := make(map[string]bool)
	multix := make(map[string]bool)

	file, err := os.Open("bgg.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var info map[string]any
		err := json.Unmarshal([]byte(scanner.Text()), &info)
		if err != nil {
			fmt.Printf("(1) ERROR: %s\n", err)
			file.Close()
			return
		}
		id, found := info["id"]
		if !found {
			continue
		}
		exists[id.(string)] = true
	}
	file.Close()

	file, err = os.Open("asg-all.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var info map[string]any
		err := json.Unmarshal([]byte(scanner.Text()), &info)
		if err != nil {
			fmt.Printf("(2) ERROR: %s\n", err)
			file.Close()
			return
		}
		id, found := info["id"]
		if !found {
			continue
		}
		if exists[id.(string)] {
			multix[id.(string)] = true
		} else {
			exists[id.(string)] = true
		}
	}
	file.Close()

	file, err = os.Open("wikia.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var info map[string]any
		text := []byte(scanner.Text())
		err := json.Unmarshal(text, &info)
		if err != nil {
			fmt.Printf("(3) ERROR: %s\n", err)
			fmt.Println(string(text))
			file.Close()
			return
		}
		id, found := info["id"]
		if !found {
			continue
		}
		if exists[id.(string)] {
			multix[id.(string)] = true
		} else {
			exists[id.(string)] = true
		}
	}
	file.Close()

	fmt.Printf("found %d unique ids and %d appearing in more than one collection", len(exists), len(multix))
}
