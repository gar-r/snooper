# Snooper

Logs the content of an uploaded file.

## Starting the server

```
docker build -t snooper .
docker run -d -p 8080:8080 -e GIN_MODE=release --restart always -v $(pwd)/snooper-data:/uploads --name snooper snooper
```

## Using the service

```
curl -F upload=@/path/to/local/file http://localhost:8080/
```