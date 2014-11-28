// Copyright (c) 2014 Joseph D Poirier
// Distributable under the terms of The simplified BSD License
// that can be found in the LICENSE file.

// Package visa wraps National Instruments VISA (Virtual Instrument Software
// Architecture) driver. The driver allows a client application to communicate
// with most instrumentation buses including GPIB, USB, Serial, and Ethernet.
// The Virtual Instrument Software Architecture (VISA) is a standard for
// configuring, programming, and troubleshooting instrumentation systems
// comprising GPIB, VXI, PXI, serial (RS232/RS485), Ethernet/LXI, and/or USB
// interfaces.
//
// The package is low level and, for the most part, is one-to-one with the
// exported C functions it wraps. Clients would typically build an instrument
// specific driver around the package but it can also be used directly.
//
// NI-VISA Drivers:
//     http://www.ni.com/downloads/ni-drivers/
//
// NI-VISA Overview:
//     http://www.ni.com/white-paper/3702/en/
//
// Operation specific sections:
// - Resource Manager Functions and Operations
// - Resource Template Operations
// - Basic I/O Operations
// - Formatted and Buffered I/O Operations
// - Memory I/O Operations
// - Shared Memory Operations
// - Interface Specific Operations

package visa

/*
#if defined(__amd64) || defined(__amd64__) || defined(__x86_64) || defined(__x86_64__) || defined(_M_X64) || defined(_M_AMD64)
#cgo linux LDFLAGS: -L/usr/local/lib -lvisa64
#cgo darwin LDFLAGS: -framework VISA
#cgo windows LDFLAGS: -LC:/WINDOWS/system32 -lvisa64
#else
#cgo linux LDFLAGS: -L/usr/local/lib -lvisa
#cgo darwin LDFLAGS: -framework VISA
#cgo windows LDFLAGS: -LC:/WINDOWS/system32 -lvisa32
#endif

#cgo linux CFLAGS:
#cgo darwin CFLAGS:
#cgo windows CFLAGS:

#include <stdlib.h>
#include <visa.h>

extern void go_cb(ViSession, ViEventType, ViEvent, ViAddr);
ViHndlr get_go_cb(void) {
	return (ViHndlr)go_cb;
}
*/
import "C"
import "unsafe"

var PackageVersion string = "v0.1"

type Status int32
type Session uint32
type Object uint32

// Platform (32 or 64 bit) specific, as determined at compile time
type BusAddress C.ViBusAddress
type PBusAddress C.ViPBusAddress
type BusSize C.ViBusSize
type AttrState C.ViAttrState
type Bool C.ViBoolean

//
type UserCallback func(instr Object, etype, eventContext uint32)
type PUserCallback *UserCallback

// -----------------------------------------
// Resource Manager Functions and Operations
// -----------------------------------------

// OpenDefaultRM returns a session to the Default Resource Manager resource.
func OpenDefaultRM() (rm Session, status Status) {
	status = Status(C.viOpenDefaultRM((*C.ViSession)(unsafe.Pointer(&rm))))
	return rm, status
}

// legacy
var GetDefaultRM = OpenDefaultRM

// FindRsrc queries a VISA system to locate the resources associated with a specified interface.
func (rm Session) FindRsrc(expr string) (findList, retCnt uint32, desc string, status Status) {
	cexpr := (*C.ViChar)(C.CString(expr))
	defer C.free(unsafe.Pointer(cexpr))
	d := make([]byte, 257)
	status = Status(C.viFindRsrc(C.ViSession(rm),
		cexpr,
		(*C.ViFindList)(unsafe.Pointer(&findList)),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt)),
		(*C.ViChar)(unsafe.Pointer(&d[0]))))
	return findList, retCnt, string(d), status
}

// FindNext gets the next resource from the list of resources found during a
// previous call to FindRsrc.
func FindNext(findList uint32) (string, Status) {
	d := make([]byte, 257)
	status := Status(C.viFindNext((C.ViFindList)(findList),
		(*C.ViChar)(unsafe.Pointer(&d[0]))))
	return string(d), status
}

// ParseRsrc parses a resource string to get the interface information.
func (rm Session) ParseRsrc(rsrcName string) (intfType, intfNum uint16, status Status) {
	crsrcName := (*C.ViChar)(C.CString(rsrcName))
	defer C.free(unsafe.Pointer(crsrcName))
	status = Status(C.viParseRsrc(C.ViSession(rm),
		crsrcName,
		(*C.ViUInt16)(unsafe.Pointer(&intfType)),
		(*C.ViUInt16)(unsafe.Pointer(&intfNum))))
	return intfType, intfNum, status
}

// ParseRsrcEx Parses a resource string to get extended interface information.
func (rm Session) ParseRsrcEx(rsrcName string) (intfType, intfNum uint16,
	rsrcClass, expandedUnaliasedName, aliasIfExists string, status Status) {

	crsrcName := (*C.ViChar)(C.CString(rsrcName))
	defer C.free(unsafe.Pointer(crsrcName))
	r := make([]byte, 257)
	e := make([]byte, 257)
	a := make([]byte, 257)
	status = Status(C.viParseRsrcEx(C.ViSession(rm),
		crsrcName,
		(*C.ViUInt16)(unsafe.Pointer(&intfType)),
		(*C.ViUInt16)(unsafe.Pointer(&intfNum)),
		(*C.ViChar)(unsafe.Pointer(&r[0])),
		(*C.ViChar)(unsafe.Pointer(&e[0])),
		(*C.ViChar)(unsafe.Pointer(&a[0]))))
	return intfType, intfNum, string(r), string(e), string(a), status
}

// Open Opens a session to the specified resource.
func (rm Session) Open(name string, mode, timeout uint32) (instr Object, status Status) {
	cname := (*C.ViChar)(C.CString(name))
	defer C.free(unsafe.Pointer(cname))
	status = Status(C.viOpen(C.ViSession(rm),
		cname,
		(C.ViAccessMode)(mode),
		(C.ViUInt32)(timeout),
		(*C.ViSession)(unsafe.Pointer(&instr))))
	return instr, status
}

// ----------------------------
// Resource Template Operations
// ----------------------------

// Close Closes the specified session.
func (rm Session) Close() Status {
	return Status(C.viClose((C.ViObject)(rm)))
}

// Close Closes the specified instrument, or find list.
func (instr Object) Close() Status {
	return Status(C.viClose((C.ViObject)(instr)))
}

// Close Closes the specified find list.
func Close(list uint32) Status {
	return Status(C.viClose((C.ViObject)(list)))
}

// SetAttribute Sets the state of an attribute.
func (instr Object) SetAttribute(attribute, attrState uint32) Status {
	return Status(C.viSetAttribute((C.ViObject)(instr),
		(C.ViAttr)(attribute),
		(C.ViAttrState)(attrState)))
}

// GetAttribute Retrieves the state of an attribute.
func (instr Object) GetAttribute(attrName uint32, addr unsafe.Pointer) Status {
	return Status(C.viGetAttribute((C.ViObject)(instr), (C.ViAttr)(attrName), addr))
}

// StatusDesc Returns a user-readable description of the status code passed to the operation.
func (instr Object) StatusDesc(status_in Status) (string, Status) {
	d := make([]byte, 257)
	status := Status(C.viStatusDesc((C.ViObject)(instr),
		(C.ViStatus)(status_in),
		(*C.ViChar)(unsafe.Pointer(&d[0]))))
	return string(d), status
}

// Terminate Requests a VISA session to terminate normal execution of an operation.
func (instr Object) Terminate(degree, jobId uint16) Status {
	return Status(C.viTerminate((C.ViObject)(instr),
		(C.ViUInt16)(degree),
		(C.ViJobId)(jobId)))
}

// Lock Establishes an access mode to the specified resource.
func (instr Object) LockExclusive(lockType, timeout uint32) Status {
	return Status(C.viLock((C.ViSession)(instr),
		(C.ViAccessMode)(lockType),
		(C.ViUInt32)(timeout),
		(*C.ViChar)(nil),
		(*C.ViChar)(nil)))
}

// Lock Establishes an access mode to the specified resource.
func (instr Object) Lock(lockType, timeout uint32, requestedKey string) (string, Status) {
	rk := (*C.ViChar)(C.CString(requestedKey))
	defer C.free(unsafe.Pointer(rk))
	a := make([]byte, 257)
	status := Status(C.viLock((C.ViSession)(instr),
		(C.ViAccessMode)(lockType),
		(C.ViUInt32)(timeout),
		rk,
		(*C.ViChar)(unsafe.Pointer(&a[0]))))
	return string(a), status
}

// Unlock Relinquishes a lock for the specified resource.
func (instr Object) Unlock() Status {
	return Status(C.viUnlock((C.ViSession)(instr)))
}

// EnableEvent Enables notification of a specified event.
func (instr Object) EnableEvent(eventType uint32, mechanism uint16, context uint32) Status {

	return Status(C.viEnableEvent((C.ViSession)(instr),
		(C.ViEventType)(eventType),
		(C.ViUInt16)(mechanism),
		(C.ViEventFilter)(context)))
}

// DisableEvent Disables notification of the specified event type(s)
// via the specified mechanism(s).
func (instr Object) DisableEvent(eventType uint32, mechanism uint16) Status {
	return Status(C.viDisableEvent((C.ViSession)(instr),
		(C.ViEventType)(eventType),
		(C.ViUInt16)(mechanism)))
}

// DiscardEvents Discards event occurrences for specified event types
// and mechanisms in a session.
func (instr Object) DiscardEvents(eventType uint32, mechanism uint16) Status {
	return Status(C.viDiscardEvents((C.ViSession)(instr),
		(C.ViEventType)(eventType),
		(C.ViUInt16)(mechanism)))
}

// WaitOnEvent Waits for an occurrence of the specified event for a given session.
func (instr Object) WaitOnEvent(inEventType, timeout uint32) (outEventType,
	outContext uint32, status Status) {

	status = Status(C.viWaitOnEvent((C.ViSession)(instr),
		(C.ViEventType)(inEventType),
		(C.ViUInt32)(timeout),
		(*C.ViEventType)(unsafe.Pointer(&outEventType)),
		(*C.ViEvent)(unsafe.Pointer(&outContext))))
	return outEventType, outContext, status
}

// InstallHandler Installs handlers for event callbacks.
func (instr Object) InstallHandler(eventType uint32, userHandle UserCallback) Status {
	return Status(C.viInstallHandler((C.ViSession)(instr),
		(C.ViEventType)(eventType),
		(C.ViHndlr)(C.get_go_cb()),
		(C.ViAddr)(unsafe.Pointer(&userHandle))))
}

// UninstallHandler Uninstalls handlers for events.
// Note that VISA identifies handlers uniquely using the userHandle reference.
func (instr Object) UninstallHandler(eventType uint32, userHandle UserCallback) Status {
	return Status(C.viUninstallHandler((C.ViSession)(instr),
		(C.ViEventType)(eventType),
		(C.ViHndlr)(C.get_go_cb()),
		(C.ViAddr)(unsafe.Pointer(&userHandle))))
}

// --------------------
// Basic I/O Operations
// --------------------

// Read Reads data from device or interface synchronously.
func (instr Object) Read(cnt uint32) (buf []byte, retCnt uint32, status Status) {
	buf = make([]byte, cnt)
	status = Status(C.viRead((C.ViSession)(instr),
		(*C.ViByte)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return buf, retCnt, status
}

// ReadAsync Reads data from device or interface asynchronously.
func (instr Object) ReadAsync(cnt uint32) (buf []byte, jobId uint32, status Status) {
	buf = make([]byte, cnt)
	status = Status(C.viReadAsync((C.ViSession)(instr),
		(*C.ViByte)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViJobId)(unsafe.Pointer(&jobId))))
	return buf, jobId, status
}

// ReadToFile Reads data synchronously and stores the transferred data in a file.
func (instr Object) ReadToFile(filename string, cnt uint32) (retCnt uint32, status Status) {
	cfilename := (*C.ViChar)(C.CString(filename))
	defer C.free(unsafe.Pointer(cfilename))
	status = Status(C.viReadToFile((C.ViSession)(instr),
		cfilename,
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return retCnt, status
}

// Write Writes data to a device or interface synchronously.
func (instr Object) Write(buf []byte, cnt uint32) (retCnt uint32, status Status) {
	status = Status(C.viWrite((C.ViSession)(instr),
		(C.ViBuf)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return retCnt, status
}

// WriteAsync Writes data to a device or interface asynchronously.
func (instr Object) WriteAsync(buf []byte, cnt uint32) (jobId uint32, status Status) {
	status = Status(C.viWriteAsync((C.ViSession)(instr),
		(C.ViBuf)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViJobId)(unsafe.Pointer(&jobId))))
	return jobId, status
}

// WriteFromFile Take data from a file and write it out synchronously.
func (instr Object) WriteFromFile(filename string, cnt uint32) (retCnt uint32, status Status) {
	cfilename := (*C.ViChar)(C.CString(filename))
	defer C.free(unsafe.Pointer(cfilename))
	status = Status(C.viWriteFromFile((C.ViSession)(instr),
		cfilename,
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return retCnt, status
}

// AssertTrigger Asserts software or hardware trigger.
func (instr Object) AssertTrigger(protocol uint16) Status {
	return Status(C.viAssertTrigger((C.ViSession)(instr),
		(C.ViUInt16)(protocol)))
}

// ReadSTB Reads a status byte of the service request.
func (instr Object) ReadSTB() (stb_stat uint16, status Status) {
	status = Status(C.viReadSTB((C.ViSession)(instr),
		(*C.ViUInt16)(unsafe.Pointer(&stb_stat))))
	return stb_stat, status
}

// Clear Clears a device.
func (instr Object) Clear() Status {
	return Status(C.viClear((C.ViSession)(instr)))
}

// -------------------------------------
// Formatted and Buffered I/O Operations
// -------------------------------------

// ViSetBuf Sets the size for the formatted I/O and/or low-level
// I/O communication buffer(s).
func (instr Object) ViSetBuf(mask uint16, size uint32) Status {
	return Status(C.viSetBuf((C.ViSession)(instr),
		(C.ViUInt16)(mask),
		(C.ViUInt32)(size)))
}

// Flush Manually flushes the specified buffers associated with
// formatted I/O operations and/or serial communication.
func (instr Object) Flush(mask uint16) Status {
	return Status(C.viFlush((C.ViSession)(instr),
		(C.ViUInt16)(mask)))
}

// BufWrite Writes data to a formatted I/O write buffer synchronously.
func (instr Object) BufWrite(buf []byte, cnt uint32) (retCnt uint32, status Status) {
	status = Status(C.viBufWrite((C.ViSession)(instr),
		(C.ViBuf)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return retCnt, status
}

// BufRead Reads data from a device or interface through the use of a formatted I/O read buffer.
func (instr Object) BufRead(cnt uint32) (buf []byte, retCnt uint32, status Status) {
	buf = make([]byte, cnt)
	status = Status(C.viBufRead((C.ViSession)(instr),
		(C.ViBuf)(unsafe.Pointer(&buf[0])),
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(unsafe.Pointer(&retCnt))))
	return buf, retCnt, status
}

// TODO: formatted IO

// Printf Converts, formats, and sends the parameters (designated by ...)
// to the device as specified by the format string.
// ViStatus _VI_FUNCC viPrintf        (ViSession vi, ViString writeFmt, ...);
// func (instr Object) Printf(writeFmt string, args ...interface{}) Status {
// 	cwriteFmt := (*C.ViChar)(C.CString(writeFmt))
// 	defer C.free(unsafe.Pointer(cwriteFmt))
// 	return Status(C.viPrintf((C.ViSession)(instr),
// 		cwriteFmt,
// 		(*C.ViVAList)(args)))
// }

// VPrintf Converts, formats, and sends the parameters designated by params
// to the device or interface as specified by the format string.
// ViStatus _VI_FUNC  viVPrintf       (ViSession vi, ViString writeFmt, ViVAList params);
// func (instr Object) VPrintf(writeFmt string, args ...interface{}) Status {
// 	cwriteFmt := (*C.ViChar)(C.CString(writeFmt))
// 	defer C.free(unsafe.Pointer(cwriteFmt))
// 	return Status(C.viVPrintf((C.ViSession)(instr),
// 		cwriteFmt,
// 		(*C.ViVAList)(args)))
// }

// SPrintf Converts, formats, and sends the parameters (designated by ...)
// to a user-specified buffer as specified by the format string.
// ViStatus _VI_FUNCC viSPrintf       (ViSession vi, ViPBuf buf, ViString writeFmt, ...);

// VSPrintf Converts, formats, and sends the parameters designated by params
// to a user-specified buffer as specified by the format string.
// ViStatus _VI_FUNC  viVSPrintf      (ViSession vi, ViPBuf buf, ViString writeFmt,
//                                     ViVAList parms);

// Scanf Reads, converts, and formats data using the format specifier.
// Stores the formatted data in the parameters (designated by ...).
// ViStatus _VI_FUNCC viScanf         (ViSession vi, ViString readFmt, ...);

// VScanf Reads, converts, and formats data using the format specifier.
// Stores the formatted data in the parameters designated by params.
// ViStatus _VI_FUNC  viVScanf        (ViSession vi, ViString readFmt, ViVAList params);

// SScanf Reads, converts, and formats data from a user-specified buffer
// using the format specifier. Stores the formatted data in the parameters
// (designated by ...).
// ViStatus _VI_FUNCC viSScanf        (ViSession vi, ViBuf buf, ViString readFmt, ...);

// VSScanf Reads, converts, and formats data from a user-specified buffer
// using the format specifier. Stores the formatted data in the parameters
// designated by params.
// ViStatus _VI_FUNC  viVSScanf       (ViSession vi, ViBuf buf, ViString readFmt,
//                                     ViVAList parms);

// Queryf Performs a formatted write and read through a single call
// to an operation.
// ViStatus _VI_FUNCC viQueryf        (ViSession vi, ViString writeFmt, ViString readFmt, ...);

// VQueryf Performs a formatted write and read through a single call
// to an operation.
// ViStatus _VI_FUNC  viVQueryf       (ViSession vi, ViString writeFmt, ViString readFmt,
//                                     ViVAList params);

// ---------------------
// Memory I/O Operations
// ---------------------

// In8 Reads in an 8-bit value from the specified memory space and offset.
func (instr Object) iIn8(space uint16, offset BusAddress) (val uint8, status Status) {
	status = Status(C.viIn8((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(*C.ViUInt8)(&val)))
	return val, status
}

// Out8 Writes an 8-bit value to the specified memory space and offset.
func (instr Object) Out8(space uint16, offset BusAddress, val uint8) Status {
	return Status(C.viOut8((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViUInt8)(val)))
}

// In16 Reads in an 16-bit value from the specified memory space and offset.
func (instr Object) In16(space uint16, offset BusAddress) (val uint16, status Status) {
	status = Status(C.viIn16((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(*C.ViUInt16)(&val)))
	return val, status
}

// Out16 Writes an 16-bit value to the specified memory space and offset.
func (instr Object) Out16(space uint16, offset BusAddress, val uint16) Status {
	return Status(C.viOut16((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViUInt16)(val)))
}

// In32 Reads in an 32-bit value from the specified memory space and offset.
func (instr Object) In32(space uint16, offset BusAddress) (val uint32, status Status) {
	status = Status(C.viIn32((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(*C.ViUInt32)(&val)))
	return val, status
}

// Out32 Writes an 32-bit value to the specified memory space and offset.
func (instr Object) Out32(space uint16, offset BusAddress, val uint32) Status {
	return Status(C.viOut32((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViUInt32)(val)))
}

// In64 Reads in an 64-bit value from the specified memory space and offset.
func (instr Object) In64(space uint16, offset BusAddress) (val uint64, status Status) {
	status = Status(C.viIn64((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(*C.ViUInt64)(&val)))
	return val, status
}

// Out64 Writes an 64-bit value to the specified memory space and offset.
func (instr Object) Out64(space uint16, offset BusAddress, val uint64) Status {
	return Status(C.viOut64((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViUInt64)(val)))
}

// MoveIn8 Moves a block of data from the specified address space and offset to local memory.
func (instr Object) MoveIn8(space uint16, offset BusAddress, length BusSize) ([]uint8, Status) {
	buf := make([]uint8, length)
	status := Status(C.viMoveIn8((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt8)(unsafe.Pointer(&buf[0]))))
	return buf, status
}

// MoveOut8 Moves a block of data from local memory to the specified
func (instr Object) MoveOut8(space uint16, offset BusAddress, length BusSize, buf []uint8) Status {
	return Status(C.viMoveOut8((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt8)(unsafe.Pointer(&buf[0]))))
}

// MoveIn16 Moves a block of data from the specified address space and offset to local memory.
func (instr Object) MoveIn16(space uint16, offset BusAddress, length BusSize) ([]uint16, Status) {
	buf := make([]uint16, length)
	status := Status(C.viMoveIn16((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt16)(unsafe.Pointer(&buf[0]))))
	return buf, status
}

// MoveOut16 Moves a block of data from local memory to the specified address space and offset.
func (instr Object) MoveOut16(space uint16, offset BusAddress, length BusSize, buf []uint16) Status {
	return Status(C.viMoveOut16((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt16)(unsafe.Pointer(&buf[0]))))
}

// MoveIn32 Moves a block of data from the specified address space and offset to local memory.
func (instr Object) MoveIn32(space uint16, offset BusAddress, length BusSize) ([]uint32, Status) {
	buf := make([]uint32, length)
	status := Status(C.viMoveIn32((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt32)(unsafe.Pointer(&buf[0]))))
	return buf, status
}

// MoveOut32 Moves a block of data from local memory to the specified address space and offset.
func (instr Object) MoveOut32(space uint16, offset BusAddress, length BusSize, buf []uint32) Status {
	return Status(C.viMoveOut32((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt32)(unsafe.Pointer(&buf[0]))))
}

// MoveIn64 Moves a block of data from the specified address space and offset to local memory.
func (instr Object) MoveIn64(space uint16, offset BusAddress, length BusSize) ([]uint64, Status) {
	buf := make([]uint64, length)
	status := Status(C.viMoveIn64((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt64)(unsafe.Pointer(&buf[0]))))
	return buf, status
}

// MoveOut64 Moves a block of data from local memory to the specified address space and offset.
func (instr Object) MoveOut64(space uint16, offset BusAddress, length BusSize, buf []uint64) Status {
	return Status(C.viMoveOut64((C.ViSession)(instr),
		(C.ViUInt16)(space),
		(C.ViBusAddress)(offset),
		(C.ViBusSize)(length),
		(C.ViAUInt64)(unsafe.Pointer(&buf[0]))))
}

// Move Moves a block of data.
func (instr Object) Move(srcSpace uint16, srcOffset BusAddress, srcWidth uint16, destSpace uint16,
	destOffset BusAddress, destWidth uint16, srcLength BusSize) Status {

	return Status(C.viMove((C.ViSession)(instr),
		(C.ViUInt16)(srcSpace),
		(C.ViBusAddress)(srcOffset),
		(C.ViUInt16)(srcWidth),
		(C.ViUInt16)(destSpace),
		(C.ViBusAddress)(destOffset),
		(C.ViUInt16)(destWidth),
		(C.ViBusSize)(srcLength)))
}

// MoveAsync Moves a block of data asynchronously.
func (instr Object) MoveAsync(srcSpace uint16, srcOffset BusAddress, srcWidth, destSpace uint16,
	destOffset BusAddress, destWidth uint16, srcLength BusSize) (jobId uint32, status Status) {

	status = Status(C.viMoveAsync((C.ViSession)(instr),
		(C.ViUInt16)(srcSpace),
		(C.ViBusAddress)(srcOffset),
		(C.ViUInt16)(srcWidth),
		(C.ViUInt16)(destSpace),
		(C.ViBusAddress)(destOffset),
		(C.ViUInt16)(destWidth),
		(C.ViBusSize)(srcLength),
		(*C.ViJobId)(unsafe.Pointer(&jobId))))
	return jobId, status
}

// MapAddress Maps the specified memory space into the process’s address space.
func (instr Object) MapAddress(mapSpace uint16, mapOffset BusAddress, mapSize BusSize,
	access uint16, suggested *byte) (address *byte, status Status) {

	status = Status(C.viMapAddress((C.ViSession)(instr),
		(C.ViUInt16)(mapSpace),
		(C.ViBusAddress)(mapOffset),
		(C.ViBusSize)(mapSize),
		(C.ViBoolean)(access),
		(C.ViAddr)(unsafe.Pointer(suggested)),
		(*C.ViAddr)(unsafe.Pointer(&address))))
	return address, status
}

// UnmapAddress Unmaps memory space previously mapped by ViMapAddress.
func (instr Object) UnmapAddress() Status {
	return Status(C.viUnmapAddress((C.ViSession)(instr)))
}

// Peek8 Reads an 8-bit value from the specified address.
func (instr Object) Peek8(address unsafe.Pointer) (val uint8) {
	C.viPeek8((C.ViSession)(instr), (C.ViAddr)(address), (*C.ViUInt8)(&val))
	return val
}

// Poke8 Writes an 8-bit value to the specified address.
func (instr Object) Poke8(address unsafe.Pointer, val uint8) {
	C.viPoke8((C.ViSession)(instr), (C.ViAddr)(address), (C.ViUInt8)(val))
}

// Peek16 Reads an 16-bit value from the specified address.
func (instr Object) Peek16(address unsafe.Pointer) (val uint16) {
	C.viPeek16((C.ViSession)(instr), (C.ViAddr)(address), (*C.ViUInt16)(&val))
	return val
}

// Poke16 Writes an 16-bit value to the specified address.
func (instr Object) Poke16(address unsafe.Pointer, val uint16) {
	C.viPoke16((C.ViSession)(instr), (C.ViAddr)(address), (C.ViUInt16)(val))
}

// Peek32 Reads an 32-bit value from the specified address.
func (instr Object) Peek32(address unsafe.Pointer) (val uint32) {
	C.viPeek32((C.ViSession)(instr), (C.ViAddr)(address), (*C.ViUInt32)(&val))
	return val
}

// Poke32 Writes an 32-bit value to the specified address.
func (instr Object) Poke32(address unsafe.Pointer, val uint32) {
	C.viPoke32((C.ViSession)(instr), (C.ViAddr)(address), (C.ViUInt32)(val))
}

// Peek64 Reads an 64-bit value from the specified address.
func (instr Object) Peek64(address unsafe.Pointer) (val uint64) {
	C.viPeek64((C.ViSession)(instr), (C.ViAddr)(address), (*C.ViUInt64)(&val))
	return val
}

// Poke64 Writes an 64-bit value to the specified address.
func (instr Object) Poke64(address unsafe.Pointer, val uint64) {
	C.viPoke64((C.ViSession)(instr), (C.ViAddr)(address), (C.ViUInt64)(val))
}

// ------------------------
// Shared Memory Operations
// ------------------------

// MemAlloc Allocates memory from a device’s memory region.
func (instr Object) MemAlloc(size BusSize) (offset BusAddress, status Status) {
	status = Status(C.viMemAlloc((C.ViSession)(instr),
		(C.ViBusSize)(size),
		(*C.ViBusAddress)(&offset)))
	return offset, status
}

// MemFree Frees memory previously allocated using the viMemAlloc() operation.
func (instr Object) MemFree(offset BusAddress) Status {
	return Status(C.viMemFree((C.ViSession)(instr), (C.ViBusAddress)(offset)))
}

// -----------------------------
// Interface Specific Operations
// -----------------------------

// GpibControlREN Controls the state of the GPIB Remote Enable (REN)
// interface line, and optionally the remote/local state of the device.
func (instr Object) GpibControlREN(mode uint16) Status {
	return Status(C.viGpibControlREN((C.ViSession)(instr),
		(C.ViUInt16)(mode)))
}

// GpibControlATN Specifies the state of the ATN line and the local
// active controller state.
func (instr Object) GpibControlATN(mode uint16) Status {
	return Status(C.viGpibControlATN((C.ViSession)(instr),
		(C.ViUInt16)(mode)))
}

// GpibSendIFC Pulse the interface clear line (IFC) for at least 100 microseconds.
func (instr Object) GpibSendIFC() Status {
	return Status(C.viGpibSendIFC((C.ViSession)(instr)))
}

// GpibCommand Write GPIB command bytes on the bus.
// ViStatus _VI_FUNC  viGpibCommand   (ViSession vi, ViBuf cmd, ViUInt32 cnt, ViPUInt32 retCnt);
func (instr Object) GpibCommand(cmd []byte, cnt uint32) (retCnt uint32, status Status) {
	status = Status(C.viGpibCommand((C.ViSession)(instr),
		(C.ViBuf)(unsafe.Pointer(&cmd[0])),
		(C.ViUInt32)(cnt),
		(*C.ViUInt32)(&retCnt)))
	return retCnt, status
}

// GpibPassControl Tell the GPIB device at the specified address to
// become controller in charge (CIC).
func (instr Object) GpibPassControl(primAddr, secAddr uint16) Status {
	return Status(C.viGpibPassControl((C.ViSession)(instr),
		(C.ViUInt16)(primAddr),
		(C.ViUInt16)(secAddr)))
}

// VxiCommandQuery Sends the device a miscellaneous command or query and/or
// retrieves the response to a previous query.
func (instr Object) VxiCommandQuery(mode uint16, cmd uint32) (response uint32, status Status) {
	status = Status(C.viVxiCommandQuery((C.ViSession)(instr),
		(C.ViUInt16)(mode),
		(C.ViUInt32)(cmd),
		(*C.ViUInt32)(&response)))
	return response, status
}

// AssertUtilSignal Asserts or deasserts the specified utility bus signal.
func (instr Object) AssertUtilSignal(line uint16) Status {
	return Status(C.viAssertUtilSignal((C.ViSession)(instr),
		(C.ViUInt16)(line)))
}

// AssertIntrSignal Asserts the specified interrupt or signal.
func (instr Object) AssertIntrSignal(mode int16, statusID uint16) Status {
	return Status(C.viAssertIntrSignal((C.ViSession)(instr),
		(C.ViInt16)(mode),
		(C.ViUInt32)(statusID)))
}

// MapTrigger Map the specified trigger source line to the specified
// destination line.
func (instr Object) MapTrigger(trigSrc, trigDest int16, mode uint16) Status {
	return Status(C.viMapTrigger((C.ViSession)(instr),
		(C.ViInt16)(trigSrc),
		(C.ViInt16)(trigDest),
		(C.ViUInt16)(mode)))
}

// UnmapTrigger Undo a previous map from the specified trigger source
// line to the specified destination line.
func (instr Object) UnmapTrigger(trigSrc, trigDest int16) Status {
	return Status(C.viUnmapTrigger((C.ViSession)(instr),
		(C.ViInt16)(trigSrc),
		(C.ViInt16)(trigDest)))
}

// UsbControlOut Performs a USB control pipe transfer to the device.
func (instr Object) UsbControlOut(bmRequestType, bRequest int16, wValue, wIndex,
	wLength uint16, buf []byte) Status {

	return Status(C.viUsbControlOut((C.ViSession)(instr),
		(C.ViInt16)(bmRequestType),
		(C.ViInt16)(bRequest),
		(C.ViUInt16)(wValue),
		(C.ViUInt16)(wIndex),
		(C.ViUInt16)(wLength),
		(*C.ViByte)(unsafe.Pointer(&buf[0]))))
}

// UsbControlIn Performs a USB control pipe transfer from the device.
func (instr Object) UsbControlIn(bmRequestType, bRequest int16, wValue, wIndex,
	wLength uint16) (buf []byte, retCnt uint16, status Status) {

	buf = make([]byte, wLength)
	status = Status(C.viUsbControlIn((C.ViSession)(instr),
		(C.ViInt16)(bmRequestType),
		(C.ViInt16)(bRequest),
		(C.ViUInt16)(wValue),
		(C.ViUInt16)(wIndex),
		(C.ViUInt16)(wLength),
		(*C.ViByte)(unsafe.Pointer(&buf[0])),
		(*C.ViUInt16)(&retCnt)))
	return buf, retCnt, status
}

// Version Returns the unformatted resource version number.
func Version() uint32 {
	return uint32(C.VI_SPEC_VERSION)
}

// VersMajor Returns the major resource version number.
func VersMajor() uint32 {
	return uint32((Version() & 0xFFF00000) >> 20)
}

// VersMinor Returns the minor resource version number.
func VersMinor() uint32 {
	return uint32((Version() & 0x000FFF00) >> 8)
}

// VersSubMinor Returns the sub-minor resource version number.
func VersSubMinor() uint32 {
	return uint32((Version() & 0x000000FF))
}

// PxiReserveTriggers Reserves multiple trigger lines that the caller can then map and/or assert.
func (instr Object) PxiReserveTriggers(cnt int16, trigBuses, trigLines *int16) (failureIndex int16, status Status) {
	status = Status(C.viPxiReserveTriggers((C.ViSession)(instr),
		(C.ViInt16)(cnt),
		(*C.ViInt16)(trigBuses),
		(*C.ViInt16)(trigLines),
		(*C.ViInt16)(&failureIndex)))
	return failureIndex, status
}

// VxiServantResponse ?
func (instr Object) VxiServantResponse(mode int16, resp uint32) Status {
	return Status(C.viVxiServantResponse((C.ViSession)(instr),
		(C.ViInt16)(mode),
		(C.ViUInt32)(resp)))
}
