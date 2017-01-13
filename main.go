package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	InputName  string
	OutputName string
	enc        = simplifiedchinese.GBK
	Isgbk      bool
)

func main() {
	iname := flag.String("i", "", "input file name")
	oname := flag.String("o", "", "output file name")
	charset := flag.String("c", "", "charset")
	flag.Parse()
	if len(*iname) == 0 {
		log.Panic("input miss")
	}
	InputName = *iname
	if len(*oname) == 0 {
		log.Panic("output miss")
	}
	if *charset == "gbk" {
		Isgbk = true
	}
	OutputName = *oname
	ifile, err := os.Open(InputName)
	if err != nil {
		log.Panic(err)
	}
	defer ifile.Close()
	ofile, err := os.Create(OutputName)
	if err != nil {
		log.Panic(err)
	}
	defer ofile.Close()
	var r *csv.Reader
	if Isgbk {
		r = csv.NewReader(transform.NewReader(ifile, enc.NewDecoder()))
	} else {
		r = csv.NewReader(ifile)
	}
	w := csv.NewWriter(ofile)
	defer w.Flush()
	var line []string
	icount := 0
	for {
		if line, err = r.Read(); err != nil {
			break
		}
		newline := make([]string, len(line))
		for i, v := range line {
			newline[i] = strings.Replace(strings.Replace(v, "\r", "", -1), "\n", "", -1)
		}
		w.Write(newline)
		icount++
	}
	if err != io.EOF {
		log.Panic(err)
	}
	log.Printf("success convert %d lines\n", icount)
}
