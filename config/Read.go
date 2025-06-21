package config

import (
	"io"
	"log"

	"github.com/ledongthuc/pdf"
)
func ReadPdf(name string) (io.Reader, error) {
  file, reading , err := pdf.Open(name)
  if err != nil{
    log.Fatalln("error opening ", err)
    return nil, err
  }
   defer file.Close()
  text, err := reading.GetPlainText()
  if err != nil{
    log.Fatalln("error reading pdf ", err)
    return nil, err
  }
  return text, nil 
}

