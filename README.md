## URL Shortener

This is a simple URL shortener service written in Golang, using Postgres for data storage and the Gin web framework for HTTP handling. The service allows users to create a shortened URL for any long URL they provide and redirect to the original URL by accessing the shortened URL.

## API Endpoints

The API provides the following endpoints:

### Redirect

```
GET /{token}
```

This endpoint retrieving the original URL from the database and redirecting the user to that URL. If the token is not found in the database, a 404 Not Found response is returned.

#### Example
```sh
curl -X GET 127.0.0.1:8888/6d0d7b
```

### Create Short URL

```
POST /shorten
```

This endpoint create a new shortened URL by taking the target_url parameter from the request body, generating a unique token, and inserting the URL and token into the database. If the insertion is successful, a 201 Created response is returned with the token in the response body. If there is an error, a 500 Internal Server Error response is returned with an error message in the response body.

#### Example
```sh
curl -X POST 127.0.0.1:8888/shorten -d '{"TargetURL": "google.com"}'
```
