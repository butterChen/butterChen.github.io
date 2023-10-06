[[aiot蔡佳喻]]
[[esp32]]
[[aiot/aiot許昌輝/C語言]]
[[dht11]]
#dht11


https://ahwuiot.notion.site/3rd-Class_Line-Notify-4b2aa48365794ded9a9131bc03c64114

----
<font color = red>原始程式碼</font>

```
#include <WiFi.h>
#include <WiFiClient.h>
#include <TridentTD_LineNotify.h>

#include <AHT10.h>
#include <Wire.h>
uint8_t readStatus = 0;
AHT10 myAHT10(AHT10_ADDRESS_0X38);

// 修改成Line Token權杖號碼
#define LINE_TOKEN "jYTanwRtmAwnwevAWbBhU35RKhEq2OXAxxxxxxxxxxx"

// 設定無線基地台SSID跟密碼
const char* ssid     = "tsvtc-w(Ruckus)";  //AP分享器SSID
const char* password = "1q2w#E$R";    //AP分享器密碼

 
float humidity, temp_f;   // 從 AHT10 讀取的值

unsigned long previousMillis = 0;        // will store last temp was read
const long interval = 2000;              // interval at which to read sensor


 
void setup(void)
{
  Serial.begin(115200);  // 設定速率 感測器
  myAHT10.begin();           // 初始化

  WiFi.mode(WIFI_STA);
  // 連接無線基地台
  WiFi.begin(ssid, password);
  Serial.print("\n\r \n\rWorking to connect");

  // 等待連線，並從 Console顯示 IP
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("DHT Weather Reading Server");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}
 
void loop(void)
{
  // 量測間等待至少 2 秒
  unsigned long currentMillis = millis();
 
  if(currentMillis - previousMillis >= interval) {
    // 將最後讀取感測值的時間紀錄下來 
    previousMillis = currentMillis;   

    // 讀取溫度大約 250 微秒!
    humidity = myAHT10.readHumidity();          // 讀取濕度(百分比)
    delay(100);
    temp_f = myAHT10.readTemperature();     // 讀取溫度(華氏)
    delay(100);
    
 
    // 檢查兩個值是否為空值
    if (isnan(humidity) || isnan(temp_f)) {
       Serial.println("Failed to read from DHT sensor!");
       return;
    }
  }

  String tempe="溫度:"+String(temp_f)+"℃";   
  String humid="濕度:"+String(humidity)+"％";

  // 顯示 Line版本
  Serial.println(LINE.getVersion());
 
  LINE.setToken(LINE_TOKEN);

  // 先換行再顯示
  LINE.notify("\n" + tempe + " ；" + humid);
      
  // 每2分鐘發送一次(delay120000)
  delay(10000);
}
```

---
<font color = red >我的</font>
```
#include <WiFi.h>

#include <WiFiClient.h>

#include <TridentTD_LineNotify.h>

#include <DHT.h>

// #include <AHT10.h>

#include <Wire.h>

#define DHTPIN 4

uint8_t readStatus = 0;

// AHT10 myAHT10(AHT10_ADDRESS_0X38);

  

// 修改成Line Token權杖號碼

#define LINE_TOKEN "TKrnbP05vNXSf8SLfGfV9WLPXYpiQ3D4LZl9VXcLZEQ"

#define DHTTYPE DHT11   // DHT 22  (AM2302), AM2321

  

// 設定無線基地台SSID跟密碼

const char* ssid     = "OPPOA775G";  //AP分享器SSID

const char* password = "123456aa";    //AP分享器密碼

  

float humidity, temp_f;   // 從 AHT10 讀取的值

  

unsigned long previousMillis = 0;        // will store last temp was read

const long interval = 2000;              // interval at which to read sensor

  

DHT dht(DHTPIN, DHTTYPE);

void setup(void)

{

  Serial.begin(115200);  // 設定速率 感測器

  // myAHT10.begin();           // 初始化

  

  WiFi.mode(WIFI_STA);

  // 連接無線基地台

  WiFi.begin(ssid, password);

  Serial.print("\n\r \n\rWorking to connect");

  

  // 等待連線，並從 Console顯示 IP

  while (WiFi.status() != WL_CONNECTED) {

    delay(500);

    Serial.print(".");

  }

  Serial.println("");

  Serial.println("DHT Weather Reading Server");

  Serial.print("Connected to ");

  Serial.println(ssid);

  Serial.print("IP address: ");

  Serial.println(WiFi.localIP());

}

void loop(void)

{

  // 量測間等待至少 2 秒

  unsigned long currentMillis = millis();

  

  if(currentMillis - previousMillis >= interval) {

    // 將最後讀取感測值的時間紀錄下來

    previousMillis = currentMillis;  

  

  // try(){

    humidity = dht.readHumidity();

    temp_f = dht.readTemperature();

  // }catch(){

  

  // }

  

    // 讀取溫度大約 250 微秒!

    // humidity = myAHT10.readHumidity();          // 讀取濕度(百分比)

    // delay(100);

    // temp_f = myAHT10.readTemperature();     // 讀取溫度(華氏)

    // delay(100);

    // 檢查兩個值是否為空值

    if (isnan(humidity) || isnan(temp_f)) {

       Serial.println("Failed to read from DHT sensor!");

       return;

    }

  }

  

  String tempe="溫度:"+String(temp_f)+"℃";  

  String humid="濕度:"+String(humidity)+"％";

  

  // 顯示 Line版本

  Serial.println(LINE.getVersion());

  LINE.setToken(LINE_TOKEN);

  

  // 先換行再顯示

  LINE.notify("\n" + tempe + " ；" + humid);

  // 每2分鐘發送一次(delay120000)

  delay(10000);

}
```