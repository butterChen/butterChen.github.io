20230914 
[[12硬體中斷與睡眠省電]]

-----------------------

bool Green = true; //宣告綠色LED狀態

void setup()

{

  pinMode(13, OUTPUT); //定義腳位功能

  pinMode(12, OUTPUT); //定義腳位功能

  pinMode(2, INPUT); //定義腳位功能

//硬體中斷宣告

  attachInterrupt(0, Int0_Green, RISING);

//寫入腳位初始狀態

  digitalWrite(12, Green);

}
void Int0_Green(){ //中斷觸發時執行函式

  Green = !Green;

    digitalWrite(12, Green);

}

void loop()

{  //使用延時方式閃爍紅色LED

  digitalWrite(13, LOW);

  delay(5000); // Wait for 5000 millisecond(s)

  digitalWrite(13, HIGH);

  delay(5000); // Wait for 5000 millisecond(s)

}