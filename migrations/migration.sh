#!/bin/bash
#set -e
echo "Sleep for 1 min..."
sleep 60
if [[ $(mysql -hgameaholic_mysql -proot -e 'show databases;' | grep gameaholic | wc -l) == 0 ]];
then
echo "Creating database..."
mysql -hgameaholic_mysql -proot -e 'create database gameaholic;'
sleep 1
echo "Restore database..."
mysql -hgameaholic_mysql -proot gameaholic < /var/run/mysqld/*_create_bootstraps.mysql.up.sql
else
echo "database already exist.."
exit 0
fi