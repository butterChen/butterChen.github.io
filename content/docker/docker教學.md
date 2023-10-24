[[docker]]
[[aiot謝耀聰]]
[[linux]]

![[Pasted image 20231020173205.png]]
使用以下指令 都需要使用 sudo 因為權限有問題
然後要先安裝
```linux
sudo apt  install docker-compose
```
在執行
```linux
	
docker-compose build  

docker-compose up
docker-compose up -d 

docker-compose up iot_mongo
docker-compose up -d  iot_mongo
docker-compose restart iot_mongo

docker-compose build  iot_data
docker-compose up  -d iot_data

docker-compose build  iot_www
docker-compose up     iot_www

docker-compose restart 
docker-compose restart  iot_data
docker-compose restart  iot_www

docker-compose ps

docker-compose stop

docker-compose down

```