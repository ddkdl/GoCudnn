package gocudnn

/*
#include <cudnn.h>
*/
import "C"

//ActivationMode is used for activation discriptor flags
type ActivationMode C.cudnnActivationMode_t

//Flags for activation mode
const (
	ActivationSigmoid     ActivationMode = C.CUDNN_ACTIVATION_SIGMOID
	ActivationRelu        ActivationMode = C.CUDNN_ACTIVATION_RELU
	ActivationTanh        ActivationMode = C.CUDNN_ACTIVATION_TANH
	ActivationClippedRelu ActivationMode = C.CUDNN_ACTIVATION_CLIPPED_RELU
	ActivationElu         ActivationMode = C.CUDNN_ACTIVATION_ELU
	ActivationIdentity    ActivationMode = C.CUDNN_ACTIVATION_IDENTITY
)

func (a *ActivationMode) c() C.cudnnActivationMode_t { return C.cudnnActivationMode_t(*a) }

//ActivationD is the activation descriptor
type ActivationD struct {
	descriptor C.cudnnActivationDescriptor_t
	coef       C.double //This is used for ceiling for clipped relu, and alpha for elu
}

//NewActivationDescriptor creates and sets and returns an activation descriptor in ActivationD and the error
func NewActivationDescriptor(mode ActivationMode, nan PropagationNAN, coef CDouble) (*ActivationD, error) {

	var descriptor C.cudnnActivationDescriptor_t
	err := Status(C.cudnnCreateActivationDescriptor(&descriptor)).error("NewActivationDescriptor-create")
	if err != nil {
		return nil, err
	}

	err = Status(C.cudnnSetActivationDescriptor(descriptor, mode.c(), nan.c(), coef.c())).error("NewActivationDescriptor-set")
	if err != nil {
		return nil, err
	}
	return &ActivationD{
		descriptor: descriptor,
		coef:       C.double(coef),
	}, err
}

//GetDescriptor gets the descriptor info for ActivationD
func (a *ActivationD) GetDescriptor() (ActivationMode, PropagationNAN, CDouble, error) {
	var coef C.double
	var mode C.cudnnActivationMode_t
	var nan C.cudnnNanPropagation_t
	err := Status(C.cudnnGetActivationDescriptor(a.descriptor, &mode, &nan, &coef)).error("GetDescriptor")
	return ActivationMode(mode), PropagationNAN(nan), CDouble(coef), err
}

//DestroyDescriptor destroys the activation descriptor
func (a *ActivationD) DestroyDescriptor() error {
	return Status(C.cudnnDestroyActivationDescriptor(a.descriptor)).error("DestroyDescriptor")
}