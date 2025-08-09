' Fuji-Llama Title Page
' Written in Atari FastBasic
' @author  Simon Young

DIM result(1024) BYTE
JSON_MODE=1

GRAPHICS 0
SETCOLOR 2,0,0
SETCOLOR 1,14,6


?"*** Welcome to Fuji-Llama ***"

?"Choose a table to join"

unit=8
URL$="N:HTTP://192.168.68.100:8080/tables"
query$="N:tables"$9B
' Open connection
NOPEN unit, 12, 0, URL$
' If not successful, then exit.
IF SERR()<>1
PRINT "Could not open connection."
NSTATUS unit
PRINT "ERROR- "; PEEK($02ED)
ELSE

PRINT "horray"


ENDIF


NCLOSE unit

REPEAT 
k=key()
UNTIL k<>0

' PROCEDURES '''''''''''''''''''''''''
PROC nprinterror
NSTATUS unit
PRINT "ERROR- "; PEEK($02ED)
ENDPROC
PROC nsetchannelmode mode
SIO $71, unit, $FC, $00, 0, $1F, 0, 12, JSON_MODE
ENDPROC
PROC nparsejson
SIO $71, unit, $50, $00, 0, $1f, 0, 12, 0
ENDPROC
PROC njsonquery
SIO $71, unit, $51, $80, &query$+1, $1f, 256, 12, 0
ENDPROC
PROC showresult
@njsonquery
'NSTATUS unit
'IF PEEK($02ED) > 0
'PRINT "Could not fetch query:"
'PRINT query$
'PRINT "ERROR- "; PEEK($02ED)
'EXIT
'ENDIF
BW=DPEEK($02EA)
NGET unit, &result, BW
BPUT #0,  &result, BW
ENDPROC

