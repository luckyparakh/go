Go’s encoding/json package encodes []byte as a base64-encoding string. The record’s value is a []byte and it should be base64 encoded like Hello1 -> SGVsbG8x.
curl -X POST localhost:7890 -d '{"record": {"value": "SGVsbG8x"}}'
curl -X POST localhost:7890 -d '{"record": {"value": "SGVsbG8y"}}'
curl -X POST localhost:7890 -d '{"record": {"value": "SGVsbG8z"}}'

curl -X GET localhost:7890 -d '{"offset": 30}' // error
curl -X GET localhost:7890 -d '{"offset": 3}'