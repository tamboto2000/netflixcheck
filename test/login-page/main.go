package main

import "github.com/tamboto2000/netflixcheck"

func main() {
	if err := netflixcheck.TestLoginPage(); err != nil {
		panic(err.Error())
	}
}
