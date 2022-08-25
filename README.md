# Snooper

Logs the content of an uploaded file.

## Starting the server

```
docker build -t snooper .
docker run -d -p 8080:8080 snooper
```

## Using the service

```
curl -F upload=@/path/to/local/file http://localhost:8080/
```