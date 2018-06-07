package gocudnn

/*
#include <cudnn.h>
*/
import "C"
import (
	"unsafe"
)

//RNNmode is used for flags use RNNModeFlag to pass them through methods
type RNNmode C.cudnnRNNMode_t

//RNNModeFlag is used to pass RNNMode flags semi safely through methods.
type RNNModeFlag struct {
}

func (r RNNmode) c() C.cudnnRNNMode_t { return C.cudnnRNNMode_t(r) }

//Relu return RNNMode(C.CUDNN_RNN_RELU)
func (r RNNModeFlag) Relu() RNNmode {
	return RNNmode(C.CUDNN_RNN_RELU)
}

//Tanh returns rnnTanh
func (r RNNModeFlag) Tanh() RNNmode {
	return RNNmode(C.CUDNN_RNN_TANH)
}

//Lstm returns rnnLstm
func (r RNNModeFlag) Lstm() RNNmode {
	return RNNmode(C.CUDNN_LSTM)
}

//Gru returns rnnGru
func (r RNNModeFlag) Gru() RNNmode {
	return RNNmode(C.CUDNN_GRU)
}

//DirectionModeFlag is used to pass DirectionModes through its methods.
type DirectionModeFlag struct {
}

//DirectionMode use DirectionModeFlag to pass them safe-ish.
type DirectionMode C.cudnnDirectionMode_t

func (r DirectionMode) c() C.cudnnDirectionMode_t { return C.cudnnDirectionMode_t(r) }

//Uni returns uniDirectional flag
func (r DirectionModeFlag) Uni() DirectionMode {
	return DirectionMode(C.CUDNN_UNIDIRECTIONAL)
}

//Bi returns biDirectional flag
func (r DirectionModeFlag) Bi() DirectionMode {
	return DirectionMode(C.CUDNN_BIDIRECTIONAL)
}

/*
 *   RNN INPUT MODE FLAGS
 */

//RNNInputModeFlag used to pass RNNInputMode Flags semi-safely through its methods.
type RNNInputModeFlag struct {
}

//RNNInputMode is used for flags
type RNNInputMode C.cudnnRNNInputMode_t

//Linear returns C.CUDNN_LINEAR_INPUT
func (r RNNInputModeFlag) Linear() RNNInputMode {
	return RNNInputMode(C.CUDNN_LINEAR_INPUT)
}

//Skip returns C.CUDNN_SKIP_INPUT
func (r RNNInputModeFlag) Skip() RNNInputMode {
	return RNNInputMode(C.CUDNN_SKIP_INPUT)
}
func (r RNNInputMode) c() C.cudnnRNNInputMode_t { return C.cudnnRNNInputMode_t(r) }

/*
 *   RNN ALGO FLAGS
 */

//RNNAlgoFlag used to pass RNNAlgo flags semi-safely.
type RNNAlgoFlag struct {
}

//RNNAlgo is used for flags
type RNNAlgo C.cudnnRNNAlgo_t

//Standard returns RNNAlgo( C.CUDNN_RNN_ALGO_STANDARD) flag
func (r RNNAlgoFlag) Standard() RNNAlgo {
	return RNNAlgo(C.CUDNN_RNN_ALGO_STANDARD)
}

//PersistStatic returns RNNAlgo( C.CUDNN_RNN_ALGO_PERSIST_STATIC) flag
func (r RNNAlgoFlag) PersistStatic() RNNAlgo {
	return RNNAlgo(C.CUDNN_RNN_ALGO_PERSIST_STATIC)
}

//PersistDynamic returns RNNAlgo( C.CUDNN_RNN_ALGO_PERSIST_DYNAMIC) flag
func (r RNNAlgoFlag) PersistDynamic() RNNAlgo {
	return RNNAlgo(C.CUDNN_RNN_ALGO_PERSIST_DYNAMIC)
}

//Count returns RNNAlgo( C.CUDNN_RNN_ALGO_COUNT) flag
func (r RNNAlgoFlag) Count() RNNAlgo {
	return RNNAlgo(C.CUDNN_RNN_ALGO_COUNT)
}

func (r RNNAlgo) c() C.cudnnRNNAlgo_t { return C.cudnnRNNAlgo_t(r) }

//AlgorithmPerformance go typed C.cudnnAlgorithmPerformance_t
type AlgorithmPerformance C.cudnnAlgorithmPerformance_t

//RNND  holdes Rnn descriptor
type RNND struct {
	descriptor C.cudnnRNNDescriptor_t
}

func rrndArrayToCarray(input []RNND) []C.cudnnRNNDescriptor_t {
	array := make([]C.cudnnRNNDescriptor_t, len(input))
	for i := 0; i < len(input); i++ {
		array[i] = input[i].descriptor
	}
	return array
}

//CreateRNNDescriptor creates an RNND descriptor
func CreateRNNDescriptor() (*RNND, error) {
	var desc C.cudnnRNNDescriptor_t
	err := Status(C.cudnnCreateRNNDescriptor(&desc)).error("CreateRNNDescriptor")
	if err != nil {
		return nil, err
	}
	return &RNND{
		descriptor: desc,
	}, nil
}

//DestroyDescriptor destroys the descriptor
func (r *RNND) DestroyDescriptor() error {
	return Status(C.cudnnDestroyRNNDescriptor(r.descriptor)).error("DestroyDescriptor-rnn")
}

//SetRNNDescriptor sets the rnndesctiptor
func (r *RNND) SetRNNDescriptor(
	handle *Handle,
	hiddenSize int32,
	numLayers int32,
	doD *DropOutD,
	inputmode RNNInputMode,
	direction DirectionMode,
	rnnmode RNNmode,
	rnnalg RNNAlgo,
	data DataType,

) error {
	return Status(C.cudnnSetRNNDescriptor(
		handle.x,
		r.descriptor,
		C.int(hiddenSize),
		C.int(numLayers),
		doD.descriptor,
		inputmode.c(),
		direction.c(),
		rnnmode.c(),
		rnnalg.c(),
		data.c(),
	)).error("SetRNNDescriptor")
}

//SetRNNProjectionLayers sets the rnnprojection layers
func (r *RNND) SetRNNProjectionLayers(
	handle *Handle,
	recProjsize int32,
	outProjSize int32,
) error {
	return Status(C.cudnnSetRNNProjectionLayers(
		handle.x,
		r.descriptor,
		C.int(recProjsize),
		C.int(outProjSize),
	)).error("SetRNNProjectionLayers")
}

//GetRNNProjectionLayers sets the rnnprojection layers
func (r *RNND) GetRNNProjectionLayers(
	handle *Handle,

) (int32, int32, error) {
	var rec, out C.int

	err := Status(C.cudnnGetRNNProjectionLayers(
		handle.x,
		r.descriptor,
		&rec,
		&out,
	)).error("SetRNNProjectionLayers")
	return int32(rec), int32(out), err
}

//SetRNNAlgorithmDescriptor sets the RNNalgorithm
func (r *RNND) SetRNNAlgorithmDescriptor(
	handle *Handle,
	algo *AlgorithmD,
) error {
	return Status(C.cudnnSetRNNAlgorithmDescriptor(handle.x, r.descriptor, algo.descriptor)).error("SetRNNAlgorithmDescriptor")
}

//GetRNNDescriptor gets algo desctiptor values returns a ton of stuff
func (r *RNND) GetRNNDescriptor(
	handle *Handle,
) (int32, int32, *DropOutD, RNNInputMode, DirectionMode, RNNmode, RNNAlgo, DataType, error) {
	var hiddensize C.int
	var numLayers C.int
	var dropoutdescriptor C.cudnnDropoutDescriptor_t
	var inputMode C.cudnnRNNInputMode_t
	var direction C.cudnnDirectionMode_t
	var mode C.cudnnRNNMode_t
	var algo C.cudnnRNNAlgo_t
	var dataType C.cudnnDataType_t
	err := Status(C.cudnnGetRNNDescriptor(
		handle.x,
		r.descriptor,
		&hiddensize,
		&numLayers,
		&dropoutdescriptor,
		&inputMode,
		&direction,
		&mode,
		&algo,
		&dataType,
	)).error("GetRNNDescriptor")
	return int32(hiddensize), int32(numLayers), &DropOutD{descriptor: dropoutdescriptor},
		RNNInputMode(inputMode), DirectionMode(direction), RNNmode(mode), RNNAlgo(algo), DataType(dataType), err
}

//SetRNNMatrixMathType Sets the math type for the descriptor
func (r *RNND) SetRNNMatrixMathType(math MathType) error {
	return Status(C.cudnnSetRNNMatrixMathType(r.descriptor, math.c())).error("SetRNNMatrixMathType")
}

//GetRNNMatrixMathType Gets the math type for the descriptor
func (r *RNND) GetRNNMatrixMathType() (MathType, error) {
	var math C.cudnnMathType_t
	err := Status(C.cudnnGetRNNMatrixMathType(r.descriptor, &math)).error("SetRNNMatrixMathType")
	return MathType(math), err
}

/* dataType in the RNN descriptor is used to determine math precision */
/* dataType in weight descriptors and input descriptors is used to describe storage */

//GetRNNWorkspaceSize gets the RNN workspace size (WOW!)
func (r *RNND) GetRNNWorkspaceSize(
	handle *Handle,
	seqLength int32,
	xD []*TensorD,
) (SizeT, error) {
	tocxD := tensorDArrayToC(xD)
	var sizeinbytes C.size_t
	err := Status(C.cudnnGetRNNWorkspaceSize(
		handle.x,
		r.descriptor,
		C.int(seqLength),
		&tocxD[0],
		&sizeinbytes,
	)).error("GetRNNWorkspaceSize")
	return SizeT(sizeinbytes), err
}

//GetRNNTrainingReserveSize gets the training reserve size
func (r *RNND) GetRNNTrainingReserveSize(
	handle *Handle,
	seqLength int32,
	xD []*TensorD,
) (SizeT, error) {
	tocxD := tensorDArrayToC(xD)
	var sizeinbytes C.size_t
	err := Status(C.cudnnGetRNNTrainingReserveSize(
		handle.x,
		r.descriptor,
		C.int(seqLength),
		&tocxD[0],
		&sizeinbytes,
	)).error("GetRNNTrainingReserveSize")
	return SizeT(sizeinbytes), err
}

//GetRNNParamsSize gets the training reserve size
func (r *RNND) GetRNNParamsSize(
	handle *Handle,
	xD *TensorD,
	data *DataType,
) (SizeT, error) {
	var sizeinbytes C.size_t
	err := Status(C.cudnnGetRNNParamsSize(
		handle.x,
		r.descriptor,
		xD.descriptor,
		&sizeinbytes,
		data.c(),
	)).error("GetRNNParamsSize")
	return SizeT(sizeinbytes), err
}

//GetRNNLinLayerMatrixParams gets the parameters of the layer matrix
func (r *RNND) GetRNNLinLayerMatrixParams(
	handle *Handle,
	pseudoLayer int32,
	/*
	   The pseudo-layer to query.
	   In uni-directional RNN-s, a pseudo-layer is the same as a "physical" layer
	   (pseudoLayer=0 is the RNN input layer, pseudoLayer=1 is the first hidden layer).

	   In bi-directional RNN-s there are twice as many pseudo-layers in comparison to "physical" layers
	   (pseudoLayer=0 and pseudoLayer=1 are both input layers;
	   pseudoLayer=0 refers to the forward part and pseudoLayer=1 refers to the backward part
	    of the "physical" input layer; pseudoLayer=2 is the forward part of the first hidden layer, and so on).

	*/
	xD *TensorD,
	wD *FilterD,
	w Memer,
	linlayerID int32,
	/*
	   Input. The linear layer to obtain information about:

	   If mode in rnnDesc was set to CUDNN_RNN_RELU or CUDNN_RNN_TANH a value of 0 references the matrix multiplication applied to the input from the previous layer,
	   a value of 1 references the matrix multiplication applied to the recurrent input.

	   If mode in rnnDesc was set to CUDNN_LSTM values of 0-3 reference matrix multiplications applied to the input from the previous layer, value of 4-7 reference matrix multiplications applied to the recurrent input.
	       Values 0 and 4 reference the input gate.
	       Values 1 and 5 reference the forget gate.
	       Values 2 and 6 reference the new memory gate.
	       Values 3 and 7 reference the output gate.
	       Value 8 references the "recurrent" projection matrix when enabled by the cudnnSetRNNProjectionLayers() function.

	   If mode in rnnDesc was set to CUDNN_GRU values of 0-2 reference matrix multiplications applied to the input from the previous layer, value of 3-5 reference matrix multiplications applied to the recurrent input.
	       Values 0 and 3 reference the reset gate.
	       Values 1 and 4 reference the update gate.
	       Values 2 and 5 reference the new memory gate.

	*/

) (FilterD, unsafe.Pointer, error) {
	var linLayerMatDesc FilterD
	var linLayerMat unsafe.Pointer
	err := Status(C.cudnnGetRNNLinLayerMatrixParams(
		handle.x,
		r.descriptor,
		C.int(pseudoLayer),
		xD.descriptor,
		wD.descriptor,
		w.Ptr(),
		C.int(linlayerID),
		linLayerMatDesc.descriptor,
		&linLayerMat,
	)).error("GetRNNLinLayerMatrixParams")
	return linLayerMatDesc, linLayerMat, err
}

//GetRNNLinLayerBiasParams gets the parameters of the layer bias
func (r *RNND) GetRNNLinLayerBiasParams(
	handle *Handle,
	pseudoLayer int32,
	/*
	   The pseudo-layer to query.
	   In uni-directional RNN-s, a pseudo-layer is the same as a "physical" layer
	   (pseudoLayer=0 is the RNN input layer, pseudoLayer=1 is the first hidden layer).

	   In bi-directional RNN-s there are twice as many pseudo-layers in comparison to "physical" layers
	   (pseudoLayer=0 and pseudoLayer=1 are both input layers;
	   pseudoLayer=0 refers to the forward part and pseudoLayer=1 refers to the backward part
	    of the "physical" input layer; pseudoLayer=2 is the forward part of the first hidden layer, and so on).

	*/
	xD *TensorD,
	wD *FilterD,
	w Memer,
	linlayerID int32,
	/*
	   Input. The linear layer to obtain information about:

	   If mode in rnnDesc was set to CUDNN_RNN_RELU or CUDNN_RNN_TANH a value of 0 references the matrix multiplication applied to the input from the previous layer,
	   a value of 1 references the matrix multiplication applied to the recurrent input.

	   If mode in rnnDesc was set to CUDNN_LSTM values of 0-3 reference matrix multiplications applied to the input from the previous layer, value of 4-7 reference matrix multiplications applied to the recurrent input.
	       Values 0 and 4 reference the input gate.
	       Values 1 and 5 reference the forget gate.
	       Values 2 and 6 reference the new memory gate.
	       Values 3 and 7 reference the output gate.
	       Value 8 references the "recurrent" projection matrix when enabled by the cudnnSetRNNProjectionLayers() function.

	   If mode in rnnDesc was set to CUDNN_GRU values of 0-2 reference matrix multiplications applied to the input from the previous layer, value of 3-5 reference matrix multiplications applied to the recurrent input.
	       Values 0 and 3 reference the reset gate.
	       Values 1 and 4 reference the update gate.
	       Values 2 and 5 reference the new memory gate.

	*/

) (FilterD, unsafe.Pointer, error) {
	var linLayerBiasDesc FilterD
	var linLayerBias unsafe.Pointer
	err := Status(C.cudnnGetRNNLinLayerBiasParams(
		handle.x,
		r.descriptor,
		C.int(pseudoLayer),
		xD.descriptor,
		wD.descriptor,
		w.Ptr(),
		C.int(linlayerID),
		linLayerBiasDesc.descriptor,
		&linLayerBias,
	)).error("GetRNNLinLayerBiasParams")
	return linLayerBiasDesc, linLayerBias, err
}
