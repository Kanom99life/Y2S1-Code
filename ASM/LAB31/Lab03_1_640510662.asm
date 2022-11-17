TITLE LAB31
.MODEL SMALL
.STACK 100H
.CODE
MAIN PROC
    MOV     BL, 80H
    MOV     CL, 1
    
PRINT_CHAR:

    MOV     AH, 2
    MOV     DL, BL
    INT     21H

    MOV     DL, 20H ;Blank space
    INT     21H
    
    INC     BL      ;Increase ASCII code
    CMP     BL, 0FFH    
    JE      END_    ;If current ASCII equal to FFH jump to end
    
    CMP     CL,10   
    JE      NEW_LINE ; If 10 characters have been displayed make a new line.
    INC     CL  ;Increase counting

    JMP     PRINT_CHAR

NEW_LINE:
    MOV     AH, 2
    MOV     DL, 0AH ;Line Feed
    INT     21H
    MOV     DL, 0DH ;Carriage Return
    INT     21H
    MOV     CL, 1   ;set counting to 1
    JMP     PRINT_CHAR

END_:
    MOV AH,4CH
    INT 21H

MAIN    ENDP
        END MAIN