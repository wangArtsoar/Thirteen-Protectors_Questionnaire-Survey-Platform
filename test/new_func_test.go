package main

//func Test__1(t *testing.T) {
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
//	// Aggregate the results.
//	var totalSize int64 = 0
//	for size := range fileSizes {
//		totalSize += size
//	}
//
//	// Print the total size.
//	fmt.Printf("Total size: %d\n", totalSize)
//
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

//func Test__2(t *testing.T) {
//	// Determine the initial directories.
//	flag.Parse()
//	roots := flag.Args()
//	if len(roots) == 0 {
//		roots = []string{"."}
//	}
//
//	// Define the channel for storing file sizes
//	fileSizes := make(chan int64)
//
//	// Spawn a goroutine to traverse the file tree
//	go func() {
//		for _, root := range roots {
//			walkDir(root, fileSizes)
//		}
//		close(fileSizes)
//	}()
//
//	// Print the results
//	for size := range fileSizes {
//		fmt.Printf("Size of %s: %d\n",
//			filepath.Base(roots[size]),
//			size)
//	}
//}
//
//func walkDir(root string, fileSizes chan<- int64) {
//	err := filepath.Walk(root, func(filePath string, info os.FileInfo, err error) error {
//		// If the file is a directory, continue
//		if info.IsDir() {
//			return nil
//		}
//
//		// Read the file
//		fileContent, err := ioutil.ReadFile(filePath)
//		if err != nil {
//			return err
//		}
//
//		// Calculate the file size and store it in the channel
//		fileSizes <- int64(len(fileContent))
//
//		return nil
//	})
//
//	if err != nil {
//		fmt.Println(err)
//	}
//}
