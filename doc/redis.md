

## 设置key

```go
err := client.Set("mykey", "myvalue", 0).Err()   // 永久存储
// 或者 
err := client.Set("mykey", "myvalue", time.Duration(0)).Err()   // 永久存储
// 或者 
err := client.Set("mykey", "myvalue", nil).Err()   // 不设置过期时间，等价于永久存储
```
