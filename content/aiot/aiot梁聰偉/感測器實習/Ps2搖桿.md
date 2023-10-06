[[aiot梁聰偉老師]]
[[arduino]]
```
int vrx = A0;

int vry = A1;

int sw = 3;

int led = 13;

bool led_state = 0;

  

void setup() {

  pinMode(sw , INPUT_PULLUP);

  pinMode(led, OUTPUT);

  Serial.begin(9600);

}

  

void loop() {

  int x = analogRead(vrx);

  int y = analogRead(vry);

  Serial.print("vrX = " + String(x) + "\t");

  Serial.print("vrY = " + String(y) + "\n");

  Serial.println(digitalRead(sw));

  if(digitalRead(sw) == LOW){

    led_state = !led_state;

  }

  digitalWrite(led, led_state);

  delay(200);

}
```

![[Pasted image 20230920150251.png]]