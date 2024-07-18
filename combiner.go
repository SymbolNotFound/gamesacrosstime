package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("bgg.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		var info map[string]any
		err := json.Unmarshal([]byte(scanner.Text()), &info)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}
		fmt.Printf("%v\n", info)
		count += 1
	}

	fmt.Printf("\nparsed %d games\n", count)
}
