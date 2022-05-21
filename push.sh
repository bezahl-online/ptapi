#/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
git pull 
REPRO="www.greisslomat.at:3307"
version=$(cat version)
PRG="ptapi"
echo "building version $(arch)_${version}"
export PTAPI_VERSION=$(arch)_$(cat version) # for docker-compose
docker login https://www.greisslomat.at:3307 --username ralph --password natural-Kennwort
(./build.sh && docker-compose build && docker push $REPRO/$PRG:$(arch)_${version} && echo -e "\n${GREEN}build successfull${NC}\n") || echo -e "\n${RED}build failed${NC}\n"
if [ "${1}" == "latest" ]; then
    docker tag $REPRO/$PRG:$(arch)_${version} $REPRO/$PRG:$(arch)_latest
    docker push $REPRO/$PRG:x86_64_latest
fi