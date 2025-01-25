# Test basic api usage
GET http://localhost:8080/api/echo/SE

HTTP 200
[Asserts]
body == "Hello, SE"
