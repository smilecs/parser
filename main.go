package main

import (
	"encoding/xml"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Sms struct {
	Body          string `xml:"body,attr"`
	Subject       string `xml:"subject,attr"`
	Address       string `xml:"address,attr"`
	ServiceCenter string `xml:"service_center,attr"`
	BankName      string
	IsCredit      bool
	IsDebit       bool
	Cateogry      string
}

type MainSms struct {
	SmsTag []Sms `xml:"sms"`
}

func main() {
	sms := MainSms{}
	//newSmsList := []Sms{}
	data, err := ioutil.ReadFile("./sms.xml")
	if err != nil {
		log.Println(err)
	}
	err = xml.Unmarshal(data, &sms)
	if err != nil {
		log.Println(err)
	}

	log.Info(sms.SmsTag[0])
}
