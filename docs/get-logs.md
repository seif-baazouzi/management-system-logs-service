# Get Logs

Used to get user workspace logs.

**URL**: `/api/v1/logs/:workspaceID`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "logs": [
        {
            "logID": "logID", 
            "label": "label", 
            "description": "description", 
            "value": 10, 
            "date": "date", 
            "userID": "userID", 
            "workspaceID": "workspaceID", 
            "createdAt": "created date" 
        },
        ...
    ]
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
