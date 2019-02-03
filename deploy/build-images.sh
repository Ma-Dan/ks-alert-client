# !/bin/bash
echo "build binary..."
CGO_ENABLED=0 go build -o alerting-ks-client  ../cmd/main.go
echo "Building images..."
cp ../pkg/swagger-ui . -r
docker build -t alerting-ks-client -f ./Dockerfile.dev .
echo "Built successfully"
#docker push carmanzhang/alerting-ks-client:latest
#echo "Push successfully"
