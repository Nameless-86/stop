package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import the pprof package for profiling
	"os"
	"runtime/pprof"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	myMapMake := make(map[string]int, 100)
	myMapLiteral := map[string]int{}

	for i := 0; i < 200; i++ {
		myMapMake[fmt.Sprintf("key%d", i)] = 1
		myMapLiteral[fmt.Sprintf("key%d", i)] = 1
	}

	// Print the lengths of both maps to confirm
	fmt.Println("Length of myMapMake:", len(myMapMake))
	fmt.Println("Length of myMapLiteral:", len(myMapLiteral))

	//memory profile
	f, err := os.Create("memprofile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}

	select {}
}
