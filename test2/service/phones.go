package service

import "fmt"

// defined collection name
const PhoneCollection = "phone"

// instance
var PhoneCache = NewCacheMap[User](PhoneCollection)

func init() {
	fmt.Println("Phone Init")
	Hook[PhoneCollection] = PhoneCache
}
