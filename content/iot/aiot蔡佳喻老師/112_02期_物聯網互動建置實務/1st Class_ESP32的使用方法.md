
[[aiot蔡佳喻]]
[[esp32]]
[[C語言]]
## 課程規劃

1. 課程介紹+ESP32基礎介紹及使用
2. WebServer DHT + ThingSpeak使用
3. MQTT + Line Notify
4. Line BOT & Google App Script

## 這門課會教什麼?

1. 透過藍芽與WIFI兩種傳輸方式，與ESP32連結，製作相關應用
2. 課程內容大多來源於網路相關內容，因此主要教學目的為:**『教導同學如何搜尋資源及將內容實現並改寫成自己想要的專題。』**

## 老師扮演的角色?

1. 相關知識提供者
2. 協助同學深入探索物聯網相關內容
3. 幫助同學解決困難

## 老師無法做到的事

- **創意**

## 評分方式

1. final100

## 課堂使用程式碼

利用按鈕控制LED

```c
int count = 0;
int LEDPin = 4, input = 22;

void setup() {
  pinMode(input, INPUT_PULLUP);
	pinMode(LEDPin, OUTPUT);
}

void loop() {
  if((digitalRead(input) == 0) &&
		(count == 0) ){
		count = count + 1;
		digitalWrite(LEDPin, count);
		delay(1000);

	}

	else if((digitalRead(input) == 0) &&
	(count > 0)){
		count = 0;
		digitalWrite(LEDPin, count);
		delay(1000);
	}

}
```

http://taichuino.blogspot.com/2013/04/arduino-pull-up.html