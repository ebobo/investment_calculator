# github.com/ebobo/investment_calculator

Investment calculator go backend

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# Build Docker image

build
docker build -f docker/server.Dockerfile -t ic-server:0.1.0 .
