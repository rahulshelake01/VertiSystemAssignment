Sample request and response:

curl --location --request POST 'http://localhost:8080/api/most-used-word' \
--header 'Content-Type: application/json' \
--data-raw '{
    "text":"Please ok ok ok ok  create a small service that accepts as input a body of text, such as that from a book, and returns the top ten most-used words along with how many times they occur in the text"
}'

Response:
{
    "success": true,
    "data": {
        "most-used-word": "ok"
    }
}
