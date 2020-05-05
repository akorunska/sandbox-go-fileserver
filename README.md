# sandbox-go-fileserver

### building

```
go build
./fileserver
```

### usage

Receiving static files:

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

Manipulating files on the fileserver:

POST request to create file:

```
curl -d '{"contents":"This content can now be found in ./files/data file."}' -H "Content-Type: application/json" -X POST http://localhost:8080/files/data
```

GET request to get file contents:

```
curl -X GET http://localhost:8080/files/data
```
