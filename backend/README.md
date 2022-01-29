# github.com/ebobo/investment_calculator

Investment calculator go backend

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image

build
docker build -f docker/server.Dockerfile -t ic-server:0.1.1 .

run
docker run -it --name ic-server -p9090:9090 ic-server:0.1.1

stop
docker stop ic-server

check container
docker inspect ic-server | grep IPAddress
