FROM golang:stretch AS build-env

WORKDIR /go/src/github.com/elysium-station/blackfury

RUN apt update
RUN apt install git -y

COPY . .

RUN make build

FROM golang:stretch

RUN apt update
RUN apt install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/elysium-station/blackfury/build/blackfuryd /usr/bin/blackfuryd

EXPOSE 26656 26657 1317 9090

CMD ["blackfuryd"]