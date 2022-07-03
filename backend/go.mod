module capstone

go 1.18

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/pa-m/sklearn v0.0.0-20200711083454-beb861ee48b1
	gonum.org/v1/gonum v0.11.0
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/pa-m/optimize v0.0.0-20190612075243-15ee852a6d9a // indirect
	github.com/pa-m/randomkit v0.0.0-20191001073902-db4fd80633df // indirect
	golang.org/x/exp v0.0.0-20191129062945-2f5052295587 // indirect
	golang.org/x/tools v0.1.9 // indirect
)

// go mod edit -replace github.com/pa-m/optimize=./optimize
replace github.com/pa-m/optimize => ./optimize
