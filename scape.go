package elementScraper

import (
	"fmt"
	"strings"
)

type tokpos struct {
	lpos, rpos int
}

type tokval struct {
	lpos, rpos int
	value      string
}

type toggle struct {
	a int
}

type termPos struct {
	term string
	lpos int
}

func (elmScrpr *elementScraper) find(lpos int) string {
	return elmScrpr.mapper[lpos]
}

// You provide a func satifying this type which gets called with the scraped data.
// you then manipulate that data the way you want
/*	Example
	func TitleScraperImpl(raw string) string {
	return strings.Title(raw)
}
*/
type replaceImpl func(string) string

// RawScraperImpl implements replaceImpl type
// This func receives the scraped input which is then returned as is
func RawScraperImpl(raw string) string {
	return raw
}

func (elmScrpr *elementScraper) replace() {
	term := elmScrpr.tpos.term
	if !containsEmptySpaces(term) {
		elmScrpr.mapper[elmScrpr.tpos.lpos] = elmScrpr.fnc(term)
		return
	}
	elmScrpr.mapper[elmScrpr.tpos.lpos] = ""

}

type elementScraper struct {
	element string
	tokee   *tokpos
	toggle  *toggle
	tpos    *termPos
	mapper  map[int]string
	toks    []tokval
	fnc     replaceImpl
	input   string
}

func (elmScrpr *elementScraper) openS() string {
	return elmScrpr.element[1:]
}

func (elmScrpr *elementScraper) closeS() string {
	return elmScrpr.lArrow() + "/" + till((elmScrpr.element[1:len(elmScrpr.element)]))
}

func (elmScrpr *elementScraper) lArrow() string {
	return elmScrpr.element[:1]
}

func (elmScrpr *elementScraper) rArrow() string {
	return elmScrpr.element[len(elmScrpr.element)-1:]
}

func containsEmptySpaces(str string) bool {
	var filled int
	for _, v := range str {
		if v != 9 && v != 32 {
			filled++
		}

	}
	return filled == 0
}

func tillFirstWhiteSpace(str string) (string, bool) {
	for i, v := range str {
		if v == 32 {
			return str[:i], true
		}
	}
	return str, false
}

func till(strr string) string {
	str, ok := tillFirstWhiteSpace(strr)
	if ok {
		return fmt.Sprint(str[:len(str)])
	}
	return fmt.Sprint(str[:len(str)-1])
}

func (elmScrpr *elementScraper) elementScraper(a string) string {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err)

		}
	}()
	for i, v := range a {
		switch elmScrpr.toggle.a {

		case 0:
			if string(v) == elmScrpr.lArrow() && string(a[i+1]) != "/" {
				if a[i+1:i+len(elmScrpr.openS())+1] != elmScrpr.openS() {
					continue
				}
				elmScrpr.toggle.a++
				elmScrpr.tokee.lpos = i + len(elmScrpr.openS()) + 1

			}

		case 1:

			if string(v) == elmScrpr.rArrow() {
				if a[i-len(elmScrpr.closeS()):i] != elmScrpr.closeS() {
					continue
				}
				elmScrpr.tokee.rpos = i - len(elmScrpr.closeS())
				elmScrpr.toggle.a--
				if elmScrpr.tokee.lpos < elmScrpr.tokee.rpos {
					bb := tokval{lpos: elmScrpr.tokee.lpos, rpos: elmScrpr.tokee.rpos, value: strings.TrimSpace(a[elmScrpr.tokee.lpos:elmScrpr.tokee.rpos])}
					elmScrpr.toks = append(elmScrpr.toks, bb)
					elmScrpr.tpos = &termPos{bb.value, elmScrpr.tokee.lpos}
					elmScrpr.replace()
				}

			}

		}
	}

	for i, j := 0, len(elmScrpr.toks)-1; i < j; i, j = i+1, j-1 {
		elmScrpr.toks[i], elmScrpr.toks[j] = elmScrpr.toks[j], elmScrpr.toks[i]
	}
	var d string
	for i, v := range elmScrpr.toks {
		if i == 0 {
			d = a[:v.lpos] + elmScrpr.find(v.lpos) + a[v.rpos:]
		} else {
			d = d[:v.lpos] + elmScrpr.find(v.lpos) + d[v.rpos:]
		}

	}
	return fmt.Sprint(d[:len(d)])
}

// Mapper contains matched 'scraped items' with key as the starting position of that element
// See NewElementScraper for example
func (elmScrpr *elementScraper) Mappers() map[int]string {
	return elmScrpr.mapper
}

// ElementFunc is optional
// See NewElementScraper for example
func (elmScrpr *elementScraper) ElementFunc(fnc replaceImpl) *elementScraper {
	elmScrpr.fnc = fnc
	return elmScrpr
}

// Run must be called after instantiating NewElementScraper, but after ElementFunc (if used)
// See NewElementScraper for example
func (elmScrpr *elementScraper) Run() string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return elmScrpr.elementScraper(elmScrpr.input)
}

/*NewElementScraper example
input := `bla bla <div class="hoogy"> hoogy </div> Noops <div> no Tile for this one </div>  <div class="hoogy"> dispensed </div> more bla`
scrptr := NewElementScraper(input, `<div class="hoogy">`).ElementFunc(TitleScraperImpl)
fmt.Println(scrptr.Run())

for k, v := range scrptr.Mappers() {
	fmt.Println(k, v)
}

Closing tag must not have any whiteSpace or tabs Eg: </script> is correct but, </script > is incorrect
Incorrect closing tags will be skipped
*/
func NewElementScraper(input, element string) *elementScraper {
	elmScrpr := &elementScraper{
		element: element,
		toggle:  &toggle{},
		tokee:   &tokpos{},
		mapper:  make(map[int]string),
		fnc:     RawScraperImpl,
		input:   input,
	}

	return elmScrpr
}
