package config

import (
	"fmt"
	"strings"
	"sync"
)

type Pdf struct {
  Indice int `json:"indice"`
  Texto string `json:"texto"`
}
const limiteCaracteres = 3000 
func Mensageiro(trechos []string, recebidor chan Pdf, chaveapi string, wg *sync.WaitGroup) {
	var textofinal[]string
  var mensagem []string
  var textoacumulado int
  var contador int
    for _ , trecho := range trechos {
		trechoLimpo := strings.ReplaceAll(trecho, "\n", "  ")
		trechoLimpo = strings.ReplaceAll(trechoLimpo, "\r", "")
    trechoLimpo = strings.ReplaceAll(trechoLimpo, "\t", " ")
      if textoacumulado + len(trechoLimpo) > limiteCaracteres {
      contador += 1
      result := strings.Join(mensagem," ")
			textofinal = append(textofinal, result)
			mensagem = []string{}
      textoacumulado = 0
      }
    mensagem = append(mensagem, trechoLimpo)
    textoacumulado += len(trechoLimpo)
  } 
if len(mensagem) > 0 {
    contador++
    result := strings.Join(mensagem," ")
    textofinal = append(textofinal, result)
  }
	wg.Add(1)
	go CallApi(recebidor,textofinal,chaveapi,wg)
}
