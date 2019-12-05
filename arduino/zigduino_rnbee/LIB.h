#include <TheThingsNetwork.h>
#include <CayenneLPP.h>

long randNumber;

// Set your AppEUI and AppKey
const char *appEui_ = "3d8f545b53302abe";
const char *appKey_ = "08e78f1ef55044e77df23763fe81aadc";
const char *devEui_ = "be2a30535b548f3d";

#define loraSerial Serial1
#define debugSerial Serial

#define freqPlan TTN_FP_EU868
#define SFF 7
#define SFW 7

// Replace REPLACE_ME with TTN_FP_EU868 or TTN_FP_US915
#define freqPlan TTN_FP_EU868

/* TTN */
TheThingsNetwork ttn(loraSerial, debugSerial, freqPlan, SFF);
CayenneLPP lpp(51);

/***************************************/
/* Sensors */

#define SENSOR_LIGHT A1
#define SENSOR_SOUND A2
#define SENSOR_PIR 4

int pir;
int light;
int sound;
