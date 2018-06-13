package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/smilecs/parser/smsparser"
)

type MainSms struct {
	SmsTag []smsparser.Sms `xml:"sms"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	sms := MainSms{}
	//	newSmsList := []medium.Sms{}
	data, err := ioutil.ReadFile("./sms.xml")
	if err != nil {
		log.Println(err)
	}
	err = xml.Unmarshal(data, &sms)
	if err != nil {
		log.Println(err)
	}
	//log.Infof("%#v", sms.SmsTag)
	//list := []smsparser.Sms{}
	//toParse, err := json.Marshal(sms.SmsTag)
	if err == nil {
		data := smsparser.GetAlertSmsList(sms.SmsTag)
		log.Infof("%#v", data)
		parsed, _ := json.Marshal(data)
		err = ioutil.WriteFile("smile.json", parsed, 0644)
		if err != nil {
			log.Error(err)
		}
	}
	/*	for _, v := range sms.SmsTag {

		alert, bankName := medium.IsAccountAlert(s.ToLower(v.Body))
		v.BankName = bankName
		if alert {
			v.BankName = v.Address
			v.IsDebit = medium.IsDebit(s.ToLower(v.Body))
			v.Cateogry = medium.TagCategory(s.ToLower(v.Body))
			v.Amount = medium.GetAmount(s.ToLower(v.Body))
			v.Date = medium.GetDate(s.ToLower(v.Body))
			v.Currency = medium.GetCurrency(v.Body)
			newSmsList = append(newSmsList, v)
			log.Infof("%#v", v)
		}

	}*/
}
