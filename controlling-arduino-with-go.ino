#include <Adafruit_NeoPixel.h>

#define BAUD 9600
#define SIZE 12
#define PIN 8

Adafruit_NeoPixel ring = Adafruit_NeoPixel(SIZE, PIN, NEO_GRB + NEO_KHZ800);

int currentLED = 0;

int currentRed = 255;
int currentGreen = 255;
int currentBlue = 255;

int currentPhase = 0;
int currentDirection = 1;
int phases[60];

String serialData = "";

void setup() {
  // init serial
  Serial.begin(BAUD);

  // init neopixel
  ring.begin();
  ring.setBrightness(100);
  ring.show();

  // map phases
  int i = 0;
  do {
    phases[i] = map(i, 0, 60, 20, 255);
    i++;
  } while (i < 60);
}

void loop() {
  if (!Serial) {
    ringPulse();

    return;
  }
  
  ringBright();

  colorUpdate();
}

void colorUpdate()
{
  while (Serial.available() > 0) {
    int red = Serial.parseInt();
    int green = Serial.parseInt();
    int blue = Serial.parseInt();

    if (Serial.read() == '\n') {
      red = constrain(red, 0, 255);
      green = constrain(green, 0, 255);
      blue = constrain(blue, 0, 255);

      currentRed = red;
      currentGreen = green;
      currentBlue = blue;

      Serial.print(red, DEC);
      Serial.print(',');
      Serial.print(green, DEC);
      Serial.print(',');
      Serial.println(blue, DEC);
    }
  }
}

void ringBright() {
  for (int i = 0; i < SIZE; i++) {
    ring.setPixelColor(i, currentRed, currentGreen, currentBlue);
  }
  ring.show();

  delay(1);
}

void ringPulse() {
  // update direction
  if (currentPhase == 0) {
    currentDirection = 1;
  }
  if (currentPhase == 59) {
    currentDirection = 0;
  }
  
  // set color
  int color = phases[currentPhase];
  
  // update ring
  for (int i = 0; i < SIZE; i++) {
    ring.setPixelColor(i, color, color, color);
  }
  ring.show();

  // update phase
  if (currentDirection == 1) {
    currentPhase += 1;
  }
  if (currentDirection == 0) {
    currentPhase -= 1;
  }
  
  delay(1);
}

void ringSpin() {
  currentLED += 1;
  if (currentLED == 12) {
    currentLED = 0;
  }

  for (int i = 0; i < SIZE; i++) {
    if (i == currentLED) {
      ring.setPixelColor(i, ring.Color(currentRed, currentGreen, currentBlue));
    } else {
      ring.setPixelColor(i, ring.Color(0, 0, 0));
    }
  }
  ring.show();

  delay(32);
}
