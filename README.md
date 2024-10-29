# AWS Local (using localstack)

## Run:

```
docker compose up -d
```

## Provision Infrsatructure

```
mv .envrc-example .envrc
vim .envrc (fill in details)
cd terraform/localstack
direnv allow
terraform plan
terraform apply
cd ../../
```

## Put something in s3 bucket

```
aws --endpoint-url=http://localhost:4566 s3 cp go.mod s3://YOUR_BUCKET_NAME
```

## Retrieve using Golang script

```
go get .
go run cmd/s3/main.go
```

## Put a message in the queue

```
go run cmd/sqs/main.go
```

## Retrieve message using CLI

```
aws --endpoint-url=http://localhost:4566 sqs list-queues (copy queue URL)
awslocal sqs receive-message --queue-url QUEUE_URL_FROM_ABOVE
```