# Order service API

## Description
Pet project which provides a simple API of order-service which allows manipulating goods via GRPC methods.\n
Main goal of a project is mastering backend skills.
The following best practices were applied for the app's code:
1) Layered architecture.
2) Logging with different layers (prod, dev) via Zap (Uber technology) logger.
3) Config in .env files and Viper for config reading.
4) GRPC Interceptor for user requests tracking.

## Installation
To run this application, you have to have access to this project's code and configure an SSH key for the local machine you run it on. 
```
git clone [https://gitlab.crja72.ru/golang/2025/spring/course/students/173343-miyto2006-course-1478](https://github.com/Meldy183/order-service)
go run ./cmd/order-app/main.go
```
## gRPC API Usage

All requests and responses are defined in `order.proto`.

---

## Available RPC Methods

| Method          | Description                                                |
|-----------------|------------------------------------------------------------|
| **CreateOrder** | Creates a new order with the specified customer and items. |
| **GetOrder**    | Retrieves a specific order by its ID.                      |
| **UpdateOrder** | Updates details of an existing order.                      |
| **DeleteOrder** | Deletes an order by its ID.                                |
| **ListOrders**  | Lists all orders.                                          |

---

## How to Test the API
1) Download Postman.
2) Enter URL (by default "localhost:50051"), but you can customize it via .env file in ./config directory.
3) Choose GRPC method and provide data (check "order.proto" file for data format).

## Environment variables
For configuration settings .env files are used. It is located under ./config directory; "env.example" file is provided for
data format; It contains two variables:
1) "GRPC_PORT" which is 50051 by default; Used for running server on this port
2) "PROTOCOL" which is tcp by default; Used for choosing a protocol type while running server
3) "ENV" which is dev by default; Used for logging level (if dev then debug logs are shown)

If you need to customize the config you should provide your own .env file and place it under ./config directory with the same format as provided in the "env/example" file.
