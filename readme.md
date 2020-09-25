# Transactions

## To run

Just run `make`.
To run the tests `make test`.

## API

### `/transactions`

#### `GET`

Get the transactions db.

##### Example response
```json
[
  {
    "amount": 88,
    "timestamp": "2020-09-25T20:05:48.539Z"
  },
  {
    "amount": 27,
    "timestamp": "2020-09-25T20:05:49.100Z"
  }
]
```

##### Return codes
- 201: If OK.
- 500: If something weird happened.

#### `POST`

Called every time a transaction is made. It is also the sole input of this rest API.

##### Parameters
Passed in the request body (JSON)
- `amount`: Arbitrary BigDecimal.
- `timestamp`: ISO 8601 timestamp.

##### Example request body
```json
{
    "amount": "12.3343",
    "timestamp": "2018-07-17T09:59:51.312Z"
}
```

##### Return codes
- 201: If OK.
- 204: If transaction older than 60 seconds
- 400: JSON is invalid
- 422: Field/s not parsable or future transaction


#### `DELETE`

Deletes all transactions.

##### Returns
- 204: If OK , it may also return the number of transactions deleted, for example.

### `/statistics`

#### `GET`

Returns the statistic based of the transactions of the last 60 seconds.

##### Return parameters
- `sum`: decimal specifying the total sum of transaction value in the last 60 seconds
- `avg`: decimal specifying the average amount of transaction value in the last 60 seconds
- `max`: decimal specifying single highest transaction value in the last 60 seconds
- `min`: decimal specifying single lowest transaction value in the last 60 seconds
- `count`: a long specifying the total number of transactions that happened in the last 60 seconds

##### Example response
```json
{
    "sum": "1000.00",
    "avg": "100.53",
    "max": "200000.49",
    "max_queue": [
        ...
    ],
    "min": "50.23",
    "min_queue": [
        ...
    ],
    "count": 10
}
```

##### Returns
- 201: If OK.

