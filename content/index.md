#firstTarget
#theme 

[[整理程式碼]]

#Target 
[[aiot蔡佳喻]]
[[參考資料]]


```
[link text](#abcd)
`<a name="abcd"></a>`
```
[[c+組語]]
[[markdown]]
[[aiot/aiot蔡佳喻老師/Obsidian]]
告一段落
[[工作分享]]


[[0.陳奶油toDolist]]

檢查git設定

禮拜一的課程陳金生需要加油

aiot學習目標
python
golang
盡量使用 python建置互聯網零件
所以會選擇esp32 有三片

並且使用golang回應?
++

需要檢查aws 目前消耗金額
<font color=red>24</font>
<font color=red>31.5 of $10 2023-10-04</font>


#code

-------------------------
整理資料進筆記本
用課表做初期匯入

------------------
先補這個禮拜的筆記
跟學習如何處理conflict

------------------------------------
大數據競賽
https://imbd2023.thu.edu.tw/
參考看看 有沒有機會去玩

----------------------
數位貨幣
****

找出需要的資料
dht11 四個實作
<a name="abcd"></a>


20231002
<font color= red>
github.com 靜態網站</font>
https://www.google.com/search?q=github+%E9%9D%9C%E6%85%8B%E7%B6%B2%E7%AB%99&rlz=1C1ONGR_zh-TWTW1064TW1065&oq=github+%E9%9D%9C%E6%85%8B%E7%B6%B2%E7%AB%99&gs_lcrp=EgZjaHJvbWUyCQgAEEUYORiABDIGCAEQRRhA0gEJMTAwNzBqMGo3qAIAsAIA&sourceid=chrome&ie=UTF-8
==從這邊處理==




quartz

https://quartz.jzhao.xyz/

esp32相關研究
https://randomnerdtutorials.com/

-------------------

- [ ] sql測試路徑
```

	D:\butter\03.go\go\test1003
	D:\butter\03.go\02Gotest\web
```
db name: gosql
table name: userinfo
4個欄位: uid, username, department, created
```mysql
CREATE TABLE `gosql`.`userinfo` 
( 
 `uid` INT(20) NOT NULL AUTO_INCREMENT ,
 `username` VARCHAR(20) NULL DEFAULT NULL ,
 `department` VARCHAR(20) NULL DEFAULT NULL ,
 `created` DATE NULL DEFAULT NULL,
 PRIMARY KEY(uid)) 
 ENGINE = InnoDB;

CREATE TABLE User (
    ID          int         NOT NULL AUTO_INCREMENT,
    Name        varchar(50) NOT NULL,
    Telephone   varchar(20),
    Age         int,
    PRIMARY KEY(ID)
);
```

- [ ] 繼續處理aws to golang
	 go env -w GO111MODULE=on      
	 go mod init Gone
	 go get github.com/go-sql-driver/mysql
- [ ] 處理quatrz ob
	考慮一下要不要加入這個社群
	- [ ] https://quartz.jzhao.xyz/
	有異常 初始化那邊不會設定
	- [ ] https://blog.csdn.net/WongSSH/article/details/127195483
	- [ ] https://github.com/butterChen/butterChen.github.io/upload/main
	- [ ] http://localhost:8080/
	- [ ] 
	![[Pasted image 20231004155025.png]]
	- [ ] 看看要不要研究一下 discord
		https://discord.com/channels/927628110009098281/927628110009098284
		
- [ ] 設定github 當成網頁
	https://butterchen.github.io/
	是目前的首頁

	接下來要怎麼做
	- [ ] 需要研究quatrz 4.0
	- [ ] 安裝以後初始化設定看不太懂
	- [ ] 設計鹹魚王教學
	- [ ] 活動規劃
	- [ ] 先解箱子直到第四輪
	- [ ] 每天買箱子
	- [ ] 買普通釣魚竿
	- [ ] 買黃金釣魚竿
	- [ ] 競技場240場
	- [ ] 參加工會活動3456
	- [ ] 小技巧
	- [ ] 每天解鹹魚鍋
- [ ] 益民查到的課程
	https://edu.tcfst.org.tw/web/tw/class/show.asp?courseidori=12C026
	
	https://edu.tcfst.org.tw/web/tw/class/show.asp?courseidori=12Q329

- [x] 幫弟弟處理出國的喘息服務
	- [ ] 等待消息
- [ ] 等帶狀況好轉重新安排新的人過來
- [ ] 國泰打電話過去找apple
	+886 (2) 7752 4883
- [ ] 測試內文超連結 不過這個功能是外部連結
[test](#Target )

![[Pasted image 20231006134443.png]]

- [ ] 參數命名方式
	再看一下
	[[Golang 變數宣告（variables）]]
	https://pjchender.dev/golang/variables/