FROM golang:1.22
# set working directory for dockerfile to /app
WORKDIR /app
# copy dependency management files into container
COPY go.mod go.sum ./
# install dependencies
RUN go mod download # && go mod verify
# copy code
COPY . .
# build image to /api-server
RUN go build -v -o /api-server
# start server
CMD ["/api-server"]

