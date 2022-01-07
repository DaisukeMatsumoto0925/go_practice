## Create Table
```
aws dynamodb create-table --endpoint-url http://localhost:8000 --table-name MyFirstTable --attribute-definitions AttributeName=MyHashKey,AttributeType=S AttributeName=MyRangeKey,AttributeType=N --key-schema AttributeName=MyHashKey,KeyType=HASH AttributeName=MyRangeKey,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
```

### list table
```
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

### run local container
```
docker run -d --name dynamodb -p 8000:8000 amazon/dynamodb-local  -jar DynamoDBLocal.jar -inMemory -sharedDb
```
