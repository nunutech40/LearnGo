package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetingssss: ")
	log.SetFlags(0)

	names := []string{"Nunu", "Dorothy", "Putri"}

	// Ambil pesan salam.
	message, err := greetings.Hellos(names)

	// Jika ada eror, cetak ke layar dan keluar dari program.
	if err != nil {
		log.Fatal(err)
	}

	// Jika tidak ada eror, cetak pesan yang dikembalikan ke layar.
	fmt.Println(message)
}

func learning1() {
	fmt.Println("I'm start learning Go Language")
}