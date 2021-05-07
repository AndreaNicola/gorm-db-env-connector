# Astrazione del connettore di Gorm.

Questo modulo fornisce un wrapper sopra il connettore di Gorm che consente di non reinventarsi la ruota ogni volta che dobbiamo connetterci ad un database.

Attualmente sono supportati i seguenti database:
- mysql

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