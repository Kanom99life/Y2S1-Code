TITLE LAB41
.MODEL SMALL
.STACK 100H
.DATA
    H   DB  " Hour(s)$"
    M   DB  " Minute(s)$"
    S   DB  " Second(s)$"
.CODE
MAIN PROC
    MOV     AX, @DATA
    MOV     DS, AX

    CALL    INDEC
    MOV     DX, 0 
    MOV     BX, 3600
    DIV     BX
    PUSH    AX
    PUSH    DX
    
    CALL    OUTDEC ;Hour in AX
    MOV     AH, 9
    LEA     DX, H
    INT     21H
    POP     DX
    
    PUSH    DX
    MOV     AH, 2
    MOV     DL, 20H
    INT     21H

    POP     AX
    MOV     BX,60
    MOV     DX, 0
    DIV     BX
    PUSH    AX
    PUSH    DX

    CALL    OUTDEC ;Minute in AX
    MOV     AH, 9
    LEA     DX, M
    INT     21H
    POP     DX
   
    PUSH    DX 
    MOV     AH, 2 
    MOV     DL, 20H
    INT     21H
    POP     AX
    
    CALL    OUTDEC ;Second in AX
    MOV     AH, 9
    LEA     DX, S
    INT     21H
    ;End
    MOV     AH, 4CH
    INT     21H

MAIN ENDP
    INCLUDE OUTDEC.asm
    INCLUDE INDEC.asm
END MAIN