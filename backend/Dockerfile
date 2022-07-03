FROM golang:1.18-alpine as develop-stage-backend

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /capstone

EXPOSE 3005

CMD [ "/capstone" ]
