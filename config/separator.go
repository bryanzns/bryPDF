package config

import (
	"io"
	"log"
	"strings"
)
func CovertingIo(r io.Reader) (string, error){
  date , err := io.ReadAll(r)
  if err != nil {
    log.Fatalln("erro ao transformar io.reader em byte: ", err)
    return "", err
  }
  return string(date), nil
}
func Separador(text io.Reader) []string {
  date, err := CovertingIo(text)
  if err != nil{
    log.Fatal(err)
  }
  data := strings.Split(data, ".")
  return data
}
