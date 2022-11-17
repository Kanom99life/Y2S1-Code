TITLE STUDENT AVERAGE
.MODEL SMALL
.STACK 100H
.DATA
    AVERAGE     DW  4 DUP(0)
    CLASS       DB 'MARY ALLEN  ', 67, 45, 98, 33
                DB 'SCOTT BAYLIS', 70, 56, 87, 44
                DB 'GEORGE FRANK', 82, 72, 89, 40
                DB 'SAM WONG    ', 78, 76, 92, 60 
.CODE
MAIN PROC
    MOV AX, @DATA
    MOV DS, AX

    LEA DI, AVERAGE              ; set DI=offset address of variable AVERAGE
    LEA SI, CLASS                ; set SI=offset address of variable CLASS 
    ADD SI, 12                   ; set SI=SI+12
    MOV CX, 4                    ; set CX=4

COMPUTE_AVERAGE:                  
    XOR AX, AX                    ; clear AX
    MOV DX, 4                     ; set DX=4

SUM:                      
    XOR BX, BX                     ; clear BX
    MOV BL, [SI]                   ; set BL=[SI]

    ADD AX, BX                     ; set AX=AX+BX

    INC SI                         ; set SI=SI+1
    DEC DX                         ; set DX=DX-1
    JNZ SUM                   

    MOV BX, 4                      ; set BX=4
    DIV BX                         ; set AX=DX:AX/BX

    MOV [DI], AX                   ; set [DI]=AX
    ADD DI, 2                      ; set DI=DI+2
    ADD SI, 12                     ; set SI=SI+12
    LOOP COMPUTE_AVERAGE        

    LEA SI, AVERAGE                ; set SI=offset address of variable AVERAGE
    LEA DI, CLASS                  ; set DI=offset address of variable CLASS
    MOV CX, 4                      ; set CX=4

PRINT_RESULT:              
    MOV BX, 12                     ; set BX=12
    MOV AH, 2                      ; set output function

NAME_:                     
    MOV DL, [DI]                   ; set DL=[DI]
    INT 21H                        ; print a character

    INC DI                         ; set DI=DI+1
    DEC BX                         ; set BX=BX-1
    JNZ NAME_                 

    MOV DL, 20H                    ; white space
    INT 21H                    

    MOV DL, ":"               
    INT 21H                   

    MOV DL, 20H                
    INT 21H                    

    XOR AH, AH                     ; clear AH
    MOV AL, [SI]                   ; set AL=[SI]

    CALL OUTDEC                
    MOV AH, 2                    
    MOV DL, 0DH                    ; carriage return
    INT 21H                    

    MOV DL, 0AH                    ; line feed
    INT 21H

    ADD SI, 2                      ; set SI=SI+2
    ADD DI, 4                      ; set DI=DI+4
    LOOP PRINT_RESULT           

    MOV AH, 4CH    
    INT     21H              
MAIN ENDP 
    INCLUDE OUTDEC.asm
END MAIN