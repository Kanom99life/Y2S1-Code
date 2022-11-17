 TITLE LAB21
.MODEL SMALL
.STACK 100H
.DATA
   ; MSG     DB      "?$"
    CH1     DB      ?
    CH2     DB      ?


.CODE
MAIN    PROC
; initailize DS
    MOV     AX, @DATA
    MOV     DS, AX

; display prompt
	MOV		AH, 2		; display character function
	MOV		DL, '?'		; character is '?'
	INT		21H

; input first
    MOV     AH, 1
    INT     21H 
    MOV     CH1, AL

; input second
    INT     21H
    MOV     CH2, AL 

;Compare
    MOV BH, CH1
    MOV BL, CH2
    CMP BH, BL
    JGE SWAP_
    
 ;display
    MOV     DL, 10
    MOV     AH, 02H
    INT     21H

    MOV 	AH, 2
	MOV 	DL, CH1
	INT 	21H

    MOV 	AH, 2
	MOV 	DL, CH2
	INT 	21H
    JMP DOS_EXIT

;CH1 <  CH2 SWAP
SWAP_:
;line feed
    MOV     DL, 10
    MOV     AH, 02H
    INT     21H
	
    MOV 	AH, 2
	MOV 	DL, CH2
	INT 	21H

	MOV 	AH, 2
	MOV 	DL, CH1
	INT 	21H


DOS_EXIT:
    MOV     AH, 4CH
    INT     21H 

MAIN    ENDP
        END MAIN