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

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
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
			panic(err)
		}

		data := make(map[string]interface{})
		err = json.Unmarshal(line, &data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v %v\n", data["time"], data["msg"])
	}
}
