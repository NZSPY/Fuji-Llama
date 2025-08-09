; Imported symbols

; Exported symbols
	.export bytecode_start

	.include "target.inc"

; TOKENS:
	.importzp	TOK_0
	.importzp	TOK_BYTE
	.importzp	TOK_CLOSE
	.importzp	TOK_CNJUMP
	.importzp	TOK_COMP_0
	.importzp	TOK_CSTRING
	.importzp	TOK_END
	.importzp	TOK_GETKEY
	.importzp	TOK_GRAPHICS
	.importzp	TOK_NUM
	.importzp	TOK_NUM_POKE
	.importzp	TOK_PEEK
	.importzp	TOK_PMGRAPHICS
	.importzp	TOK_PRINT_STR
	.importzp	TOK_VAR_STORE
;-----------------------------
; Macro to get variable ID from name
	.import __HEAP_RUN__
.macro makevar name
	.byte <((.ident (.concat ("fb_var_", name)) - __HEAP_RUN__)/2)
.endmacro
; Variables
	.segment "HEAP"
	.export fb_var____DEBUG_KEY
fb_var____DEBUG_KEY:	.res 2	; Word variable
;-----------------------------
; Bytecode
	.segment "BYTECODE"
bytecode_start:
@FastBasic_LINE_6:	; LINE 6
	.byte	TOK_0
	.byte	TOK_PMGRAPHICS
	.byte	TOK_BYTE
	.byte	6
	.byte	TOK_CLOSE
	.byte	TOK_0
	.byte	TOK_GRAPHICS
@FastBasic_LINE_7:	; LINE 7
	.byte	TOK_0
	.byte	TOK_NUM_POKE
	.word	710
@FastBasic_LINE_8:	; LINE 8
	.byte	TOK_BYTE
	.byte	230
	.byte	TOK_NUM_POKE
	.word	709
@FastBasic_LINE_10:	; LINE 10
	.byte	TOK_CSTRING
	.byte	19, "*** FUJI-LLAMA ***", 155
	.byte	TOK_PRINT_STR
@FastBasic_LINE_16:	; LINE 16
jump_lbl_1:
@FastBasic_LINE_17:	; LINE 17
	.byte	TOK_NUM
	.word	644
	.byte	TOK_PEEK
	.byte	TOK_COMP_0
	.byte	TOK_CNJUMP
	.word	jump_lbl_1
@FastBasic_LINE_20:	; LINE 20
	.byte	TOK_GETKEY
	.byte	TOK_VAR_STORE
	makevar	"___DEBUG_KEY"
	.byte	TOK_END
