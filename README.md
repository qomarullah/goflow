#### Synopsis

GoFlow is platform for building custom ordering flow based on predefined format XML and JSON

#### What Already and Not Yet
 1. Read and parse JSON
 2. Call function by name from json 
 3. Define If , break, loop function => not yet
 4. Read and parse XML => not yet
 5. Define more function => not yet
 - db
 - http

#### Installation and Development

1. install golang
2. clone this project
3. go get github.com/astaxie/beego 
4. update config in 'conf/App.conf'
5. run bee run -downdoc=true -gendoc=true 
    or bee run
    or go run godb.go
6. test API http://localhost:8080/swagger/
7. dashboard & monitoring http://localhost:8088

#### API Reference
- https://beego.me
- https://github.com/elgs/gosqljson


#### Tests

Describe and show how to run the tests with code examples.

#### Contributors

Let people know how they can dive into the project, include important links to things like issue trackers, irc, twitter accounts if applicable.

#### License

A short snippet describing the license (MIT, Apache, etc.)
