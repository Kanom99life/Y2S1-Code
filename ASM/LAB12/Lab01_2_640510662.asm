TITLE LAB12
.MODEL SMALL
.STACK 100H
.DATA
    MSG     DB      "ENTER THEREE INITIALS: $"
    CH1     DB      ?
    CH2     DB      ?
    CH3     DB      ?

.CODE
MAIN    PROC
; initailize DS
    MOV     AX, @DATA
    MOV     DS, AX

; print user prompt
    LEA     DX, MSG
    MOV     AH, 9
    INT     21H

; input first
    MOV     AH, 1
    INT     21H 
    MOV     CH1, AL

; input middle
    INT     21H
    MOV     CH2, AL 

; input last
    INT     21H 
    MOV     CH3, AL

; line feed
    MOV     DL, 10
    MOV     AH, 02H
    INT     21H
 
; display First char
	MOV 	AH, 2
	MOV 	DL, CH1
	INT 	21H

; line feed
    MOV     DL, 10
    MOV     AH, 02H
    INT     21H

; display Middle char
	MOV 	AH, 2
	MOV 	DL, CH2
	INT 	21H

; line feed
    MOV     DL, 10
    MOV     AH, 02H
    INT     21H

; display Last char
	MOV 	AH, 2
	MOV 	DL, CH3
	INT 	21H

; DOS exit
    MOV     AH, 4CH
    INT     21H 

MAIN    ENDP
        END MAIN