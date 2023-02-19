package main

import "math/rand"

var CANDIDATE = "abcdefghijklmnopqrstuvwxyz"
var LENGTH = 4

func generateToken(t ContentType) string {
	for true {
		token := generateTokenCore()
		if !Contains(token, t) {
			return token
		}
	}
	return ""
}

func generateTokenCore() string {
	token := ""
	for i := 0; i < LENGTH; i++ {
		index := rand.Intn(len(CANDIDATE))
		token += CANDIDATE[index: index+1]
	}
	return token
}