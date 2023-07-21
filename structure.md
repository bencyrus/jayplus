.
├── README.md
├── cmd
│   └── app
│       └── main.go
├── go.mod
├── pkg
│   ├── app
│   │   ├── application.go
│   │   └── router.go
│   │
│   ├── authentication
│   │   ├── auth.go
│   │   ├── handlers.go
│   │   ├── routes.go
│   │   ├── middlewares.go
│   │   ├── repository
│   │   │   ├── authRepository.go
│   │   │   ├── authRepository_test.go
│   │   │   └── postgres
│   │   │       ├── postgresAuthRepository.go
│   │   │       └── postgresAuthRepository_test.go
│   │   └── tests
│   │       ├── auth_test.go
│   │       ├── handlers_test.go
│   │       ├── routes_test.go
│   │       └── middlewares_test.go
│   ├── booking
│   │   ├── handlers
│   │   │   ├── customerHandlers.go
│   │   │   ├── adminHandlers.go
│   │   │   └── tests
│   │   │       ├── customerHandlers_test.go
│   │   │       └── adminHandlers_test.go
│   │   ├── routes
│   │   │   ├── customerRoutes.go
│   │   │   ├── adminRoutes.go
│   │   │   └── tests
│   │   │       ├── customerRoutes_test.go
│   │   │       └── adminRoutes_test.go
│   │   ├── middleware
│   │   │   ├── bookingMiddleware.go
│   │   │   └── tests
│   │   │       └── bookingMiddleware_test.go
│   │   ├── services
│   │   │   ├── paymentService.go # Stripe logic here
│   │   │   └── tests
│   │   │       └── paymentService_test.go
│   │   └── repository
│   │       ├── bookingRepository.go
│   │       ├── postgres
│   │       │   ├── postgresBookingRepository.go
│   │       │   └── tests
│   │       │       └── postgresBookingRepository_test.go
│   └── messaging  # Twilio package
│       ├── smsService.go
│       └── tests
│           └── smsService_test.go
├── models
│   ├── booking.go
│   ├── service.go
│   ├── vehicle.go
│   ├── invoice.go
│   ├── payment.go
│   ├── employee.go
│   ├── analytics.go
│   └── user.go
├── config
│   └── config.go
└── internal
    └── db
        ├── db.go
        ├── migrations
        └── tests
            └── db_test.go

