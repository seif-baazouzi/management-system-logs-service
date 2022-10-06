# Create Log

Used to create new log.

**URL**: `/api/v1/logs/:workspaceID`

**Method**: `POST`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "label": "label", 
    "description": "description", // optional
    "value": value, // unsigned number
    "date": "date",
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success",
    "logID": "logID"
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "label": "label error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
