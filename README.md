# test_backend

## Development
Setup you $GOPATH.

Please make sure that you install [Glide
  package manager](https://github.com/Masterminds/glide) in the environment.

```
cd $GOPATH/src/github.com/nickhsine/test_backend
glide install                           # Install packages and dependencies
go run main.go                          # Run without live-reloading
```

## RESTful API
`test_backend` is a RESTful API built by golang.

It provides several RESTful web services, including
- Read events
- Create a event
- View a event


### Read events
- URL: `/v1/new-alarm-events/`
- Method: `GET`
- URL param:
  * Optional:
  `
  offset=[integer]
  limit=[integer]
  `
  * Explain:
  `offset`: the number you want to skip
  `limit`: the number you want server to return

- Response:
  * **Code:** 200 <br />
    **Content:**
    ```
    {
      "data": {
        total: 100,
        limit: 10,
        offset: 0,
        events: [...]
      },
       "status": "ok"
    }
    ```
  * **Code:** 500 <br />
  **Content:** `{"status": "error", "message": "${here_goes_error_msg}"}`

### Create a event
- URL: `/v1/new-alarm-events/`
- Content-Type of Header: `application/json`
- Method: `POST`
- Data Params:
```
{
   "camera_id": "camera_id_1",
   "prediction": "people",
   "starting_timestamp": 1508060290,
   "is_viewed": false,
   "thumbnail": "https://host/path/to/thumbnail"
}
```

- Response: 
  * **Code:** 201 <br />
    **Content:**
    ```
    {
        "status": "ok"
    }
    ```
  * **Code:** 400 <br />
  **Content:** `{"status": "error", "message": "${here_goes_error_msg}"}`
  * **Code:** 500 <br />
  **Content:** `{"status": "error", "message": "${here_goes_error_msg}"}`
 
### View a event
- URL: `/v1/event-viewed/event-id/`
- Content-Type of Header: `application/json`
- Method: `POST`
- Data Params:
```
{
   "event_id": "event_id"
}
```

- Response: 
  * **Code:** 200 <br />
    **Content:**
    ```
    {
        "status": "ok"
    }
    ```
  * **Code:** 400 <br />
  **Content:** `{"status": "error", "message": "${here_goes_error_msg}"}`
  * **Code:** 500 <br />
  **Content:** `{"status": "error", "message": "${here_goes_error_msg}"}`
