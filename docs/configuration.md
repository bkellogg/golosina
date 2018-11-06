# Configuration

The configuration of the platform is done by using this library https://github.com/joho/godotenv. It will require you to copy the .env to the root of your

### Basic

The basic App configurations have the following .env constants:

```
APP_NAME="Your app name"
APP_HOST=localhost
APP_PORT=8000
```

### Database

These are the available configuration values we can set for our database connection:

```
DB_ENABLED=false
DB_DRIVER="mysql"
DB_HOST="localhost"
DB_PORT=3306
DB_USER="your-database-user"
DB_PASSWORD="some-dev-password"
DB_DATABASE="the-schema-name"
```

!> It is important to remember that the `DB_ENABLED` value will initialize the connection to the database when the value is `true`. We have it defaulted to `false` because usually when you start coding you dont have a db ready.
