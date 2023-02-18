package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func useage(name string) {
	println("Useage: ")
	println(name + " -A <hash algorithm> (-i|--input <input>)|(-if|--input-file <input file>) [-of|--output-file <output file>]")
}
func Sha1Byte(data []byte) string {
	s := sha1.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}
func Sha256Byte(data []byte) string {
	s := sha256.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}
func Md5Byte(data []byte) string {
	s := md5.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}
func main() {
	name := os.Args[0]
	if len(os.Args) == 0 {
		useage(name)
		return
	}
	args := os.Args[1:]
	var (
		algorithm     string = ""
		inputfile     string = ""
		hasinputfile  bool   = false
		inputs        string = ""
		hasinput      bool   = false
		outputfile    string = ""
		hasoutputfile bool   = false
		c             bool   = false
	)
	for k, v := range args {
		if c {
			c = false
			continue
		}
		if v == "-A" || v == "--algorithm" {
			algorithm = args[k+1]
			c = true
			continue
		}
		if v == "-if" || v == "--input-file" {
			inputfile = args[k+1]
			hasinputfile = true
			c = true
			continue
		}
		if v == "-i" || v == "--input" {
			inputs = args[k+1]
			hasinput = true
			c = true
			continue
		}
		if v == "-of" || v == "--output-file" {
			outputfile = args[k+1]
			hasoutputfile = true
			c = true
			continue
		}
	}
	if hasinput == hasinputfile {
		useage(name)
	}
	var inputbyte []byte
	if hasinput {
		inputbyte = []byte(inputs)
	} else if hasinputfile {
		f, err := os.OpenFile(inputfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fb, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		inputbyte = fb
	}
	switch algorithm {
	case "sha1":
		hashcode := Sha1Byte(inputbyte)
		if hasoutputfile {
			WriteString(outputfile, hashcode)
		} else {
			println(hashcode)
		}
	case "sha256":
		hashcode := Sha256Byte(inputbyte)
		if hasoutputfile {
			WriteString(outputfile, hashcode)
		} else {
			println(hashcode)
		}
	case "md5":
		hashcode := Sha1Byte(inputbyte)
		if hasoutputfile {
			WriteString(outputfile, hashcode)
		} else {
			println(hashcode)
		}
	}
}
