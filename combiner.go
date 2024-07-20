package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func main() {
	asg := make(map[string]bool)
	wikia := make(map[string]bool)

	file, err := os.Open("asg-all.jsonl")
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
			fmt.Printf("(2) ERROR: %s\n", err)
			file.Close()
			return
		}
		id, found := info["id"]
		if !found {
			continue
		}
		asg[id.(string)] = true
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
		wikia[id.(string)] = true
	}
	file.Close()

	file, err = os.Open("bgg.jsonl")
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
			fmt.Printf("(1) ERROR: %s\n", err)
			file.Close()
			return
		}
		id, found := info["id"]
		if !found {
			continue
		}

		_, found = asg[id.(string)]
		if !found {
			_, found = wikia[id.(string)]
		}
		if found {
			filedir := fmt.Sprintf("./g/%s", id.(string))
			content := []byte(fmt.Sprintf("# %s\n\nThis page is a stub", info["name"].(string)))
			os.Mkdir(filedir, 0755)
			err := os.WriteFile(path.Join(filedir, "index.md"), content, 0644)
			//err := os.WriteFile(path.Join(filedir, "meta.cue"), ..., 0644)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	file.Close()
}
