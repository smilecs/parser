package medium

import (
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
	Currency      string
}

const DEBIT = "debit"

func IsAccountAlert(value string) (bool, string) {
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

func IsDebit(body string) bool {
	return s.Contains(body, DEBIT)
}

func TagCategory(body string) string {
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

func GetAmount(body string) float64 {
	r, _ := regexp.Compile(`(((\d+,\d+))|\d+\.\d+)`)
	values := r.FindAllString(body, -1)
	a, err := strconv.ParseFloat(s.Replace(values[0], ",", "", -1), 64)
	if err != nil {
		log.Info(err)
	}
	return a
}

func GetDate(body string) string {
	r, _ := regexp.Compile(`((\d{2}|\d)-([a-z]{3}|[a-z]+)-\d+)`)
	values := r.FindAllString(body, -1)
	return values[0]
}

func GetCurrency(body string) string {
	var key = "unknown"
	for _, v := range Currency {
		for _, vv := range v {
			if s.Contains(body, vv) {
				key = vv
				break
			}
		}
	}
	return key
}
