# !/bin/bash
echo "build binary..."
CGO_ENABLED=0 go build -o alerting-ks-client  ../cmd/main.go
echo "Building images..."
docker build -t carmanzhang/alerting-ks-client:latest -f ./Dockerfile.dev .
echo "Built successfully"
docker push carmanzhang/alerting-ks-client:latest
echo "Push successfully"
