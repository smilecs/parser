package main

import (
	"encoding/xml"
	"io/ioutil"
	"regexp"
	"strconv"
	s "strings"

	log "github.com/sirupsen/logrus"
)

type Sms struct {
	Body          string `xml:"body,attr"`
	Subject       string `xml:"subject,attr"`
	Address       string `xml:"address,attr"`
	ServiceCenter string `xml:"service_center,attr"`
	BankName      string
	IsDebit       bool
	Cateogry      string
	Date          string
	Amount        float64
}

type MainSms struct {
	SmsTag []Sms `xml:"sms"`
}

const DEBIT = "debit"

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	sms := MainSms{}
	// newSmsList := []Sms{}
	data, err := ioutil.ReadFile("./sms.xml")
	if err != nil {
		log.Println(err)
	}
	err = xml.Unmarshal(data, &sms)
	if err != nil {
		log.Println(err)
	}
	for _, v := range sms.SmsTag {
		alert, bankName := isAccountAlert(s.ToLower(v.Body))
		v.BankName = bankName
		if alert {
			v.IsDebit = isDebit(s.ToLower(v.Body))
			v.Cateogry = tagCategory(s.ToLower(v.Body))
			v.Amount = getAmount(s.ToLower(v.Body))
			log.Infof("%#v", v)
		}

	}
}

func isAccountAlert(value string) (bool, string) {
	var bankName string
	bankMatcher, _ := regexp.MatchString(`(\d{3}[A-Za-z]{4}\d{3})`, value)
	for _, k := range BankNames {
		if s.Contains(value, k) && bankMatcher {
			bankName = k
			break
		}
	}
	return bankMatcher, bankName
}

func isDebit(body string) bool {
	return s.Contains(body, DEBIT)
}

func tagCategory(body string) string {
	var key = "other"
	for k, v := range Categories {
		for _, vv := range v {
			if s.Contains(body, vv) {
				key = k
				break
			}
		}
	}
	return key
}

//your acct \d+x+\d+[a-z\s]+(\d+\.?\d{0,2}) on (\d{2}-[a-z]{3}-\d+).*

func getAmount(body string) float64 {
	r, _ := regexp.Compile(`(((\d+,\d+))|\d+\.\d+)`)
	values := r.FindAllString(body, -1)
	a, err := strconv.ParseFloat(s.Replace(values[0], ",", "", -1), 64)
	if err != nil {
		log.Info(err)
	}
	return a
}
