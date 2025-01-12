FROM golang:1.22-bullseye AS build

WORKDIR /
COPY . .

RUN go mod download

RUN go build -o /eventT7


FROM debian:bullseye-slim
LABEL maintainer="duy.nguyen17@mservice.com.vn"

RUN apt update

COPY . .
COPY --from=build /eventT7 /eventT7

EXPOSE 8001

CMD [ "/eventT7" ]