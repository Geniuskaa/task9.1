package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err, _ := execute(); err != nil {
		os.Exit(1)
	}

	_, data := execute()
	var decoded Curriencies
	err := xml.Unmarshal(data, &decoded) // преобразовали данные в тип Curriencies
	//log.Println(decoded)
	log.Println(err)

	fileForJson := transferFromXmlToJson(decoded)
	log.Println(ExportJson(fileForJson))

}

type Client struct {
	Transport http.RoundTripper
	CheckRedirect func (rew *http.Request, via []*http.Request) error
	Jar http.CookieJar
	Timeout time.Duration
}

var DefaultClient *Client = &Client{}


func execute() (err error, dataImport []byte) { //получаем данные с сервера
	reqUrl := "https://raw.githubusercontent.com/netology-code/bgo-homeworks/master/10_client/assets/daily.xml"
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Println(err)
		return err, nil
	}

	log.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Println(cerr)
			err = cerr
		}
	}()

	data := body
	//log.Println(string(body))

	return err, data

}

type CurriencyXML struct {
	ID string `xml:"ID,attr"`
	NumCode int64 `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal int64 `xml:"Nominal"`
	Name string `xml:"Name"`
	Value float64 `xml:"Value"`
}

type Curriencies struct {
	//XMLName xml.Name `xml:"ValCurs"`
	XMLName xml.Name `xml:"ValCurs"`
	Date string `xml:"Date,attr"`
	Name string `xml:"name,attr"`
	ValueIds []CurriencyXML `xml:"Valute"`
}

type CurrienciesForStore struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Value float64 `json:"value"`
}

type CurriencesJSON struct {
	Code string
	Name string
	Value float64
}

func transferFromXmlToJson(file Curriencies) (anotherFile []CurrienciesForStore) {
	//length :=len(file.ValueIds)
	var sliceForStore []CurrienciesForStore //:= //make([]CurrienciesForStore, length)
	for m := range file.ValueIds {
		sliceForStore = append(sliceForStore, CurrienciesForStore{
			Code:  string(file.ValueIds[m].CharCode),
			Name:  file.ValueIds[m].Name,
			Value: file.ValueIds[m].Value,
		})
	}
	return sliceForStore
}




func ExportJson(sliceOfTransactions []CurrienciesForStore)  ([]byte, error) { //sliceOFTransactions []Transaction

	file, err := os.Create("currencies.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(file)

	rawDataIn, err := ioutil.ReadFile("currencies.json")
	if err != nil {
		log.Println("Cannot load settings:", err)
	}
	log.Println(rawDataIn)

	encoded, err := json.Marshal(sliceOfTransactions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = ioutil.WriteFile("currencies.json", encoded, 0)
	if err != nil {
		log.Println("Cannot write updated settings file:", err)
	}

	/*var decoded []Transaction
	// Важно: передаём указатель, чтобы функция смогла записать данные
	err = json.Unmarshal(encoded, &decoded)
	log.Printf("%#v", decoded)*/


	return encoded, nil
}