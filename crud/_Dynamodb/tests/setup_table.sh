## Creating Table
local_dynamodb="http://localhost:8000"

echo "Creating Table in local dynamodb: ${local_dynamodb}"

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
    --endpoint-url "${local_dynamodb}"

echo "Tables: "

aws dynamodb list-tables --endpoint-url "${local_dynamodb}"