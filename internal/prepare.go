package internal

import "os"

func CheckRequiredDirectories(dirs ...string) error {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
