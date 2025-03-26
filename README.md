# CRUD WITH GOLANG

Practicing with golang standard library & mux for web development.

Copy / Extract all file, and then run the command prompt on files directory `go run *` or `go run main.go`

## STATIC HTML

`HOMEPAGE` http://localhost:8080/

## GET ALL MOVIE

`GET` http://localhost:8080/movies

## GET MOVIE BY ID

`GET` http://localhost:5001/movie/{id}

## CREATE NEW MOVIE

`POST` http://localhost:5001/movies

### Body

```JSON
{
    "title": string,
    "isbn": string,
    "director" : {
        "firstname": string,
        "lastname": string,
    }
}
```

## UPDATE A MOVIE

`PUT` http://localhost:5001/movies/{id}

```
{
    "title": string,
    "isbn": string,
    "firstname": string,
    "lastname": string
}
```

## DELETE A MOVIE

`DELETE` http://localhost:8080/movies/{id}
