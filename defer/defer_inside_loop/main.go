package main

import "os"

// Стоит помнить, что defer не отрабатывает
// внутри области видимости, а после того,
// как произойдет выход из функции.

func readFiles(paths []string) error {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		// reading file...
		defer file.Close()
	}

	return nil
}

func main() {
	_ = readFiles([]string{"text1.txt", "text2.txt", "text3.txt"})
}
