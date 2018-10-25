package dir

import (
	"io/ioutil"
)

func List(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, file.Name()+"/")
		} else {
			paths = append(paths, file.Name())
		}
	}

	return paths, nil
}
