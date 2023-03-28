FROM golang:1.20
WORKDIR /app
COPY ./ ./
RUN go mod download
EXPOSE 8080
ENV LISTEN_PORT=:8080
ENV API_URL_CURRENT=https://api.openweathermap.org/data/2.5/weather?q=%s&appid=fbf18251cec73bc7b51fcf91b9c2abe7&units=metric
ENV API_URL_FORECAST=https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=fbf18251cec73bc7b51fcf91b9c2abe7&units=metric
RUN go build ./cmd/storage/main.go
CMD ["./main"]
