# To structure this project I have followed the Domain Driven, Data Oriented Architecture.

```
.
├── app
│   ├── api
│   │   ├── hanlder
│   │      ├── device
│   ├── main.go
│   
├── business
│   ├── business
│       ├── device
│         ├── model.go
│         ├── business.go
│         ├── store
│            ├── postgress
│            ├── mocks
├── docker-compose.yml   
├── Makefile             
├── go.mod               
├── go.sum               
└── .env 
```


## How to run
To run the project, you can use the `Makefile` which contains predefined commands to build and run the application. Here are some common commands:

- `make build`: Compiles the application.
- `make run`: Runs the application.
- To start database from docker
    - `make docker-up`: run database in docker.
    - `make docker-down`: stop database in docker.




## Device API

#### Get Device
- **URL:** `/api/v1/devices/{id}`
- **Method:** `GET`
- **URL Params:** 
    - `id=[uuid]` (required)
- **Success Response:**
    - **Code:** 200
    - **Content:** 
        ```json
        {
            "id": "6f24bfd3-64ee-47c4-b903-0ba81852ac6e",
            "name": "name",
            "brand": "brand",
            "created_at": "2024-11-18T09:10:21.174436+01:00"
        }
        ```
- **Error Response:**
    - **Code:** 404 NOT FOUND
    - **Content:** 
        ```json
        {
            "error": "device not found"
        }
        ```

#### Create Device
- **URL:** `/api/v1/devices`
- **Method:** `POST`
- **Data Params:** 
    ```json
    {
        "name": "Device Name",
        "brand": "brand"
    }
    ```
- **Success Response:**
    - **Code:** 201
    - **Content:** 
        ```json
        {
             "id": "6f24bfd3-64ee-47c4-b903-0ba81852ac6e",
            "name": "Device Name",
            "brand": "brand",
            "created_at": "2024-11-18T09:10:21.174436+01:00"
        }
        ```
- **Error Response:**
    - **Code:** 400 BAD REQUEST
    - **Content:** 
        ```json
        {
            "error": "name is required"
        }
        ```

#### Update Device
- **URL:** `/api/v1/devices/{id}`
- **Method:** `PUT`
- **URL Params:** 
    - `id=[uuid]` (required)
- **Data Params:** 
    ```json
    {
        "name": "Updated Device Name",
        "brand": "inactive"
    }
    ```
- **Success Response:**
    - **Code:** 200
    - **Content:** 
        ```json
        {
            "message": "device updated"
        }
        ```
- **Error Response:**
    - **Code:** 404 NOT FOUND
    - **Content:** 
        ```json
        {
            "error": "device not found"
        }
        ```

#### Delete Device
- **URL:** `/api/v1/devices/{id}`
- **Method:** `DELETE`
- **URL Params:** 
    - `id=[uuid]` (required)
- **Error Response:**
    - **Code:** 404 NOT FOUND
    - **Content:** 
        ```json
        {
            "error": "device not found"
        }
        ```
        #### Get All Devices
        - **URL:** `/api/v1/devices`
        - **Method:** `GET`
        - **URL Params:** 
            - `offset=[integer]` (optional, default is 0)
            - `limit=[integer]` (optional, default is 10)
            - `brand=[string]` (optional)
        - **Success Response:**
            - **Code:** 200
            - **Content:** 
                ```json
                [
                    {
                        "id": "6f24bfd3-64ee-47c4-b903-0ba81852ac6e",
                        "name": "Device Name",
                        "brand": "brand",
                        "created_at": "2024-11-18T09:10:21.174436+01:00"
                    },
                    {
                        "id": "7g35cge4-75ff-58d5-c014-1cb92963bd7f",
                        "name": "Another Device",
                        "brand": "another brand",
                        "created_at": "2024-11-19T10:11:22.185547+01:00"
                    }
                ]
                ```
        