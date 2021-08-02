package main

import log "github.com/sirupsen/logrus"

func check_name(s string) bool {
	if s == " " {
		log.Warn("Empty Name field found")
		return false
	} else {
		return true
	}

}
func check_phone(s string) bool {
	if len(s) == 10 {
		return true
	} else {
		log.Warn("Invalid Phone Number found")
		return false

	}

}
func check_email(email string) bool {
	if email == " " {
		log.Warn("Empty Mail field found")
		return false
	} else {
		return true
	}
}
func check_validity(name string, phonenumber string, email string) bool {

	if check_name(name) && check_phone(phonenumber) && check_email(email) {
		return true
	} else {
		return false
	}
}
