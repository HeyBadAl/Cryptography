FROM golang:latest AS build 

WORKDIR /app 

COPY go.mod  ./
COPY . .

RUN go build -o myapp 


# Create the final image
FROM golang:latest

WORKDIR /app

COPY --from=build /app/myapp .

# EXPOSE 8080

CMD ["./myapp"]
