# Delete Log

Used to delete log.

**URL**: `/api/v1/logs/:logID`

**Method**: `DELETE`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
