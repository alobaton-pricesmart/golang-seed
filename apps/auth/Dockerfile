
FROM golang:latest 

# Protos

# App files
COPY apps/auth apps/auth

# Lint
RUN go tool vet apps

# Build app
RUN go install ./apps/auth/cmd/auth

FROM launcher.gcr.io/google/debian9:latest

RUN apt-get update && \
    apt-get install -y ca-certificates

WORKDIR /opt/ac
COPY /go/bin/auth .

CMD ["/opt/ac/auth"]
