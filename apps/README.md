# apps

Put here your applications.

```
<APP_NAME>
├── cmd
│   └── <APP_NAME>
│       └── cmd
|           └── main.go <- Main file
├── pkg 
|   ├── config
|   |   └── settings.go <- Application settings
|   ├── models <- Application models
|   |   └── repo.go <- Database connection
|   └── services <- Application services
└── Dockerfile
```