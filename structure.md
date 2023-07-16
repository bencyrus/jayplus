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
│   │   ├── handlers
│   │   │   ├── authHandler.go
│   │   │   └── tests
│   │   │       └── authHandler_test.go
│   │   ├── routes
│   │   │   ├── authRoutes.go
│   │   │   └── tests
│   │   │       └── authRoutes_test.go
│   │   ├── middlewares
│   │   │   ├── authMiddleware.go
│   │   │   └── tests
│   │   │       └── authMiddleware_test.go
│   │   └── repository
│   │       ├── authRepository.go
│   │       └── postgres
│   │           ├── postgresAuthRepository.go
│   │           └── tests
│   │               └── postgresAuthRepository_test.go
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
│   ├── user.go
│   └── tests
│       ├── booking_test.go
│       ├── service_test.go
│       ├── vehicle_test.go
│       ├── invoice_test.go
│       ├── payment_test.go
│       ├── employee_test.go
│       ├── analytics_test.go
│       └── user_test.go
├── config
│   └── config.go
└── internal
    └── db
        ├── db.go
        ├── migrations
        └── tests
            └── db_test.go

