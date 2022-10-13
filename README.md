# sulifa

# run

```sh
make run
```

Auth section:

Send post request to register endpoint(body should contain data: username, password)

Server will add user to "user" collection, not login yet(not authorized)

then:

Send post request to login endpoint(body should contain data: username, password)

Server will add user to "authorized" collection means that register and login stages done
Response: return user struct


# websockets


