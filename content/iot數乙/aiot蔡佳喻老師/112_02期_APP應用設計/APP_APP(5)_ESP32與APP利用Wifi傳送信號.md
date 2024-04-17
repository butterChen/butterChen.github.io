使用ngrok將內網傳到外部網路
讓外部網路可以連線至內部網路
缺點需要一直去申請(約一天一次)

20230912_使用WIFI控制ESP32上開關燈的

app端
[[112_02期_APP應用設計]]
[[esp32]]
https://dashboard.ngrok.com/get-started/your-authtoken


## 這堂課教什麼?

---

- ESP32與APP利用Wifi傳送信號

## 🐣課程PPT

---

[APP程式設計與應用_week5.pptx](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/bb2cd565-2bcc-4c26-8166-81d52b8fde8f/APP%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88%E8%88%87%E6%87%89%E7%94%A8_week5.pptx)

## 🔆素材 & APP檔案

---

[WIFI_LED.aia](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/94453b4f-e114-4400-bee3-a2e8534acf47/WIFI_LED.aia)

[素材.zip](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/58ced981-2b04-4649-afa7-20ecc9cc6208/%E7%B4%A0%E6%9D%90.zip)

APP說明：

請設計使用WIFI控制ESP32上開關燈的APP

![[Pasted image 20230925103619.png]]
![[Pasted image 20230925103628.png]]

```c
#include <WiFi.h>
// Replace with your network credentials
const char* ssid     = "tsvtc-w(Ruckus)";  //AP分享器SSID
const char* password = "1q2w#E$R";    //AP分享器密碼
WiFiServer server(80);

// Variable to store the HTTP request
String header;

// Auxiliar variables to store the current output state
String output22State = "off";
String output23State = "off";

// Assign output variables to GPIO pins
const int output22 = 22;
const int output23 = 23;

// Current time
unsigned long currentTime = millis();
// Previous time
unsigned long previousTime = 0; 
// Define timeout time in milliseconds (example: 2000ms = 2s)
const long timeoutTime = 2000;

void setup() {
  Serial.begin(115200);
  // Initialize the output variables as outputs
  pinMode(output22, OUTPUT);
  pinMode(output23, OUTPUT);
  // Set outputs to LOW
  digitalWrite(output22, LOW);
  digitalWrite(output23, LOW);

  // Connect to Wi-Fi network with SSID and password
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  // Print local IP address and start web server
  Serial.println("");
  Serial.println("WiFi connected.");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
  server.begin();
}

void loop(){
  WiFiClient client = server.available();   // Listen for incoming clients
  if (client) {                             // If a new client connects,
    Serial.println("New Client.");          // print a message out in the serial port
    String currentLine = "";                // make a String to hold incoming data from the client
    currentTime = millis();
    previousTime = currentTime;
    while (client.connected() && currentTime - previousTime <= timeoutTime) { // loop while the client's connected
      currentTime = millis();         
      if (client.available()) {             // if there's bytes to read from the client,
        char c = client.read();             // read a byte, then
        Serial.write(c);                    // print it out the serial monitor
        header += c;
        if (c == '\n') {                    // if the byte is a newline character
          // if the current line is blank, you got two newline characters in a row.
          // that's the end of the client HTTP request, so send a response:
          if (currentLine.length() == 0) {
        // HTTP headers always start with a response code (e.g. HTTP/1.1 200 OK)
        // and a content-type so the client knows what's coming, then a blank line:
            client.println("HTTP/1.1 200 OK");
            client.println("Content-type:text/html");
            client.println("Connection: close");
            client.println();
            
            // turns the GPIOs on and off
            if (header.indexOf("GET /22/on") >= 0) {
              Serial.println("GPIO 22 on");
              output22State = "on";
              digitalWrite(output22, HIGH);
            } else if (header.indexOf("GET /22/off") >= 0) {
              Serial.println("GPIO 22 off");
              output22State = "off";
              digitalWrite(output22, LOW);
            } else if (header.indexOf("GET /23/on") >= 0) {
              Serial.println("GPIO 23 on");
              output23State = "on";
              digitalWrite(output23, HIGH);
            } else if (header.indexOf("GET /23/off") >= 0) {
              Serial.println("GPIO 23 off");
              output23State = "off";
              digitalWrite(output23, LOW);
            }
            
            // Display the HTML web page
            client.println("<!DOCTYPE html><html>");
            client.println("<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">");
            client.println("<link rel=\"icon\" href=\"data:,\">");
            // CSS to style the on/off buttons 
            // Feel free to change the background-color and font-size attributes to fit your preferences
            client.println("<style>html { font-family: Helvetica; display: inline-block; margin: 0px auto; text-align: center;}");
            client.println(".button { background-color: #195B6A; border: none; color: white; padding: 16px 40px;");
            client.println("text-decoration: none; font-size: 30px; margin: 2px; cursor: pointer;}");
            client.println(".button2 {background-color: #77878A;}</style></head>");
            
            // Web Page Heading
            client.println("<body><h1>ESP32 Web Server</h1>");
            
            // Display current state, and ON/OFF buttons for GPIO 22  
            client.println("<p>GPIO 22 - State " + output22State + "</p>");
            // If the output22State is off, it displays the ON button       
            if (output22State=="off") {
              client.println("<p><a href=\"/22/on\"><button class=\"button\">ON</button></a></p>");
            } else {
              client.println("<p><a href=\"/22/off\"><button class=\"button button2\">OFF</button></a></p>");
            } 
               
            // Display current state, and ON/OFF buttons for GPIO 23  
            client.println("<p>GPIO 23 - State " + output23State + "</p>");
            // If the output23State is off, it displays the ON button       
            if (output23State=="off") {
              client.println("<p><a href=\"/23/on\"><button class=\"button\">ON</button></a></p>");
            } else {
              client.println("<p><a href=\"/23/off\"><button class=\"button button2\">OFF</button></a></p>");
            }
            client.println("</body></html>");
            
            // The HTTP response ends with another blank line
            client.println();
            // Break out of the while loop
            break;
          } else { // if you got a newline, then clear currentLine
            currentLine = "";
          }
        } else if (c != '\r') {  // if you got anything else but a carriage return character,
          currentLine += c;      // add it to the end of the currentLine
        }
      }
    }
    // Clear the header variable
    header = "";
    // Close the connection
    client.stop();
    Serial.println("Client disconnected.");
    Serial.println("");
  }
}

```

./ngrok.exe http 自己的esp32 IP:80 --authtoken 金鑰

https://ngrok.com/

	**ssh -R 80:localhost:3000 [serveo.net](http://serveo.net)**
	
http://serveo.net/#intro