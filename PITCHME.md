Meetup: Cloud Native Applications 
===

* Subject: Getting Started with Cloud Native Go
* Date: 2018-01-20, Saturday, 13:00~16:00
* Place: 15F, ASEM Tower.
* Younggyu Kim (younggyu.kim@oracle.com) 
* OCAP (Oracle Cloud Adoption Platform) Team
* Principal Sales Consultant

---
# Agenda

* Chapter 1 : Introduction to Cloud Native Apps and Microservices
* Chapter 2 : Simple Go Microservices
* Chapter 3 : Introduction to Docker and Go Microservice Containerization
* Chapter 4 : Introduction to Kubernetes and Go Microservice Orchestration
---
![https://github.com/cncf/landscape](https://raw.githubusercontent.com/cncf/landscape/master/landscape/CloudNativeLandscape_latest.png)

---
## Simple Go Microservices

* Simple HTTP server implementation in Go
* JSON marshalling/unmarshalling of Go structs
* Simple REST API implementation

---
### Simple Go HTTP Server Implementation

* Using the Go net/http package
* Implementing and start a simple HTTP server
* Defining simple handler functions

---
#### Initialize Project

```sh
$ echo $GOPATH

# if you haven't set GOPATH variable yet, run these 3 commands
$ mkdir -p ~/go
$ echo "export GOPATH=~/go" >> ~/.bash_profile
$ source ~/.bash_profile

$ mkdir -p ${GOPATH}/src/GettingStartedWithCloudNativeGo
$ cd ${GOPATH}/src/GettingStartedWithCloudNativeGo
$ mkdir -p src/chapter2_1
$ cd src/chapter2_1
$ vi microservice.go
```
---
#### microservice.go
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go.")
}
```

---
#### Run & Test the service
```sh
$ go build -o microservice
$ ./microservice &
$ lsof -n -i tcp:8080
```

And then call the service with this url: [http://localhost:8080](http://localhost:8080)

---
```go
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
```
@[4](import os package)
@[9](call port() function)
@[12-19](write port() function)

---
#### Run & Test the service
```sh
$ go clean
$ go build -o microservice
$ export PORT=9090
$ ./microservice
```
Now you can call the service with this url: [http://localhost:9090](http://localhost:9090)

---
```go
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.ListenAndServe(port(), nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, message)
}
```
@[3](add echo handler)
@[7-12](implement echo() function)

---
#### Run & Test the service

```sh
$ go clean
$ go build -o microservice
$ export PORT=8080
$ ./microservice
```
[http://localhost:8080/api/echo?message=Cloud+Native+Go](http://localhost:8080/api/echo?message=Cloud+Native+Go)

---
### JSON Marshalling/Unmarshalling of Go Structs

* Using the Go JSON package
* JSON marshalling/unmarshalling of Go structs
* Defining Go structs with additional JSON metadata
* Adding simple REST endpoint with JSON response

---
#### Init Project

```sh
$ mkdir src/chapter2_2
$ cp src/chapter2_1/microservice.go src/chapter2_2
$ cd src/chapter2_2
$ mkdir api
$ vi api/book.go
```

---
#### src/chapter2_2/api/book.go
```go
package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
    Title string
    Author string
    ISBN string
}

func (b Book) ToJSON() []byte {
	return nil
}

func FromJSON(data []byte) Book {
	return Book{}
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
}
```
@[3-6](import required packages)
@[8-9](Book struct)
@[11-13](ToJSON function)
@[15-17](FromJSON function)
@[19-20](API handler function)


---

```go
type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}
```

---
#### Test 
```sh
$ echo $GOPATH

$ export GOPATH=~/go
$ go get github.com/stretchr/testify/assert
$ ls $GOPATH/src/github.com/

$ vi api/book_test.go
```

---
#### src/chapter2_2/api/book_test.go
```go
package api

import (
	"testing"
)

func TestBookToJSON(t *testing.T) {
}

func TestBookFromJSON(t *testing.T) {
}
```

---
#### Run Test code

```sh
$ go test ./api -v
```

```text
=== RUN   TestBookToJSON
--- PASS: TestBookToJSON (0.00s)
=== RUN   TestBookFromJSON
--- PASS: TestBookFromJSON (0.00s)
```

---
```go
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"0123456789"}`,
		string(json), "Book JSON marshalling wrong.")
}
```

```sh
$ go test ./api -v
```

---
#### book.go - FromJSON

```go
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}
``` 

---
#### book_test.go - 

```go
func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"Title":"Cloud Native Go","Author":"M.-L. Reimer","ISBN":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
		book, "Book JSON unmarshalling wrong.")
}
```

```sh
$ go test ./api -v
```

---
#### book_test.go - to lower case

```go
func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
		book, "Book JSON unmarshalling wrong.")
}
```
@[5](Title - title, Author - author, ISBN - isbn)
@[10](Title - title, Author - author, ISBN - isbn)

---

#### Run Test
```sh
$ go test ./api -v
```

```text
=== RUN   TestBookToJSON
--- FAIL: TestBookToJSON (0.00s)
        Error Trace:    book_test.go:13
        Error:          Not equal:
                        expected: "{\"title\":\"Cloud Native Go\",\"author\":\"M.-L. Reimer\",\"isbn\":\"0123456789\"}"
                        actual  : "{\"Title\":\"Cloud Native Go\",\"Author\":\"M.-L. Reimer\",\"ISBN\":\"0123456789\"}"
        Test:           TestBookToJSON
        Messages:       Book JSON marshalling wrong.
=== RUN   TestBookFromJSON
--- PASS: TestBookFromJSON (0.00s)
FAIL
exit status 1
FAIL    _/Users/nicholas/Documents/ShapeMyFuture/Getting_Started_with_Cloud_Native_Go/GettingStartedWithCloudNativeGo/src/chapter2_2/api      0.017s
```

---
#### book.go - Book struct
add json tag like below

```go
type Book struct {
	Title  string	`json:"title"`
	Author string	`json:"author"`
	ISBN   string	`json:"isbn"`
}
```

```sh
$ go test ./api -v
```

---
#### book.go - BooksHandleFunc

```go
var Books = []Book {
	Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
```
@[1-4](add sample book data)
@[6-14](BooksHandleFunc function)
---
## Resources

* [https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028](https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028)
