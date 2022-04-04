package two

import "math/rand"

var addressMap map[int]*AddressBook = make(map[int]*AddressBook)

type AddressBook struct {
	name      string
	address   string
	telephone string
}

func insertAddress(name string, address string, telephone string) *AddressBook {
	addressBook := &AddressBook{name, address, telephone}

	id := rand.Int()
	addressMap[id] = addressBook

	return addressBook
}

func getAddress(id int) *AddressBook {
	return addressMap[id]
}

func deleteAddress(id int) {
	delete(addressMap, id)
}
