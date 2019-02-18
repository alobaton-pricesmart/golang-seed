
FROM golang:latest 

# Protos

# App files
COPY apps/settings apps/settings

# Lint
RUN go tool vet apps

# Build app
RUN go install ./apps/settings/cmd/settings

FROM launcher.gcr.io/google/debian9:latest

RUN apt-get update && \
    apt-get install -y ca-certificates

WORKDIR /opt/ac
COPY /go/bin/settings .

CMD ["/opt/ac/settings"]
