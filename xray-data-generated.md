
# root
- `curl localhost:9001`

# ping
- `curl localhost:9001/ping`

# new 
- `curl -X POST localhost:9001/new`

# get all
- `curl -X GET localhost:9001/all`

# del 
- `curl -X DELETE localhost:9001/del -d '{"id":1}'`

# err 400 
- `curl -X GET localhost:9001/error/400`

# err panic
- `curl -X GET localhost:9001/error/panic`