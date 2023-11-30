package file

import (
	"log"
	"os"
	"strings"
)

func ReadFile(path string) (map[int][]byte) {
	fileData := make(map[int][]byte)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		fileData[i] = []byte(line)
	}
	return fileData
}
