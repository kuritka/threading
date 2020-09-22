### Stopping infinite loop after specified time:



```go
const abort = 1 * time.Minute
ctx, _ := context.WithTimeout(context.Background(), abort)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ABORT")
			return
		default:
		}
    //code within the loop
   }
```
