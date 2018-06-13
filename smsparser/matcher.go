package smsparser

var (
	bankNames = map[string]string{
		"1":  "firstbank",
		"2":  "fcmb",
		"3":  "uba",
		"4":  "gtbank",
		"5":  "keystone",
		"6":  "unionbank",
		"7":  "zenith",
		"8":  "sterling",
		"9":  "wema",
		"10": "alat",
	}

	categories = map[string][]string{
		"eatery":        {"chicken republic", "sweet sensation", "uq restaurant", "kfc", "debonairs", "dominos", "cold stone"},
		"leisure":       {"cinema", " game", "plaza", "beer hug"},
		"transport":     {"uber", "taxify", "lyft", "uberm"},
		"restaurant":    {"cafe neo", "neo", "cafe"},
		"grocery":       {"hubmart", "spar", "domino"},
		"savings":       {"piggybank", "cowrywise", "piggy"},
		"personal":      {"cosmetics"},
		"loans":         {"paylater", "branch", "aella", "onefinance", "onefi"},
		"entertainment": {"netflix", "iroko", "iflix", "googleplay", "ecenter", "center"},
		"online":        {"paystack", "flutterwave"},
		"atm":           {"atm", "withdrawal"},
		"transfer":      {"cr"},
	}

	currency = map[string][]string{
		"Naira": {"NGN", "N"},
	}
)
