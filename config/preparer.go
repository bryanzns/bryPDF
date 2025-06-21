package config

import (
	"fmt"
	"strings"
	"sync"
)

type Pdf struct {
  Index  int `json:"indice"`
  Text string `json:"texto"`
}
const characterlimit  = 3000 
func messenger(excerpts  []string, receiver  chan Pdf, APIkey string, wg *sync.WaitGroup) {
	var finaltext  []string
  var message  []string
  var accumulatedtext   int
  var counter int
    for _ , excerpt  := range excerps {
		cleansection  := strings.ReplaceAll(trecho, "\n", "  ")
		cleansection = strings.ReplaceAll(trechoLimpo, "\r", "")
    cleansection = strings.ReplaceAll(trechoLimpo, "\t", " ")
      if  accumulatedtext + len(cleansection) > characterlimit {
      counter += 1
      result := strings.Join(mensagem," ")
			finaltext = append(finaltext, result)
			message = []string{}
      accumulatedtext = 0
      }
    message = append(message, cleansection)
    textoacumulado += len(cleansection)
  } 
if len(message) > 0 {
    counter++
    result := strings.Join(mensagem," ")
    finaltext = append(finaltext, result)
  }
	wg.Add(1)
	go CallApi(receiver ,finaltext ,APIkey,wg)
}
