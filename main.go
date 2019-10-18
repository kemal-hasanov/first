package main

import "fmt"

func asdasd(ime string, pozdrav string) {
	print(pozdrav)
	print(" ")

	print(ime)
	print(" ")

	print(familiq(ime))
	println("!")
}
func familiq(ime string) string {
	return ime + "ov"
}

func main() {
	var ime string
	var pozdrav string
	fmt.Scanln(&ime)
	fmt.Scanln(&pozdrav)
	asdasd(ime, pozdrav)
	asdasd("Kemal", pozdrav)
	asdasd(ime, "chao")

}
