// app/test/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/taadis/letgo/handler/home"

	"github.com/taadis/letgo/ghttp"
	//"github.com/taadis/letgo/handler/company"
	//"github.com/taadis/letgo/handler/home"
	//"github.com/taadis/letgo/handler/middleware"
	//"github.com/taadis/letgo/handler/user"
)

func main() {
	fmt.Println("ready to start test app")
	mux := ghttp.NewServeMux()
	mux.Get("/", home.HandleFunc)
	//mux.HandleFunc("/system/user/", user.HandleFunc)
	//mux.Handle("/system/user/", midleware.PanicAndRecover(http.HandlerFunc(user.HandleFunc)))
	//mux.HandleFunc("/company/", company.HandleFunc)
	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
	//fmt.Println("golb running in: ", "http://localhost:5903")
}