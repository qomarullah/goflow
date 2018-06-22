#### Synopsis

GoDB provides convenient API to query from any DB and with custom SQL and parameter, built with golang (beego)

#### API Example
  ```
  curl -X GET "http://localhost:8080/v1/query/select?ds=localhost&sqlid=user" -H "accept: application/json"
  ```
  ##### Configuration App.conf
  ```
  #ds
  ds.localhost= "apps:aplikasi@tcp(localhost:3306)/test|mysql|10|10|120000"
  sqlid.user="SELECT * from `user` limit 10"
  ```
  
  ```
  result:
 {
  "count": 10,
  "data": [
    {
      "email": "slene",
      "id": "1",
      "name": "testing"
    },
    {
      "email": "someemail@someemailprovider.com",
      "id": "2",
      "name": "First"
    },
    {
      "email": "someemail@someemailprovider.com",
      "id": "3",
      "name": "First"
    }
    ....
  ],
  "desc": "-",
  "status": false
}
```
```
curl -X GET "http://localhost:8080/v1/query/select?ds=localhost2&sqlid=user2&id=1" -H "accept: application/json"
```
#### Configuration App.conf
```
ds.localhost2= "apps:aplikasi@tcp(localhost:3306)/mfs|mysql|10|10|120000"
sqlid.user2="SELECT * from `user` where id=[id]"
```
```
Result
{
  "count": 1,
  "data": [
    {
      "email": "slene",
      "id": "1",
      "name": "testing"
    }
  ],
  "desc": "-",
  "success": true
}
```

#### Installation or Development

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
