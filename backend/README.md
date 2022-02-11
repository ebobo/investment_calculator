# github.com/ebobo/investment_calculator

Investment calculator go backend

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image / push image to docker hub

build
docker build -f docker/server.Dockerfile -t ic-server:0.1.1 .

build for docker hub
docker build -f docker/server.Dockerfile -t xbobo/investment_calculator:server_1.1 .

tag image
docker tag ic-server:0.1.1 xbobo/investment_calculator:server_1.1

push image to docker hub
docker push xbobo/investment_calculator:server_1.1

run
docker run -it --name ic-server -p9090:9090 ic-server:0.1.1

run (pull image from docker hub)
docker run -d --name ic-server -p9090:9090 xbobo/investment_calculator:server_1.1

stop
docker stop ic-server

check container
docker inspect ic-server | grep IPAddress

clean builder layer image
docker image prune --filter label=stage=builder
