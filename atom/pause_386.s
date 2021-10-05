#include "textflag.h"

// void pause(cnt int)
TEXT ·pause(SB), NOSPLIT, $0-4
	MOVL    cnt+0(FP), CX
again:
	PAUSE
	LOOP    again
	RET
