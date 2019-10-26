package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough args")
	}

	err := processFile(os.Args[1], []string{"time", "request-id", "path", "msg", "workstation-id"})
	if err != nil {
		fmt.Println(err)
	}
}

func processFile(path string, fields []string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line []byte
	for {
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		data := make(map[string]interface{})
		err = json.Unmarshal(line, &data)
		if err != nil {
			return err
		}
		var result string
		for _, f := range fields {
			result += fmt.Sprintf("%v ", data[f])
		}
		fmt.Println(result)
	}

	return nil
}
