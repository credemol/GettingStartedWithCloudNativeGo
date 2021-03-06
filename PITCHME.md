Meetup: Cloud Native Applications 
===

---
## Getting Started with Cloud Native Go
* Subject: Getting Started with Cloud Native Go
* Date: 2018-01-20, Saturday, 13:00~16:00
* Place: 15F, ASEM Tower.
* Younggyu Kim (younggyu.kim@oracle.com) 
* OCAP (Oracle Cloud Adoption Platform) Team
* Principal Sales Consultant

---
## Resource

- GIT: [https://github.com/credemol/GettingStartedWithCloudNativeGo](https://github.com/credemol/GettingStartedWithCloudNativeGo)
- PT: [https://gitpitch.com/credemol/GettingStartedWithCloudNativeGo/master](https://gitpitch.com/credemol/GettingStartedWithCloudNativeGo/master)

---
### Postman Configuration

```sh
$ mkdir -p ~/tmp/cloud-native-go
$ cd ~/tmp/cloud-native-go
$ ls

# Open GettingStartedWithCloudNativeGo folder from Visual Studion Code

# run postman application
# Import GettingStartedWithCloudNativeGo
# Import Cloud-Native-Go Env 
```


---
## Agenda

1. Introduction to Cloud Native Apps and Microservices
1. Simple Go Microservices
1. Introduction to Docker and Go Microservice Containerization
1. Introduction to Kubernetes and Go Microservice Orchestration
---
![https://github.com/cncf/landscape](https://raw.githubusercontent.com/cncf/landscape/master/landscape/CloudNativeLandscape_latest.png)

---
## 2. Simple Go Microservices

* Simple HTTP server implementation in Go
* JSON marshalling/unmarshalling of Go structs
* Simple REST API implementation

---
### 2.1 Simple Go HTTP Server Implementation

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
# Open GettingStartedWithCloudNativeGo folder with Visual Studio Code
$ mkdir -p chapter2_1
$ cd chapter2_1
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
# go build -o microservice.exe
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
# go build -o microservice.exe
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
### 2.2 JSON Marshalling/Unmarshalling of Go Structs

* Using the Go JSON package
* JSON marshalling/unmarshalling of Go structs
* Defining Go structs with additional JSON metadata
* Adding simple REST endpoint with JSON response

---
#### Initialize Project

```sh
$ cd ${GOPATH}/src/GettingStartedWithCloudNativeGo
$ cp -R chapter2_1 mkdir chapter2_2
$
$ cd chapter2_2
$ mkdir api
$ vi api/book.go
```

---
#### chapter2_2/api/book.go
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
@[8-12](Book struct)
@[14-16](ToJSON function)
@[18-20](FromJSON function)
@[22-23](API handler function)


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
#### Unit Test 
```sh
$ echo $GOPATH

$ export GOPATH=~/go
$ go get github.com/stretchr/testify/assert
$ ls $GOPATH/src/github.com/

$ vi api/book_test.go
```

---
#### chapter2_2/api/book_test.go
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
#### chapter2_2/microservice.go

```go
package main

import (
	"GettingStartedWithCloudNativeGo/chapter2_2/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}
```
@[4](import book api)
@[14](add HandlerFunc for '/api/books')

---
#### Run & Test microservice

```sh
$ go build -o microservice
$ ./microservice
```

With Postman, send a request to [http://localhost:8080/api/books](http://localhost:8080/api/books)

---
### 2.3 Simple REST API Implementation

* Using the request path and query parameters
* Using HTTP status codes
* Using HTTP verbs: GET, PUT, POST, DELETE

---
#### Initialize Project

```sh
$ cd ${GOPATH}/src/GettingStartedWithCloudNativeGo
$ cp -R chapter2_2 chapter2_3
$ cd chapter2_3
```

---
#### chapter2_3/api/book.go

```go
package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = map[string]Book{
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
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
@[12](add description field, with omitempty)
@[32-35](using map in order to mimic an in-memory data store)

---
#### book.go - AllBooks(), writeJSON() functions

```go
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0

	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
```
@[1-10](AllBooks function)
@[12-19](writeJSON function)

---
#### book.go - BooksHandleFunc

```go
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}
```

---
#### chapter2_3/microservice.go

```go
import (
	"GettingStartedWithCloudNativeGo/chapter2_3/api"
	"fmt"
	"net/http"
	"os"
)
```
@[2](chapter2_2 -> chapter2_3)

---
#### Run & Test microservice.go

```sh
$ cd chapter2_3
$ go build -o microservice
$ ./microservice
```

With Postman, send a request to [http://localhost:8080/api/books](http://localhost:8080/api/books) with GET and POST

---
#### book.go - CreateBook

```go
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}
```

---
#### book.go - BooksHandleFunc

```go
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}
```
@[6-19](Create Book)

---
#### Run & Test microservice

```sh
$ go build -o microservice
$ ./microservice
```
With Postman, send a request to http://localhost:8080/api/books with POST.
The request payload should be like below

```json
{
	"title": "New Book",
	"author": "Postman",
	"isbn": "1234567890"
}
```

---
### Get One, Update, Delete a book

---
#### book.go - BookHandleFunc
```go
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	isbn := path[len("/api/books/"):]
	fmt.Println("isbn", isnb)

	switch method := r.Method; method {
	case http.MethodGet:
		fmt.Println("get")
	case http.MethodPut:
		fmt.Println("put")
	case http.MethodDelete:
		fmt.Println("delete")
	default:
		fmt.Println("method:", method)
		w.WriteHeader(http.StatusBadRequest)
	}
}
```

---
#### microservice.go 

```go
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}
```
@[6](add BookHandleFunc)

---
#### Test & Run microservice

```sh
$ go build -o microservice
$ ./microservice
```
You can send a request to [http://localhost:8080/api/books/0000000000](http://localhost:8080/api/books/0000000000)

---
#### book.go

```go
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

func DeleteBook(isbn string) {
	delete(books, isbn)
}
```
@[1-4](GetBook function)
@[6-12](UpdateBook function)
@[14-16](DeleteBook function)

---
#### book.go - BookHandleFunc

```go
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("path:", path)

	isbn := path[len("/api/books/"):]
	fmt.Println("isbn:", isbn)

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}
```
@[9-15](Read a Book)
@[16-28](Update a Book)
@[29-31](Delete a Book)
@[32-34](Bad Request)

---
### Test & Run

```sh
$ go build -o microservice
$ ./microservice
```

call this url [http://localhost:8080/api/books/0000000000](http://localhost:8080/api/books/0000000000) with
HTTP GET, PUT, DELETE method.
when updating the book, use this payload

```json
{
    "title": "Cloud Natvie Go(Update)",
    "author": "M.-Leander Reimer",
    "isbn": "0000000000"
}
```

---
## Chapter 3

> Introduction to Docker and Microservice Containerization

1. Basic Docker Workflow and Docker Commands
1. Build Naive Docker Image for Go Microservice
1. Running Containerized Go Microservice Locally
1. Improved Docker Image and Docker Compose

---
Initilize Project

```sh
$ cd $GOPATH/src/GettingStartedWithCloudNativeGo
$ cp -R chapter2_3 chapter3_1
```
---
### 3.1 Basic Docker Workflow and Docker Commands

* Hardware versus OS virtualization
* Docker images and containers
* The Docker workflow
* Some basic Docker commands

---
#### Hardware versus OS Virtualization 
![Hardware versus OS Virtualization](https://i.ytimg.com/vi/jDivl-tCWJM/maxresdefault.jpg)

---
#### Docker images and Containers
(Image)

---
#### Docker Hub

1. hub.docker.com
1. keyword: golang
1. click Search button

---
#### The Docker Workflow
![docker-stages](https://user-images.githubusercontent.com/5771924/34644160-5939b5a0-f385-11e7-8dfd-457c965e7298.png)

---
#### Some Baisc Docker Commands

Command                                      | Action
---------------------------------------------|-------------------------------------------------
docker build -t <tag> .                      | Build a docker image with the given tag from the current directory.
docker images<br/>docker image ls            | Prints all local images
docker run<br/>  -d<br/>  -e `<environment variable>`<br/>  -p `<host-port>:<container-port>`<br/>  `<image> <entrypoint process>` | Run a Docker image: Creates and runs a container.<ul><li>in background</li><li>with defined EVN variable</li><li>with port forwarding from Host to Container</li><li>image name and entrypoint process</li></ul>
docker run <br/>  -ti<br/>  `<image> /bin/sh`| Run a Docker image and open a shell within the container <ul><li>with forwarding of local terminal</li><li>image name and shell (or "/bin/bash")</li></ul> 
docker ps --all<br/>docker container ls -a   | Prints all containers (without --all prints only running containers)
docker kill `<container>`                    | Terminate container (send SIGKILL to entrypoint process)
docker rm `<container>`<br/>docker container rm `<container>` | Remove container.
docker rim -f `<image>`<br/>docker image rm `<image>` | Remove local image(use -f to force removal) 

---

```sh
$ docker images
$ docker image ls
$ docker image ls | grep golang
```
visit [https://hub.docker.com/_/golang/](https://hub.docker.com/_/golang/) and see its versions and description

```sh
$ docker pull golang
$ docker image ls | grep golang
```

---
### 3.2 Build Naive Docker Image for Go Microservice

* Writing a Dockerfile for a Go microservice
* Building the Docker image from the Dockerfile
* Tagging and pushing image to Docker Hub

---
#### Writing a Dockerfile

```sh
$ cd $GOPATH/src/GettingStartedWithCloudNativeGo
$ cp -R chapter2_3 chapter3_2
$ cd chapter3_2

$ go clean 
$ grep -rl 'chapter2_3' ./ | xargs sed -i '' -e 's/\/chapter2_3//g' 
$ vi Dockerfile
```

---
#### chapter3_2/Dockerfile
```dockerfile
FROM golang:latest
MAINTAINER YG Kim(credemol@gmail.com)

ENV SOURCES /go/src/GettingStartedWithCloudNativeGo/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT GettingStartedWithCloudNativeGo
```

---
#### Build Docker Image

```sh
docker build -t cloud-native-go:1.0.0 .
docker image ls | grep cloud-native-go
```

---
#### Alpine Version

chapter3_2/Dockerfile
```dockerfile
FROM golang:1.9.2-alpine3.7
```

```sh
docker build -t cloud-native-go:1.0.1 .
docker image ls | grep cloud-native-go
```

---
#### Push Docker image to Docker Hub

```sh
$ docker tag cloud-native-go:1.0.1 credemol/cloud-native-go:1.0.1
$ docker image ls | grep cloud-native-go
$ docker login
$ docker push credemol/cloud-native-go:1.0.1

```
see if your docker image is registered to Docker Hub
[https://hub.docker.com/r/credemol/cloud-native-go/](https://hub.docker.com/r/credemol/cloud-native-go/)
In this case, credemol is my username of Docker Hub, so please change it with your own username.

---
### 3.3 Running Containerized Go Microservice Locally

* Running Docker image locally
* Specify environment variables and ports
* Starting, stopping, and restarting containers
* Adding CPU and memory constraints


#### Running Docker image locally

```sh
$ cd $GOPATH/src/GettingStartedWithCloudNativeGo
$ docker image ls | grep cloud-native-go
$ docker run -it -p 8080:8080 cloud-native-go:1.0.1
```

Send a GET request to [http://docker-host:8080/api/books](http://docker-host:8080/api/books), in this case docker-host can be localhost if the Operating System of your laptop, or if you are working on Windows 7 or 8, please run _docker-machine ls_ to see what docker host ip is. That is configured as _192.168.99.100_ by default. 

> Press Ctrl+C to stop the container

---
#### Running Docker image locally (Continued)

```sh
$ docker run -it -e "PORT=9090" -p 8080:9090 cloud-native-go:1.0.1
$ Ctrl+C

$ docker ps 
$ docker ps --all (or -a)
$ docker container ls
$ docker container ls -a 
$ docker run --name cloud-native-go -d -p 8080:8080 cloud-native-go:1.0.1
$ docker container ls
$ docker container ls -a
```
Through Postman, you can see that everythinng is working as expected

---
#### Running Docker image locally (Continued)

```sh
$ docker stop cloud-native-go
$ docker container ls
$ docker container ls -a
$ docker start cloud-native-go
$ docker container ls
$ docker kill cloud-native-go
$ docker container rm cloud-native-go
$ docker run --name cloud-native-go --cpu-quota 50000 --memory 64m --memory-swappiness 0 -d -p 8080:8080 cloud-native-go:1.0.1
$ docker container rm -f cloud-native-go
$ docker container rm $(docker container ls -aq)
```

---
### 3.4 Improved Docker Image and Docker Compose

* Writing improved Dockerfile for smaller images
* Using Docker Compose to build and run the image

---
#### Writing improved Dockerfile for smaller images

```sh
$ cp $GOPATH/src/GettingStartedWithCloudNativeGo
$ cp -R chapter2_3 chapter3_4
$ cd chapter3_4

$ go clean 
$ grep -rl 'chapter2_3' ./ | xargs sed -i '' -e 's/chapter2_3/chapter3_4/g' 

$ vi Dockerfile
```

---
#### chapter3_4/Dockerfile

```dockerfile
FROM alpine:3.5
MAINTAINER YG Kim(credemol@gmail.com)

COPY ./Cloud-Native-Go /app/Cloud-Native-Go
RUN chmod +x /app/Cloud-Native-Go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/Cloud-Native-Go
```

---
#### Go build

```sh
$ GOOS=linux GOARCH=amd64 go build -o Cloud-Native-Go
$ ls

$ docker build -t cloud-native-go:1.0.1-alpine .
$ docker image ls | grep cloud-native-go

$ docker run --name cloud-native-go -d -p 8080:8080 cloud-native-go:1.0.1-alpine

$ docker tag cloud-native-go:1.0.1-alpine credemol/cloud-native-go:1.0.1-alpine
$ docker login
$ docker push credemol/cloud-native-go:1.0.1-alpine


```

---
#### Docker Compose

```sh
$ docker container stop $(docker container ls -q)
$ docker container rm $(docker container ls -aq)
$ docker image rm cloud-native-go:1.0.1-alpine
$ vi docker-compose.yml
```

#### chapter3_4/docker-compose.yml

```yaml
version: '2'
services:
  microservice:
    build: .
    image: cloud-native-go:1.0.1-alpine
    environment:
    - PORT=9090
    ports:
    - "9090:9090"
```

---
#### docker-compose build 

```sh
$ docker-compose build
$ docker image ls

```

---
#### docker-compose.yml - multi containers

```yaml
version: '2'
services:
  microservice:
    build: .
    image: cloud-native-go:1.0.1-alpine
    environment:
    - PORT=9090
    ports:
    - "9090:9090"
  nginx:
    image: "nginx:1.11.9"
    ports:
    - "8080:80"
    links:
    - microservice  
```
@[10-15](nginx container)

---
#### docker-compose - build multi containers

```sh
$ docker-compose up -d
```

send requests to URLs below
* [http://localhost:8080](http://localhost:8080) or [http://docker-machine-ip:8080](http://docker-machine-ip:8080)
* [http://localhost:9090](http://localhost:9090) or [http://docker-machine-ip:9090](http://docker-machine-ip:9090)

```sh
$ doker-compose down
$ docker container ls -a
```

---
## 4 Introduction to Kubernetes and Go Microservice Orchestration

- Overview of Kubernetes architecture and main concepts
- Deploy a Go microservice to Kubernetes locally
- Implement Deployment and Service descriptors 
- Scale Deployments and perform Rolling updates

---
### 4.1 Overview of Kubernetes architecture and main concepts

- Basic architecture of Kubernetes cluster
- Key concepts and building blocks
- Setting up Kubernetes in the cloud and locally
- Introducing the Kubernetes CLI

---
#### Kubernetes is One of Several Cluster Operating Systems
![cluster-operating-systems](https://user-images.githubusercontent.com/5771924/34965743-b96394c6-fa99-11e7-814e-8b0eefa8f034.png)

---
![Basic Architecture of Kubernetes Cluster](https://i.ytimg.com/vi/fEZxrwNlLyM/maxresdefault.jpg)

---
![Key Concepts and Building Blocks](https://user-images.githubusercontent.com/5771924/34968800-b7f6f140-faae-11e7-963d-d3b60907c228.png)

---
#### Setup Kubernetes in the Cloud and Locally

- echo - The default provider is GCE
  - export KUBERNETES_PROVIDER=gce
  - export KUBE_GCE_ZONE=europe-west1-d
  - export NUM_NODES=4

- echo - Another possible provider is AWS
  - export KUBERNETES_PROVIDER=aws
  - export KUBE_AWS_ZONE=eu-central-1a
  - export NUM_SIZE=t2.small

- curl -sSL https://get.k8s.io | bash    

---
#### Working with minikube

```sh
$ minikube start
$ minikube status
$ kubectl
$ kubectl cluster-info
$ minikube dashboard
$ kubectl get nodes
$ kubectl config get-contexts
```
---
### 4.2 Deploy a Go microservice to Kubernetes locally

- Writing simple YAML descriptor for a pod
- Connecting to the pod through port forwarding
- Displaying and assigning labels to pods
- Using namespaces to organize pods
- Stopping and removing pods

#### Minikube - Docker Env

```
$ minikube docker-env
$ docker images
$ env | grep DOCKER
$ eval $(minikube docker-env)
$ env | grep DOCKER
$ docker images | grep cloud-native-go
# you can't see your cloud-native-go:1.0.1-alpine docker image 
```

---
```sh
$ cp -R chapter3_4 chapter4_2
$ cd chapter4_2
$ docker build -t cloud-native-go:1.0.1-alpine .

#$ docker images | grep cloud-native-go
#$ docker login
#$ docker tag cloud-native-go:1.0.1-alpine credemol/cloud-native-go:1.0.1-alpine
#$ docker images | grep cloud-native-go
#$ docker push credemol/cloud-native-go:1.0.1-alpine
```
> It is necessary to push docker image before you create a pod.
---
#### Writing simple YAML descriptor for a pod

chapter4_2/k8s-pod.yml

```yaml
apiVersion: v1
kind: Pod
metadata: 
  name: cloud-native-go
  labels:
    language: go
    env: test
spec:
  containers:
  - name: cloud-native-go
    image: cloud-native-go:1.0.1
    ports:
    - containerPort: 8080
      protocol: TCP
```

---
```sh
$ kubectl create -f k8s-pod.yml
$ kubectl get pods
$ kubectl describe pod cloud-native-go
$ kubectl port-forward cloud-native-go 8080:8080
```

---
#### work with labels 

```sh
# Ctrl - C to stop port-forwarding
$ kubectl get pods --show-labels
$ kubectl get pods -o wide -L env,language
$ kubectl label pod cloud-native-go hello=world
$ kubectl get pods --show-labels
$ kubectl label pod cloud-native-go env=prod --overwrite
$ kubectl get pods --show-labels
$ kubectl label pod cloud-native-go hello-
$ kubectl get pods --show-labels
```

---
#### work with namespaces

```sh
$ kubectl get ns
$ kubectl get pods --namespaces kube-system
$ kubectl vi k8s-namespace.yml
```

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: cloud-native-go
```

---

```sh
$ kubectl create -f k8s-namespace.yml
$ kubectl get ns

$ kubectl create -f k8s-pod.yml --namespace cloud-native-go
$ kubectl get pods
$ kubectl get pods --namespace cloud-native-go

# following two commands are same.
$ kubectl delete ns cloud-native-go 
$ kubectl delete -f k8s-namespace.yml 

# all resources in cloud-native-go namespace have been deleted
$ kubectl get pods --namespace cloud-native-go

# delete cloud-native-go pod 
$ kubectl delete pod cloud-native-go
```

---
### 4.3 Implement Deployment and Service descriptors 

- Writing simple YAML descriptor for a deployment
- Assigning CPU and RAM resources to a container
- Writing simple YAML descriptor for a service
- Adding liveness and readiness probes
- Connecting to the service through a node port

---
#### Initialize Project

```sh
$ cp -R chapter4_2 chapter4_3
$ cd chapter4_3
$ vi k8s-deployment.yaml

```

---
#### k8s-deployment.yaml

```sh
apiVersion: extensions/v1beta1
kind: Deployment
metadata: 
  name: cloud-native-go
  labels:
    app: cloud-native-go
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: cloud-native-go
        tier: service
    spec:
      containers:
      - name: cloud-native-go
        image: "cloud-native-go:1.0.1-alpine"
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
```

---

```sh
$ kubectl get deployments,pods,rs

$ kubectl create -f k8s-deployments.yaml
$ kubectl get deployments,pods,rs
```

---
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata: 
  name: cloud-native-go
  labels:
    app: cloud-native-go
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: cloud-native-go
        tier: service
    spec:
      containers:
      - name: cloud-native-go
        image: "cloud-native-go:1.0.1-alpine"
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        resources:
          requests:
            memory: "64Mi"
            cpu: "125m"
          limits:
            memory: "128Mi"
            cpu: "250m"
```
@[23-29](Managing Compute Resources for Container)
---
[https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/)

---
#### apply deployment

```sh
$ kubectl apply -f k8s-deployment.yaml

$ kubectl describe deployment cloud-native-go
```

---
#### service

```sh
$ vi k8s-service.yml
```

---
#### k8s-service.yml

```yaml
apiVersion: v1
kind: Service
metadata: 
  name: cloud-native-go
  labels:
    app: cloud-native-go
    tier: service
spec:
  type: NodePort
  ports:
  - port: 8080
  selector:
    app: cloud-native-go
```

---
#### create service

```sh
$ kubectl create -f k8s-service.yml
$ kubectl get svc 
$ kubectl describe svc cloud-native-go

$ minikube service cloud-native-go --url
$ curl $(minikube service cloud-native-go --url)
```

---
#### Adding liveness and readiness probes (k8s-deployment.yaml)

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata: 
  name: cloud-native-go
  labels:
    app: cloud-native-go
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: cloud-native-go
        tier: service
    spec:
      containers:
      - name: cloud-native-go
        image: "cloud-native-go:1.0.1-alpine"
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        resources:
          requests:
            memory: "64Mi"
            cpu: "125m"
          limits:
            memory: "128Mi"
            cpu: "250m"
        readinessProbe:
          httpGet:
            path: /
            port: 8080
          initialDelaySeconds: 5
          timeoutSeconds: 5
        livenessProbe:
          httpGet: 
            path: /
            port: 8080
          initialDelaySeconds: 5
          timeoutSeconds: 5
```
@[30-41] (readinessProbe and livenessProbe)

---
#### apply deployment

```sh
$ kubectl apply -f k8s-deployment.yaml

$ kubectl describe deployment cloud-native-go
$ kubectl get pods
```

---
### 4.4 Scale Deployments and perform Rolling updates

- Scaling Kubernetes deployments
- Performing rolling updates of Kubernetes deployment

---
#### Initialize project

```sh
$ cd $GOPATH/src/GettingStartedWithCloudNativeGo
$ cp -R chapter4_3 chapter4_4
$ cd chapter4_4
$ kubectl delete all -l app=cloud-native-go

```

---
#### Scaling Kubernetes deployments (k8s-deployment.yaml)

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata: 
  name: cloud-native-go
  labels:
    app: cloud-native-go
spec:
  replicas: 2
  revisionHistoryLimit: 10
  minReadySeconds: 5
  strategy: 
    type: RollingUpdate
    rollingUpdate: 
      maxUnavailable: 1
      maxSurge: 2
  template:
    metadata:
      labels:
        app: cloud-native-go
        tier: service
    spec:
      containers:
      - name: cloud-native-go
        image: "cloud-native-go:1.0.1-alpine"
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        resources:
          requests:
            memory: "64Mi"
            cpu: "125m"
          limits:
            memory: "128Mi"
            cpu: "250m"
        readinessProbe:
          httpGet:
            path: /
            port: 8080
          initialDelaySeconds: 5
          timeoutSeconds: 5
        livenessProbe:
          httpGet: 
            path: /
            port: 8080
          initialDelaySeconds: 5
          timeoutSeconds: 5
```
@[9-15](Rolling Update)

---
#### get pods with watch mode 

Terminal 1
```sh
$ kubectl get pods -w
```

Terminal 2
```sh
$ kubectl create -f k8s-deployment.yaml --record=true
$ kubectl get pods
$ kubectl scale deployment cloud-native-go --repllicas=5
$ kubectl get pods
$ kubectl scale deployment cloud-native-go --repllicas=3
$ kubectl get pods
$ kubectl rollout history deployment cloud-native-go
```

---
#### Version Upgrade.

```sh
$ kubectl create -f k8s-service.yml
$ curl $(minikube service cloud-native-go --url)

$ minikube docker-env
$ eval $(minikube docker-env)

$ vi microservice.go 
```

---
#### microservice.go

```go
import (
	"GettingStartedWithCloudNativeGo/chapter4_4/api"
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go.(Update)")
}
```

---
#### build and run

```sh
$ rm Cloud-Native-Go
$ GOOS=linux GOARCH=amd64 go build -o Cloud-Native-Go
$ ./Cloud-Native-Go

# Run curl command in a new terminal
$ curl http://localhost:8080  

# Ctrl-C
$ docker image ls | grep cloud-native-go
$ docker build -t cloud-native-go:1.0.2-alpine .
$ docker image ls | grep cloud-native-go


```

---
#### Rolling update And Undo

```sh
$ curl $(minikube service cloud-native-go --url)

$ kubectl set image deployment cloud-native-go cloud-native-go=cloud-native-go:1.0.2-alpine
$ curl $(minikube service cloud-native-go --url)

$ kubectl rollout undo deployment cloud-native-go
$ curl $(minikube service cloud-native-go --url)
```

----
#### Performing rolling updates of Kubernetes deployment

---
## Resources

* [https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028](https://www.slideshare.net/QAware/a-hitchhikers-guide-to-the-cloud-native-stack-77181028)
