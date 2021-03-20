#include "textflag.h"

// void lock(int32 *ptr, int32 old, int32 new)
TEXT ·lock(SB), NOSPLIT, $0-16
	MOVQ	ptr+0(FP), BX
	MOVL	old+8(FP), AX
	MOVL	new+12(FP), CX
again:
	LOCK
	CMPXCHGL	CX, 0(BX)
	JE		ok
	PAUSE
	JMP		again
ok:
	RET

// void unlock(int32 *ptr, int32 val)
TEXT ·unlock(SB), NOSPLIT, $0-12
	MOVQ	ptr+0(FP), BX
	MOVL	val+8(FP), AX
	XCHGL	AX, 0(BX)
	RET
