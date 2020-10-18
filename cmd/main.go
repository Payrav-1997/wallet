package main

import (
	// "io/ioutil"
	// "path/filepath"
	"os"
	"log"
	"fmt"
	"io"
)

func main() {
	// svc := &wallet.Service{}
	// svc.RegisterAccount("+992915224442")
	// svc.ImportFromFile("data/import.txt")

	// abs, err := filepath.Abs(".")
	// if err != nil{
	// 	log.Print(err)
	// 	return
	// }
	// log.Print(abs)

	//
	src, err := os.Open("data/export.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		if cerr := src.Close(); cerr != nil {
			log.Print(err)
			return
		}
	}()
	stats, err := src.Stat()
	if err != nil {
		log.Print(err)
		return
	}
	dst, err := os.Create("data/copy.txt")
	if err != nil{
		log.Print(err)
		return
	}
	defer func(){
		if cerr := dst.Close(); cerr != nil{
			log.Print(cerr)
		}
	}()

	written,err := io.Copy(dst,src)
	if err != nil{
		log.Print(err)
		return
	}
	if written != stats.Size(){
		log.Print(fmt.Errorf("copied size: %d, original size : %d",written,stats.Size()))
		return
	}

}

// func WriteFile(filename string, data []byte,perm os.FileMode)error{
// 	f,  err := os.OpenFile(filename,as.O_WORNLY|os.O_CREATE|os.O_TRUNC,perm)
// 	if err !=nil{
// 		return err
// 	}
// 	_,err = f.Write(data)
// 	if err1 := f.Close(); err ==nil{
// 		err=err1
// 	}
// 	return err
// }
