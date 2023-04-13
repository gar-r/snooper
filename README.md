# Snooper

Logs the content of an uploaded file.

## Starting the server

```
docker build -t snooper .
docker run -d -p 8080:8080 -e GIN_MODE=release --restart always -v $(pwd)/snooper-data:/snooper/uploads --name snooper snooper
```

## Using the service

POST a file:

```
curl -F upload=@local-file http://localhost:8080/
```

GET the uploaded file:

```
curl http://localhost:8080/local-file
```

...or download the attachment as a local file:

```
curl -O -J http://localhost:8080/local-file
```