package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type JSON struct {
	Data Glossary `json:"glossary"`
}

type Glossary struct {
	Title    string `json:"title"`
	GlossDiv struct {
		Title     string `json:"title"`
		GlossList struct {
			GlossEntry struct {
				ID     string `json:"ID"`
				SortAs string `json:"SortAs"`
				// GlossTerm string   `json:"GlossTerm"`
				// Acronym   string   `json:"Acronym"`
				// Abbrev    string   `json:"Abbrev"`
				// GlossDef  GlossDef `json:"GlossDef"`
				// GlossSee  string   `json:"GlossSee"`
			} `json:"GlossEntry"`
		} `json:"GlossList"`
	} `json:"GlossDiv"`
}

type GlossDef struct {
	Para         string   `json:"para"`
	GlossSeeAlso []string `json:"GlossSeeAlso"`
}

func generateExampleJson() []byte {
	return []byte(`
	{
		"glossary": {
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": ["GML", "XML"]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}
	`)
}

func generateExampleJsonArray() []byte {
	return []byte(`
	[
		{
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": ["GML", "XML"]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	]
	`)
}

func MAKE_JSON_FORMAT() {
	str := `response status code : 400, body : {"status":400,"result":false,"message":"The 123123 has already been used."}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(str[strings.Index(str, "{"):]), &data)
	if err != nil {
		panic(err)
	}
	message, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(strings.Contains(data["message"].(string), "has already been used."))
	fmt.Println(string(message))
}

func USAGE_JSON_Decoder() {
	var err error
	var serializedData []byte = generateExampleJson()

	var decoder *json.Decoder = json.NewDecoder(bytes.NewBuffer(serializedData))
	var data interface{}
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println((data.(map[string]interface{})["glossary"]).(map[string]interface{})["title"])
}

func USAGE_JSON_MARSHALLING() {
	var err error
	var serializedData []byte = generateExampleJson()
	fmt.Println(serializedData)
	raw := &json.RawMessage{}
	err = raw.UnmarshalJSON(serializedData)
	if err != nil {
		panic(err)
	}
	m, err := raw.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%p\n", m)              // 0xc00010c400
	fmt.Printf("%p\n", serializedData) //0xc00010c000
	if err != nil {
		panic(err)
	}
	m1, _ := raw.MarshalJSON()
	fmt.Printf("%p\n", m1) // 0xc00010c400
	m2, _ := raw.MarshalJSON()
	fmt.Printf("%p\n", m2) // 0xc00010c400
	m3, _ := raw.MarshalJSON()
	fmt.Printf("%p\n", m3) // 0xc00010c400
	var data JSON = JSON{}

	b1 := serializedData
	fmt.Printf("%p\n", b1) // 0xc00010c000
	b2 := serializedData
	fmt.Printf("%p\n", b2) // 0xc00010c000
	b3 := serializedData
	fmt.Printf("%p\n", b3) // 0xc00010c000

	err = json.Unmarshal(*raw, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func USAGE_JSON_MARSHALLING_ARRAY() {
	var err error
	var serializedData []byte = generateExampleJsonArray()
	fmt.Println(serializedData)
	raw := &json.RawMessage{}
	err = raw.UnmarshalJSON(serializedData)
	if err != nil {
		panic(err)
	}
	var data []Glossary = []Glossary{}
	err = json.Unmarshal(*raw, &data)
	if err != nil {
		panic(err)
	}
	reserialized, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reserialized))
}

func main() {
	USAGE_JSON_MARSHALLING()
}
