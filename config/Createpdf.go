package config

import (
	"log"
	"sort"
	"sync"

	"github.com/jung-kurt/gofpdf"
)

func CriarPdf(mensageiro chan Pdf, wg *sync.WaitGroup){
  var texto []Pdf
  for text := range mensageiro {
  texto = append(texto, Pdf{
    Indice: text.Indice,
    Texto: text.Texto,
  })
  }
  sort.Slice(texto, func(i int, j int) bool {
    return texto[i].Indice < texto[j].Indice
  })
  var textoordenado []string
  for _, trecho := range texto {
    textoordenado = append(textoordenado, trecho.Texto)
  }
  err := SalvarPdf(textoordenado, "saida.pdf")
  if err != nil {
    log.Fatal("erro ao criar a saida do pdf", err)
  }
	wg.Done()
}

func SalvarPdf(textoordenado []string, caminhoArquivo string) error {
    pdf := gofpdf.New("P", "mm", "A4", "") 
    pdf.SetFont("Arial", "", 12)

    pdf.AddPage()

    for _, trecho := range textoordenado {
        pdf.MultiCell(0, 10, trecho, "", "L", false)
        pdf.Ln(5) 
    }

    err := pdf.OutputFileAndClose(caminhoArquivo)
    if err != nil {
        log.Printf("Erro ao criar PDF: %v", err)
        return err
    }

    return nil
}
