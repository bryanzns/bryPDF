package main

import (
	"fmt"
	"sync"

	"github.com/brypdf/config"
)

func main() {
  var wg sync.WaitGroup
  receivingchannel  := make(chan config.Pdf)
    yellow := "\033[1;33m" 
    reset := "\033[0m"

    fmt.Println(yellow + `
▗▄▄▖ ▗▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄ ▗▄▄▄▖
▐▌ ▐▌▐▌ ▐▌▝▚▞▘ ▐▌ ▐▌▐▌  █▐▌   
▐▛▀▚▖▐▛▀▚▖ ▐▌  ▐▛▀▘ ▐▌  █▐▛▀▀▘
▐▙▄▞▘▐▌ ▐▌ ▐▌  ▐▌   ▐▙▄▄▀▐▌   
                              
---------------------------------- 
                              
` + reset)
	// put the path here
	file, err := config.ReadPdf("..") //   example:  /home/..../Downloads/DOC-20250614-WA0027..pdf"
  if err != nil {
    fmt.Println("erro ao ler pdf")
    return
  } 
  Separatetext  := config.separator(file)
  wg.Add(1)
  go config.CreatePdf(receivingchannel, &wg)
  APIkey  := ""  //put the key here
  config.Messenger(Separatetext,receivingchannel,APIkey, &wg)
 wg.Wait()
}
