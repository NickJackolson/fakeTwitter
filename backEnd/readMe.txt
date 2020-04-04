GO API as Backend:
-Connects to a local sqlite3 database which is stored in data.db file.
-Has two tables: user and article
-user = {id,username,password} values(int,text,text)
-article = {id,title,content,ptime,author}, values(int,nvarchar,nvarchar,datetime,nvarchar)
-Listens to port 8081.
-Has all CRUD operation functions ready.
-Some functions or queries may not be best practices.
-Passwords are not encrypted.
-Database objects do not have entity relations.
