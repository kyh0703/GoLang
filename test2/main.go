package main

import "fmt"

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// timer := time.NewTimer(time.Second * 10)
	// go func() {
	// 	<-timer.C
	// 	fmt.Println("timer ON")
	// 	u, ok := service.UserCache.Get("621c1055a89c8e6470dbb194")
	// 	fmt.Println(u.ID, u.Tenant, ok)
	// 	wg.Done()
	// }()
	// wg.Add(1)
	// go service.StartHook()
	// wg.Wait()

	maps := make(map[string]int)
	maps["test"] = 1
	fmt.Println(maps)

	mapOfMap := make(map[string]map[string]int)
	inner := make(map[string]int)
	mapOfMap["test"] = inner
	mapOfMap["test"]["test"] = 1
	fmt.Println(mapOfMap)
}
