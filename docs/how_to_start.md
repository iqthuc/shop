# luồng khởi động ứng dụng
```
LUỒNG KHỞI TẠO:
---------------
1. main.go
   └── gọi app.Bootstrap()
       │
       ├── 2. Bootstrap() gọi config.Load() để nạp cấu hình
       │
       └── 3. Bootstrap() gọi app.New(cfg) với cấu hình đã nạp
           │
           └── 4. app.New() khởi tạo các dependency chính
               │   - Khởi tạo logger
               │   - Khởi tạo database
               │   - Khởi tạo server
               │
               └── 5. app.Run() gọi setupRoutes()
                   │
                   ├── 6. setupRoutes() khởi tạo và đăng ký handlers
                   │   - Khởi tạo auth handler và đăng ký auth router
                   │   - Khởi tạo users handler và đăng ký users router
                   │   - Khởi tạo products handler và đăng ký products router
                   │
                   └── 7. app.Run() khởi động server và chờ signal để shutdown


QUAN HỆ PHỤ THUỘC MẪU:
-------------------
cmd/api/main.go
    └── internal/app/bootstrap.go
        ├── internal/config/config.go
        └── internal/app/app.go
            ├── pkg/logger/logger.go  ─────┐
            │                              │
            ├── pkg/database/database.go ──┤
            │                              ├── internal/config/config.go
            ├── internal/server/server.go ─┘
            │
            ├── internal/auth/router.go    ┐
            │                              │
            ├── internal/users/router.go   ├── pkg/database/database.go
            │                              ├── pkg/logger/logger.go
            └── internal/products/router.go┘


```
