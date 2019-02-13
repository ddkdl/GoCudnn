package gocudnn

/*
#include <string.h>
#include <cuda.h>
#include <cuda_runtime_api.h>

const void ** voiddptrnull = NULL;
const size_t ptrSize = sizeof(void *);
const size_t maxArgSize = 8;
const CUjit_option * nullJitOptions = NULL;


*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

//Kernel is a function stored on the gpu.
type Kernel struct {
	name string
	f    C.CUfunction
	m    *Module
	args []unsafe.Pointer
	mux  sync.Mutex
}

//Module are used to hold kernel functions on the device that is in use
type Module struct {
	m      C.CUmodule
	loaded bool
}

func (m *Module) c() C.CUmodule { return m.m }

//extern __host__ cudaError_t CUDARTAPI cudaLaunchKernel(const void *func, dim3 gridDim, dim3 blockDim, void **args, size_t sharedMem, cudaStream_t stream);

//NewModule creates a module in the current context.
func (cu Cuda) NewModule(filename string) (module *Module, err error) {
	var mod C.CUmodule
	fname := C.CString(filename)
	x := C.cuModuleLoad(&mod, fname)
	module = &Module{
		m:      mod,
		loaded: true,
	}
	return module, newErrorDriver("NewModule", x)
}

//NewModuleEx takes a string of the ptx data
func (cu Cuda) NewModuleEx(Ptx string) (*Module, error) {
	var mod C.CUmodule
	cptx := unsafe.Pointer(C.CString(Ptx))
	x := C.cuModuleLoadDataEx(&mod, cptx, 0, C.nullJitOptions, C.voiddptrnull)
	return &Module{
		m:      mod,
		loaded: true,
	}, newErrorDriver("NewModule", x)
}

//UnLoad Loads a Module
func (m *Module) UnLoad() error {
	x := C.cuModuleUnload(m.m)
	m.loaded = false
	return newErrorDriver("UnLoadModule", x)

}

//Load Loads a Module with a ptx file name
func (m *Module) Load(filename string) error {
	if m.loaded == true {
		return errors.New("(*Module)Load: Goside: Already Loaded")
	}
	fname := C.CString(filename)
	x := C.cuModuleLoad(&m.m, fname)
	return newErrorDriver("Load", x)
}

//LoadEx loads the ptx string straightup
func (m *Module) LoadEx(ptx string) error {
	if m.loaded == true {
		return errors.New("(*Module)LoadEx: Goside: Already Loaded")
	}
	cptx := unsafe.Pointer(C.CString(ptx))

	x := C.cuModuleLoadDataEx(&m.m, cptx, 0, C.nullJitOptions, C.voiddptrnull)
	return newErrorDriver("LoadEx", x)
}

//MakeKernel makes a kernel.  (basically a function in cuda) If module is unloaded, then the kernel returned won't work
func (cu Cuda) MakeKernel(kname string, m *Module) (k *Kernel, err error) {
	var kern C.CUfunction
	if m.loaded == false {
		return nil, errors.New("MakeKernel: Module Not Loaded")
	}

	name := C.CString(kname)
	err = newErrorDriver("MakeKernel", C.cuModuleGetFunction(&kern, m.m, name))
	if err != nil {
		return nil, err
	}
	k = &Kernel{
		name: kname,
		m:    m,
		f:    kern,
	}
	if setfinalizer {
		runtime.SetFinalizer(k, destroycudakernel)
	}
	return k, nil
}

//Launch launches the kernel that is stored in the hidden module in this struct.
//gx,gy,gz are the grid values.
//bx,by,bz are the block values.
//shared is the shared memory size
//stream is the cuda stream this module will use.
//args are the arguments that are used for the kernel. kernels are written in a .cu file.  You will need to make a .ptx file in order to use it.
//you can check out the kernels package to get an idea of how to make the stuff.

const pointerSize = 8

func offSet(ptr unsafe.Pointer, i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + pointerSize*uintptr(i))
}

//Launch will launch a kernal that is in it
func (k *Kernel) Launch(gx, gy, gz, bx, by, bz, shared uint32, stream *Stream, args ...interface{}) error {

	err := k.ifacetounsafecomplete(args)
	if err != nil {
		return err
	}
	var shold C.cudaStream_t

	if stream == nil {
		shold = nil
	} else {
		shold = stream.stream
	}
	if setkeepalive {
		k.keepsalive()

	}
	return newErrorDriver("cuLaunchKernel",
		C.cuLaunchKernel(k.f,
			C.uint(gx),
			C.uint(gy),
			C.uint(gz),
			C.uint(bx),
			C.uint(by),
			C.uint(bz),
			C.uint(shared),
			shold,
			&k.args[0],
			C.voiddptrnull,
		))
}

//LaunchOld will launch a kernal that is in it was running into a seg fault
func (k *Kernel) LaunchOld(gx, gy, gz, bx, by, bz, shared uint32, stream *Stream, args ...interface{}) error {

	kernelParams, err := k.ifacetounsafe(args)

	if err != nil {
		return err
	}

	//Start straight up took this from gorgonia/cu
	argv := C.malloc(C.size_t(len(kernelParams) * pointerSize))
	argp := C.malloc(C.size_t(len(kernelParams) * pointerSize))
	defer C.free(argv)
	defer C.free(argp)
	for i := range kernelParams {
		*((*unsafe.Pointer)(offSet(argp, i))) = offSet(argv, i) // argp[i] = &argv[i]
		holder := *((*uint64)(kernelParams[i]))
		*((*uint64)(offSet(argv, i))) = holder // argv[i] = *kernelParams[i]
	}

	if err != nil {

		return err
	}
	var shold C.cudaStream_t

	if stream == nil {
		shold = nil
	} else {
		shold = stream.stream
	}
	if setkeepalive {
		k.keepsalive()

	}
	return newErrorDriver("cuLaunchKernel",
		C.cuLaunchKernel(k.f,
			C.uint(gx),
			C.uint(gy),
			C.uint(gz),
			C.uint(bx),
			C.uint(by),
			C.uint(bz),
			C.uint(shared),
			shold,
			(*unsafe.Pointer)(argp),
			(*unsafe.Pointer)(nil),
		))
}

//KernelArguments need to be loaded before using LaunchV2.  Handy if most of the parameters are not changing. Also, if you want to parallelize it then you can use this.
type KernelArguments struct {
	gx, gy, gz CUInt
	bx, by, bz CUInt
	shared     CUInt
	stream     *Stream
	args       []interface{}
}

//SetGrid sets the grid dims
func (k *KernelArguments) SetGrid(gx, gy, gz uint32) {
	k.gx, k.gy, k.gz = CUInt(gx), CUInt(gy), CUInt(gz)

}

//GetGrid returns the grid values
func (k *KernelArguments) GetGrid() (uint32, uint32, uint32) {
	return uint32(k.gx), uint32(k.gy), uint32(k.gz)
}

//SetBlock sets the block dims
func (k *KernelArguments) SetBlock(bx, by, bz uint32) {
	k.bx, k.by, k.bz = CUInt(bx), CUInt(by), CUInt(bz)

}

//GetBlock returns the block values
func (k *KernelArguments) GetBlock() (uint32, uint32, uint32) {
	return uint32(k.bx), uint32(k.by), uint32(k.bz)

}

//SetShared sets the shared dims
func (k *KernelArguments) SetShared(sharedsize uint32) {
	k.shared = CUInt(sharedsize)
}

//GetShared returns the shared memory size value
func (k *KernelArguments) GetShared() uint32 {
	return uint32(k.shared)
}

//SetArguments sets the arguments
func (k *KernelArguments) SetArguments(args ...interface{}) {
	k.args = args
}

//GetArguments returns the empty interface array of arguments
func (k *KernelArguments) GetArguments() []interface{} {
	return k.args
}
func (k *Kernel) ifacetounsafefirst(args []interface{}) error {

	k.args = make([]unsafe.Pointer, len(args))
	for i := range args {
		k.args[i] = unsafe.Pointer(C.malloc(C.maxArgSize))
		switch x := args[i].(type) {
		case nil:
			C.memcpy(k.args[i], unsafe.Pointer(&x), C.ptrSize) //This might need to be (C.voiddptrnull)

		case *Malloced:
			if x == nil {
				C.memcpy(k.args[i], unsafe.Pointer(&x), C.ptrSize)
			}
			C.memcpy(k.args[i], unsafe.Pointer(&x.ptr), C.ptrSize)
			if setkeepalive {
				keepsalivebuffer(x)
			}
		default:
			scalar := CScalarConversion(x)
			if scalar == nil {
				return fmt.Errorf("Kernel Launch - type %T not supported .. %+v", x, x)
			}
			C.memcpy(k.args[i], scalar.CPtr(), scalar.SizeT().c())

		}

	}

	return nil
}
func (k *Kernel) ifacetounsafecomplete(args []interface{}) error {
	if k.args == nil {
		return k.ifacetounsafefirst(args)
	}
	if len(k.args) == 0 {
		return k.ifacetounsafefirst(args)
	}
	if len(k.args) != len(args) {
		k.destroyargs()
		return k.ifacetounsafefirst(args)
	}
	for i := range args {
		switch x := args[i].(type) {
		case nil:
			C.memcpy(k.args[i], unsafe.Pointer(&x), C.ptrSize) //This might need to be (C.voiddptrnull)

		case *Malloced:
			if x == nil {
				C.memcpy(k.args[i], unsafe.Pointer(&x), C.ptrSize)
			}
			C.memcpy(k.args[i], unsafe.Pointer(&x.ptr), C.ptrSize)
			if setkeepalive {
				keepsalivebuffer(x)
			}
		default:
			scalar := CScalarConversion(x)
			if scalar == nil {
				return fmt.Errorf("Kernel Launch - type %T not supported .. %+v", x, x)
			}
			C.memcpy(k.args[i], scalar.CPtr(), scalar.SizeT().c())

		}

	}

	return nil
}
func (k *Kernel) destroyargs() {
	for i := range k.args {
		C.free(k.args[i])
	}
}
func (k *Kernel) keepsalive() {
	runtime.KeepAlive(k)

}
func destroycudakernel(k *Kernel) {
	k.destroyargs()
}

//Destroy destroys the argument array
func (k *Kernel) Destroy() {
	k.destroyargs()
}
func (k *Kernel) ifacetounsafe(args []interface{}) ([]unsafe.Pointer, error) {
	array := make([]unsafe.Pointer, len(args))
	for i := 0; i < len(args); i++ {

		switch x := args[i].(type) {
		case nil:
			array[i] = unsafe.Pointer(C.voiddptrnull)
		case *Malloced:
			if x == nil {
				array[i] = unsafe.Pointer(C.voiddptrnull)
			}
			if x.ptr == nil {
				return nil, errors.New("Memory Doesn't Have A Pointer")
			}
			array[i] = unsafe.Pointer(&x.ptr)
		case *GoPointer:
			return nil, errors.New("*GoPointer not supported")
		default:
			scalar := CScalarConversion(x)
			if scalar == nil {
				return nil, errors.New("Not a supported value")
			}
			array[i] = scalar.CPtr()
		}

	}
	return array, nil
}
