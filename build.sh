#/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
(git pull && CGO_ENABLED=0 go build -o ptapiserver && echo -e "\n${GREEN}build successfull${NC}\n") || echo -e "\n${RED}build failed${NC}\n"
