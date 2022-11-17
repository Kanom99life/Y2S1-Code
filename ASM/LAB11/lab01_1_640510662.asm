TITLE LAB11
.MODEL SMALL
.STACK 100H
.DATA
	;data definitions go here
	CR		EQU 	0DH
	LF		EQU 	0AH
	MSG 	DB 		CR, LF, 'THE SUM OF '
	NUM1	DB		?
	MSG1 	DB 		' AND '
	NUM2 	DB		?
	MSG2 	DB 		' IS '
	VALUE	DB		?, '$'


.CODE
MAIN	PROC
; initailize DS 
	MOV 	AX, @DATA
	MOV		DS, AX 		; intialize DS
	
; display prompt
	MOV		AH, 2		; display character function
	MOV		DL, '?'		; character is '?'
	INT		21H

; input number1
	MOV 	AH, 1	;input function
	INT		21H
	MOV		NUM1, AL	;move input to variable

; input number2
	INT		21H
	MOV		NUM2, AL ;move input to variable
	
;sum
	MOV 	AL, NUM1  ;move value of NUM1 to AL
	ADD 	VALUE, AL ;add AL with VALUE

	MOV 	AL, NUM2 ;move value of NUM2 to AL
	ADD 	VALUE, AL ;add AL with VALUE

	SUB 	VALUE, 30H ;sub value of ascii code

;display prompt
	LEA	DX,MSG ; get message
	MOV	AH, 9 ; display string function
	INT	21H	; display message

; DOS exit
	MOV		AH, 4CH
	INT		21H		; DOS exit
MAIN	ENDP
		END MAIN