package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	/*r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}*/
	fmt.Fprintf(w, "hello chain!")

}

func computeBall() {
	fmt.Println(33 * 33 * 33 * 33 * 33 * 33 * 16) //20663487504

	redBall := []int{1, 4, 17, 31, 16, 5}
	//blueBall := []int{9}
	//01020304050601
	//01020304050616

	//33323130292816

	redBallStart := 1
	redBallEnd := 33
	//blueBallStart := 1
	//blueBallEnd := 16
	rand.Seed(time.Now().UnixNano())

	redBallLen := len(redBall)
	//blueBallLen := len(blueBall)
	sum := 0
	same := false
	for sum = 0; sum < 1000000000 && same == false; sum++ {
		for i := 0; i < redBallLen; i++ {
			redValue := redBallStart + rand.Intn(redBallEnd)
			if redValue != redBall[i] {
				break
			}
			if i == redBallLen-1 {
				same = true
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
