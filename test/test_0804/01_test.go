package test_0804

import (
	"fmt"
	"strconv"
	"testing"
)

//func Test01(t *testing.T) {
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
//	var nfiles, nbytes int64
//	for size := range fileSizes {
//		nfiles++
//		nbytes += size
//	}
//	printDiskUsage(nfiles, nbytes)
//}
//
//func printDiskUsage(nfiles, nbytes int64) {
//	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
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
//		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
//		return nil
//	}
//	return entries
//}

//func Test01FirstAnswer(t *testing.T) {
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
//	// Print the results.
//	for result := range fileSizes {
//		fmt.Printf("Size of %s: %d bytes\n", result.Name(), result.Size())
//	}
//}
//
//func walkDir(dirPath string, fileSizes chan int64) {
//	// Check for the directory entry of the file.
//	entry, err := os.Stat(dirPath)
//	if err != nil {
//		return
//	}
//
//	// Create a map key for the directory entry.
//	key := filepath.Join(dirPath, entry.Name())
//
//	totalSize := int64(0)
//	if entry.IsDir() {
//		for _, filePath := range Walk(dirPath, entry) {
//			fileSize, err := ioutil.Stat(filePath)
//			totalSize += fileSize.Size()
//			fileSizes <- NewEntry(key, fileSize.Size(), fileSize.Mode())
//		}
//	} else {
//		totalSize = fileSizes[key].totalSize + fileSizes[key].usedSize
//		fileSizes <- NewEntry(key, totalSize, entry.Mode())
//	}
//
//	for {
//		size, ok := <-fileSizes
//		if !ok {
//			break
//		}
//		key := size.key
//		totalSize := size.totalSize + size.usedSize
//		fileSizes <- NewEntry(key, totalSize, size.Mode())
//	}
//}
//
//// Package `filepath` provides many file- and directory-related functions.
//// Walk returns a list of Pathnames for the files and directories at or below the given path.
//func Walk(startPath string, info os.FileInfo, walkFn filepath.WalkFunc) ([]Pathname, error) {
//	var paths []Pathname
//
//	filepath.WalkDir(info, walkFn, startPath, func(path string, d os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if entry := filepath.Base(info.Name()); startPath != path {
//			filePath := path
//			if entry != "" && filepath.Base(entry) == "." {
//				filePath = filepath.Join(path, entry)
//			}
//			paths = append(paths, Pathname{filePath})
//		}
//		return nil
//	})
//
//	return paths, nil
//}
//
//// Package `fmt` provides types and functions for formatting data and objects as outputs (e. g. printing to the console).
//// NewEntry is a result type to keep track of the used and total size for a file.
//type NewEntry struct {
//	key   string        // The directory name for which we've collected file sizes.
//	totalSize int64        // The total size.
//	usedSize  int64        // The estimated size.
//	Mode     os.FileMode  // The file mode.
//}
//

//func Test01SecondAnswer(t *testing.T) {
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
//	for fileSize := range fileSizes {
//		fmt.Printf("文件 的大小是: %d bytes\n", fileSize)
//	}
//}
//
//func walkDir(root string, fileSizes chan int64) {
//	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if err != nil || info.IsDir() {
//			return err
//		}
//		fi, err := os.Open(path)
//		if err != nil {
//			return err
//		}
//		defer fi.Close()
//		stat, err := fi.Stat()
//		if err != nil {
//			return err
//		}
//		size := stat.Size()
//		fileSizes <- size
//		return nil
//	})
//	if err != nil {
//		fmt.Println("遍历目录时发生了错误:", err)
//	}
//}

//func Test01ThirdAnswer(t *testing.T) {
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
//		fmt.Printf("The size of  is %d bytes\n", size)
//	}
//}
//
//// walkDir recursively walks the directory tree rooted at root,
//// calling walkFunc for each file and directory in the tree.
//func walkDir(root string, fileSizes chan int64) {
//	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if err != nil || info.IsDir() {
//			return err
//		}
//		size, err := ioutil.ReadFile(path)
//		if err != nil {
//			return err
//		}
//		fileSizes <- int64(len(size))
//		return nil
//	})
//}

//func Test01FourthAnswer(t *testing.T) {
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
//		fmt.Printf("The size of %s is %d bytes\n", filepath.Base(size.Dir), size.Size)
//	}
//}
//
//// Walk the directory and calculate the file size.
//func walkDir(root string, fileSizes chan int64) {
//	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if err != nil || info.IsDir() {
//			return err
//		}
//		size, err := strconv.ParseInt(info.Size(), 10, 64)
//		if err != nil {
//			return err
//		}
//		fileSizes <- size
//		return nil
//	})
//	if err != nil {
//		fmt.Println("Error walking the directory:", err)
//		return
//	}
//}

//func Test01FifthAnswer(t *testing.T) {
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
//		fmt.Printf("The size of %s is %d bytes\n", filepath.Base(size.Dir), size.Size)
//	}
//}
//
//// Walk the directory and collect the file sizes.
//func walkDir(root string, fileSizes chan int64) {
//	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if err != nil || info.IsDir() {
//			return err
//		}
//		size, err := getFileSize(path)
//		if err != nil {
//			return err
//		}
//		fileSizes <- size
//		return nil
//	})
//	if err != nil {
//		fmt.Println("Error walking the directory:", err)
//	}
//}
//
//// Get the file size of a path.
//func getFileSize(path string) (int64, error) {
//	file, err := os.Open(path)
//	if err != nil {
//		return 0, err
//	}
//	defer file.Close()
//
//	stat, err := file.Stat()
//	if err != nil {
//		return 0, err
//	}
//
//	return stat.Size(), nil
//}

func TestJava(t *testing.T) {
	num := coins(100)
	fmt.Println(num) // 输出 27
}

func coins(n int) int {
	s := make(map[string]string)
	for i := 0; i <= 100/7; i++ {
		for j := 0; j <= 100/3; j++ {
			for k := 50; k >= 0; k-- {
				if i*7+j*3+k*2 == n {
					str := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
					s[str] = str
				}
			}
		}
	}
	fmt.Println(s["1401"])
	return len(s)
}
