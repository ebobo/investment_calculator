# github.com/ebobo/investment_calculator

Investment calculator go backend

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image / push image to docker hub

build
docker build -f docker/server.Dockerfile -t ic-server:0.1.1 .

build for docker hub
docker build -f docker/server.Dockerfile -t xbobo/investment_calculator:server_1.1 .

build for linux amd64 from m1 mac
docker build --platform=linux/amd64 -f docker/server.Dockerfile -t xbobo/investment_calculator:server_1.1 .

tag image
docker tag ic-server:0.1.1 xbobo/investment_calculator:server_1.1

push image to docker hub
docker push xbobo/investment_calculator:server_1.1

run
docker run -it --name ic-server -p9090:9090 ic-server:0.1.1

run (pull image from docker hub)
docker run -d --name ic-server -p9090:9090 xbobo/investment_calculator:server_1.1

run (with environment variable)
docker run -it -e MS_GRPC_ADDR=172.17.0.3:9094 --name ic-server -p9090:9090 -p9092:9092 ic-server:0.1.1

stop
docker stop ic-server

check container
docker inspect ic-server | grep IPAddress

clean builder layer image
docker image prune --filter label=stage=builder

# Debug Methods

if exec to container is needed
"CMD tail -f /dev/null" --makes the container not quit
then use "docker exec -it ic-server bash" to use bash in container

"bash: "**_" No such file or directory" error -- means that either the executable binary itself or one of the libraries it needs does not exist
use "ldd _**" to print shared library dependencies

eg: "linux-vdso.so.1 (0x00007ffcf6559000) libc.musl-x86_64.so.1 => not found"

use "apt-get install musl-dev" to install libc.musl-x86_64.so.1
