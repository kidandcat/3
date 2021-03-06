package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for resize kernel
var resize_code cu.Function

// Stores the arguments for resize kernel invocation
type resize_args_t struct {
	arg_dst    unsafe.Pointer
	arg_Dx     int
	arg_Dy     int
	arg_Dz     int
	arg_src    unsafe.Pointer
	arg_Sx     int
	arg_Sy     int
	arg_Sz     int
	arg_layer  int
	arg_scalex int
	arg_scaley int
	argptr     [11]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for resize kernel invocation
var resize_args resize_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	resize_args.argptr[0] = unsafe.Pointer(&resize_args.arg_dst)
	resize_args.argptr[1] = unsafe.Pointer(&resize_args.arg_Dx)
	resize_args.argptr[2] = unsafe.Pointer(&resize_args.arg_Dy)
	resize_args.argptr[3] = unsafe.Pointer(&resize_args.arg_Dz)
	resize_args.argptr[4] = unsafe.Pointer(&resize_args.arg_src)
	resize_args.argptr[5] = unsafe.Pointer(&resize_args.arg_Sx)
	resize_args.argptr[6] = unsafe.Pointer(&resize_args.arg_Sy)
	resize_args.argptr[7] = unsafe.Pointer(&resize_args.arg_Sz)
	resize_args.argptr[8] = unsafe.Pointer(&resize_args.arg_layer)
	resize_args.argptr[9] = unsafe.Pointer(&resize_args.arg_scalex)
	resize_args.argptr[10] = unsafe.Pointer(&resize_args.arg_scaley)
}

// Wrapper for resize CUDA kernel, asynchronous.
func k_resize_async(dst unsafe.Pointer, Dx int, Dy int, Dz int, src unsafe.Pointer, Sx int, Sy int, Sz int, layer int, scalex int, scaley int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("resize")
	}

	resize_args.Lock()
	defer resize_args.Unlock()

	if resize_code == 0 {
		resize_code = fatbinLoad(resize_map, "resize")
	}

	resize_args.arg_dst = dst
	resize_args.arg_Dx = Dx
	resize_args.arg_Dy = Dy
	resize_args.arg_Dz = Dz
	resize_args.arg_src = src
	resize_args.arg_Sx = Sx
	resize_args.arg_Sy = Sy
	resize_args.arg_Sz = Sz
	resize_args.arg_layer = layer
	resize_args.arg_scalex = scalex
	resize_args.arg_scaley = scaley

	args := resize_args.argptr[:]
	cu.LaunchKernel(resize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("resize")
	}
}

// maps compute capability on PTX code for resize kernel.
var resize_map = map[int]string{0: "",
	20: resize_ptx_20,
	30: resize_ptx_30,
	35: resize_ptx_35,
	50: resize_ptx_50,
	52: resize_ptx_52,
	53: resize_ptx_53}

// resize PTX code for various compute capabilities.
const (
	resize_ptx_20 = `
.version 4.3
.target sm_20
.address_size 64

	// .globl	resize

.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<55>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r8, [resize_param_1];
	ld.param.u32 	%r14, [resize_param_2];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r9, [resize_param_5];
	ld.param.u32 	%r10, [resize_param_6];
	ld.param.u32 	%r11, [resize_param_8];
	ld.param.u32 	%r12, [resize_param_9];
	ld.param.u32 	%r13, [resize_param_10];
	mov.u32 	%r15, %ctaid.x;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r18, %r16, %r15, %r17;
	mov.u32 	%r19, %ntid.y;
	mov.u32 	%r20, %ctaid.y;
	mov.u32 	%r21, %tid.y;
	mad.lo.s32 	%r22, %r19, %r20, %r21;
	setp.lt.s32	%p1, %r18, %r8;
	setp.lt.s32	%p2, %r22, %r14;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB0_10;
	bra.uni 	BB0_1;

BB0_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r13, 1;
	@%p4 bra 	BB0_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r12, 1;
	@%p5 bra 	BB0_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r23, 0;
	cvta.to.global.u64 	%rd6, %rd5;
	mov.u32 	%r54, %r23;

BB0_4:
	mul.lo.s32 	%r51, %r12, %r18;
	mul.lo.s32 	%r33, %r13, %r22;
	mad.lo.s32 	%r34, %r11, %r10, %r33;
	mad.lo.s32 	%r35, %r9, %r34, %r51;
	mad.lo.s32 	%r36, %r9, %r54, %r35;
	mul.wide.s32 	%rd7, %r36, 4;
	add.s64 	%rd11, %rd6, %rd7;
	mov.u32 	%r53, %r23;

BB0_5:
	mov.u32 	%r4, %r53;
	mad.lo.s32 	%r41, %r22, %r13, %r54;
	setp.lt.s32	%p6, %r41, %r10;
	setp.lt.s32	%p7, %r51, %r9;
	and.pred  	%p8, %p6, %p7;
	@!%p8 bra 	BB0_7;
	bra.uni 	BB0_6;

BB0_6:
	ld.global.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB0_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r51, %r51, 1;
	add.s32 	%r6, %r4, 1;
	setp.lt.s32	%p9, %r6, %r12;
	mov.u32 	%r53, %r6;
	@%p9 bra 	BB0_5;

	add.s32 	%r54, %r54, 1;
	setp.lt.s32	%p10, %r54, %r13;
	@%p10 bra 	BB0_4;

BB0_9:
	mad.lo.s32 	%r50, %r22, %r8, %r18;
	cvta.to.global.u64 	%rd8, %rd4;
	mul.wide.s32 	%rd9, %r50, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB0_10:
	ret;
}


`
	resize_ptx_30 = `
.version 4.3
.target sm_30
.address_size 64

	// .globl	resize

.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<55>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r8, [resize_param_1];
	ld.param.u32 	%r14, [resize_param_2];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r9, [resize_param_5];
	ld.param.u32 	%r10, [resize_param_6];
	ld.param.u32 	%r11, [resize_param_8];
	ld.param.u32 	%r12, [resize_param_9];
	ld.param.u32 	%r13, [resize_param_10];
	mov.u32 	%r15, %ctaid.x;
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r17, %tid.x;
	mad.lo.s32 	%r18, %r16, %r15, %r17;
	mov.u32 	%r19, %ntid.y;
	mov.u32 	%r20, %ctaid.y;
	mov.u32 	%r21, %tid.y;
	mad.lo.s32 	%r22, %r19, %r20, %r21;
	setp.lt.s32	%p1, %r18, %r8;
	setp.lt.s32	%p2, %r22, %r14;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB0_10;
	bra.uni 	BB0_1;

BB0_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r13, 1;
	@%p4 bra 	BB0_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r12, 1;
	@%p5 bra 	BB0_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r23, 0;
	cvta.to.global.u64 	%rd6, %rd5;
	mov.u32 	%r54, %r23;

BB0_4:
	mul.lo.s32 	%r51, %r12, %r18;
	mul.lo.s32 	%r33, %r13, %r22;
	mad.lo.s32 	%r34, %r11, %r10, %r33;
	mad.lo.s32 	%r35, %r9, %r34, %r51;
	mad.lo.s32 	%r36, %r9, %r54, %r35;
	mul.wide.s32 	%rd7, %r36, 4;
	add.s64 	%rd11, %rd6, %rd7;
	mov.u32 	%r53, %r23;

BB0_5:
	mov.u32 	%r4, %r53;
	mad.lo.s32 	%r41, %r22, %r13, %r54;
	setp.lt.s32	%p6, %r41, %r10;
	setp.lt.s32	%p7, %r51, %r9;
	and.pred  	%p8, %p6, %p7;
	@!%p8 bra 	BB0_7;
	bra.uni 	BB0_6;

BB0_6:
	ld.global.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB0_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r51, %r51, 1;
	add.s32 	%r6, %r4, 1;
	setp.lt.s32	%p9, %r6, %r12;
	mov.u32 	%r53, %r6;
	@%p9 bra 	BB0_5;

	add.s32 	%r54, %r54, 1;
	setp.lt.s32	%p10, %r54, %r13;
	@%p10 bra 	BB0_4;

BB0_9:
	mad.lo.s32 	%r50, %r22, %r8, %r18;
	cvta.to.global.u64 	%rd8, %rd4;
	mul.wide.s32 	%rd9, %r50, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB0_10:
	ret;
}


`
	resize_ptx_35 = `
.version 4.3
.target sm_35
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	resize
.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<43>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r17, [resize_param_1];
	ld.param.u32 	%r23, [resize_param_2];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r18, [resize_param_5];
	ld.param.u32 	%r19, [resize_param_6];
	ld.param.u32 	%r20, [resize_param_8];
	ld.param.u32 	%r21, [resize_param_9];
	ld.param.u32 	%r22, [resize_param_10];
	mov.u32 	%r24, %ntid.x;
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r25, %r24, %r1, %r2;
	mov.u32 	%r3, %ntid.y;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %tid.y;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	setp.lt.s32	%p1, %r25, %r17;
	setp.lt.s32	%p2, %r6, %r23;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB6_10;
	bra.uni 	BB6_1;

BB6_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r22, 1;
	@%p4 bra 	BB6_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r21, 1;
	@%p5 bra 	BB6_9;

	cvta.to.global.u64 	%rd1, %rd6;
	mul.lo.s32 	%r7, %r6, %r22;
	mul.lo.s32 	%r8, %r21, %r25;
	mul.lo.s32 	%r30, %r22, %r6;
	mad.lo.s32 	%r31, %r20, %r19, %r30;
	mad.lo.s32 	%r9, %r18, %r31, %r8;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r26, 0;
	mov.u32 	%r42, %r26;

BB6_4:
	mad.lo.s32 	%r33, %r18, %r42, %r9;
	mul.wide.s32 	%rd7, %r33, 4;
	add.s64 	%rd11, %rd1, %rd7;
	add.s32 	%r11, %r42, %r7;
	mov.u32 	%r39, %r8;
	mov.u32 	%r41, %r26;

BB6_5:
	mov.u32 	%r13, %r41;
	mov.u32 	%r12, %r39;
	setp.lt.s32	%p6, %r12, %r18;
	setp.lt.s32	%p7, %r11, %r19;
	and.pred  	%p8, %p7, %p6;
	@!%p8 bra 	BB6_7;
	bra.uni 	BB6_6;

BB6_6:
	ld.global.nc.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB6_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r14, %r12, 1;
	add.s32 	%r15, %r13, 1;
	setp.lt.s32	%p9, %r15, %r21;
	mov.u32 	%r39, %r14;
	mov.u32 	%r41, %r15;
	@%p9 bra 	BB6_5;

	add.s32 	%r42, %r42, 1;
	setp.lt.s32	%p10, %r42, %r22;
	@%p10 bra 	BB6_4;

BB6_9:
	mad.lo.s32 	%r38, %r6, %r17, %r25;
	cvta.to.global.u64 	%rd8, %rd5;
	mul.wide.s32 	%rd9, %r38, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB6_10:
	ret;
}


`
	resize_ptx_50 = `
.version 4.3
.target sm_50
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	resize
.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<43>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r17, [resize_param_1];
	ld.param.u32 	%r23, [resize_param_2];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r18, [resize_param_5];
	ld.param.u32 	%r19, [resize_param_6];
	ld.param.u32 	%r20, [resize_param_8];
	ld.param.u32 	%r21, [resize_param_9];
	ld.param.u32 	%r22, [resize_param_10];
	mov.u32 	%r24, %ntid.x;
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r25, %r24, %r1, %r2;
	mov.u32 	%r3, %ntid.y;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %tid.y;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	setp.lt.s32	%p1, %r25, %r17;
	setp.lt.s32	%p2, %r6, %r23;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB6_10;
	bra.uni 	BB6_1;

BB6_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r22, 1;
	@%p4 bra 	BB6_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r21, 1;
	@%p5 bra 	BB6_9;

	cvta.to.global.u64 	%rd1, %rd6;
	mul.lo.s32 	%r7, %r6, %r22;
	mul.lo.s32 	%r8, %r21, %r25;
	mul.lo.s32 	%r30, %r22, %r6;
	mad.lo.s32 	%r31, %r20, %r19, %r30;
	mad.lo.s32 	%r9, %r18, %r31, %r8;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r26, 0;
	mov.u32 	%r42, %r26;

BB6_4:
	mad.lo.s32 	%r33, %r18, %r42, %r9;
	mul.wide.s32 	%rd7, %r33, 4;
	add.s64 	%rd11, %rd1, %rd7;
	add.s32 	%r11, %r42, %r7;
	mov.u32 	%r39, %r8;
	mov.u32 	%r41, %r26;

BB6_5:
	mov.u32 	%r13, %r41;
	mov.u32 	%r12, %r39;
	setp.lt.s32	%p6, %r12, %r18;
	setp.lt.s32	%p7, %r11, %r19;
	and.pred  	%p8, %p7, %p6;
	@!%p8 bra 	BB6_7;
	bra.uni 	BB6_6;

BB6_6:
	ld.global.nc.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB6_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r14, %r12, 1;
	add.s32 	%r15, %r13, 1;
	setp.lt.s32	%p9, %r15, %r21;
	mov.u32 	%r39, %r14;
	mov.u32 	%r41, %r15;
	@%p9 bra 	BB6_5;

	add.s32 	%r42, %r42, 1;
	setp.lt.s32	%p10, %r42, %r22;
	@%p10 bra 	BB6_4;

BB6_9:
	mad.lo.s32 	%r38, %r6, %r17, %r25;
	cvta.to.global.u64 	%rd8, %rd5;
	mul.wide.s32 	%rd9, %r38, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB6_10:
	ret;
}


`
	resize_ptx_52 = `
.version 4.3
.target sm_52
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	resize
.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<43>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r17, [resize_param_1];
	ld.param.u32 	%r23, [resize_param_2];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r18, [resize_param_5];
	ld.param.u32 	%r19, [resize_param_6];
	ld.param.u32 	%r20, [resize_param_8];
	ld.param.u32 	%r21, [resize_param_9];
	ld.param.u32 	%r22, [resize_param_10];
	mov.u32 	%r24, %ntid.x;
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r25, %r24, %r1, %r2;
	mov.u32 	%r3, %ntid.y;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %tid.y;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	setp.lt.s32	%p1, %r25, %r17;
	setp.lt.s32	%p2, %r6, %r23;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB6_10;
	bra.uni 	BB6_1;

BB6_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r22, 1;
	@%p4 bra 	BB6_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r21, 1;
	@%p5 bra 	BB6_9;

	cvta.to.global.u64 	%rd1, %rd6;
	mul.lo.s32 	%r7, %r6, %r22;
	mul.lo.s32 	%r8, %r21, %r25;
	mul.lo.s32 	%r30, %r22, %r6;
	mad.lo.s32 	%r31, %r20, %r19, %r30;
	mad.lo.s32 	%r9, %r18, %r31, %r8;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r26, 0;
	mov.u32 	%r42, %r26;

BB6_4:
	mad.lo.s32 	%r33, %r18, %r42, %r9;
	mul.wide.s32 	%rd7, %r33, 4;
	add.s64 	%rd11, %rd1, %rd7;
	add.s32 	%r11, %r42, %r7;
	mov.u32 	%r39, %r8;
	mov.u32 	%r41, %r26;

BB6_5:
	mov.u32 	%r13, %r41;
	mov.u32 	%r12, %r39;
	setp.lt.s32	%p6, %r12, %r18;
	setp.lt.s32	%p7, %r11, %r19;
	and.pred  	%p8, %p7, %p6;
	@!%p8 bra 	BB6_7;
	bra.uni 	BB6_6;

BB6_6:
	ld.global.nc.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB6_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r14, %r12, 1;
	add.s32 	%r15, %r13, 1;
	setp.lt.s32	%p9, %r15, %r21;
	mov.u32 	%r39, %r14;
	mov.u32 	%r41, %r15;
	@%p9 bra 	BB6_5;

	add.s32 	%r42, %r42, 1;
	setp.lt.s32	%p10, %r42, %r22;
	@%p10 bra 	BB6_4;

BB6_9:
	mad.lo.s32 	%r38, %r6, %r17, %r25;
	cvta.to.global.u64 	%rd8, %rd5;
	mul.wide.s32 	%rd9, %r38, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB6_10:
	ret;
}


`
	resize_ptx_53 = `
.version 4.3
.target sm_53
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	resize
.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<21>;
	.reg .b32 	%r<43>;
	.reg .b64 	%rd<12>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r17, [resize_param_1];
	ld.param.u32 	%r23, [resize_param_2];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r18, [resize_param_5];
	ld.param.u32 	%r19, [resize_param_6];
	ld.param.u32 	%r20, [resize_param_8];
	ld.param.u32 	%r21, [resize_param_9];
	ld.param.u32 	%r22, [resize_param_10];
	mov.u32 	%r24, %ntid.x;
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r25, %r24, %r1, %r2;
	mov.u32 	%r3, %ntid.y;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %tid.y;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	setp.lt.s32	%p1, %r25, %r17;
	setp.lt.s32	%p2, %r6, %r23;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB6_10;
	bra.uni 	BB6_1;

BB6_1:
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p4, %r22, 1;
	@%p4 bra 	BB6_9;

	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	setp.lt.s32	%p5, %r21, 1;
	@%p5 bra 	BB6_9;

	cvta.to.global.u64 	%rd1, %rd6;
	mul.lo.s32 	%r7, %r6, %r22;
	mul.lo.s32 	%r8, %r21, %r25;
	mul.lo.s32 	%r30, %r22, %r6;
	mad.lo.s32 	%r31, %r20, %r19, %r30;
	mad.lo.s32 	%r9, %r18, %r31, %r8;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r26, 0;
	mov.u32 	%r42, %r26;

BB6_4:
	mad.lo.s32 	%r33, %r18, %r42, %r9;
	mul.wide.s32 	%rd7, %r33, 4;
	add.s64 	%rd11, %rd1, %rd7;
	add.s32 	%r11, %r42, %r7;
	mov.u32 	%r39, %r8;
	mov.u32 	%r41, %r26;

BB6_5:
	mov.u32 	%r13, %r41;
	mov.u32 	%r12, %r39;
	setp.lt.s32	%p6, %r12, %r18;
	setp.lt.s32	%p7, %r11, %r19;
	and.pred  	%p8, %p7, %p6;
	@!%p8 bra 	BB6_7;
	bra.uni 	BB6_6;

BB6_6:
	ld.global.nc.f32 	%f17, [%rd11];
	add.f32 	%f20, %f20, %f17;
	add.f32 	%f19, %f19, 0f3F800000;

BB6_7:
	add.s64 	%rd11, %rd11, 4;
	add.s32 	%r14, %r12, 1;
	add.s32 	%r15, %r13, 1;
	setp.lt.s32	%p9, %r15, %r21;
	mov.u32 	%r39, %r14;
	mov.u32 	%r41, %r15;
	@%p9 bra 	BB6_5;

	add.s32 	%r42, %r42, 1;
	setp.lt.s32	%p10, %r42, %r22;
	@%p10 bra 	BB6_4;

BB6_9:
	mad.lo.s32 	%r38, %r6, %r17, %r25;
	cvta.to.global.u64 	%rd8, %rd5;
	mul.wide.s32 	%rd9, %r38, 4;
	add.s64 	%rd10, %rd8, %rd9;
	div.rn.f32 	%f18, %f20, %f19;
	st.global.f32 	[%rd10], %f18;

BB6_10:
	ret;
}


`
)
