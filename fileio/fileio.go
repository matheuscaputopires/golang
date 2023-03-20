package fileio

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fileReadAndCreate(fileName string) {
	fmt.Println("Creating and writing to a file:")
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	iPrimeArr := []int{1, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		fmt.Println("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}

func fileAppend(fileName string) {
	/*
		One of these must be specified:

		O_RDONLY: open the file read-only
		O_WRONLY: open the file write-only
		O_RDWR: open the file read-write

		And these parameters can be or'ed:

		O_APPEND: append data to the filen when writing
		O_CREATE: create a new file if non existing
		O_EXCL: used with O_CREATE, file must not exist
		O_SYNC: open for synchronous I/O
		O_TRUNC: truncate regular writable file when opened
	*/
	_, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
		return
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("13\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("This is intended to go through file reading and writing")
	fileName := "data.txt"
	fileAppend(fileName)
	fileReadAndCreate(fileName)
}
