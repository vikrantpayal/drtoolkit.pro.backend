# ABOUT drtooklit.pro.backend
Go programs to query and update drtoolkit.pro db

## DATABASE
postgresql14

## BACKEND
golang v1.17.7

## HOWTO
### DB
To install postgres, use homebrew
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Remember to upgrade brew before installing
```
brew upgrade
```

Install postgres and start the service
```
brew install postgresql
```
Start the service
```
brew services start postgresql
```

*Remember, it is services not service. Don't sit there and scratch your head.*

The path to postgresql files for MAC is
> /usr/local/opt/postgresql@14

Create database with 
```
createdb mydb
```

Access db with 
```
psql mydb
```

## CONNECTING  GO to POSTGRESQL
Install the pq package 

```
go get github.com/lib/pq
```

Init the new progres
```
go mod init example/hello
```

Connection and CRUD code (go>postgresql)
> https://golangdocs.com/golang-postgresql-example


## REACT to GOLANG
https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/
