package npp

//#include <nppdefs.h>
import "C"
import "unsafe"

//Flags is a special struct that contains all the flag types npp uses.
//Even though these types are flags themselves.
//They can return flags with different values within the same type through methods
type Flags struct {
	InterpolationMode  InterpolationMode
	BayerGridPosition  BayerGridPosition
	MaskSize           MaskSize
	DifferentialKernel DifferentialKernel
	Axis               Axis
	CmpOp              CmpOp
	RoundMode          RoundMode
	BorderType         BorderType
	HintAlgorithm      HintAlgorithm
	AlphaOp            AlphaOp
	ZCType             ZCType
	HuffmanTableType   HuffmanTableType
	Norm               Norm
}

//GetFlags returns a Flags struct.  This struct has all the flag types within it.
func GetFlags() Flags {
	return Flags{}
}

//InterpolationMode is a wrapper for interpolation flags
type InterpolationMode C.NppiInterpolationMode

func (n InterpolationMode) c() C.NppiInterpolationMode {
	return (C.NppiInterpolationMode)(n)
}

func (n InterpolationMode) cint() C.int {
	return (C.int)(n)
}

//UNDEFINED returns InterpolationMode(C.NPPI_INTER_UNDEFINED)
func (n InterpolationMode) UNDEFINED() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_UNDEFINED)
}

//NN returns InterpolationMode(C.NPPI_INTER_NN)
/**<  Nearest neighbor filtering. */
func (n InterpolationMode) NN() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_NN)
}

//LINEAR returns InterpolationMode(C.NPPI_INTER_LINEAR)
/**<  Linear interpolation. */
func (n InterpolationMode) LINEAR() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_LINEAR)
}

//CUBIC returns InterpolationMode(C.NPPI_INTER_CUBIC)
/**<  Cubic interpolation. */
func (n InterpolationMode) CUBIC() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_CUBIC)
}

//BSPLINE returns InterpolationMode(C.NPPI_INTER_CUBIC2P_BSPLINE)
/**<  Two-parameter cubic filter (B=1, C=0) */
func (n InterpolationMode) BSPLINE() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_CUBIC2P_BSPLINE)
}

//CATMULLROM returns InterpolationMode(C.NPPI_INTER_CUBIC2P_CATMULLROM)
/**<  Two-parameter cubic filter (B=0, C=1/2) */
func (n InterpolationMode) CATMULLROM() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_CUBIC2P_CATMULLROM)
}

//B05C03 returns InterpolationMode(C.NPPI_INTER_CUBIC2P_B05C03)
/**<  Two-parameter cubic filter (B=1/2, C=3/10) */
func (n InterpolationMode) B05C03() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_CUBIC2P_B05C03)
}

//SUPER returns InterpolationMode(C.NPPI_INTER_SUPER)
/**<  Super sampling. */
func (n InterpolationMode) SUPER() InterpolationMode { return InterpolationMode(C.NPPI_INTER_SUPER) }

//LANCZOS returns InterpolationMode(C.NPPI_INTER_LANCZOS)
/**<  Lanczos filtering. */
func (n InterpolationMode) LANCZOS() InterpolationMode { return InterpolationMode(C.NPPI_INTER_LANCZOS) }

//LANCZ0S3ADVANCED returns InterpolationMode(C.NPPI_INTER_LANCZOS3_ADVANCED)
/**<  Generic Lanczos filtering with order 3. */
func (n InterpolationMode) LANCZ0S3ADVANCED() InterpolationMode {
	return InterpolationMode(C.NPPI_INTER_LANCZOS3_ADVANCED)
}

//SMOOTHEDGE returns InterpolationMode(C.NPPI_SMOOTH_EDGE)
/**<  Smooth edge filtering. */
func (n InterpolationMode) SMOOTHEDGE() InterpolationMode {
	return InterpolationMode(C.NPPI_SMOOTH_EDGE)
}

//BayerGridPosition is used as flags. Contains methods for different flags
type BayerGridPosition C.NppiBayerGridPosition

func (b BayerGridPosition) c() C.NppiBayerGridPosition { return C.NppiBayerGridPosition(b) }

//BGGR returns BayerGridPosition(C.NPPI_BAYER_BGGR)
func (b BayerGridPosition) BGGR() BayerGridPosition { return BayerGridPosition(C.NPPI_BAYER_BGGR) }

//RGGB returns BayerGridPosition(NPPI_BAYER_RGGB)
func (b BayerGridPosition) RGGB() BayerGridPosition { return BayerGridPosition(C.NPPI_BAYER_RGGB) }

//GRBG 	returns BayerGridPosition(C.NPPI_BAYER_GRBG)
func (b BayerGridPosition) GRBG() BayerGridPosition { return BayerGridPosition(C.NPPI_BAYER_GRBG) }

//GBRG 	returns BayerGridPosition(NPPI_BAYER_GBRG)
func (b BayerGridPosition) GBRG() BayerGridPosition { return BayerGridPosition(C.NPPI_BAYER_GBRG) }

//MaskSize has methods that are flags that are Fixed filter-kernel sizes.
type MaskSize C.NppiMaskSize

func (m MaskSize) c() C.NppiMaskSize { return C.NppiMaskSize(m) }

//Size1x3 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size1x3() MaskSize { return MaskSize(C.NPP_MASK_SIZE_1_X_3) }

//Size1x5 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size1x5() MaskSize { return MaskSize(C.NPP_MASK_SIZE_1_X_5) }

//Size3x1 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size3x1() MaskSize { return MaskSize(C.NPP_MASK_SIZE_3_X_1) }

//Size5x1 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size5x1() MaskSize { return MaskSize(C.NPP_MASK_SIZE_5_X_1) }

//Size3x3 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size3x3() MaskSize { return MaskSize(C.NPP_MASK_SIZE_3_X_3) }

//Size5x5 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size5x5() MaskSize { return MaskSize(C.NPP_MASK_SIZE_5_X_5) }

//Size7x7 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size7x7() MaskSize { return MaskSize(C.NPP_MASK_SIZE_7_X_7) }

//Size9x9 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size9x9() MaskSize { return MaskSize(C.NPP_MASK_SIZE_9_X_9) }

//Size11x11 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size11x11() MaskSize { return MaskSize(C.NPP_MASK_SIZE_11_X_11) }

//Size13x13 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size13x13() MaskSize { return MaskSize(C.NPP_MASK_SIZE_13_X_13) }

//Size15x15 returns the NPP flag wrapped in MaskSize
func (m MaskSize) Size15x15() MaskSize { return MaskSize(C.NPP_MASK_SIZE_15_X_15) }

/**
 * Differential Filter types
 */

//DifferentialKernel wraps a C.NppiDifferentialKernel
type DifferentialKernel C.NppiDifferentialKernel

func (d DifferentialKernel) c() C.NppiDifferentialKernel { return C.NppiDifferentialKernel(d) }

//SOBEL returns  DifferentialKernel(C.NPP_FILTER_SOBEL) flag
func (d DifferentialKernel) SOBEL() DifferentialKernel { return DifferentialKernel(C.NPP_FILTER_SOBEL) }

//SCHARR returns DifferentialKernel(C.NPP_FILTER_SCHARR) flag
func (d DifferentialKernel) SCHARR() DifferentialKernel {
	return DifferentialKernel(C.NPP_FILTER_SCHARR)
}

/*
 *
 * Float32
 *
 */

//Float32 is a float32 used by npp.  A pointer of this type could be in cuda memory.
type Float32 C.Npp32f /**<  32-bit (IEEE) floating-point numbers */

func (n *Float32) cptr() *C.Npp32f {
	return (*C.Npp32f)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Float32) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Float32) c() C.Npp32f {
	return C.Npp32f(n)
}
func (n Float32) isnpp() bool {
	return true
}

/*
 *
 * Float64
 *
 */

//Float64 is a float64. A pointer of this type could be in cuda memory.
type Float64 C.Npp64f /**<  64-bit floating-point numbers */

func (n *Float64) cptr() *C.Npp64f {
	return (*C.Npp64f)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Float64) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n Float64) c() C.Npp64f {
	return C.Npp64f(n)
}

/*
 *
 * Uint8
 *
 */

//Uint8 is an uint8 for npp. A pointer of this type could be in cuda memory.
type Uint8 C.Npp8u /**<  8-bit unsigned chars */
func (n *Uint8) cptr() *C.Npp8u {
	return (*C.Npp8u)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Uint8) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Uint8) c() C.Npp8u {
	return C.Npp8u(n)
}

/*
 *
 * Int8
 *
 */

//Int8 is a int8 for Npp.  A pointer of this type could be in cuda memory.
type Int8 C.Npp8s /**<  8-bit signed chars */

func (n *Int8) cptr() *C.Npp8s {
	return (*C.Npp8s)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Int8) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Int8) c() C.Npp8s {
	return C.Npp8s(n)
}

/*
 *
 * Uint16
 *
 */

//Uint16 is a uint16.  A pointer of this type could be in cuda memory.
type Uint16 C.Npp16u /**<  16-bit unsigned integers */

func (n *Uint16) cptr() *C.Npp16u {
	return (*C.Npp16u)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Uint16) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n Uint16) c() C.Npp16u {
	return C.Npp16u(n)
}

/*
 *
 * Int16
 *
 */

//Int16 is a  int16.  A pointer of this type could be in cuda memory.
type Int16 C.Npp16s /**<  16-bit signed integers */

func (n *Int16) cptr() *C.Npp16s {
	return (*C.Npp16s)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Int16) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Int16) c() C.Npp16s {
	return C.Npp16s(n)
}

/*
 *
 * Uint32
 *
 */

//Uint32 is a uint32.  A pointer of this type could be in cuda memory.
type Uint32 C.Npp32u /**<  32-bit unsigned integers */

func (n *Uint32) cptr() *C.Npp32u {
	return (*C.Npp32u)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Uint32) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Uint32) c() C.Npp32u {
	return C.Npp32u(n)
}

/*
 *
 * Int32
 *
 */

//Int32 is a int32.  A pointer of this type could be in cuda memory.
type Int32 C.Npp32s /**<  32-bit signed integers */

func (n *Int32) cptr() *C.Npp32s {
	return (*C.Npp32s)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Int32) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Int32) c() C.Npp32s {
	return C.Npp32s(n)
}

/*
 *
 * Uint64
 *
 */

//Uint64 is a uint64.  A pointer of this type could be in cuda memory.
type Uint64 C.Npp64u /**<  64-bit unsigned integers */

func (n *Uint64) cptr() *C.Npp64u {
	return (*C.Npp64u)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Uint64) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Uint64) c() C.Npp64u {
	return C.Npp64u(n)
}

/*
 *
 * Int64
 *
 */

//Int64 is a int64.  A pointer of this type could be in cuda memory.
type Int64 C.Npp64s /**<  64-bit signed integers */

func (n *Int64) cptr() *C.Npp64s {
	return (*C.Npp64s)(n)
}

//Unsafe returns an unsafe pointer to this variable location. This is so it can be used with other cuda libraries like (cudnn, cudart, cuda, and such)
func (n *Int64) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(n)
}
func (n Int64) c() C.Npp64s {
	return C.Npp64s(n)
}

func convertUint64toCNpp64uarray(x []Uint64) []C.Npp64u {
	y := make([]C.Npp64u, len(x))
	for i := range x {
		y[i] = C.Npp64u(x[i])
	}
	return y
}
func convertCNpp64utoUint64array(x []C.Npp64u) []Uint64 {
	y := make([]Uint64, len(x))
	for i := range x {
		y[i] = Uint64(x[i])
	}
	return y
}
func convertUint32toCNpp32uarray(x []Uint32) []C.Npp32u {
	y := make([]C.Npp32u, len(x))
	for i := range x {
		y[i] = C.Npp32u(x[i])
	}
	return y
}
func convertCNpp32utoUint32array(x []C.Npp32u) []Uint32 {
	y := make([]Uint32, len(x))
	for i := range x {
		y[i] = Uint32(x[i])
	}
	return y
}

func convertUint16toCNpp16uarray(x []Uint16) []C.Npp16u {
	y := make([]C.Npp16u, len(x))
	for i := range x {
		y[i] = C.Npp16u(x[i])
	}
	return y
}
func convertCNpp16utoUint16array(x []C.Npp16u) []Uint16 {
	y := make([]Uint16, len(x))
	for i := range x {
		y[i] = Uint16(x[i])
	}
	return y
}

func convertNpp8utoCNpp8uarray(x []Uint8) []C.Npp8u {
	y := make([]C.Npp8u, len(x))
	for i := range x {
		y[i] = C.Npp8u(x[i])
	}
	return y
}
func convertCNpp8utoNpp8uarray(x []C.Npp8u) []Uint8 {
	y := make([]Uint8, len(x))
	for i := range x {
		y[i] = Uint8(x[i])
	}
	return y
}

/*Uint8Complex  Complex Number
 * This struct represents an unsigned char complex number.
 */
type Uint8Complex C.Npp8uc

//Set sets the real and imaginary vals
func (n *Uint8Complex) Set(real, imaginary Uint8) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Uint8Complex) Get() (real, imaginary Uint8) {
	real = (Uint8)(n.re)
	imaginary = (Uint8)(n.im)
	return real, imaginary
}

/*Uint16Complex - See below
 * Complex Number
 * This struct represents an unsigned short complex number.
 */
type Uint16Complex C.Npp16uc

//Set sets the real and imaginary vals
func (n *Uint16Complex) Set(real, imaginary Uint16) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Uint16Complex) Get() (real, imaginary Uint16) {
	real = (Uint16)(n.re)
	imaginary = (Uint16)(n.im)
	return real, imaginary
}

/*Int16Complex - See below
 * * Complex Number
 * This struct represents a short complex number.
 */
type Int16Complex C.Npp16sc

//Set sets the real and imaginary vals
func (n *Int16Complex) Set(real, imaginary Int16) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Int16Complex) Get() (real, imaginary Int16) {
	real = (Int16)(n.re)
	imaginary = (Int16)(n.im)
	return real, imaginary
}

/*Uint32Complex - See below
 * * Complex Number
 * This struct represents an unsigned int complex number.
 */
type Uint32Complex C.Npp32uc

//Set sets the real and imaginary vals
func (n *Uint32Complex) Set(real, imaginary Uint32) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Uint32Complex) Get() (real, imaginary Uint32) {
	real = (Uint32)(n.re)
	imaginary = (Uint32)(n.im)
	return real, imaginary
}

/*Int32Complex - Complex Number
 * This struct represents a signed int complex number.
 */
type Int32Complex C.Npp32sc

//Set sets the real and imaginary vals
func (n *Int32Complex) Set(real, imaginary Int32) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Int32Complex) Get() (real, imaginary Int32) {
	real = (Int32)(n.re)
	imaginary = (Int32)(n.im)
	return real, imaginary
}

//Float32Complex This struct represents a single floating-point complex number.
type Float32Complex C.Npp32fc

func (n *Float32Complex) c() C.Npp32fc {
	return C.Npp32fc(*n)
}

//Set sets the real and imaginary vals
func (n *Float32Complex) Set(real, imaginary Float32) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Float32Complex) Get() (real, imaginary Float32) {
	real = (Float32)(n.re)
	imaginary = (Float32)(n.im)
	return real, imaginary
}

// Int64Complex struct represents a long long complex number.
type Int64Complex C.Npp64sc

//Set sets the real and imaginary vals
func (n *Int64Complex) Set(real, imaginary Int64) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Int64Complex) Get() (real, imaginary Int64) {
	real = (Int64)(n.re)
	imaginary = (Int64)(n.im)
	return real, imaginary
}

//Float64Complex struct represents a double floating-point complex number.
type Float64Complex C.Npp64fc

//Set sets the real and imaginary vals
func (n *Float64Complex) Set(real, imaginary Float64) {
	n.re = real.c()
	n.im = imaginary.c()
}

//Get gets the real and imaginary vals
func (n *Float64Complex) Get() (real, imaginary Float64) {
	real = (Float64)(n.re)
	imaginary = (Float64)(n.im)
	return real, imaginary
}

/*
typedef struct NPP_ALIGN_16
{
    Npp64f  re;
    Npp64f  im;
} Npp64fc;
*/

//#define NPP_MIN_8U      ( 0 )                        /**<  Minimum 8-bit unsigned integer */
//#define NPP_MAX_8U      ( 255 )                      /**<  Maximum 8-bit unsigned integer */
//#define NPP_MIN_16U     ( 0 )                        /**<  Minimum 16-bit unsigned integer */
//#define NPP_MAX_16U     ( 65535 )                    /**<  Maximum 16-bit unsigned integer */
//#define NPP_MIN_32U     ( 0 )                        /**<  Minimum 32-bit unsigned integer */
//#define NPP_MAX_32U     ( 4294967295U )              /**<  Maximum 32-bit unsigned integer */
//#define NPP_MIN_64U     ( 0 )                        /**<  Minimum 64-bit unsigned integer */
//#define NPP_MAX_64U     ( 18446744073709551615ULL )  /**<  Maximum 64-bit unsigned integer */

//#define NPP_MIN_8S      (-127 - 1 )                  /**<  Minimum 8-bit signed integer */
//#define NPP_MAX_8S      ( 127 )                      /**<  Maximum 8-bit signed integer */
//#define NPP_MIN_16S     (-32767 - 1 )                /**<  Minimum 16-bit signed integer */
//#define NPP_MAX_16S     ( 32767 )                    /**<  Maximum 16-bit signed integer */
//#define NPP_MIN_32S     (-2147483647 - 1 )           /**<  Minimum 32-bit signed integer */
//#define NPP_MAX_32S     ( 2147483647 )               /**<  Maximum 32-bit signed integer */
//#define NPP_MAX_64S     ( 9223372036854775807LL )    /**<  Maximum 64-bit signed integer */
//#define NPP_MIN_64S     (-9223372036854775807LL - 1) /**<  Minimum 64-bit signed integer */

//#define NPP_MINABS_32F  ( 1.175494351e-38f )         /**<  Smallest positive 32-bit floating point value */
//#define NPP_MAXABS_32F  ( 3.402823466e+38f )         /**<  Largest  positive 32-bit floating point value */
//#define NPP_MINABS_64F  ( 2.2250738585072014e-308 )  /**<  Smallest positive 64-bit floating point value */
//#define NPP_MAXABS_64F  ( 1.7976931348623158e+308 )  /**<  Largest  positive 64-bit floating point value */

//Point is a 2d point
type Point C.NppiPoint

//Set sets the Point
func (n *Point) Set(x, y int32) {
	n.x = (C.int)(x)
	n.y = (C.int)(y)
}

//Get gets the Point's x and y
func (n *Point) Get() (x, y int32) {
	return (int32)(n.x), (int32)(n.y)
}

/*
typedef struct
{
    int x;
    int y;
} NppiPoint;
*/

//PolarPoint is a 2D Polar Point
type PolarPoint C.NppPointPolar

//Set sets the polar cordinates
func (n *PolarPoint) Set(rho, theta Float32) {
	n.rho = (C.Npp32f)(rho)
	n.theta = (C.Npp32f)(theta)
}

//Get gets the polar coordinates
func (n *PolarPoint) Get() (rho, theta Float32) {
	return (Float32)(n.rho), (Float32)(n.theta)
}

/*

typedef struct {
    Npp32f rho;
    Npp32f theta;
} NppPointPolar;
*/

//Size -2D Size represents the size of a a rectangular region in two space.
type Size C.NppiSize

func (n Size) c() C.NppiSize {
	return (C.NppiSize)(n)
}
func (n *Size) cptr() *C.NppiSize {
	return (*C.NppiSize)(n)
}

/*
typedef struct
{
    int width;
    int height;
} NppiSize;
*/

//Set sets the width and Height
func (n *Size) Set(w, h int32) {
	n.width = C.int(w)
	n.height = C.int(h)

}

//Get returns the width and Height
func (n *Size) Get() (w, h int32) {
	w = int32(n.width)
	h = int32(n.height)
	return w, h
}

//Rect - 2D Rectangle This struct contains position and size information of a rectangle in two space.
// The rectangle's position is usually signified by the coordinate of its upper-left corner.
type Rect C.NppiRect

func (n Rect) c() C.NppiRect {
	return (C.NppiRect)(n)
}
func (n *Rect) cptr() *C.NppiRect {
	return (*C.NppiRect)(n)
}

//Set sets the NppiRect's values
func (n *Rect) Set(x, y, w, h int32) {
	n.x = (C.int)(x)
	n.y = (C.int)(y)
	n.width = (C.int)(w)
	n.height = (C.int)(h)
}

//Get gets the NppiRect's values
func (n *Rect) Get() (x, y, w, h int32) {
	x = (int32)(n.x)
	y = (int32)(n.y)
	w = (int32)(n.width)
	h = (int32)(n.height)
	return x, y, w, h
}

/*
typedef struct
{
    int x;         //  x-coordinate of upper left corner (lowest memory address).
    int y;       // y-coordinate of upper left corner (lowest memory address).
    int width;    // Rectangle width.
    int height;   // Rectangle height.
} NppiRect;
*/

//Axis enums NpiiAxis
type Axis C.NppiAxis

//Horizontal chooses the Horizontal Axis
func (n Axis) Horizontal() Axis {
	return Axis(C.NPP_HORIZONTAL_AXIS)
}

//Vertical chooses the verticle Axis
func (n Axis) Vertical() Axis {
	return Axis(C.NPP_VERTICAL_AXIS)
}

//Both chooses both
func (n Axis) Both() Axis {
	return Axis(C.NPP_BOTH_AXIS)
}
func (n Axis) c() C.NppiAxis {
	return (C.NppiAxis)(n)
}

//CmpOp is a flag type used for comparisons
type CmpOp C.NppCmpOp

//Less is <
func (n CmpOp) Less() CmpOp {
	return CmpOp(C.NPP_CMP_LESS)
}

//LessEq is <=
func (n CmpOp) LessEq() CmpOp {
	return CmpOp(C.NPP_CMP_LESS_EQ)
}

//Eq is =
func (n CmpOp) Eq() CmpOp {
	return CmpOp(C.NPP_CMP_EQ)
}

//GreaterEq is >=
func (n CmpOp) GreaterEq() CmpOp {
	return CmpOp(C.NPP_CMP_GREATER_EQ)
}

//Greater is >
func (n CmpOp) Greater() CmpOp {
	return CmpOp(C.NPP_CMP_GREATER)
}

/**
NppRoundMode go wrapper for roundimg modes description from original header
 * Rounding Modes
 *
 * The enumerated rounding modes are used by a large number of NPP primitives
 * to allow the user to specify the method by which fractional values are converted
 * to integer values. Also see \ref rounding_modes.
 *
 * For NPP release 5.5 new names for the three rounding modes are introduced that are
 * based on the naming conventions for rounding modes set forth in the IEEE-754
 * floating-point standard. Developers are encouraged to use the new, longer names
 * to be future proof as the legacy names will be deprecated in subsequent NPP releases.
 *
*/

//RoundMode methods return the rounding mode flags
type RoundMode C.NppRoundMode

//RndNear will round to the nearest number
func (r RoundMode) RndNear() RoundMode {
	return RoundMode(C.NPP_RND_NEAR)
}

/*RndFinancial -From Original Header
 * Round according to financial rule.
 * All fractional numbers are rounded to their nearest integer. The ambiguous
 * cases (i.e. \<integer\>.5) are rounded away from zero.
 * E.g.
 * - roundFinancial(0.4)  = 0
 * - roundFinancial(0.5)  = 1
 * - roundFinancial(-1.5) = -2
 */
func (r RoundMode) RndFinancial() RoundMode {
	return RoundMode(C.NPP_RND_FINANCIAL)
}

/*RndZero - From Original Header
 * Round towards zero (truncation).
 * All fractional numbers of the form \<integer\>.\<decimals\> are truncated to
 * \<integer\>.
 * - roundZero(1.5) = 1
 * - roundZero(1.9) = 1
 * - roundZero(-2.5) = -2
 */
func (r RoundMode) RndZero() RoundMode {
	return RoundMode(C.NPP_RND_ZERO)
}

/*
 * Other rounding modes supported by IEEE-754 (2008) floating-point standard:
 *
 * - NPP_ROUND_TOWARD_INFINITY // ceiling
 * - NPP_ROUND_TOWARD_NEGATIVE_INFINITY // floor
 *
 */

//BorderType is a flag type used to set the type of boarder.  Flags are passed through methods
type BorderType C.NppiBorderType

//Undefined returns BorderType(C.NPP_BORDER_UNDEFINED)
func (b BorderType) Undefined() BorderType {
	return BorderType(C.NPP_BORDER_UNDEFINED)
}

//None returns BorderType(C.NPP_BORDER_NONE)
func (b BorderType) None() BorderType {
	return BorderType(C.NPP_BORDER_NONE)
}

//Constant returns BorderType(C.NPP_BORDER_CONSTANT)
func (b BorderType) Constant() BorderType {
	return BorderType(C.NPP_BORDER_CONSTANT)
}

//Replicate returns  BorderType(C.NPP_BORDER_REPLICATE)
func (b BorderType) Replicate() BorderType {
	return BorderType(C.NPP_BORDER_REPLICATE)
}

//Wrap returns BorderType(C.NPP_BORDER_WRAP)
func (b BorderType) Wrap() BorderType {
	return BorderType(C.NPP_BORDER_WRAP)
}

//Mirror returns BorderType(C.NPP_BORDER_MIRROR)
func (b BorderType) Mirror() BorderType {
	return BorderType(C.NPP_BORDER_MIRROR)
}

//HintAlgorithm are flags
type HintAlgorithm C.NppHintAlgorithm

//None returns HintAlgorithm(C.NPP_ALG_HINT_NONE)
func (h HintAlgorithm) None() HintAlgorithm { return HintAlgorithm(C.NPP_ALG_HINT_NONE) }

//Fast returns HintAlgorithm(C.NPP_ALG_HINT_FAST)
func (h HintAlgorithm) Fast() HintAlgorithm { return HintAlgorithm(C.NPP_ALG_HINT_FAST) }

//Accurate returns HintAlgorithm(C.NPP_ALG_HINT_ACCURATE)
func (h HintAlgorithm) Accurate() HintAlgorithm { return HintAlgorithm(C.NPP_ALG_HINT_ACCURATE) }

/*
 * Alpha composition controls.
 */

//AlphaOp contains methods used to pass flags for composition controlls
type AlphaOp C.NppiAlphaOp

//AlphaOver returns AlphaOp(C.NPPI_OP_ALPHA_OVER)}
func (a AlphaOp) AlphaOver() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_OVER) }

//AlphaIn returns AlphaOp(C.NPPI_OP_ALPHA_IN)}
func (a AlphaOp) AlphaIn() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_IN) }

//AlphaOut returns AlphaOp(C.NPPI_OP_ALPHA_OUT)}
func (a AlphaOp) AlphaOut() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_OUT) }

//AlphaAtop returns AlphaOp(C.NPPI_OP_ALPHA_ATOP)}
func (a AlphaOp) AlphaAtop() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_ATOP) }

//AlphaXOR returns AlphaOp(C.NPPI_OP_ALPHA_XOR)}
func (a AlphaOp) AlphaXOR() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_XOR) }

//AlphaPlus returns AlphaOp(C.NPPI_OP_ALPHA_PLUS)}
func (a AlphaOp) AlphaPlus() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_PLUS) }

//AlphaOverPremul returns AlphaOp(C.NPPI_OP_ALPHA_OVER_PREMUL)}
func (a AlphaOp) AlphaOverPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_OVER_PREMUL) }

//AlphaInPremul returns AlphaOp(C.NPPI_OP_ALPHA_IN_PREMUL)}
func (a AlphaOp) AlphaInPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_IN_PREMUL) }

//AlphaOutPremul returns AlphaOp(C.NPPI_OP_ALPHA_OUT_PREMUL)}
func (a AlphaOp) AlphaOutPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_OUT_PREMUL) }

//AlphaAtopPremul returns AlphaOp(C.NPPI_OP_ALPHA_ATOP_PREMUL)}
func (a AlphaOp) AlphaAtopPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_ATOP_PREMUL) }

//AlphaXORPremul returns AlphaOp(C.NPPI_OP_ALPHA_XOR_PREMUL)}
func (a AlphaOp) AlphaXORPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_XOR_PREMUL) }

//AlphaPlusPremul returns AlphaOp(C.NPPI_OP_ALPHA_PLUS_PREMUL)}
func (a AlphaOp) AlphaPlusPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_PLUS_PREMUL) }

//AlphaPremul returns AlphaOp(C.NPPI_OP_ALPHA_PREMUL)}
func (a AlphaOp) AlphaPremul() AlphaOp { return AlphaOp(C.NPPI_OP_ALPHA_PREMUL) }

//HOGConfig type defines the configuration parameters for the HOG descriptor
type HOGConfig C.NppiHOGConfig

func (h HOGConfig) c() C.NppiHOGConfig {
	return C.NppiHOGConfig(h)
}

func (h *HOGConfig) cptr() *C.NppiHOGConfig {
	return (*C.NppiHOGConfig)(h)
}

//Set sets the HOGConfig inner types
func (h *HOGConfig) Set(cellSize, histogramBlockSize, nHistogramBins int32, detectionWindowSize Size) {
	h.cellSize = (C.int)(cellSize)
	h.histogramBlockSize = (C.int)(histogramBlockSize)
	h.histogramBlockSize = (C.int)(nHistogramBins)
	h.detectionWindowSize = detectionWindowSize.c()
}

//Get gets the inner type values
func (h *HOGConfig) Get() (cellSize, histogramBlockSize, nHistogramBins int32, detectionWindowSize Size) {
	cellSize = (int32)(h.cellSize)
	histogramBlockSize = (int32)(h.histogramBlockSize)
	nHistogramBins = (int32)(h.nHistogramBins)
	detectionWindowSize = (Size)(h.detectionWindowSize)
	return cellSize, histogramBlockSize, nHistogramBins, detectionWindowSize
}

/*
typedef struct
{
    int      cellSize;              //square cell size (pixels).
    int      histogramBlockSize;    //square histogram block size (pixels).
    int      nHistogramBins;        //required number of histogram bins.
    NppiSize detectionWindowSize;   //detection window size (pixels).
} NppiHOGConfig;
*/

//#define NPP_HOG_MAX_CELL_SIZE                          (16) /**< max horizontal/vertical pixel size of cell.   */
//#define NPP_HOG_MAX_BLOCK_SIZE                         (64) /**< max horizontal/vertical pixel size of block.  */
//#define NPP_HOG_MAX_BINS_PER_CELL                      (16) /**< max number of histogram bins. */
//#define NPP_HOG_MAX_CELLS_PER_DESCRIPTOR              (256) /**< max number of cells in a descriptor window.   */
//#define NPP_HOG_MAX_OVERLAPPING_BLOCKS_PER_DESCRIPTOR (256) /**< max number of overlapping blocks in a descriptor window.   */
//#define NPP_HOG_MAX_DESCRIPTOR_LOCATIONS_PER_CALL     (128) /**< max number of descriptor window locations per function call.   */

//HaarClassifier32f is a structure used in Haar Classification
type HaarClassifier32f C.NppiHaarClassifier_32f

func (h HaarClassifier32f) c() C.NppiHaarClassifier_32f {
	return C.NppiHaarClassifier_32f(h)
}

func (h *HaarClassifier32f) cptr() *C.NppiHaarClassifier_32f {
	return (*C.NppiHaarClassifier_32f)(h)
}

//Set sets the HOGConfig inner types
func (h *HaarClassifier32f) Set(classifiers []*Int32, classifierStep int, classifierSize Size, counterDevice *Int32) {
	h.numClassifiers = (C.int)(len(classifiers))
	h.classifiers = classifiers[0].cptr()
	h.classifierStep = (C.size_t)(classifierStep)
	h.classifierSize = classifierSize.c()
	h.counterDevice = counterDevice.cptr()
}

/*
//Get gets the inner type values
func (h *HaarClassifier32f) Get()(classifiers []*Int32, classifierStep int, classifierSize Size, counterDevice *Int32) {

numofclassifiers :=	(C.int)h.numClassifiers
classifiers=make([]Int32,	numofclassifiers)
h.classifiers
	classifiers[0].cptr() = h.classifiers =
	h.classifierStep = (C.size_t)(classifierStep)
	h.classifierSize = classifierSize.c()
	h.counterDevice = counterDevice.cptr()
*/
/*
typedef struct
{
    int      numClassifiers;     // number of classifiers
    Npp32s * classifiers;        // packed classifier data 40 bytes each
    size_t   classifierStep;
    NppiSize classifierSize;
    Npp32s * counterDevice;
} NppiHaarClassifier_32f;
*/

//HaarBuffer is a buffer for algorithms that require a HaarBuffer
type HaarBuffer C.NppiHaarBuffer

func (h HaarBuffer) c() C.NppiHaarBuffer {
	return C.NppiHaarBuffer(h)
}

func (h *HaarBuffer) cptr() *C.NppiHaarBuffer {
	return (*C.NppiHaarBuffer)(h)
}

//Get gets the HaarBuffer inner values
func (h HaarBuffer) Get() (BufferSize int32, Buffer *Int32) {
	BufferSize = (int32)(h.haarBufferSize)
	Buffer = (*Int32)(h.haarBuffer)
	return BufferSize, Buffer
}

//Set sets the HaarBuffer inner values
func (h *HaarBuffer) Set(buffsize int32, buffer *Int32) {
	h.haarBufferSize = (C.int)(buffsize)
	h.haarBuffer = buffer.cptr()
}

/*
typedef struct
{
    int      haarBufferSize;     //size of the buffer
    Npp32s * haarBuffer;        //buffer

} NppiHaarBuffer;
*/

//ZCType is a type that holds flags through methods
type ZCType C.NppsZCType

//ZCR sign change
func (z ZCType) ZCR() ZCType { return ZCType(C.nppZCR) }

//ZCXor sign change XOR
func (z ZCType) ZCXor() ZCType { return ZCType(C.nppZCXor) }

//ZCC sign change count_0
func (z ZCType) ZCC() ZCType { return ZCType(C.nppZCC) }

//HuffmanTableType is a type used for HuffmanTableType flags flags are passed by methods
type HuffmanTableType C.NppiHuffmanTableType

//DCTable - DC Table flag
func (h HuffmanTableType) DCTable() HuffmanTableType { return HuffmanTableType(C.nppiDCTable) }

//ACTable - AC Table flag
func (h HuffmanTableType) ACTable() HuffmanTableType { return HuffmanTableType(C.nppiACTable) }

//Norm is used for norm flags where needed Norm will return flags through methods
type Norm C.NppiNorm

//Inf maximum
func (n Norm) Inf() Norm { return Norm(C.nppiNormInf) }

//L1 sum
func (n Norm) L1() Norm { return Norm(C.nppiNormL1) }

//L2 square root of sum of squares
func (n Norm) L2() Norm { return Norm(C.nppiNormL2) }