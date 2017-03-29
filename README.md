# elementScraper
Scrape text between html styled tags

````
input := `bla bla <div class="hoogy"> hoogy </div> Noops <div> no Tile for this one </div>  <div class="hoogy"> dispensed </div> more bla`
scrptr := NewElementScraper(input, `<div class="hoogy">`).ElementFunc(TitleScraperImpl)
fmt.Println(scrptr.Run())

for k, v := range scrptr.Mappers() {
	fmt.Println(k, v)
}

Closing tag must not have any whiteSpace or tabs Eg: </script> is correct but, </script > is incorrect
Incorrect closing tags will be skipped
````
