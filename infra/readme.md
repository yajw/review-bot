# MySQL setup
`brew install percona-server`

```bash
# start
mysql.server start

# change root password
mysqladmin -u root -p'<temp_password>' password '<new_password>'

# connect
mysql -uroot -p2024

# review_bot
alter user 'review_bot'@'localhost' identified by '7z$8K@k7';
grant all privileges on review.* to 'review_bot'@'localhost';
```

