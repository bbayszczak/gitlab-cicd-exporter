package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func BytesToFile(path string, filePrefix string, content []byte) error {
	var out bytes.Buffer
	err := json.Indent(&out, content, "", "  ")
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s/%s_%d.json", path, filePrefix, time.Now().Unix())
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	out.WriteTo(writer)
	writer.Flush()
	return nil
}
