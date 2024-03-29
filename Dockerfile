FROM golang:1.21-alpine
WORKDIR /src
COPY go.mod go.sum  ./
COPY main.go  ./
# COPY controllers db graph dbModels helpers rssfeed tools schedule eventHandler ./
COPY controllers ./controllers
COPY db ./db
COPY graph ./graph
COPY dbModels ./dbModels
COPY helpers ./helpers
COPY rssfeed ./rssfeed
COPY tools ./tools
COPY schedule ./schedule
COPY eventHandler ./eventHandler
RUN go build -o /bin/crispypod

FROM alpine:latest
WORKDIR /crispypod
VOLUME [ "/crispypod/UploadFile" ]
COPY --from=0 /bin/crispypod /bin/crispypod
ENV GIN_MODE=release
CMD [ "crispypod" ]