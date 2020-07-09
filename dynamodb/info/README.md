# Welcome to DynamoDB!

This is for dynamodb cli references

## Useful commands

### Create Table
- `aws dynamodb create-table --table-name MusicCollection --attribute-definitions AttributeName=Artist,AttributeType=S AttributeName=SongTitle,AttributeType=S --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --profile sbx-acc`

### Describe Table
- `aws dynamodb describe-table --table-name MusicCollection --profile sbx-acc`

### Batch Write
- `aws dynamodb batch-write-item --request-items file://request-items.json --profile sbx-acc`

### Batch Get Item
- `aws dynamodb batch-get-item --request-items file://request-items.json --profile sbx-acc`

### Get Item
- `aws dynamodb get-item --table-name MusicCollection --key file://key.json --profile sbx-acc`

### Put Item
- `aws dynamodb put-item --table-name MusicCollection --item file://item.json --return-consumed-capacity TOTAL --profile sbx-acc`

### Scan
- `aws dynamodb scan --table-name MusicCollection --profile sbx-acc`

### Query
- `aws dynamodb query --table-name MusicCollection --key-condition-expression "Artist = :v1" --expression-attribute-values file://expression-attributes.json --profile sbx-acc`
