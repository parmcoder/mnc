# Transaction-tracker
This submission consist of 2 parts
1. server - the server that could handle requests
2. client module - we called it `mnc`, which will be used inside server
## MNC - mock node client
This is the module we used to connect to the server

### Quick start
```go
    // This is how you can use the service to inject the dependency
	ac := mnc.CreateAPIConnector()
    
    // You can broadcast transaction
    var createBroadcast mnc.CreateBroadcast
    response, err := ac.Broadcast(&createBroadcast)

    // You can also monitor transaction 
	var getTransaction mnc.GetTransaction
	response, err := ac.Monitor(&getTransaction)
```
### Functions
1. **Broadcast Transaction**:
 Construct a JSON payload and send a POST request to `https://mock-node-wgqbnxruha-as.a.run.app/broadcast`. The payload structure is as follows:

    ```go
    type CreateBroadcast struct {
        Symbol    string `json:"symbol"`
        Price     uint64 `json:"price"`
        Timestamp uint64 `json:"timestamp"`
    }
    ```

    Example usage

    ```go
        var message mnc.CreateBroadcast

        response, err := (*c.apiConnector).Broadcast(&message)
        if err != nil {
            return err
        }
    ```

2. **Transaction Status Monitoring**: Utilize the transaction hash obtained from the response to periodically issue GET requests to `https://mock-node-wgqbnxruha-as.a.run.app/check/<tx_hash>`.

    ```go
    type GetTransaction struct {
	    TxHash string `json:"tx_hash"`
    }
    ```
    Example usage

    ```go
	    var message mnc.GetTransaction

	    response, err := (*c.apiConnector).Monitor(&message)
        if err != nil {
            return err
        }
    ```

## Server - Service integration
Change your directory to `server` folder and run the server using
```
go run main.go
```

### Endpoints
1. **Broadcast Transaction**
- METHOD: POST
- PATH: /mnc
- BODY:  
    ```json
    {
        "symbol": "string",     // Transaction symbol, e.g., BTC
        "price": 2,             // Symbol price, e.g., 100000
        "timestamp": 1          // Timestamp of price retrieval
    }
    ```
- Example:
    ```
    curl --location 'localhost:1323/mnc' \
    --header 'Content-Type: application/json' \
    --data '{
        "symbol": "ETH",
        "price": 45000,
        "timestamp": 167891234235
    }'
    ```
2. **Transaction Status Monitoring**
- METHOD: GET
- PATH: /mnc/<tx_hash>
- Example:
    ```
    curl --location 'localhost:1323/mnc/6774392a2a2e67efd2cd61b366dfcc2046bdec4ffc9a6f77ed98b88995f5f241'
    ```
