FROM golang:1.16.3-alpine3.13 as build
RUN mkdir /app
COPY ./src/main.go /app
WORKDIR /app
RUN go mod init github.com/evilsnow/otus_lesson && go get -u github.com/gorilla/mux && go build /app/main.go

FROM alpine:3.13.5 as prod
RUN mkdir /app
COPY --from=build /app/main /app
EXPOSE 8000
CMD ["/app/main"]