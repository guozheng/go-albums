# A demo album management API using Gin

The code is based on Golang tutorial [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin).

## APIs

/albums
   * GET – Get a list of all albums, returned as JSON.
   * POST – Add a new album from request data sent as JSON.

/albums/:id
   * GET – Get an album by its ID, returning the album data as JSON.

## Usage

Start the service:
```
> go run .
```

Get all the albums:
```
> curl http://localhost:8080/albums
```

Get an album by id:
```
> curl http://localhost:8080/albums/2
```

Add a new album:
```
> curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## References
   * [Gin Web Framework](https://gin-gonic.com/docs/)