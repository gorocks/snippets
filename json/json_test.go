package json_test

import (
	"encoding/json"
	"testing"
)

// jsonText comes from http://json.org/example.html
var jsonText = []byte(`
{
   "glossary":{
      "title":"example glossary",
      "GlossDiv":{
         "title":"S",
         "GlossList":{
            "GlossEntry":{
               "ID":"SGML",
               "SortAs":"SGML",
               "GlossTerm":"Standard Generalized Markup Language",
               "Acronym":"SGML",
               "Abbrev":"ISO 8879:1986",
               "GlossDef":{
                  "para":"A meta-markup language, used to create markup languages such as DocBook.",
                  "GlossSeeAlso":[
                     "GML",
                     "XML"
                  ]
               },
               "GlossSee":"markup"
            }
         }
      }
   }
}`)

type glossary struct {
	Glossary struct {
		Title    string `json:"title"`
		GlossDiv struct {
			Title     string `json:"title"`
			GlossList struct {
				GlossEntry struct {
					ID        string `json:"ID"`
					SortAs    string `json:"SortAs"`
					GlossTerm string `json:"GlossTerm"`
					Acronym   string `json:"Acronym"`
					Abbrev    string `json:"Abbrev"`
					GlossDef  struct {
						Para         string   `json:"para"`
						GlossSeeAlso []string `json:"GlossSeeAlso"`
					} `json:"GlossDef"`
					GlossSee string `json:"GlossSee"`
				} `json:"GlossEntry"`
			} `json:"GlossList"`
		} `json:"GlossDiv"`
	} `json:"glossary"`
}

type glossarySectional struct {
	Glossary struct {
		Title    string `json:"title"`
		GlossDiv struct {
			Title     string          `json:"title"`
			GlossList json.RawMessage `json:"GlossList"` // diff: delay JSON decoding
		} `json:"GlossDiv"`
	} `json:"glossary"`
}

func benchmarkJSONUnmarshal(f func(), b *testing.B) {
	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkJSONUnmarshal_0(b *testing.B) {
	benchmarkJSONUnmarshal(func() {
		var g glossary
		json.Unmarshal(jsonText, &g)
	}, b)
}

func BenchmarkJSONUnmarshal_1(b *testing.B) {
	benchmarkJSONUnmarshal(func() {
		var g glossarySectional
		json.Unmarshal(jsonText, &g)
	}, b)
}

// BenchmarkJSONUnmarshal_0-8   	  200000	     10565 ns/op
// BenchmarkJSONUnmarshal_1-8   	  200000	      7699 ns/op

// benchmark                    old ns/op     new ns/op     delta
// BenchmarkJSONUnmarshal-8     10298         7591          -26.29%
