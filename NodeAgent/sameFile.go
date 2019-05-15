package main

import "os"

func sameFile(file1, file2 os.FileInfo) bool {
	return os.SameFile(file1, file2)
}
