# sandbox-go-fileserver

## building

```
go build
./fileserver
```

## usage

### Receiving static files:

```
> curl --request GET localhost:8080/
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Sandbox file server</title>
  </head>
  <body>
    <h1>Hello from the golang sandbox file server!</h1>
  </body>
</html>

> curl --request GET localhost:8080/test/test
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Sandbox file server</title>
  </head>
  <body>
    <h1>TEST!</h1>
  </body>
</html>

```

### Manipulating files on the fileserver:

File routes can only be accessed via basic auth. One user is hardcoded in `constans.go`.
Add `--user username:123123` to each curl request to avoid `authorization failed` error.

#### POST
POST request to create file:

```
curl -d '{"contents":"This content can now be found in ./files/data file."}' -H "Content-Type: application/json" -X POST http://localhost:8080/files/data --user username:123123
```
When running this command twice without changing the filename following error message will be returned:
```
curl -d '{"contents":"42"}' -H "Content-Type: application/json" -X POST http://localhost:8080/files/data --user username:123123
file already exists
```

#### GET
GET request to get file contents:

```
curl -X GET http://localhost:8080/files/data --user username:123123                                                                                           
This content can now be found in ./files/data file.
```

If there is no such file on the server, respective message would be returned:
```
curl -X GET http://localhost:8080/files/foobar --user username:123123
file does not exist
```

### PATCH
PATCH request to update existing file:
```
curl -d '{"contents":"Ha! ./files/data file is now totaly patched."}' -H "Content-Type: application/json" -X PATCH http://localhost:8080/files/data --user username:123123

curl -X GET http://localhost:8080/files/data --user username:123123                                                                                    
Ha! ./files/data file is now totaly patched.
```
If there is no file to be patched, error will be returned:
```
curl -d '{"contents":"Attempting to patch"}' -H "Content-Type: application/json" -X PATCH http://localhost:8080/files/foobar --user username:123123
file does not exist
```
