### Commands

```sh
brew services start mariadb # to start mariadb service
mysql -uroot # to connect to mariadb as rootuser, no password when installed
ps -x # to get information from cmdline
px -u # to get information from cmdline
px aux # to get information from cmdline
```

## Tree structure
- CMD holds the main entry of my application
- Internal

## Dependencies

- [chi](https://github.com/go-chi/chi) chi is a lightweight, idiomatic and composable router for building Go HTTP services


## abbreviations
mux --> multiplexor

### sql queries

```sql
grant all on widgets.* to 'name'@'%' identified by 'password';
```

### Issues

1. Not being able to connect as root

```sql
sudo mysql

-- for MySQL
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';

-- for MariaDB
ALTER USER 'root'@'localhost' IDENTIFIED VIA mysql_native_password USING PASSWORD('root');
```