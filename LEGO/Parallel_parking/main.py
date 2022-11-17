#!/usr/bin/env pybricks-micropython


from pybricks.hubs import EV3Brick
from pybricks.ev3devices import Motor, UltrasonicSensor
from pybricks.parameters import Port
from pybricks.tools import wait
from pybricks.robotics import DriveBase


ev3 = EV3Brick()

obstacle_sensor = UltrasonicSensor(Port.S4)


left_motor = Motor(Port.B)
right_motor = Motor(Port.C)

robot = DriveBase(left_motor, right_motor, wheel_diameter=55.5, axle_track=104)

count1 = 0
count2 = 0
ev3.speaker.beep()
while True:

    robot.drive(200, 0)

    # Wait until an obstacle is detected. This is done by repeatedly
    # doing nothing (waiting for 10 milliseconds) while the measured
    # distance is still greater than 300 mm.
    #1st box
    if obstacle_sensor.distance() <= 100:
        ev3.screen.clear()
        count1 += 1
        #print("1")
        #robot.drive(100, 0)
        ev3.screen.draw_text(0, 50, "Distance: {}mm".format(obstacle_sensor.distance()))
        #wait(10)
        
    
    if count1 > 0:
        if obstacle_sensor.distance() >= 100:
        #print("2")
            ev3.screen.clear()
            ev3.screen.draw_text(0, 50, "Distance: {}mm".format(obstacle_sensor.distance()))
            count2 += 1
            #robot.drive(200, 0)
            #wait(50)
        
            #2nd box
    if count2 > 0:
        if obstacle_sensor.distance() <= 100:
            ev3.screen.clear()
            ev3.screen.draw_text(0, 50, "Distance: {}mm".format(obstacle_sensor.distance()))
            ev3.speaker.beep()
            robot.straight(-38)
            robot.turn(85)
            
            
            #while  obstacle_sensor.distance() >= 200:
            #ev3.screen.clear()
            #ev3.screen.draw_text(0, 50, "Distance: {}mm".format(obstacle_sensor.distance()))
            #robot.drive(-100, 0)
            #if obstacle_sensor.distance() >= 200:
            robot.straight(-245)
            robot.turn(-85)
            robot.straight(25)
                
            break
                
                
            # Drive backward for 300 millimeters.
    
    
    
    #robot.straight(-300)

    # Turn around by 120 degrees
    #robot.turn(120)
