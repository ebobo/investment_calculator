# github.com/ebobo/investment_calulator_record

Investment calculator record save micro-service

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image

build
docker build -f docker/record.Dockerfile -t record-ms:0.1.0 .

run
docker run -it --name record-ms record-ms:0.1.0

stop
docker stop record-ms

check container
docker inspect record-ms | grep IPAddress

clean builder layer image
docker image prune --filter label=stage=builder
