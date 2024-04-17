```sql
•為 phpMyAdmin 創建一個超級用戶（superuser ）帳號。

•進入mysql進入mysql

•sudo mysql

•添加一個新的 MySQL 用戶

•CREATE USER 'superui'@'%' IDENTIFIED BY 'password_here’;

•給予新用戶 superui 超級用戶權限

•GRANT ALL PRIVILEGES ON *.* TO 'superui'@'%’;

•退出 MySQL

•exit

•檢視password policy

•SHOW VARIABLES LIKE 'validate_password%';

•修改 MySQL 的環境變數(立即生效 但下次失效)

•set global validate_password.length=6;

•SET GLOBAL validate_password.policy=LOW;

•查詢使用者

•SELECT User, Host FROM mysql.user;

•刪除使用者

•DROP USER 'superui'@'%';

•變更密碼

ALTER USER 'superui'@'%' IDENTIFIED BY 'NEW_USER_PASSWORD';
```