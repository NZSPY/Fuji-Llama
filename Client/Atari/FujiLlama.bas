' Fuji-Llama Title Page
' Written in Atari FastBasic
' @author  Simon Young

' Fuji-Net Setup Variblies 
UNIT=1
DIM RESULT(1024) BYTE
JSON_MODE=1
URL$="N:HTTP://192.168.68.100:8080/tables"
QUERY$=""


' Draw the opening screen 
GRAPHICS 0
SETCOLOR 2,0,0
SETCOLOR 1,14,6
?"*** Welcome to Fuji-Llama ***"
?"Choose a table to join"

? "press and key to open the connection"
@Wait
@openconnection

? "press and key to setup JSON"
@Wait

@nsetchannelmode 
@nparsejson
IF SErr()<>1
'PRINT "Could not parse JSON."
@nprinterror
@Wait
ENDIF

@getresult
@Wait

? $(&RESULT)

? "done"
@Wait
NCLOSE UNIT


PROC Wait
K=0
POKE 764,255
REPEAT 
K=key()
UNTIL K<>0
ENDPROC

' PROCEDURES to get Json data and load into the Var Result

PROC openconnection ' Open the connection, or throw error and end program
NOPEN UNIT, 12, 0, URL$
' If not successful, then exit.
IF SERR()<>1
PRINT "Could not open connection."
@nprinterror
EXIT
ELSE
PRINT "Horray"
ENDIF
ENDPROC

PROC nsetchannelmode ' Set the channel mode to the JSON_mode
SIO $71, UNIT, $FC, $00, 0, $1F, 0, 12, JSON_MODE
ENDPROC

PROC nparsejson ' send Parse to the FujiNet, so it parses the JSON value set by teh URL$
SIO $71, UNIT, $50, $00, 0, $1f, 0, 12, 0
ENDPROC

PROC njsonquery ' Querey the JSON data that has been parsesed base on the attributes in $query
SIO $71, UNIT, $51, $80, &query$+1, $1f, 256, 12, 0
ENDPROC

PROC nprinterror ' get the current eror and display on screen 
NSTATUS UNIT
PRINT "ERROR- "; PEEK($02ED)
ENDPROC

PROC getresult
@njsonquery
NSTATUS UNIT
IF PEEK($02ED) > 128
PRINT "Could not fetch query:"
PRINT QUERY$
PRINT "ERROR- "; PEEK($02ED)
EXIT
ENDIF
BW=DPEEK($02EA)
NGET UNIT, &RESULT, BW
' BPUT #0,  &RESULT, BW 
ENDPROC

