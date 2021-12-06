# chess_move_go
chess valid move (only for knight move)

go run main

## if need change default port
export CHESS_PORT=8080

## to get a valid moves list
curl http://localhost:8080/knight/c2

## to list all table positions
curl http://localhost:8080/knight/

## save a position (if not exists)
curl -X POST -H 'content-type: application/json' --data '{"id":"a2","position":"a2"}'  http://localhost:8080/knight/


## to run in docker
...
