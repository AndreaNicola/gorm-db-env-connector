# Gorm connector abstraction

This module implements a stupid wrapper over gorm database connection. The goal is go DRY when you need to get your connection params from the process environment.

Actually i wrapped those databases:
- mysql
- postgresql

## Examples

There are 2 way to use this wrapper:
- connect to your database using environment variables
- connect directly with connection params

### Esempio 1

In this example we are creating a struct containing:
- the env variables names containing connection parameters
- the default values just in case our variables don't have a value

```go
log.Println("Database connection initialization...")

dbEnv := connector.MySqlEnv{
		DbUrlEnvVar:       "DATABASE_URL",
		DbSchemaEnvVar:    "DATABASE_SCHEMA",
		DbUsernameEnvVar:  "DATABASE_USERNAME",
		DbPasswordEnvVar:  "DATABASE_PASSWORD",
		DbUrlDefault:      "mysql.products.svc.cluster.local:3306",
		DbSchemaDefault:   "default-schema",
		DbUsernameDefault: "default-username",
		DbPasswordDefault: "default-password",
	}

db = connector.MySQLConnectEnv(dbEnv)
log.Println("Database connections initialized")
```

## Esempio 2

```go
log.Println("Database connection initialization...")

dbParams := connector.MySqlParams{
		DbUrl:      "mysql.products.svc.cluster.local:3306",
		DbSchema:   "schema",
		DbUsername: "username",
		DbPassword: "password",
	}

db = connector.MySQLConnect(dbParams)
log.Println("Database connections initialized")
```

# Astrazione del connettore di Gorm.

Questo modulo fornisce un wrapper sopra il connettore di Gorm che consente di non reinventarsi la ruota ogni volta che dobbiamo connetterci ad un database.

Attualmente sono supportati i seguenti database:
- mysql
- postgresql

## Esempio di utilizzo

Ci sono 2 modalità di utilizzo:
- indicare i nomi delle variabili di ambiente che contengono i parametri di connessione
- indicare direttamente i parametri di connessione

### Esempio 1

Nell'esempio seguente possiamo vedere che:
- vengono passati i nomi delle variabili di ambiente che contengono i parametri di connessione (i primi 4)
- vengono passati eventuali parametri di default che subentrano nel caso le variabili di ambiente non siano valorizzate.

```go
log.Println("Database connection initialization...")

dbEnv := connector.MySqlEnv{
		DbUrlEnvVar:       "DATABASE_URL",
		DbSchemaEnvVar:    "DATABASE_SCHEMA",
		DbUsernameEnvVar:  "DATABASE_USERNAME",
		DbPasswordEnvVar:  "DATABASE_PASSWORD",
		DbUrlDefault:      "mysql.products.svc.cluster.local:3306",
		DbSchemaDefault:   "default-schema",
		DbUsernameDefault: "default-username",
		DbPasswordDefault: "default-password",
	}

db = connector.MySQLConnectEnv(dbEnv)
log.Println("Database connections initialized")
```

## Esempio 2

```go
log.Println("Database connection initialization...")

dbParams := connector.MySqlParams{
		DbUrl:      "mysql.products.svc.cluster.local:3306",
		DbSchema:   "schema",
		DbUsername: "username",
		DbPassword: "password",
	}

db = connector.MySQLConnect(dbParams)
log.Println("Database connections initialized")
```
