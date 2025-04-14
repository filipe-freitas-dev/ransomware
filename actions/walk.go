package actions

import (
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
)

func Walk(rootPath string, flag string) error {
	files := make(chan string, 10)

	var wg sync.WaitGroup

	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range files {
				if strings.ToLower(flag)[0] == 'e' {
					Encrypt(file)
				} else {
					Decrypt(file)
				}
			}
		}()
	}

	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files <- path
		}
		return nil
	})

	close(files)
	wg.Wait()

	return err
}
