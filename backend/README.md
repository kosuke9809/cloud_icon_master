# aws_icon_quiz_backend

## アーキテクチャ
- Onion Architectureを採用
    ```
    .
    ├── cmd
    │   ├── app
    │   └── migrate
    ├── internal
    │   ├── domain
    │   │   ├── model
    │   │   ├── repository
    │   │   ├── service
    │   │   └── validation
    │   ├── infra
    │   │   ├── database
    │   │   │   ├── config
    │   │   │   ├── migration
    │   │   │   └── seed
    │   │   └── oauth
    │   ├── interfaces
    │   │   ├── controller
    │   │   └── router
    │   └── usecase
    └── pkg

    ```