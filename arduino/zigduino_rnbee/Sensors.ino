boolean isPeopleDetected()
{
  int sensorValue = digitalRead(SENSOR_PIR);
  if(sensorValue == HIGH)//if the sensor value is HIGH?
  {
    return true;//yes,return true
  }
  else
  {
    return false;//no,return false
  }
}

boolean read_pir() 
{
  if(isPeopleDetected())//if it detects the moving people?
    pir=1;
  else
    pir=0;
  debugSerial.print("pir=");
  debugSerial.print(pir);  
}

void read_light()
{
  light=analogRead(SENSOR_LIGHT);
  debugSerial.print("light=");
  debugSerial.print(light);
}

void read_sound()
{
  sound=analogRead(SENSOR_SOUND);
  debugSerial.print("sound=");
  debugSerial.print(sound);
}

void update_sensors()
{
  debugSerial.print("[IBAT];");
//  read_temp_hum();
//  serialDebug.print(";");
  read_light();
  debugSerial.print(";");
  read_sound();
  debugSerial.print(";");
  read_pir();
  debugSerial.println();
}
