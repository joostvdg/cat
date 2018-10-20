FROM golang:1.11 AS build
WORKDIR /src
ENV LAST_UPDATE=20181020
COPY . /src
RUN go get -d -v -t
# Yes, shame on me
# TODO: write tests and enable
# RUN go test --cover ./...
RUN go build -v -tags netgo -o cat

FROM alpine:3.8
ENV LAST_UPDATE=20180921
LABEL authors="Joost van der Griendt <joostvdg@gmail.com>"
LABEL version="0.1.0"
LABEL description="Docker image for CAT"
CMD ["cat"]
COPY --from=build /src/cat /usr/local/bin/cat
RUN chmod +x /usr/local/bin/cat
