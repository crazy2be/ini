// A simple package to read ini files (actually just name=value files, [sections] are not supported yet because we have no need for them.
package ini

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Load an ini file. Pass a filename, returns a map of all of the name=value pairs within the file, and an error if applicable.
// TODO: Read by line, rather than buffering the whole file into memory.
func Load(filename string) (map[string]string, error) {
	settings := make(map[string]string, 10)
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(fileContents, []byte("\n"))
	for _, line := range lines {
		splitLine := bytes.SplitN(line, []byte("="), 2)
		if len(splitLine) < 2 {
			//No = in line or value is blank. Skipping.
			break
		}
		settings[string(splitLine[0])] = string(splitLine[1])
	}
	return settings, nil
}

// Save a map of settings to filename in name=value format. Returns nil on success, error otherwise. Creates any necessary directories to save the file according to the path given.
func Save(filename string, settings map[string]string) (error) {
	dirname, _ := path.Split(filename)
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}

	var file *os.File
	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	for key, value := range settings {
		fmt.Fprintf(file, "%s=%s\n", key, value)
	}
	return nil
}
