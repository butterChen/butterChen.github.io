[[0.toDolist]]
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

### 進入docker
```cli
docker ps
```
結果
```
CONTAINER ID   IMAGE           COMMAND                  CREATED        STATUS                          PORTS                         NAMES
5284755f28b6   iot_data:112    "sh /x/appRun.sh"        21 hours ago   Restarting (1) 34 seconds ago                                 docker_iot_data_1
50e3643e08b3   iot_www:112     "sh /x/appRun.sh"        21 hours ago   Restarting (1) 34 seconds ago                                 docker_iot_www_1
122720f3dc2a   iot_mongo:112   "docker-entrypoint.s…"   21 hours ago   Up 2 hours                      172.17.0.1:21627->27017/tcp   docker_iot_mongo_1
```

```cli
docker exec -it 122720f3dc2a bash
```

