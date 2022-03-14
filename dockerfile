FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 4271:4271
RUN go run .