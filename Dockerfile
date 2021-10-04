FROM golang:1.17 AS builder
COPY . /workspace
WORKDIR /workspace
RUN make deps
RUN make cross