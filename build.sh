#/bin/bash
(git pull && CGO_ENABLED=0 go build -o gm65server && echo "build successfull") || echo "build failed"
