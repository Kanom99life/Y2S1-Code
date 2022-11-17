#Group 16 Sec 1
#Armita Phothiklang 640510636
#Thiraphat Nilsiri 640510662
#Suphakit Ng 640510682

import random
from pybricks.hubs import EV3Brick
from pybricks.ev3devices import Motor, UltrasonicSensor, GyroSensor
from pybricks.parameters import Port
from pybricks.tools import wait
from pybricks.robotics import DriveBase

# Initialize the EV3 Brick.
ev3 = EV3Brick()

# Initialize the Ultrasonic Sensor. It is used to detect
# obstacles as the robot drives around.
obstacle_sensor = UltrasonicSensor(Port.S4)

# Initialize two motors with default settings on Port B and Port C.
# These will be the left and right motors of the drive base.
left_motor = Motor(Port.B)
right_motor = Motor(Port.C)
gyro_sensor = GyroSensor(Port.S1)

# The DriveBase is composed of two motors, with a wheel on each motor.
# The wheel_diameter and axle_track values are used to make the motors
# move at the correct speed when you give a motor command.
# The axle track is the distance between the points where the wheels
# touch the ground.
robot = DriveBase(left_motor, right_motor, wheel_diameter=55.5, axle_track=104)

# Play a sound to tell us when we are ready to start moving
ev3.speaker.beep()

# The following loop makes the robot drive forward until it detects an
# obstacle. Then it backs up and turns around. It keeps on doing this
# until you stop the program.
while True:
    # Begin driving forward at 200 millimeters per second.
    robot.drive(200, 0)


    while obstacle_sensor.distance() > 180:
        ev3.screen.clear()
        ev3.screen.draw_text(0, 50, "Distance: {}mm".format(obstacle_sensor.distance()))
        wait(20)
    
    ev3.speaker.beep()
        
    gyroValue = gyro_sensor.angle()
    
    mylist = ["L", "R"]
    LR = random.choice(mylist)
    if (LR == "L"):
    
        robot.turn(-115)
    
        if gyroValue < -90:
            ev3.screen.draw_text(0, 80, "Turn left {} degrees".format(abs(gyroValue)))
            robot.turn(gyroValue - 90)
            gyroValue = 0
            wait(10)
        else:
            ev3.screen.draw_text(0, 80, "Turn left {} degrees".format(abs(gyroValue)))
            gyroValue = 0
            wait(10)
        
    elif (LR == "R"):
        
        robot.turn(115)
    
        if gyroValue < 90:
            ev3.screen.draw_text(0, 80, "Turn right {} degrees".format(gyroValue))
            robot.turn(abs(gyroValue - 90))
            gyroValue = 0
            wait(10)
        else:
            ev3.screen.draw_text(0, 80, "Turn right {} degrees".format(gyroValue))
            gyroValue = 0
            wait(10)
    