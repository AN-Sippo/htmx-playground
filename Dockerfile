FROM golang
WORKDIR .
COPY . .
RUN go mod download 
RUN go build -o /out
EXPOSE 8080
CMD ["/out"]