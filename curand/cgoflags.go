package curand

//#cgo LDFLAGS:-L/usr/local/cuda/lib64 -lcurand -lcuda -lcudart
//#cgo CFLAGS: -I/usr/local/cuda/include/

//#cgo LDFLAGS:-L/usr/local/cuda-10.1/lib64 -lcurand -lcuda -lcudart
//#cgo CFLAGS: -I/usr/local/cuda-10.1/include/
import "C"
