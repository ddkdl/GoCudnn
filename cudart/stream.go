package cudart

/*

#include <cuda_runtime_api.h>
const cudaStream_t gocunullstream = NULL;
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

//Stream holds a C.cudaStream_t
type Stream struct {
	stream C.cudaStream_t
}

//Ptr returns an unsafe pointer to the hidden stream.
//This allows stream to be used with other cuda libraries in other go packages
//so if a C function calls for a Pointer then you can type case the unsafe pointer
//into a (C.cudaStream_t)(unsafe.Pointer)
func (s *Stream) Ptr() unsafe.Pointer {
	if s == nil {
		return unsafe.Pointer(C.cudaStream_t(C.NULL))
	}
	return unsafe.Pointer(s.stream)
}
func (s *Stream) c() C.cudaStream_t {
	if s == nil {
		return C.cudaStream_t(C.NULL)
	}
	return s.stream
}

//ExternalWrapper is used for other packages that might return a C.cudaStream_t
func ExternalWrapper(x unsafe.Pointer) *Stream {
	return &Stream{
		stream: C.cudaStream_t(x),
	}
}

//CreateBlockingStream creats an asyncronus stream stream for the user
func CreateBlockingStream() (*Stream, error) {
	s := new(Stream)
	err := s.create(false, false, 0)

	runtime.SetFinalizer(s, destroystream)

	return s, err
}

//CreateNonBlockingStream creates a blocking stream
func CreateNonBlockingStream() (*Stream, error) {
	s := new(Stream)
	err := s.create(true, false, 0)

	runtime.SetFinalizer(s, destroystream)

	return s, err

}

//CreateNonBlockingPriorityStream creates a non blocking Priority Stream
func CreateNonBlockingPriorityStream(priority int32) (*Stream, error) {
	s := new(Stream)
	err := s.create(true, true, priority)

	runtime.SetFinalizer(s, destroystream)

	return s, err

}

//CreateBlockingPriorityStream creates a blocking stream
func CreateBlockingPriorityStream(priority int32) (*Stream, error) {
	s := new(Stream)
	err := s.create(true, true, priority)

	runtime.SetFinalizer(s, destroystream)

	return s, err

}

//Sync Syncronizes the stream
func (s *Stream) Sync() error {
	if s == nil {
		return newErrorRuntime("Sync", C.cudaStreamSynchronize(nil))
	}
	return newErrorRuntime("Sync", C.cudaStreamSynchronize(s.stream))
}

//Create creates a Stream
func (s *Stream) create(blocking, priority bool, rank int32) error {
	if blocking == true && priority == false {
		x := C.cudaStreamCreateWithFlags(&s.stream, C.cudaStreamDefault)

		return newErrorRuntime("cudaStreamCreate", x)
	}
	if blocking == false && priority == false {
		x := C.cudaStreamCreateWithFlags(&s.stream, C.cudaStreamNonBlocking)

		return newErrorRuntime("cudaStreamCreate", x)
	}
	if blocking == true && priority == true {
		x := C.cudaStreamCreateWithPriority(&s.stream, C.cudaStreamDefault, C.int(rank))

		return newErrorRuntime("cudaStreamCreate", x)
	}
	if blocking == false && priority == true {
		x := C.cudaStreamCreateWithPriority(&s.stream, C.cudaStreamNonBlocking, C.int(rank))
		return newErrorRuntime("cudaStreamCreate", x)
	}
	return errors.New("CreateStream: Unreachable: How did this Happen")
}
func destroystream(s *Stream) error {

	err := newErrorRuntime("Destroy", C.cudaStreamDestroy(s.stream))
	s = nil
	return err
}
