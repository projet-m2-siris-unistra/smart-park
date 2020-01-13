#include "LIB.h"

int sensorPin = A6;    // select the input pin for the potentiometer
int sensorValue = 0;  // variable to store the value coming from the sensor

// test presentation
int pirState = LOW;
int inputPin = 2;
int val = 0;                    // variable for reading the pin status

//

void pinsInit()
{
  pinMode(SENSOR_PIR, INPUT);
}

// the setup routine runs once when you press reset:
void setup() 
{       
  pinsInit();  

  pinMode(inputPin, INPUT);
  
  loraSerial.begin(57600);
  debugSerial.begin(9600);

  // Wait a maximum of 10s for Serial Monitor
  while (!debugSerial && millis() < 10000)
    ;

  debugSerial.println("-- STATUS");

  ttn.showStatus();

  debugSerial.println("-- JOIN");
  ttn.sendMacSet(1,devEui_);

  
  ttn.sendMacSet(2, appEui_);
  ttn.sendMacSet(5, appKey_);
  
  ttn.saveState();
  ttn.configureEU868();
  ttn.sendMacSet(10, "7");
  ttn.sendMacSet(7, "5");

  ttn.sendJoinSet(0);
}

// the loop routine runs over and over again forever:
void loop() 
{
  debugSerial.println("-- LOOP");
  
  lpp.reset();
  val = digitalRead(inputPin);  // read input value
  Serial.println(val);
//  lpp.addPresence(4, read_pir());
  lpp.addPresence(4, val);
  //  lpp.addLuminosity(4, val);

  // read the value from the sensor:
  sensorValue = analogRead(sensorPin);
  Serial.println(sensorValue);
  
  
  //float voltage = sensorValue * (5.0 / 1023.0);
  float voltage = sensorValue / 9.8;
  //int pour = (int) voltage;
  Serial.println(voltage);
  //Serial.println(pour);
  lpp.addTemperature(5, voltage);


  // Send it off
  ttn.sendBytes(lpp.getBuffer(), lpp.getSize());


  delay(15000);
}
