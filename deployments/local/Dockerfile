FROM golang:1.24-bullseye

ENV CGO_ENABLED=1

RUN apt-get update && apt-get install -y gcc build-essential ca-certificates git-core ssh

# Add the private key arg
ARG SSH_PRIVATE_KEY
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

# Make the root foler for our ssh 
RUN mkdir -p /root/.ssh && \
chmod 0700 /root/.ssh && \
ssh-keyscan github.com > /root/.ssh/known_hosts

WORKDIR /home/app

COPY . .

RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
chmod 600 /root/.ssh/id_rsa && \
go mod download && \
rm -rf /root/.ssh/

# RUN go mod tidy

# Builds the consumer binary
RUN go build -o hotel-service cmd/api/main.go

CMD ["./hotel-service"]