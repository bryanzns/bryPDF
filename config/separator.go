package config

import (
	"io"
	"log"
	"strings"
)
func CovertingIo(r io.Reader) (string, error){
  data, err := io.ReadAll(r)
  if err != nil {
    log.Fatalln("erro ao transformar io.reader em byte: ", err)
    return "", err
  }
  return string(data), nil
}
func Separador(text io.Reader) []string {
  data, err := CovertingIo(text)
  if err != nil{
    log.Fatal(err)
  }
  dado := strings.Split(data, ".")
  return dado
}
