FROM golang:1.21-alpine
WORKDIR /src
COPY main.go go.mod go.sum  ./
COPY migrations ./migrations
COPY helpers ./helpers
COPY rssFeed ./rssFeed
RUN go build -o /bin/crispypod


FROM alpine:latest

# ARG PB_VERSION=0.21.3

WORKDIR /crispypod
COPY --from=0 /bin/crispypod ./

# RUN apk add --no-cache \
#     unzip \
#     ca-certificates

# # download and unzip PocketBase
# ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
# RUN unzip /tmp/pb.zip -d /pb/

# uncomment to copy the local pb_migrations dir into the image
# COPY ./pb_migrations /pb/pb_migrations

# uncomment to copy the local pb_hooks dir into the image
# COPY ./pb_hooks /pb/pb_hooks
ENV DISABLE_PB_WEBUI=1

EXPOSE 8080

# start PocketBase
CMD ["/crispypod/crispypod", "serve", "--http=0.0.0.0:8080"]