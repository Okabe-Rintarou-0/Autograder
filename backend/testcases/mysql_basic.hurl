# Test database initialization
POST http://localhost:5000/db/execute
{
    "database": "ebookstore",
    "sql": "select * from users where users.username like 'admin'"
}
HTTP 200
[Asserts]
jsonpath "$.result[0].username" == "admin"
