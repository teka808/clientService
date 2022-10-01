**ClientService**

```
    GO
    PostgreSQL
    RabbitMQ
    Gorilla/mux
```

```bash
    The ClientService calling DeviceService fot getting data
```

***GET DEVICES:***

```
   /devices?page=0&limit=10
```

```
    RESPONSE: 
    [
        {
            "id": "c44dt2ecie6h7m4lih70",
            "type": "awair-element",
            "status": "disconnected",
            "timezone": "Asia/Tokyo",
            "Coordinates": {
                "latitude": 130.707891,
                "longitude": 32.8031
            }
        },
    ]

```

***GET DEVICES BY STATUS:***

```
    GET DEVICES BY STATUS:
    /devices/status/{status}?page=0&limit=10
```

```
    RESPONSE: 
    [
        {
            "id": "c44dt2ecie6h7m4lih70",
            "type": "awair-element",
            "status": "disconnected",
            "timezone": "Asia/Tokyo",
            "Coordinates": {
                "latitude": 130.707891,
                "longitude": 32.8031
            }
        },
    ]

```

***GET DEVICES BY TYPE:***

```
    /devices/type/{type}?page=0&limit=10
```

```
    RESPONSE: 
    [
        {
            "id": "c44dt2ecie6h7m4lih70",
            "type": "awair-element",
            "status": "disconnected",
            "timezone": "Asia/Tokyo",
            "Coordinates": {
                "latitude": 130.707891,
                "longitude": 32.8031
            }
        },
    ]

```

***GET DEVICE BY ID:***
```
    /devices/{id}
```

```
    RESPONSE: 
   
    {
        "id": "c44dt2ecie6h7m4lih70",
        "type": "awair-element",
        "status": "disconnected",
        "timezone": "Asia/Tokyo",
        "Coordinates": {
            "latitude": 130.707891,
            "longitude": 32.8031
        }
    }
    
```