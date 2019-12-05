#include "LIB.h"

void pinsInit()
{
  pinMode(SENSOR_PIR, INPUT);
}

// the setup routine runs once when you press reset:
void setup() 
{       
  pinsInit();  
  
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
  lpp.addTemperature(1, 22.5);
  lpp.addBarometricPressure(2, 1073.21);
  lpp.addGPS(3, 52.37365, 4.88650, 2);
  lpp.addPresence(4, read_pir());

  // Send it off
  ttn.sendBytes(lpp.getBuffer(), lpp.getSize());
  
  delay(15000); 
}
