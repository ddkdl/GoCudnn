package gocudnn

/*
#include <cudnn.h>
*/
import (
	"C"
)

//OpTensorFlag is used pass OpTensorOp Flags Semi Safely using methods
type OpTensorFlag struct {
}

//OpTensorOp is used for flags for the Optensor functions
type OpTensorOp C.cudnnOpTensorOp_t

//Add returns 	 OpTensorOp(C.CUDNN_OP_TENSOR_ADD)
func (o OpTensorFlag) Add() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_ADD)
}

//Mul returns  OpTensorOp(C.CUDNN_OP_TENSOR_MUL)
func (o OpTensorFlag) Mul() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_MUL)
}

//Min returns OpTensorOp(C.CUDNN_OP_TENSOR_MIN)
func (o OpTensorFlag) Min() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_MIN)
}

//Max returns OpTensorOp(C.CUDNN_OP_TENSOR_MAX)
func (o OpTensorFlag) Max() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_MAX)
}

//Sqrt returns OpTensorOp(C.CUDNN_OP_TENSOR_SQRT)
func (o OpTensorFlag) Sqrt() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_SQRT)
}

//Not returns OpTensorOp(C.CUDNN_OP_TENSOR_NOT)
func (o OpTensorFlag) Not() OpTensorOp {
	return OpTensorOp(C.CUDNN_OP_TENSOR_NOT)
}
func (o OpTensorOp) c() C.cudnnOpTensorOp_t { return C.cudnnOpTensorOp_t(o) }

//OPTensorD holds OP Tensor information
type OPTensorD struct {
	descriptor C.cudnnOpTensorDescriptor_t
}

//NewOpTensorDescriptor creates and sets an OpTensor
func NewOpTensorDescriptor(opTensOp OpTensorOp, opTensorCompType DataType, opTensorNanOpt PropagationNAN) (*OPTensorD, error) {
	var descriptor C.cudnnOpTensorDescriptor_t
	err := Status(C.cudnnCreateOpTensorDescriptor(&descriptor)).error("NewOpTensorDescriptor-Create")
	if err != nil {
		return nil, err
	}
	err = Status(C.cudnnSetOpTensorDescriptor(descriptor, opTensOp.c(), C.cudnnDataType_t(opTensorCompType), C.cudnnNanPropagation_t(opTensorNanOpt))).error("NewOpTensorDescriptor-set")
	if err != nil {
		return nil, err
	}
	return &OPTensorD{descriptor: descriptor}, nil
}

//GetDescriptor returns the descriptor information with error
func (t *OPTensorD) GetDescriptor() (OpTensorOp, DataType, PropagationNAN, error) {
	var tensop C.cudnnOpTensorOp_t
	var datatype C.cudnnDataType_t
	var nanprop C.cudnnNanPropagation_t

	x := C.cudnnGetOpTensorDescriptor(t.descriptor, &tensop, &datatype, &nanprop)
	return OpTensorOp(tensop), DataType(datatype), PropagationNAN(nanprop), Status(x).error("GetOpTensorDescriptor")
}
