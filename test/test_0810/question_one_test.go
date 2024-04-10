package test_0810

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//func TestTheOneFromAnswerByFirst(t *testing.T) {
//	// Determine the initial directories.
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//
//	// Traverse the file tree.
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results.
//	var totalSize int64
//	for size := range fileSizes {
//		totalSize += size
//		fmt.Println(size)
//	}
//	fmt.Printf("Total size of files in the tree: %d\n", totalSize)
//}
//
//func walkDir(dir string, fileSizes chan int64) {
//	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			fmt.Printf("Error accessing directory %s: %v\n", dir, err)
//			return err
//		}
//		if info.IsDir() {
//			return nil
//		} else {
//			fmt.Println(info.Name())
//			size := info.Size()
//			fileSizes <- size
//		}
//		return nil
//	})
//}

// 21
// 21
// 20
// 22
// Determine the initial directories.
//func TestTheOneFromAnswerBySecond(t *testing.T) {
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//
//	// Traverse the file tree.
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results.
//	for size := range fileSizes {
//		fmt.Println(size)
//	}
//}

// walkDir recursively traverses the directory tree rooted at the passed directory,
// and populates a slice of file sizes.
//func walkDir(dir string, fileSizes chan<- int64) {
//	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if !info.IsDir() {
//			fileSizes <- int64(len(path))
//		}
//		return nil
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

//// 579
//// 869
//// 2977
//// 852
//// Total size: 5277
//func TestTheOneFromAnswerByClaude(t *testing.T) {
//
//	// Determine the initial directories.
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//
//	// Traverse the file tree.
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results.
//	var total int64 = 0
//	for size := range fileSizes {
//		total += size
//		fmt.Println(size)
//	}
//	fmt.Printf("Total size: %d\n", total)
//}
//
//func walkDir(dir string, fileSizes chan<- int64) {
//	for _, entry := range dirents(dir) {
//		if entry.IsDir() {
//			subdir := filepath.Join(dir, entry.Name())
//			walkDir(subdir, fileSizes)
//		} else {
//			fileSizes <- entry.Size()
//		}
//	}
//}
//
//func dirents(dir string) []os.FileInfo {
//	entries, err := ioutil.ReadDir(dir)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "du: %v\n", err)
//		return nil
//	}
//	return entries
//}

//func TestTheOneFromAnswerByNewBing(t *testing.T) {
//	// Determine the initial directories.
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//	// Traverse the file tree.
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results.
//	var total int64 = 0
//	for size := range fileSizes {
//		total += size
//		fmt.Println(size)
//	}
//	fmt.Printf("Total size: %d\n", total)
//}
//
//func walkDir(dir string, fileSizes chan<- int64) {
//	for _, entry := range dirents(dir) {
//		if entry.IsDir() {
//			subdir := filepath.Join(dir, entry.Name())
//			walkDir(subdir, fileSizes)
//		} else {
//			info, err := entry.Info()
//			if err != nil {
//				return
//			}
//			fmt.Println(info.Name())
//			fileSizes <- info.Size()
//		}
//	}
//}
//
//func dirents(dir string) []os.DirEntry {
//	entries, err := os.ReadDir(dir)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "du: %v\n", err)
//		return nil
//	}
//	return entries
//}

//func TestAnswerBy0818First(t *testing.T) {
//	// Determine the initial directories.
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//
//	// Traverse the file tree.
//	fileSizes := make(chan int64)
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results.
//	for size := range fileSizes {
//		fmt.Printf(" %d bytes\n", size)
//	}
//}
//
//func walkDir(dir string, fileSizes chan int64) {
//	files, err := ioutil.ReadDir(dir)
//	if err != nil {
//		fmt.Printf("Failed to read directory: %v\n", err)
//		return
//	}
//
//	for _, file := range files {
//		filePath := filepath.Join(dir, file.Name())
//		fileSize, err := ioutil.ReadFile(filePath)
//		if err != nil {
//			fmt.Printf("Failed to read file: %v\n", err)
//			return
//		}
//
//		fileSizes <- int64(len(fileSize))
//	}
//}

func TestAnswerBy0818Second(t *testing.T) {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	for size := range fileSizes {
		fmt.Printf("Size of %d bytes in %s\n", size, roots[len(roots)-1])
	}

	// Clean up.
	for _, root := range roots {
		os.RemoveAll(root)
	}

	fmt.Println("All files and directories have been traversed and their sizes have been printed.")
}

// Helper function to traverse a directory recursively.
func walkDir(dirPath string, fileSizes chan int64) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories.
		if info.IsDir() && info.Name()[0] == '.' {
			return filepath.SkipDir
		}

		// Get the file size.
		fileSize, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		//fileSize = int64(fileSize)

		// Add the file size to the channel.
		fileSizes <- int64(len(fileSize))

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %s\n", err)
	}
}

// Helper function to print the file sizes.
func printFileSizes(sizes chan int64) {
	for size := range sizes {
		fmt.Printf("Size of %d bytes\n", size)
	}
}
