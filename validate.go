package main

import (
	"regexp"
	"strconv"
	"strings"
)

func validateProduceItem(produceItem ProduceItem) bool {
	var isProduceCodeValid = validateProduceCode(produceItem.ProduceCode)
	var isNameValid = validateName(produceItem.Name)
	var isUnitPriceValid = validateUnitPrice(produceItem.UnitPrice)

	if isProduceCodeValid && isNameValid && isUnitPriceValid {
		return true
	}

	return false

}

func validateProduceCode(code string) bool {
	//validate Produce Code string criteria
	if code == "" {
		APIValidationError += " Produce Code cannot be empty. "
		return false
	}

	pattern := regexp.MustCompile("[^A-Za-z0-9/-]+")
	loc := pattern.FindIndex([]byte(code))
	if loc != nil {
		APIValidationError += " Produce Code can only contain alphanumeric characters and '-'. "
		return false
	}

	if len(code) != 19 {
		APIValidationError += " Produce Code must be 19 digits including all '-'. "
		return false
	}

	firstFourDigit := string(code[0:4])
	firstDash := string(code[4:5])
	secondFourDigits := string(code[5:9])
	secondDash := string(code[9:10])
	thirdFourDigits := string(code[10:14])
	thirdDash := string(code[14:15])
	fourthFourDigits := string(code[15:19])

	var sliceArray [7]string
	sliceArray = [7]string{firstFourDigit, firstDash, secondFourDigits, secondDash, thirdFourDigits, thirdDash, fourthFourDigits}
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			if validateCodeDigits(sliceArray[i]) == false {
				return false
			}
		} else {
			if sliceArray[i] != "-" {
				APIValidationError += " Produce Code must be in XXXX-XXXX-XXXX-XXXX format. "
				return false
			}
		}
	}

	for i := 0; i < len(Inventory); i++ {
		if strings.ToUpper(code) == strings.ToUpper(Inventory[i].ProduceCode) {
			APIValidationError += " Cannot add duplicate Produce Code to Inventory. "
			return false
		}

	}
	return true
}

func validateCodeDigits(code string) bool {
	pattern := regexp.MustCompile("[^A-Za-z0-9]+")
	loc := pattern.FindIndex([]byte(code))
	if loc != nil {
		APIValidationError += " Produce Code must be in XXXX-XXXX-XXXX-XXXX format. "
		return false
	}
	return true
}

func validateName(name string) bool {
	//will allow name to be any non-empty string(assumption)
	if name == "" {
		APIValidationError += " Name cannot be empty. "
		return false
	}
	return true
}

func validateUnitPrice(price string) bool {
	if price == "" {
		APIValidationError += " Unit Price cannot be empty. "
		return false
	}

	pattern := regexp.MustCompile("[^0-9/.]+")
	loc := pattern.FindIndex([]byte(price))
	if loc != nil {
		APIValidationError += " Unit Price can only contain numbers and decimal point. "
		return false
	}
	//two approved cases(assumptions) number without decimal and number.two digits
	//will assume number with decimal is .00

	//first check if decimal
	var decimalIndex = strings.Index(price, ".")

	if decimalIndex > -1 {
		//Decimal place is first character
		if decimalIndex == 0 {
			APIValidationError += " Unit Price is in incorrect format. "
			return false
		}
		//Decimal Place is last character
		if len(price)-1 == decimalIndex {
			APIValidationError += " Unit Price is in incorrect format. "
			return false
		}
		var tempPrice = string(price[0:decimalIndex] + price[(decimalIndex+1):(len(price)-1)])
		var additionalDecimalIndex = strings.Index(tempPrice, ".")
		if additionalDecimalIndex > -1 {
			APIValidationError += " Unit Price is in incorrect format. "
			return false
		}

		if decimalIndex != len(price)-3 {
			APIValidationError += " Unit Price is in incorrect format. "
			return false
		}

	} else {
		if _, err := strconv.Atoi(price); err == nil {

		} else {
			APIValidationError += " Unit Price is must be a valid number. "
			return false
		}
	}
	return true
}

func parseUnitPrice(price string) string {
	var parsedPrice = "$" + price
	return parsedPrice
}
