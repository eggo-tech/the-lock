#include "textflag.h"

// void lock(int32 *ptr, int32 old, int32 new)
TEXT ·lock(SB), NOSPLIT, $0-12
	MOVL	ptr+0(FP), BX
	MOVL	old+4(FP), AX
	MOVL	new+8(FP), CX
again:
	LOCK
	CMPXCHGL	CX, 0(BX)
	JE		ok
	PAUSE
	JMP		again
ok:
	RET

// void unlock(int32 *ptr, int32 val)
TEXT ·unlock(SB), NOSPLIT, $0-8
	MOVL	ptr+0(FP), BX
	MOVL	val+4(FP), AX
	XCHGL	AX, 0(BX)
	RET
