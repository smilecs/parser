package smsparser

import (
	"errors"
	s "strings"

	log "github.com/sirupsen/logrus"
)

func GetAlertSmsList(data []Sms) []Sms {
	//smsList := []Sms{}
	alerts := []Sms{}
	//err := json.Unmarshal(data, &smsList)
	/*if err != nil {
		log.Error(err)
	}*/
	//log.Infof("%#v", data)
	for _, v := range data {
		sms, err := addmetaData(v)
		//log.Infof("%#v", sms)
		if err == nil {
			alerts = append(alerts, sms)
		}
	}
	//dat, _ := json.Marshal(alerts)
	return alerts
}

func GetAlert(data Sms) Sms {
	/*var alertData []byte
	err := json.Unmarshal(data, &sms)
	if err != nil {
		log.Error(err)
	}*/
	dat, err := addmetaData(data)
	if err == nil {
		log.Error(err)
		//alertData, _ = json.Marshal(dat)
	}
	return dat
}

func addmetaData(v Sms) (Sms, error) {
	err := errors.New("text not transaction")
	alert := IsAccountAlert(s.ToLower(v.Body))
	if alert {
		v.BankName = v.Address
		v.IsDebit = IsDebit(s.ToLower(v.Body))
		v.Cateogry = TagCategory(s.ToLower(v.Body))
		v.Amount = GetAmount(s.ToLower(v.Body))
		v.Date = GetDate(s.ToLower(v.Body))
		v.Currency = GetCurrency(v.Body)
		err = nil
	}
	return v, err
}
