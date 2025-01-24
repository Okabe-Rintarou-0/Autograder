# Test basic api usage
GET http://localhost:8080/api/echo/SE

HTTP 200
[Asserts]
body == "Hello, SE"

# Test database initialization
POST http://localhost:5000/execute
{
    "database": "ebookstore",
    "sql": "select * from users where users.username like 'admin'"
}
HTTP 200
[Asserts]
jsonpath "$.result[0].username" == "admin"
