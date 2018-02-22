// +build !cl10

package cl

// #include "cl.h"
import "C"
import "unsafe"

const FPConfigCorrectlyRoundedDivideSqrt FPConfig = C.CL_FP_CORRECTLY_ROUNDED_DIVIDE_SQRT

func init() {
	fpConfigNameMap[FPConfigCorrectlyRoundedDivideSqrt] = "CorrectlyRoundedDivideSqrt"
}

func (d *Device) BuiltInKernels() string {
	return d.mustGetInfoString(C.CL_DEVICE_BUILT_IN_KERNELS)
}

// Is CL_FALSE if the implementation does not have a linker available. Is CL_TRUE if the linker is available. This can be CL_FALSE for the embedded platform profile only. This must be CL_TRUE if CL_DEVICE_COMPILER_AVAILABLE is CL_TRUE
func (d *Device) LinkerAvailable() bool {
	return d.mustGetInfoBool(C.CL_DEVICE_LINKER_AVAILABLE)
}

func (d *Device) ParentDevice() *Device {
	var deviceId C.cl_device_id
	if err := C.clGetDeviceInfo(d.id, C.CL_DEVICE_PARENT_DEVICE, C.size_t(unsafe.Sizeof(deviceId)), unsafe.Pointer(&deviceId), nil); err != C.CL_SUCCESS {
		panic("ParentDevice failed")
	}
	if deviceId == nil {
		return nil
	}
	return &Device{id: deviceId}
}

// Max number of pixels for a 1D image created from a buffer object. The minimum value is 65536 if CL_DEVICE_IMAGE_SUPPORT is CL_TRUE.
func (d *Device) ImageMaxBufferSize() int {
	return int(d.mustGetInfoSize(C.CL_DEVICE_IMAGE_MAX_BUFFER_SIZE))
}

// Max number of images in a 1D or 2D image array. The minimum value is 2048 if CL_DEVICE_IMAGE_SUPPORT is CL_TRUE
func (d *Device) ImageMaxArraySize() int {
	return int(d.mustGetInfoSize(C.CL_DEVICE_IMAGE_MAX_ARRAY_SIZE))
}
