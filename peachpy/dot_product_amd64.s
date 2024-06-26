// +build !noasm
// Generated by PeachPy 0.2.0 from dot_product.py


// func DotProduct(x *float32, y *float32, length uint) float32
TEXT ·DotProduct(SB),4,$0-28
	MOVQ x+0(FP), AX
	MOVQ y+8(FP), BX
	MOVQ length+16(FP), CX
	BYTE $0xC5; BYTE $0xF8; BYTE $0x57; BYTE $0xC0 // VXORPS xmm0, xmm0, xmm0
	BYTE $0xC5; BYTE $0xF0; BYTE $0x57; BYTE $0xC9 // VXORPS xmm1, xmm1, xmm1
	BYTE $0xC5; BYTE $0xE8; BYTE $0x57; BYTE $0xD2 // VXORPS xmm2, xmm2, xmm2
	BYTE $0xC5; BYTE $0xE0; BYTE $0x57; BYTE $0xDB // VXORPS xmm3, xmm3, xmm3
	BYTE $0xC5; BYTE $0xD8; BYTE $0x57; BYTE $0xE4 // VXORPS xmm4, xmm4, xmm4
	BYTE $0xC5; BYTE $0xD0; BYTE $0x57; BYTE $0xED // VXORPS xmm5, xmm5, xmm5
	SUBQ $48, CX
	JCS vector_loop_end
vector_loop_begin:
		BYTE $0xC5; BYTE $0xFC; BYTE $0x10; BYTE $0x30 // VMOVUPS ymm6, [rax]
		BYTE $0xC5; BYTE $0xFC; BYTE $0x10; BYTE $0x78; BYTE $0x20 // VMOVUPS ymm7, [rax + 32]
		BYTE $0xC5; BYTE $0x7C; BYTE $0x10; BYTE $0x40; BYTE $0x40 // VMOVUPS ymm8, [rax + 64]
		BYTE $0xC5; BYTE $0x7C; BYTE $0x10; BYTE $0x48; BYTE $0x60 // VMOVUPS ymm9, [rax + 96]
		BYTE $0xC5; BYTE $0x7C; BYTE $0x10; BYTE $0x90; BYTE $0x80; BYTE $0x00; BYTE $0x00; BYTE $0x00 // VMOVUPS ymm10, [rax + 128]
		BYTE $0xC5; BYTE $0x7C; BYTE $0x10; BYTE $0x98; BYTE $0xA0; BYTE $0x00; BYTE $0x00; BYTE $0x00 // VMOVUPS ymm11, [rax + 160]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x4D; BYTE $0x98; BYTE $0x03 // VFMADD132PS ymm0, ymm6, [rbx]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x45; BYTE $0x98; BYTE $0x4B; BYTE $0x20 // VFMADD132PS ymm1, ymm7, [rbx + 32]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x3D; BYTE $0x98; BYTE $0x53; BYTE $0x40 // VFMADD132PS ymm2, ymm8, [rbx + 64]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x35; BYTE $0x98; BYTE $0x5B; BYTE $0x60 // VFMADD132PS ymm3, ymm9, [rbx + 96]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x2D; BYTE $0x98; BYTE $0xA3; BYTE $0x80; BYTE $0x00; BYTE $0x00; BYTE $0x00 // VFMADD132PS ymm4, ymm10, [rbx + 128]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x25; BYTE $0x98; BYTE $0xAB; BYTE $0xA0; BYTE $0x00; BYTE $0x00; BYTE $0x00 // VFMADD132PS ymm5, ymm11, [rbx + 160]
		ADDQ $192, AX
		ADDQ $192, BX
		SUBQ $48, CX
		JCC vector_loop_begin
vector_loop_end:
	BYTE $0xC5; BYTE $0xFC; BYTE $0x58; BYTE $0xC1 // VADDPS ymm0, ymm0, ymm1
	BYTE $0xC5; BYTE $0xEC; BYTE $0x58; BYTE $0xD3 // VADDPS ymm2, ymm2, ymm3
	BYTE $0xC5; BYTE $0xDC; BYTE $0x58; BYTE $0xE5 // VADDPS ymm4, ymm4, ymm5
	BYTE $0xC5; BYTE $0xFC; BYTE $0x58; BYTE $0xC2 // VADDPS ymm0, ymm0, ymm2
	BYTE $0xC5; BYTE $0xFC; BYTE $0x58; BYTE $0xC4 // VADDPS ymm0, ymm0, ymm4
	BYTE $0xC5; BYTE $0xF0; BYTE $0x57; BYTE $0xC9 // VXORPS xmm1, xmm1, xmm1
	ADDQ $48, CX
	JEQ scalar_loop_end
scalar_loop_begin:
		BYTE $0xC5; BYTE $0xFA; BYTE $0x10; BYTE $0x10 // VMOVSS xmm2, [rax]
		BYTE $0xC4; BYTE $0xE2; BYTE $0x69; BYTE $0x99; BYTE $0x0B // VFMADD132SS xmm1, xmm2, [rbx]
		ADDQ $4, AX
		ADDQ $4, BX
		SUBQ $1, CX
		JNE scalar_loop_begin
scalar_loop_end:
	BYTE $0xC5; BYTE $0xFC; BYTE $0x58; BYTE $0xC1 // VADDPS ymm0, ymm0, ymm1
	BYTE $0xC4; BYTE $0xE3; BYTE $0x7D; BYTE $0x19; BYTE $0xC1; BYTE $0x01 // VEXTRACTF128 xmm1, ymm0, 1
	BYTE $0xC5; BYTE $0xF8; BYTE $0x58; BYTE $0xC1 // VADDPS xmm0, xmm0, xmm1
	BYTE $0xC5; BYTE $0xFB; BYTE $0x7C; BYTE $0xC0 // VHADDPS xmm0, xmm0, xmm0
	BYTE $0xC5; BYTE $0xFB; BYTE $0x7C; BYTE $0xC0 // VHADDPS xmm0, xmm0, xmm0
	MOVSS X0, ret+24(FP)
	BYTE $0xC5; BYTE $0xF8; BYTE $0x77 // VZEROUPPER
	RET
