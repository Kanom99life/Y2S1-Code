TITLE LAB42
.MODEL SMALL
.STACK 100H
.DATA
    MSG     DB  0DH
    HALF    DB  ?,' Half', 0AH
    QUA     DB  ?,' Quater', 0AH
    DIME    DB  ?,' Dimes', 0AH
    NIC     DB  ?,' Nickels', 0AH
    PEN     DB  ?,' Pennies','$'
.CODE
MAIN PROC
    MOV     AX, @DATA
    MOV     DS, AX

    CALL    INDEC  ;Input cent in AX
    ;Find Dollar
    MOV     BL, 50  
    DIV     BL
    ADD     AL, 30H ;Change to digit
    MOV     HALF, AL
    ;Find Quater
    MOV     AL,AH
    MOV     AH,0
    MOV     BL, 25
    DIV     BL
    ADD     AL, 30H
    MOV     QUA, AL
    ;Find Dime
    MOV     AL,AH
    MOV     AH, 0
    MOV     BL, 10 
    DIV     BL
    ADD     AL, 30H
    MOV     DIME, AL
    ;Find Nickel
    MOV     AL, AH
    MOV     AH, 0
    MOV     BL, 5 
    DIV     BL 
    ADD     AL, 30H
    MOV     NIC, AL
    ;Find Penny
    ADD     AH, 30H
    MOV     PEN, AH
    ;Print prompt
    MOV     AH, 9
    LEA     DX,MSG
    INT     21H
    ;End
    MOV     AH,4CH
    INT     21H

MAIN ENDP 
    INCLUDE INDEC.asm
END MAIN