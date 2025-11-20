[[aiot蔡佳喻]]
[[esp32]]
[[googlesheet]]

https://ahwuiot.notion.site/Google-APP-Script-ESP32-1aa0abdd669a417e9f58fa004a362d5d

```app script
function getLightStatus() {
  var sheet = SpreadsheetApp.getActiveSpreadsheet().getActiveSheet();
  var status = sheet.getRange("A1").getValue();
  return status;
}
function doGet(e) {
  var status = getLightStatus();
  return ContentService.createTextOutput(status);
}
```

```c
//Include required libraries
#include "WiFi.h"
#include <HTTPClient.h>
// WiFi credentials
const char* ssid     = "tsvtc-w(Ruckus)";  //AP分享器SSID
const char* password = "1q2w#E$R";    //AP分享器密碼
// Google script ID and required credentials
String GOOGLE_SCRIPT_ID = "自己的部署ID";    // change Gscript ID
void setup() {
  delay(1000);
  pinMode(2, OUTPUT);
  Serial.begin(115200);
  delay(1000);
  // connect to WiFi
  Serial.println();
  Serial.print("Connecting to wifi: ");
  Serial.println(ssid);
  Serial.flush();
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
}
void loop() {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String url = "https://script.google.com/macros/s/" + GOOGLE_SCRIPT_ID + "/exec?";
    Serial.println("Making a request");
    http.begin(url.c_str()); //Specify the URL and certificate
    http.setFollowRedirects(HTTPC_STRICT_FOLLOW_REDIRECTS);
    int httpCode = http.GET();
    String payload;
    if (httpCode > 0) { //Check for the returning code
      payload = http.getString();
      Serial.println(httpCode);
      Serial.println(payload);
      if(payload == "ON"){
        digitalWrite(2, HIGH);
        Serial.print("LED is ON");
      }
      if(payload == "OFF"){
        digitalWrite(2, LOW);
        Serial.print("LED is OFF");
      }
    }
    else {
      Serial.println("Error on HTTP request");
    }
    http.end();
  }
  delay(15000);
}
```

```
char msg[] = "1,20,300,4000,50000";

void setup()
{
  Serial.begin(115200);
  Serial.print(F("Parsing String: "));
  Serial.println(msg);
  char* ptr = strtok(msg, ",");
  byte i = 0;
  Serial.println(F("index\ttext\tNumeric Value"));
  while (ptr) {
    Serial.print(i);
    Serial.print(F("\t\""));
    Serial.print(ptr); // this is the ASCII text we want to transform into an integer
    Serial.print(F("\"\t"));
    Serial.println(atol(ptr)); // atol(ptr) will be your long int you could store in your array at position i. atol() info at http://www.cplusplus.com/reference/cstdlib/atol
    ptr = strtok(NULL, ",");
    i++;
  }
}

void loop() {}
```

<font color = red>strcmp(msg[0] ,"ON") == 0</font>

```
#include "WiFi.h"
#include <HTTPClient.h>

const char* ssid     = "tsvtc-w(Ruckus)";  // AP分享器SSID
const char* password = "1q2w#E$R";        // AP分享器密碼
String GOOGLE_SCRIPT_ID = "AKfycbytikgODO1T6-j8OS_mxCWE60Bgho9cGsF5Nf4CZiDtugdohg_AIzwUA6_qYWZx0KzuHA";  // change Gscript ID

void setup() {
  delay(1000);
  pinMode(2, OUTPUT);
  pinMode(4, OUTPUT);
  pinMode(18, OUTPUT);
  Serial.begin(115200);
  delay(1000);
  Serial.println();
  Serial.print("Connecting to wifi: ");
  Serial.println(ssid);
  Serial.flush();
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
}

void loop() {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String url = "https://script.google.com/macros/s/"+ GOOGLE_SCRIPT_ID + "/exec?";
    Serial.println("Making a request");
    http.begin(url.c_str());
    http.setFollowRedirects(HTTPC_STRICT_FOLLOW_REDIRECTS);
    int httpCode = http.GET();
    String payload;

    if (httpCode > 0) {
      payload = http.getString();
      char* ptr = strtok(const_cast<char*>(payload.c_str()), ",");
      byte i = 0;
      char* msg[100]; 
      while (ptr) {
        // 设定足够大的数组来存储子字符串
        msg[i] = strdup(ptr);
        ptr = strtok(NULL, ",");
        i++;
       
      }
      Serial.print("msg[0]");
      Serial.println(msg[0]);
      Serial.print("msg[1]");
      Serial.println(msg[1]);
      Serial.print("msg[2]");
      Serial.println(msg[2]);
      Serial.println(httpCode);
      Serial.println(payload);
      if (strcmp(msg[0] ,"ON") == 0) {
        digitalWrite(2, HIGH);
        Serial.print("LED2 is ON");
      }
      if (strcmp(msg[0] ,"OFF") == 0) {
        digitalWrite(2, LOW);
        Serial.print("LED2 is OFF");
      }
      if (strcmp(msg[1] ,"ON") == 0) {
        digitalWrite(4, HIGH);
        Serial.print("LED4 is ON");
      }
      if (strcmp(msg[1] ,"OFF") == 0) {
        digitalWrite(4, LOW);
        Serial.print("LED4 is OFF");
      }
      if (strcmp(msg[2] ,"ON") == 0) {
        digitalWrite(18, HIGH);
        Serial.print("LED18 is ON");
      }
      if (strcmp(msg[2] ,"OFF") == 0) {
        digitalWrite(18, LOW);
        Serial.print("LED18 is OFF");
      }
    } else {
      Serial.println("Error on HTTP request");
    }
    http.end();
  }
  delay(10000);
}

```
==觀察連線異常可能讓程式碼出現問題==
經過測試為同學網路端程式有異常
```c
//Include required libraries

#include "WiFi.h"

#include <HTTPClient.h>

// WiFi credentials

const char* ssid     = "OPPOA775G";  //AP分享器SSID

const char* password = "123456aa";    //AP分享器密碼

// Google script ID and required credentials

String GOOGLE_SCRIPT_ID = "AKfycbw33wtPJ6jLJUmijTSyGTYR_m0QaD36EIWuwG3UA3gkF1iHxSCvG3UZ_YBLZqq70EPE";    // change Gscript ID

void setup() {

  delay(1000);

  pinMode(2, OUTPUT);

  Serial.begin(115200);

  delay(1000);

  // connect to WiFi

  Serial.println();

  Serial.print("Connecting to wifi: ");

  Serial.println(ssid);

  Serial.flush();

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {

    delay(500);

    Serial.print(".");

  }

}

void loop() {

  if (WiFi.status() == WL_CONNECTED) {

    HTTPClient http;

    String url = "https://script.google.com/macros/s/" + GOOGLE_SCRIPT_ID + "/exec?";

    Serial.println("Making a request");

    http.begin(url.c_str()); //Specify the URL and certificate

    http.setFollowRedirects(HTTPC_STRICT_FOLLOW_REDIRECTS);

    int httpCode = http.GET();

    String payload;

    if (httpCode > 0) { //Check for the returning code

      payload = http.getString();

      Serial.println(httpCode);

      Serial.println(payload);

      if(payload == "ON"){

        digitalWrite(2, HIGH);

        Serial.print("LED is ON");

      }

      if(payload == "OFF"){

        digitalWrite(2, LOW);

        Serial.print("LED is OFF");

      }

    }

    else {

      Serial.println("Error on HTTP request");

    }

    http.end();

  }

  delay(1000);

}
```