# FROM golang:latest as dev

# WORKDIR /work

FROM golang:latest as build

WORKDIR /videos
COPY . /videos/
RUN go build -o videos
CMD [ "./videos" ]
