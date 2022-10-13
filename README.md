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


Send post request to createroom endpoint(body should contain ID who want to create a room)

Server will add new room and player who created the room to "rooms" collection and set up status: waiting (means need one more player)
Response: return room id

Send post request to ощшткщщь endpoint(body should contain room ID and username who want to join a room)

Server will check room if empty new and player who joined the room will be added to player2 "rooms" collection and set up status: 1 (means room is full)