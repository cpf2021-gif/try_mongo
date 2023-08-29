# 使用go存储md文件

## 尝试使用mongodb存储md文件
1. 使用mongo-driver连接mongodb
```go
    // connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// disconnect
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
```
2. 测试连接
```go
	// test connection
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
```
3. 创建/获取集合
```go
    // create collection
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection, err := client.Database("bookdb").CreateCollection(ctx, "mdfiles")
	if err != nil {
		panic(err)
	}

    // get collection
    collection = client.Database("bookdb").Collection("mdfiles")
``` 

