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
@[3-6](import required packages)
@[8-9](Book struct)
@[11-13](ToJSON function)
@[15-17](FromJSON function)
@[19-20](API handler function)
```

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
$ go get github.com/stretchr/testify/assert
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




---
## Resources

* [https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028](https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028)
