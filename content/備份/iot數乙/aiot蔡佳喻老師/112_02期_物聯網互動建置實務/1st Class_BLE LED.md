[[aiot蔡佳喻]]
[[esp32]]
[[ble]]

> 這個Project不好操作，並且需要耐心，以及恆心

- 請先安裝以下APP
    
    安卓用戶
    
    [nRF Connect for Mobile - Apps on Google Play](https://play.google.com/store/apps/details?id=no.nordicsemi.android.mcp&gl=US&pli=1)
    
    蘋果用戶
    
    [‎LightBlue®](https://apps.apple.com/tw/app/lightblue/id557428110)

```c
#include <BLEDevice.h> 
#include <BLEServer.h> 
#include <BLEUtils.h> 
#include <BLE2902.h> 
#include <String.h>

int count = 0;
int LEDPin = 4;
BLECharacteristic *pCharacteristic;
bool deviceConnected = false;
uint8_t txValue = 0;
long lastMsg = 0;
String rxload="Test\n";

/*為BLE供應商定義指定的 UUID 編號*/
#define SERVICE_UUID           "7d6159d2-881e-11ed-a1eb-0242ac120002" 
#define CHARACTERISTIC_UUID_RX "7d6159d2-881e-11ed-a1eb-0242ac120002"
#define CHARACTERISTIC_UUID_TX "7d6159d2-881e-11ed-a1eb-0242ac120002"
/*為BLE服務器編寫一個Callback函式來管理BLE的連接。*/
class MyServerCallbacks: public BLEServerCallbacks {
    void onConnect(BLEServer* pServer) {
      deviceConnected = true;
    };
    void onDisconnect(BLEServer* pServer) {
      deviceConnected = false;
    }
};

/*具有BLE功能的Callback函式。 
 *調用時，移動終端向 ESP32 發送數據時，會將其存儲到 reload 中。
 */
class MyCallbacks: public BLECharacteristicCallbacks {
    void onWrite(BLECharacteristic *pCharacteristic) {
      std::string rxValue = pCharacteristic->getValue();
      if (rxValue.length() > 0) {
        rxload="";
        for (int i = 0; i < rxValue.length(); i++){
          rxload +=(char)rxValue[i];
        }
      }
    }
};

/*創建BLE server流程：
 *1. Create a BLE Server
 *2. Create a BLE Service
 *3. Create a BLE Characteristic on the Service
 *4. Create a BLE Descriptor on the characteristic 
 *5. Start the service
 *6. Start advertising
 */
void setupBLE(String BLEName){
  const char *ble_name=BLEName.c_str();
  BLEDevice::init(ble_name);
  BLEServer *pServer = BLEDevice::createServer();
  pServer->setCallbacks(new MyServerCallbacks());
  BLEService *pService = pServer->createService(SERVICE_UUID); 
  pCharacteristic= pService->createCharacteristic(CHARACTERISTIC_UUID_TX,BLECharacteristic::PROPERTY_NOTIFY);
  pCharacteristic->addDescriptor(new BLE2902());
  BLECharacteristic *pCharacteristic = pService->createCharacteristic(CHARACTERISTIC_UUID_RX,BLECharacteristic::PROPERTY_WRITE);
  pCharacteristic->setCallbacks(new MyCallbacks()); 
  pService->start();
  pServer->getAdvertising()->start();
  Serial.println("Waiting a client connection to notify...");
}

void setup(){
	pinMode(LEDPin, OUTPUT);
	Serial.begin(115200);
	setupBLE("ESP32_BLE");//請每個人按照座號修改成不一樣的=>EX:1號的ESP01_BLE
}

void loop(){
	long now = millis();
  if (now - lastMsg > 100) {
		if (deviceConnected&&rxload.length()>0) {
			Serial.println(rxload);
			// Serial.println("rxload=");
			String Light_state;
			String str;
			if(rxload == "hi"){
				count = count + 1;
				digitalWrite(LEDPin, HIGH);
				delay(500);
			}
			else if(rxload == "bye"){
				digitalWrite(LEDPin, LOW);
				count = 0;
				delay(500);
			}
		else if(rxload == "print"){
			if(digitalRead(LEDPin) == HIGH){
			Light_state = "On";
		}
		else if(digitalRead(LEDPin) == LOW){
			Light_state = "OFF";            
		} 
		str = "LED is " + Light_state;
		const char *newValue=str.c_str();
		pCharacteristic->setValue(newValue);
		pCharacteristic->notify();

		}
		rxload="";

		}
		lastMsg = now;
	}
}

```