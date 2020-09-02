package transactions

import (
	"encoding/json"
	"log"
)
type Service struct {

}
type CurriencyXML struct {
	XMLName string `xml:"Valute ID"`
	//ID string `xml:"Id"`
	NumCode int `xml:"Num_code"`
	CharCode string `xml:"Char_code"`
	Nominal int `xml:"Nominal"`
	Name string `xml:"Name"`
	Value float64 `xml:"Value"`
}

type Curriencies struct {
	XMLName string `xml:"ValCurs"`
	Currencies []CurriencyXML // заменить!!!
}

type CurrienciesForStore struct {
	Code string
	Name string
	Value float64
}

type CurriencesJSON struct {
	Code string
	Name string
	Value float64
}






func ExportJson(sliceOfTransactions []CurriencyXML)  ([]byte, error) { //sliceOFTransactions []Transaction

	encoded, err := json.Marshal(sliceOfTransactions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(encoded))

	/*var decoded []Transaction
	// Важно: передаём указатель, чтобы функция смогла записать данные
	err = json.Unmarshal(encoded, &decoded)
	log.Printf("%#v", decoded)*/


	return encoded, nil
}
