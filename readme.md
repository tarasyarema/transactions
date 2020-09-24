# Transactions

## API

### `/transactions`

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

##### Return parameters
None.

##### Returns
- 201: If OK.
- 204: If transaction older than 60 seconds
- 400: JSON is invalid
- 422: Field/s not parsable or future transaction


#### `DELETE`

Deletes all transactions. 

##### Parameters
Accepts empty request body.

##### Example request body
```json
{}
```

##### Return parameters
None.

##### Returns
- 204: If OK , it may also return the number of transactions deleted, for example.

It may also return some other code if there where some problem deleting the transactions.


### `/statistics`

#### `GET`

Returns the statistic based of the transactions of the last 60 seconds.

##### Parameters
Accepts empty request body.

##### Example request body
```json
{}
```

##### Return parameters
- `sum`: BigDecimal specifying the total sum of transaction value in the last 60 seconds
- `avg`: BigDecimal specifying the average amount of transaction value in the last 60 seconds
- `max`: BigDecimal specifying single highest transaction value in the last 60 seconds
- `min`: BigDecimal specifying single lowest transaction value in the last 60 seconds
- `count`: a long specifying the total number of transactions that happened in the last 60 seconds

All BigDecimal values always contain exactly two decimal places and use HALF_ROUND_UP rounding. eg: 10.345 is returned as 10.35, 10.8 is returned as 10.80.

##### Example response
```json
{
    "sum": "1000.00",
    "avg": "100.53",
    "max": "200000.49",
    "min": "50.23",
    "count": 10
}
```

##### Returns
- 201: If OK.

