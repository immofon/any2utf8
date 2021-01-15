package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Must(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(3)
	}
}

func Any2UTF8(data []byte) []byte {
	res, err := chardet.NewTextDetector().DetectBest(data)
	if err != nil {
		return data
	}

	switch res.Charset {
	case "GB-18030":
		reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())
		d, err := ioutil.ReadAll(reader)
		if err != nil {
			return data
		}
		return d
	case "UTF-8", "ISO-8859-1":
		return data
	default:
		return data
	}
}
