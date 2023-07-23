## bash

## Creating Table

aws dynamodb create-table \
    --table-name User \
    --attribute-definitions \
        AttributeName=Name,AttributeType=S \
        AttributeName=Email,AttributeType=S \
    --key-schema \
        AttributeName=Name,KeyType=HASH \
        AttributeName=Email,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

+------------+---------------------+
| Partition  | Sort                |
| Key (Name) | Key (Email)         |
+------------+---------------------+
|            |                     |
|            |                     |
|            |                     |
+------------+---------------------+


## aws cli 

aws dynamodb describe-table --table-name User | grep TableStatus

ç