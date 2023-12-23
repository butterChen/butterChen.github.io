[[12硬體中斷與睡眠省電]]
//Arduino睡眠範例

#include <Enerlib.h>  //載入能源函式庫

Energy energy;        //使用函式庫宣告物件 energy

int x = 0;

  

void setup() {

  Serial.begin(9600);   //啟用終端機

  pinMode(2, INPUT);    //定義D2腳位為輸入

  attachInterrupt(0, wakeup, RISING);  //使用中斷函式宣告中斷腳位0為上升緣觸發

  //觸發中斷時呼叫函式wakeup

}

void wakeup(){  //僅做喚醒呼叫用，故為空函式

}

  

void loop() {

  x = x + 1;

  delay(100);

  Serial.print("x = ");

  Serial.println(x);

  if( x%100 == 0 ){   //條件為計數到100之倍數時成立

    energy.PowerDown(); //使用物件energy呼叫斷電函式進入睡眠

  }

}