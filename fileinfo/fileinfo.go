package fileinfo

import (
	"fmt"
	"os"
)

func ReadFileinfo(filename string) {

	// FileInfo
	fi, err := os.Stat(filename)

	if err != nil {
		panic(err)
	}

	// FileMode
	fm := fi.Mode()

	fmt.Printf("Information about a file %v:\n", fi.Name())

	b := fi.Size()
	KiB := b / 1024
	MiB := KiB / 1024
	GiB := MiB / 1024

	fmt.Printf("Size: %v bytes, %v KiB, %v MiB, %v GiB\n", b, KiB, MiB, GiB)

	if fm.IsDir() {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}

	if fm.IsRegular() {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	fmt.Printf("Has Unix permission bits: %v\n", fm.Perm())

	if fm&os.ModeAppend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	if fm&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}

	if fm&os.ModeCharDevice != 0 {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not a Unix character device")
	}

	if fm&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
}
