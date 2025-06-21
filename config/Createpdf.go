package config

import (
	"log"
	"sort"
	"sync"

	"github.com/jung-kurt/gofpdf"
)

func CreatePdf(messenger chan Pdf, wg *sync.WaitGroup){
  var texto []Pdf
  for text := range messenger {
  text = append(text, Pdf{
    Index : text.Indice,
    Text : text.Texto,
  })
  }
  sort.Slice(text, func(i int, j int) bool {
    return text[i].Indice < text[j].Indice
  })
  var orderedtext []string
  for _, excerpt  := range text {
   orderedtext  = append(orderedtext, excerpt.Text)
  }
  err := SavePdf(orderedtext, "exit.pdf")
  if err != nil {
    log.Fatal("error creating pdf output ", err)
  }
	wg.Done()
}

func SalvarPdf(orderedtext []string, pathFile  string) error {
    pdf := gofpdf.New("P", "mm", "A4", "") 
    pdf.SetFont("Arial", "", 12)

    pdf.AddPage()

    for _, excerpt := range orderedtext {
        pdf.MultiCell(0, 10, excerpt, "", "L", false)
        pdf.Ln(5) 
    }

    err := pdf.OutputFileAndClose(pathFile)
    if err != nil {
        log.Printf("Erro ao criar PDF: %v", err)
        return err
    }

    return nil
}
