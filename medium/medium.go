package medium

import (
	"encoding/json"
	"errors"
	s "strings"

	log "github.com/sirupsen/logrus"
)

func GetAlertSmsList(data []byte) []byte {
	smsList := []Sms{}
	alerts := []Sms{}
	err := json.Unmarshal(data, &smsList)
	if err != nil {
		log.Error(err)
	}
	for _, v := range smsList {
		sms, err := addmetaData(v)
		if err == nil {
			alerts = append(alerts, sms)
		}
	}
	log.Infof("%#v", alerts)
	dat, _ := json.Marshal(alerts)
	return dat
}

func GetAlert(data []byte) []byte {
	sms := Sms{}
	var alertData []byte
	err := json.Unmarshal(data, &sms)
	if err != nil {
		log.Error(err)
	}
	dat, err := addmetaData(sms)
	if err == nil {
		alertData, _ = json.Marshal(dat)
	}
	return alertData
}

func addmetaData(v Sms) (Sms, error) {
	err := errors.New("text not transaction")
	alert, _ := IsAccountAlert(s.ToLower(v.Body))
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
