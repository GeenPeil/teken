package storage

import "fmt"

// fileFolderByNumber returns the filename and foldername by ID
func fileFolderByNumber(n int) (string, string) {
	return fmt.Sprintf("%d.gp", n), fmt.Sprintf("%d", n/1000+1)
}
