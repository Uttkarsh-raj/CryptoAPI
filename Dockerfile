FROM golang:1.20.5-alpine

ENV MONGODB_URL = your_url_here

ENV CLUSTER = your_cluster_name_here

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
