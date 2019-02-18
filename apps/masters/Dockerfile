
FROM golang:latest 

# Protos

# App files
COPY apps/masters apps/masters

# Lint
RUN go tool vet apps

# Build app
RUN go install ./apps/masters/cmd/masters

FROM launcher.gcr.io/google/debian9:latest

RUN apt-get update && \
    apt-get install -y ca-certificates

WORKDIR /opt/ac
COPY /go/bin/masters .

CMD ["/opt/ac/masters"]
