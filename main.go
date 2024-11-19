package main

var ENV string

func main() {

	keys := getKeys()
	_, addresses, _ := deployContracts(keys)
	testkeys(addresses, keys, ENV)

}
