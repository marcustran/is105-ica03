package fileinfo

import (
	"fmt"
<<<<<<< HEAD
=======
	"log"
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
	"os"
)

func ReadFileinfo(filename string) {

<<<<<<< HEAD
	// FileInfo
	fi, err := os.Stat(filename)

	if err != nil {
		panic(err)
	}

	// FileMode
	fm := fi.Mode()

	fmt.Printf("Information about a file %v:\n", fi.Name())

=======
	//filenavn
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	file := fi.Mode()
	fmt.Printf("Info about file %v:\n", fi.Name())
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
	b := fi.Size()
	KiB := b / 1024
	MiB := KiB / 1024
	GiB := MiB / 1024

<<<<<<< HEAD
	fmt.Printf("Size: %v bytes, %v KiB, %v MiB, %v GiB\n", b, KiB, MiB, GiB)

	if fm.IsDir() {
=======
	fmt.Printf("Size: %v bytes, %v KiB, %v Mib, %v GiB\n", b, KiB, MiB, GiB)

	if file.IsDir() {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}

<<<<<<< HEAD
	if fm.IsRegular() {
=======
	if file.IsRegular() {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

<<<<<<< HEAD
	fmt.Printf("Has Unix permission bits: %v\n", fm.Perm())

	if fm&os.ModeAppend != 0 {
=======
	fmt.Printf("Has Unix permission bits: %v\n", file.Perm)

	if file&os.ModeAppend != 0 {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

<<<<<<< HEAD
	if fm&os.ModeDevice != 0 {
=======
	if file&os.ModeDevice != 0 {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}

<<<<<<< HEAD
	if fm&os.ModeCharDevice != 0 {
=======
	if file&os.ModeCharDevice != 0 {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not a Unix character device")
	}

<<<<<<< HEAD
	if fm&os.ModeSymlink != 0 {
=======
	if file&os.ModeSymlink != 0 {
>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
<<<<<<< HEAD
=======

>>>>>>> e86768995ff80744a7e8911650dff120ab8d0213
}
