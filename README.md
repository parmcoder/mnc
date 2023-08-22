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

    ```json
    {
        "symbol": "string",          // Transaction symbol, e.g., BTC
        "price": uint64,             // Symbol price, e.g., 100000
        "timestamp": uint64          // Timestamp of price retrieval
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

    Example usage

    ```go
	    var message mnc.GetTransaction

	    response, err := (*c.apiConnector).Monitor(&message)
        if err != nil {
            return err
        }
    ```
