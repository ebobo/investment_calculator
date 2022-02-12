# github.com/ebobo/investment_calulator_record

Investment calculator record save micro-service

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image / push image to docker hub

build
docker build -f docker/record.Dockerfile -t record-ms:0.1.0 .

build for docker hub
docker build -f docker/record.Dockerfile -t xbobo/investment_calculator:record_1.0 .

tag image
docker tag record-ms:0.1.0 xbobo/investment_calculator:record_1.0

push image to docker hub
docker push xbobo/investment_calculator:record_1.0

run
docker run -it --name record-ms record-ms:0.1.0

run (with environment variable local docker)
docker run -it --name record-ms -e IC_GRPC_ADDR=172.17.0.2:9092 -p9094:9094 xbobo/investment_calculator:record_1.0

stop
docker stop record-ms

check container
docker inspect record-ms | grep IPAddress

clean builder layer image
docker image prune --filter label=stage=builder
