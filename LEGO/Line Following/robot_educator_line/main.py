#!/usr/bin/env pybricks-micropython

"""
Example LEGO® MINDSTORMS® EV3 Robot Educator Color Sensor Down Program
----------------------------------------------------------------------

This program requires LEGO® EV3 MicroPython v2.0.
Download: https://education.lego.com/en-us/support/mindstorms-ev3/python-for-ev3

Building instructions can be found at:
https://education.lego.com/en-us/support/mindstorms-ev3/building-instructions#robot
"""
from pybricks.hubs import EV3Brick
from pybricks.ev3devices import Motor, ColorSensor, GyroSensor
from pybricks.parameters import Port,Direction, Color
from pybricks.tools import wait,DataLog, StopWatch
from pybricks.robotics import DriveBase


def main():
    PID()

def line_follower():
    

    # Initialize the motors.
    left_motor = Motor(Port.B)
    right_motor = Motor(Port.C)

    # Initialize the color sensor.
    line_sensor = ColorSensor(Port.S3)

    # Initialize the drive base.
    robot = DriveBase(left_motor, right_motor, wheel_diameter=55.5, axle_track=104)

    # Calculate the light threshold. Choose values based on your measurements.
    color = color_calibration()
    BLACK = color[1]
    WHITE = color[0]
    threshold = (BLACK + WHITE) / 2

    # Set the drive speed at 100 millimeters per second.
    DRIVE_SPEED = 100

    # Set the gain of the proportional line controller. This means that for every
    # percentage point of light deviating from the threshold, we set the turn
    # rate of the drivebase to 1.2 degrees per second.

    # For example, if the light value deviates from the threshold by 10, the robot
    # steers at 10*1.2 = 12 degrees per second.
    PROPORTIONAL_GAIN = 1.2

    # Start following the line endlessly.
    while True:
        # Calculate the deviation from the threshold.
        deviation = line_sensor.reflection() - threshold

        # Calculate the turn rate.
        turn_rate = PROPORTIONAL_GAIN * deviation

        # Set the drive base speed and turn rate.
        robot.drive(DRIVE_SPEED, turn_rate)

        # You can wait for a short time or do other things in this loop.
        wait(10)

def color_calibration():
    ev3 = EV3Brick()
    color_sensor = ColorSensor(Port.S3)
    # print("Press to calibrate")
    ev3.screen.draw_text(0, 50, 'Press button')
    while not any(ev3.buttons.pressed()):
        continue
    wait(1000)
    white_color = color_sensor.reflection()
    ev3.speaker.beep()
    ev3.screen.clear()
    ev3.screen.draw_text(0, 20, 'White color: {}'.format(white_color))
    # print("White color::", white_color)

    while not any(ev3.buttons.pressed()):
        continue
    wait(1000)
    black_color = color_sensor.reflection()
    ev3.speaker.beep()
    ev3.screen.draw_text(0, 40, 'Black color: {}'.format(black_color))
    
    #while not any(ev3.buttons.pressed()):
     #   continue
    #wait(1000)
    #green_color = color_sensor.reflection()
    #ev3.speaker.beep()
    #ev3.screen.draw_text(0, 60, 'Start color: {}'.format(green_color))
    # print("Black color:", black_color)
    
    #while not any(ev3.buttons.pressed()):
     #   continue
    #wait(1000)
    #brown_color = color_sensor.reflection()
    #ev3.speaker.beep()
    #ev3.screen.draw_text(0, 80, 'Stop color: {}'.format(brown_color))
    #list_color = [white_color, black_color, green_color, brown_color]
    list_color =  [white_color, black_color]
    return list_color

def PID():
    ev3 = EV3Brick()
    line_sensor = ColorSensor(Port.S3)
    left_motor = Motor(Port.B,Direction.CLOCKWISE)
    right_motor = Motor(Port.C,Direction.CLOCKWISE)
    gyro_sensor = GyroSensor(Port.S1)
    
   
    #motor_speedlog = DataLog('error', 'integral', 'derivative', 'turn_rate', name = 'motor_speedlog')
    data = DataLog('time', 'angle')
    

    
    color = color_calibration()
    EPSILON = 2.0
    WHITE = color[0]
    BLACK = color[1]
    #START = color[2]
    #STOP = color[3]
    while not any(ev3.buttons.pressed()):
        continue
    wait(1000)
    
    thresholod1 = int((WHITE - BLACK) / 3) + BLACK
    thresholod2 = WHITE - int((WHITE - BLACK) / 3)
    threshold = (BLACK + WHITE) / 2
    DRIVE_SPEED = 100
    PROPORTIONAL_GAIN = 4.2
    INTEGRAL_GAIN = 0.008
    DERIVATIVE_GAIN = 0.01
    integral = 0
    derivative = 0
    last_error = 0
    count = 0
    
    watch = StopWatch()
    while True:
        
        #if (line_sensor.reflection() == START): #or (line_sensor <= abs(EPSILON - GREEN)) or (line_sensor >= EPSILON + GREEN):
            # count += 1
        
        angle = gyro_sensor.angle()
        time = watch.time()
        if (line_sensor.color() == Color.RED): #: #or (line_sensor <= abs(BROWN -EPSILON)) or (line_sensor >= EPSILON + GREEN):
            watch.pause()
            data.log(time, angle)
            break
        
        if (line_sensor.reflection() < thresholod1) or (line_sensor.reflection() > thresholod2):
            error = line_sensor.reflection()-threshold
            integral += error
            derivative = error - last_error
        
            turn_rate = PROPORTIONAL_GAIN * error + INTEGRAL_GAIN * integral + DERIVATIVE_GAIN * derivative
            #motor_speedlog.log(error, integral, derivative, turn_rate)
            left_motor.run(DRIVE_SPEED + turn_rate)
            right_motor.run(DRIVE_SPEED - turn_rate)
            data.log(time, angle)
            # data.log(time, angle)
        
            # if count >= 1:
                # data.log(time, angle)
        else:
            gyro_sensor.reset_angle(0)
            error = line_sensor.reflection()-threshold
            integral += error
            derivative = error - last_error
            turn_rate = PROPORTIONAL_GAIN * error + INTEGRAL_GAIN * integral + DERIVATIVE_GAIN * derivative
           
            left_motor.run(200 + turn_rate)
            right_motor.run(200 - turn_rate)
            data.log(time, angle)
             
        
        
        wait(10)
        
        
if __name__ == '__main__':
    main()