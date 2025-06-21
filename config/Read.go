package config

import (
	"io"
	"log"

	"github.com/ledongthuc/pdf"
)
func ReadPdf(name string) (io.Reader, error) {
  file, leitura, err := pdf.Open(name)
  if err != nil{
    log.Fatalln("erro ao abrir", err)
    return nil, err
  }
   defer file.Close()
  texto, err := leitura.GetPlainText()
  if err != nil{
    log.Fatalln("erro ao ler o pdf", err)
    return nil, err
  }
  return texto, nil 
}

