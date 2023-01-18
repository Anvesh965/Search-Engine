FROM golang:alpine

WORKDIR /projects/data/

COPY go.mod go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o se.exe ./cmd/

EXPOSE 4000

CMD [ "./se.exe" ]




