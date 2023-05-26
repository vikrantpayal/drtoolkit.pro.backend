# ABOUT drtooklit.pro.backend
Go programs to query and update drtoolkit.pro db

## DATABASE
postgresql14

## BACKEND
golang

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

### CONNECTING FROM GO
Install the pq package 

```
go get github.com/lib/pq

```



