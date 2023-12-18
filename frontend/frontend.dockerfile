# build a tiny Docker image for running the application
FROM alpine:latest

# create a directory /app in the container
RUN mkdir /app

# copy the compiled binary from the builder stage to /app in the current container
COPY frontApp /app
COPY cmd/web/templates /app/templates

# specify the command to run when the container starts
CMD ["/app/frontApp"]
