#!/usr/bin/env pybricks-micropython
from pybricks import ev3brick as brick
from pybricks.hubs import EV3Brick
from pybricks.ev3devices import (Motor, TouchSensor, ColorSensor,
                                 InfraredSensor, UltrasonicSensor, GyroSensor)
from pybricks.parameters import Port, Stop, Direction, Button, Color
from pybricks.tools import wait, StopWatch, DataLog
from pybricks.robotics import DriveBase
from pybricks.media.ev3dev import SoundFile, ImageFile


# This program requires LEGO EV3 MicroPython v2.0 or higher.
# Click "Open user guide" on the EV3 extension tab for more information.


# Create your objects here.
ev3 = EV3Brick()


# Write your program here.
ev3.speaker.beep()

brick.sound.beep()
# Initialize a mortor at port B.
left = Motor(Port.B)

test_motor.run_target(500, 90)

# Play another beep sound.
# This time with a higher pitch(1000 Hz) and longer duration(500 ms)
brick.sound.beep(1000, 500)
