## 🍩 go module init 

1. init module
```shell
go mod init top.lel/ginapp01
```
2. create ginapp01.go
3. write your code!
4. go modules config
```markdown
GO111MODULE=on
```

## 🌾 project describe
1. this is a project based on gin and gorm...
2. a simple demo use MySQL.
3. write an easy date fmt struct in `tool.Date`.
4. do a test.
- just save a user.
- [http://localhost:8080/v1/users/addUser](http://localhost:8080/v1/users/addUser "post a user")
```json
{
	"name": "康娜",
	"age": 12,
	"birthday": "2019-01-08T11:00:00Z",
	"email": "99999@gmail.com",
	"activiedAt": "2019-01-08 11:00:00"
}
```