FROM golang:1.17-alpine3.15

# Set destination for COPY
WORKDIR /app
COPY . ./

# Download Go modules and build
RUN apk --no-cache add build-base \
    && go mod download \
    && go build -o /go-invoices \
    && rm -rf /go \
    && apk del build-base 

EXPOSE 5002

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Tell docker that all future commands should run as the appuser user
USER appuser

# Run
CMD [ "/go-invoices" ]
