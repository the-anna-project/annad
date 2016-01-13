# api
The api documentation describes all interfaces.

### network

###### http
HTTP is a rather legacy tainted protocol. It does not fit all business
requirements of API's. Deciding when to use what request method or
response status code is a pain we simply can prevent. For simplicity reasons we
go with a very straight forward solution. See the details below.

---

###### versioning
All network API's will provide one single version. This is because of
simplicity. Versioning and API changes will be handled internally anyway.
Clients don't need to struggle with that. API's are cleaner.

---

###### request method
All network API's are only accepting `POST` requests. This is because of
simplicity. Clients don't need to struggle with that.

---

###### response code
All HTTP responses will have one of the following HTTP status codes.
```
200
500
```

As long as the system operates as expected a `200` is returned. This is also
the case if a bad request was made. The returned data format is something like
JSON in all cases, besides critical system errors. `500` is only returned if
the system is not properly able to answer. Here only plain text is returned.

---

###### response body
All network responses will have the following body. Fields are maybe omitted
when they are empty.
```
{
  "code": "<code>",
  "text": "<text>"
}
```
