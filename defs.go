package visa

/*
#cgo windows CFLAGS: -I.
#include <stdlib.h>
#include <visa.h>
*/
import "C"

const (
	VI_SPEC_VERSION = C.VI_SPEC_VERSION
	// Attributes (platform independent size)
	VI_ATTR_RSRC_CLASS                  = C.VI_ATTR_RSRC_CLASS
	VI_ATTR_RSRC_NAME                   = C.VI_ATTR_RSRC_NAME
	VI_ATTR_RSRC_IMPL_VERSION           = C.VI_ATTR_RSRC_IMPL_VERSION
	VI_ATTR_RSRC_LOCK_STATE             = C.VI_ATTR_RSRC_LOCK_STATE
	VI_ATTR_MAX_QUEUE_LENGTH            = C.VI_ATTR_MAX_QUEUE_LENGTH
	VI_ATTR_USER_DATA_32                = C.VI_ATTR_USER_DATA_32
	VI_ATTR_FDC_CHNL                    = C.VI_ATTR_FDC_CHNL
	VI_ATTR_FDC_MODE                    = C.VI_ATTR_FDC_MODE
	VI_ATTR_FDC_GEN_SIGNAL_EN           = C.VI_ATTR_FDC_GEN_SIGNAL_EN
	VI_ATTR_FDC_USE_PAIR                = C.VI_ATTR_FDC_USE_PAIR
	VI_ATTR_SEND_END_EN                 = C.VI_ATTR_SEND_END_EN
	VI_ATTR_TERMCHAR                    = C.VI_ATTR_TERMCHAR
	VI_ATTR_TMO_VALUE                   = C.VI_ATTR_TMO_VALUE
	VI_ATTR_GPIB_READDR_EN              = C.VI_ATTR_GPIB_READDR_EN
	VI_ATTR_IO_PROT                     = C.VI_ATTR_IO_PROT
	VI_ATTR_DMA_ALLOW_EN                = C.VI_ATTR_DMA_ALLOW_EN
	VI_ATTR_ASRL_BAUD                   = C.VI_ATTR_ASRL_BAUD
	VI_ATTR_ASRL_DATA_BITS              = C.VI_ATTR_ASRL_DATA_BITS
	VI_ATTR_ASRL_PARITY                 = C.VI_ATTR_ASRL_PARITY
	VI_ATTR_ASRL_STOP_BITS              = C.VI_ATTR_ASRL_STOP_BITS
	VI_ATTR_ASRL_FLOW_CNTRL             = C.VI_ATTR_ASRL_FLOW_CNTRL
	VI_ATTR_RD_BUF_OPER_MODE            = C.VI_ATTR_RD_BUF_OPER_MODE
	VI_ATTR_RD_BUF_SIZE                 = C.VI_ATTR_RD_BUF_SIZE
	VI_ATTR_WR_BUF_OPER_MODE            = C.VI_ATTR_WR_BUF_OPER_MODE
	VI_ATTR_WR_BUF_SIZE                 = C.VI_ATTR_WR_BUF_SIZE
	VI_ATTR_SUPPRESS_END_EN             = C.VI_ATTR_SUPPRESS_END_EN
	VI_ATTR_TERMCHAR_EN                 = C.VI_ATTR_TERMCHAR_EN
	VI_ATTR_DEST_ACCESS_PRIV            = C.VI_ATTR_DEST_ACCESS_PRIV
	VI_ATTR_DEST_BYTE_ORDER             = C.VI_ATTR_DEST_BYTE_ORDER
	VI_ATTR_SRC_ACCESS_PRIV             = C.VI_ATTR_SRC_ACCESS_PRIV
	VI_ATTR_SRC_BYTE_ORDER              = C.VI_ATTR_SRC_BYTE_ORDER
	VI_ATTR_SRC_INCREMENT               = C.VI_ATTR_SRC_INCREMENT
	VI_ATTR_DEST_INCREMENT              = C.VI_ATTR_DEST_INCREMENT
	VI_ATTR_WIN_ACCESS_PRIV             = C.VI_ATTR_WIN_ACCESS_PRIV
	VI_ATTR_WIN_BYTE_ORDER              = C.VI_ATTR_WIN_BYTE_ORDER
	VI_ATTR_GPIB_ATN_STATE              = C.VI_ATTR_GPIB_ATN_STATE
	VI_ATTR_GPIB_ADDR_STATE             = C.VI_ATTR_GPIB_ADDR_STATE
	VI_ATTR_GPIB_CIC_STATE              = C.VI_ATTR_GPIB_CIC_STATE
	VI_ATTR_GPIB_NDAC_STATE             = C.VI_ATTR_GPIB_NDAC_STATE
	VI_ATTR_GPIB_SRQ_STATE              = C.VI_ATTR_GPIB_SRQ_STATE
	VI_ATTR_GPIB_SYS_CNTRL_STATE        = C.VI_ATTR_GPIB_SYS_CNTRL_STATE
	VI_ATTR_GPIB_HS488_CBL_LEN          = C.VI_ATTR_GPIB_HS488_CBL_LEN
	VI_ATTR_CMDR_LA                     = C.VI_ATTR_CMDR_LA
	VI_ATTR_VXI_DEV_CLASS               = C.VI_ATTR_VXI_DEV_CLASS
	VI_ATTR_MAINFRAME_LA                = C.VI_ATTR_MAINFRAME_LA
	VI_ATTR_MANF_NAME                   = C.VI_ATTR_MANF_NAME
	VI_ATTR_MODEL_NAME                  = C.VI_ATTR_MODEL_NAME
	VI_ATTR_VXI_VME_INTR_STATUS         = C.VI_ATTR_VXI_VME_INTR_STATUS
	VI_ATTR_VXI_TRIG_STATUS             = C.VI_ATTR_VXI_TRIG_STATUS
	VI_ATTR_VXI_VME_SYSFAIL_STATE       = C.VI_ATTR_VXI_VME_SYSFAIL_STATE
	VI_ATTR_WIN_BASE_ADDR_32            = C.VI_ATTR_WIN_BASE_ADDR_32
	VI_ATTR_WIN_SIZE_32                 = C.VI_ATTR_WIN_SIZE_32
	VI_ATTR_ASRL_AVAIL_NUM              = C.VI_ATTR_ASRL_AVAIL_NUM
	VI_ATTR_MEM_BASE_32                 = C.VI_ATTR_MEM_BASE_32
	VI_ATTR_ASRL_CTS_STATE              = C.VI_ATTR_ASRL_CTS_STATE
	VI_ATTR_ASRL_DCD_STATE              = C.VI_ATTR_ASRL_DCD_STATE
	VI_ATTR_ASRL_DSR_STATE              = C.VI_ATTR_ASRL_DSR_STATE
	VI_ATTR_ASRL_DTR_STATE              = C.VI_ATTR_ASRL_DTR_STATE
	VI_ATTR_ASRL_END_IN                 = C.VI_ATTR_ASRL_END_IN
	VI_ATTR_ASRL_END_OUT                = C.VI_ATTR_ASRL_END_OUT
	VI_ATTR_ASRL_REPLACE_CHAR           = C.VI_ATTR_ASRL_REPLACE_CHAR
	VI_ATTR_ASRL_RI_STATE               = C.VI_ATTR_ASRL_RI_STATE
	VI_ATTR_ASRL_RTS_STATE              = C.VI_ATTR_ASRL_RTS_STATE
	VI_ATTR_ASRL_XON_CHAR               = C.VI_ATTR_ASRL_XON_CHAR
	VI_ATTR_ASRL_XOFF_CHAR              = C.VI_ATTR_ASRL_XOFF_CHAR
	VI_ATTR_WIN_ACCESS                  = C.VI_ATTR_WIN_ACCESS
	VI_ATTR_RM_SESSION                  = C.VI_ATTR_RM_SESSION
	VI_ATTR_VXI_LA                      = C.VI_ATTR_VXI_LA
	VI_ATTR_MANF_ID                     = C.VI_ATTR_MANF_ID
	VI_ATTR_MEM_SIZE_32                 = C.VI_ATTR_MEM_SIZE_32
	VI_ATTR_MEM_SPACE                   = C.VI_ATTR_MEM_SPACE
	VI_ATTR_MODEL_CODE                  = C.VI_ATTR_MODEL_CODE
	VI_ATTR_SLOT                        = C.VI_ATTR_SLOT
	VI_ATTR_INTF_INST_NAME              = C.VI_ATTR_INTF_INST_NAME
	VI_ATTR_IMMEDIATE_SERV              = C.VI_ATTR_IMMEDIATE_SERV
	VI_ATTR_INTF_PARENT_NUM             = C.VI_ATTR_INTF_PARENT_NUM
	VI_ATTR_RSRC_SPEC_VERSION           = C.VI_ATTR_RSRC_SPEC_VERSION
	VI_ATTR_INTF_TYPE                   = C.VI_ATTR_INTF_TYPE
	VI_ATTR_GPIB_PRIMARY_ADDR           = C.VI_ATTR_GPIB_PRIMARY_ADDR
	VI_ATTR_GPIB_SECONDARY_ADDR         = C.VI_ATTR_GPIB_SECONDARY_ADDR
	VI_ATTR_RSRC_MANF_NAME              = C.VI_ATTR_RSRC_MANF_NAME
	VI_ATTR_RSRC_MANF_ID                = C.VI_ATTR_RSRC_MANF_ID
	VI_ATTR_INTF_NUM                    = C.VI_ATTR_INTF_NUM
	VI_ATTR_TRIG_ID                     = C.VI_ATTR_TRIG_ID
	VI_ATTR_GPIB_REN_STATE              = C.VI_ATTR_GPIB_REN_STATE
	VI_ATTR_GPIB_UNADDR_EN              = C.VI_ATTR_GPIB_UNADDR_EN
	VI_ATTR_DEV_STATUS_BYTE             = C.VI_ATTR_DEV_STATUS_BYTE
	VI_ATTR_FILE_APPEND_EN              = C.VI_ATTR_FILE_APPEND_EN
	VI_ATTR_VXI_TRIG_SUPPORT            = C.VI_ATTR_VXI_TRIG_SUPPORT
	VI_ATTR_TCPIP_ADDR                  = C.VI_ATTR_TCPIP_ADDR
	VI_ATTR_TCPIP_HOSTNAME              = C.VI_ATTR_TCPIP_HOSTNAME
	VI_ATTR_TCPIP_PORT                  = C.VI_ATTR_TCPIP_PORT
	VI_ATTR_TCPIP_DEVICE_NAME           = C.VI_ATTR_TCPIP_DEVICE_NAME
	VI_ATTR_TCPIP_NODELAY               = C.VI_ATTR_TCPIP_NODELAY
	VI_ATTR_TCPIP_KEEPALIVE             = C.VI_ATTR_TCPIP_KEEPALIVE
	VI_ATTR_4882_COMPLIANT              = C.VI_ATTR_4882_COMPLIANT
	VI_ATTR_USB_SERIAL_NUM              = C.VI_ATTR_USB_SERIAL_NUM
	VI_ATTR_USB_INTFC_NUM               = C.VI_ATTR_USB_INTFC_NUM
	VI_ATTR_USB_PROTOCOL                = C.VI_ATTR_USB_PROTOCOL
	VI_ATTR_USB_MAX_INTR_SIZE           = C.VI_ATTR_USB_MAX_INTR_SIZE
	VI_ATTR_PXI_DEV_NUM                 = C.VI_ATTR_PXI_DEV_NUM
	VI_ATTR_PXI_FUNC_NUM                = C.VI_ATTR_PXI_FUNC_NUM
	VI_ATTR_PXI_BUS_NUM                 = C.VI_ATTR_PXI_BUS_NUM
	VI_ATTR_PXI_CHASSIS                 = C.VI_ATTR_PXI_CHASSIS
	VI_ATTR_PXI_SLOTPATH                = C.VI_ATTR_PXI_SLOTPATH
	VI_ATTR_PXI_SLOT_LBUS_LEFT          = C.VI_ATTR_PXI_SLOT_LBUS_LEFT
	VI_ATTR_PXI_SLOT_LBUS_RIGHT         = C.VI_ATTR_PXI_SLOT_LBUS_RIGHT
	VI_ATTR_PXI_TRIG_BUS                = C.VI_ATTR_PXI_TRIG_BUS
	VI_ATTR_PXI_STAR_TRIG_BUS           = C.VI_ATTR_PXI_STAR_TRIG_BUS
	VI_ATTR_PXI_STAR_TRIG_LINE          = C.VI_ATTR_PXI_STAR_TRIG_LINE
	VI_ATTR_PXI_SRC_TRIG_BUS            = C.VI_ATTR_PXI_SRC_TRIG_BUS
	VI_ATTR_PXI_DEST_TRIG_BUS           = C.VI_ATTR_PXI_DEST_TRIG_BUS
	VI_ATTR_PXI_MEM_TYPE_BAR0           = C.VI_ATTR_PXI_MEM_TYPE_BAR0
	VI_ATTR_PXI_MEM_TYPE_BAR1           = C.VI_ATTR_PXI_MEM_TYPE_BAR1
	VI_ATTR_PXI_MEM_TYPE_BAR2           = C.VI_ATTR_PXI_MEM_TYPE_BAR2
	VI_ATTR_PXI_MEM_TYPE_BAR3           = C.VI_ATTR_PXI_MEM_TYPE_BAR3
	VI_ATTR_PXI_MEM_TYPE_BAR4           = C.VI_ATTR_PXI_MEM_TYPE_BAR4
	VI_ATTR_PXI_MEM_TYPE_BAR5           = C.VI_ATTR_PXI_MEM_TYPE_BAR5
	VI_ATTR_PXI_MEM_BASE_BAR0_32        = C.VI_ATTR_PXI_MEM_BASE_BAR0_32
	VI_ATTR_PXI_MEM_BASE_BAR1_32        = C.VI_ATTR_PXI_MEM_BASE_BAR1_32
	VI_ATTR_PXI_MEM_BASE_BAR2_32        = C.VI_ATTR_PXI_MEM_BASE_BAR2_32
	VI_ATTR_PXI_MEM_BASE_BAR3_32        = C.VI_ATTR_PXI_MEM_BASE_BAR3_32
	VI_ATTR_PXI_MEM_BASE_BAR4_32        = C.VI_ATTR_PXI_MEM_BASE_BAR4_32
	VI_ATTR_PXI_MEM_BASE_BAR5_32        = C.VI_ATTR_PXI_MEM_BASE_BAR5_32
	VI_ATTR_PXI_MEM_BASE_BAR0_64        = C.VI_ATTR_PXI_MEM_BASE_BAR0_64
	VI_ATTR_PXI_MEM_BASE_BAR1_64        = C.VI_ATTR_PXI_MEM_BASE_BAR1_64
	VI_ATTR_PXI_MEM_BASE_BAR2_64        = C.VI_ATTR_PXI_MEM_BASE_BAR2_64
	VI_ATTR_PXI_MEM_BASE_BAR3_64        = C.VI_ATTR_PXI_MEM_BASE_BAR3_64
	VI_ATTR_PXI_MEM_BASE_BAR4_64        = C.VI_ATTR_PXI_MEM_BASE_BAR4_64
	VI_ATTR_PXI_MEM_BASE_BAR5_64        = C.VI_ATTR_PXI_MEM_BASE_BAR5_64
	VI_ATTR_PXI_MEM_SIZE_BAR0_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR0_32
	VI_ATTR_PXI_MEM_SIZE_BAR1_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR1_32
	VI_ATTR_PXI_MEM_SIZE_BAR2_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR2_32
	VI_ATTR_PXI_MEM_SIZE_BAR3_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR3_32
	VI_ATTR_PXI_MEM_SIZE_BAR4_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR4_32
	VI_ATTR_PXI_MEM_SIZE_BAR5_32        = C.VI_ATTR_PXI_MEM_SIZE_BAR5_32
	VI_ATTR_PXI_MEM_SIZE_BAR0_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR0_64
	VI_ATTR_PXI_MEM_SIZE_BAR1_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR1_64
	VI_ATTR_PXI_MEM_SIZE_BAR2_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR2_64
	VI_ATTR_PXI_MEM_SIZE_BAR3_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR3_64
	VI_ATTR_PXI_MEM_SIZE_BAR4_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR4_64
	VI_ATTR_PXI_MEM_SIZE_BAR5_64        = C.VI_ATTR_PXI_MEM_SIZE_BAR5_64
	VI_ATTR_PXI_IS_EXPRESS              = C.VI_ATTR_PXI_IS_EXPRESS
	VI_ATTR_PXI_SLOT_LWIDTH             = C.VI_ATTR_PXI_SLOT_LWIDTH
	VI_ATTR_PXI_MAX_LWIDTH              = C.VI_ATTR_PXI_MAX_LWIDTH
	VI_ATTR_PXI_ACTUAL_LWIDTH           = C.VI_ATTR_PXI_ACTUAL_LWIDTH
	VI_ATTR_PXI_DSTAR_BUS               = C.VI_ATTR_PXI_DSTAR_BUS
	VI_ATTR_PXI_DSTAR_SET               = C.VI_ATTR_PXI_DSTAR_SET
	VI_ATTR_PXI_ALLOW_WRITE_COMBINE     = C.VI_ATTR_PXI_ALLOW_WRITE_COMBINE
	VI_ATTR_TCPIP_HISLIP_OVERLAP_EN     = C.VI_ATTR_TCPIP_HISLIP_OVERLAP_EN
	VI_ATTR_TCPIP_HISLIP_VERSION        = C.VI_ATTR_TCPIP_HISLIP_VERSION
	VI_ATTR_TCPIP_HISLIP_MAX_MESSAGE_KB = C.VI_ATTR_TCPIP_HISLIP_MAX_MESSAGE_KB
	VI_ATTR_TCPIP_IS_HISLIP             = C.VI_ATTR_TCPIP_IS_HISLIP

	VI_ATTR_JOB_ID              = C.VI_ATTR_JOB_ID
	VI_ATTR_EVENT_TYPE          = C.VI_ATTR_EVENT_TYPE
	VI_ATTR_SIGP_STATUS_ID      = C.VI_ATTR_SIGP_STATUS_ID
	VI_ATTR_RECV_TRIG_ID        = C.VI_ATTR_RECV_TRIG_ID
	VI_ATTR_INTR_STATUS_ID      = C.VI_ATTR_INTR_STATUS_ID
	VI_ATTR_STATUS              = C.VI_ATTR_STATUS
	VI_ATTR_RET_COUNT_32        = C.VI_ATTR_RET_COUNT_32
	VI_ATTR_BUFFER              = C.VI_ATTR_BUFFER
	VI_ATTR_RECV_INTR_LEVEL     = C.VI_ATTR_RECV_INTR_LEVEL
	VI_ATTR_OPER_NAME           = C.VI_ATTR_OPER_NAME
	VI_ATTR_GPIB_RECV_CIC_STATE = C.VI_ATTR_GPIB_RECV_CIC_STATE
	VI_ATTR_RECV_TCPIP_ADDR     = C.VI_ATTR_RECV_TCPIP_ADDR
	VI_ATTR_USB_RECV_INTR_SIZE  = C.VI_ATTR_USB_RECV_INTR_SIZE
	VI_ATTR_USB_RECV_INTR_DATA  = C.VI_ATTR_USB_RECV_INTR_DATA
	VI_ATTR_PXI_RECV_INTR_SEQ   = C.VI_ATTR_PXI_RECV_INTR_SEQ
	VI_ATTR_PXI_RECV_INTR_DATA  = C.VI_ATTR_PXI_RECV_INTR_DATA

	// Attributes (platform dependent size)
	VI_ATTR_USER_DATA = C.VI_ATTR_USER_DATA
	VI_ATTR_RET_COUNT = C.VI_ATTR_RET_COUNT

	VI_ATTR_WIN_BASE_ADDR     = C.VI_ATTR_WIN_BASE_ADDR
	VI_ATTR_WIN_SIZE          = C.VI_ATTR_WIN_SIZE
	VI_ATTR_MEM_BASE          = C.VI_ATTR_MEM_BASE
	VI_ATTR_MEM_SIZE          = C.VI_ATTR_MEM_SIZE
	VI_ATTR_PXI_MEM_BASE_BAR0 = C.VI_ATTR_PXI_MEM_BASE_BAR0
	VI_ATTR_PXI_MEM_BASE_BAR1 = C.VI_ATTR_PXI_MEM_BASE_BAR1
	VI_ATTR_PXI_MEM_BASE_BAR2 = C.VI_ATTR_PXI_MEM_BASE_BAR2
	VI_ATTR_PXI_MEM_BASE_BAR3 = C.VI_ATTR_PXI_MEM_BASE_BAR3
	VI_ATTR_PXI_MEM_BASE_BAR4 = C.VI_ATTR_PXI_MEM_BASE_BAR4
	VI_ATTR_PXI_MEM_BASE_BAR5 = C.VI_ATTR_PXI_MEM_BASE_BAR5
	VI_ATTR_PXI_MEM_SIZE_BAR0 = C.VI_ATTR_PXI_MEM_SIZE_BAR0
	VI_ATTR_PXI_MEM_SIZE_BAR1 = C.VI_ATTR_PXI_MEM_SIZE_BAR1
	VI_ATTR_PXI_MEM_SIZE_BAR2 = C.VI_ATTR_PXI_MEM_SIZE_BAR2
	VI_ATTR_PXI_MEM_SIZE_BAR3 = C.VI_ATTR_PXI_MEM_SIZE_BAR3
	VI_ATTR_PXI_MEM_SIZE_BAR4 = C.VI_ATTR_PXI_MEM_SIZE_BAR4
	VI_ATTR_PXI_MEM_SIZE_BAR5 = C.VI_ATTR_PXI_MEM_SIZE_BAR5

	// Event Types
	VI_EVENT_IO_COMPLETION    = C.VI_EVENT_IO_COMPLETION
	VI_EVENT_TRIG             = C.VI_EVENT_TRIG
	VI_EVENT_SERVICE_REQ      = C.VI_EVENT_SERVICE_REQ
	VI_EVENT_CLEAR            = C.VI_EVENT_CLEAR
	VI_EVENT_EXCEPTION        = C.VI_EVENT_EXCEPTION
	VI_EVENT_GPIB_CIC         = C.VI_EVENT_GPIB_CIC
	VI_EVENT_GPIB_TALK        = C.VI_EVENT_GPIB_TALK
	VI_EVENT_GPIB_LISTEN      = C.VI_EVENT_GPIB_LISTEN
	VI_EVENT_VXI_VME_SYSFAIL  = C.VI_EVENT_VXI_VME_SYSFAIL
	VI_EVENT_VXI_VME_SYSRESET = C.VI_EVENT_VXI_VME_SYSRESET
	VI_EVENT_VXI_SIGP         = C.VI_EVENT_VXI_SIGP
	VI_EVENT_VXI_VME_INTR     = C.VI_EVENT_VXI_VME_INTR
	VI_EVENT_PXI_INTR         = C.VI_EVENT_PXI_INTR
	VI_EVENT_TCPIP_CONNECT    = C.VI_EVENT_TCPIP_CONNECT
	VI_EVENT_USB_INTR         = C.VI_EVENT_USB_INTR

	VI_ALL_ENABLED_EVENTS = C.VI_ALL_ENABLED_EVENTS

	// Completion and Error Codes
	VI_SUCCESS_EVENT_EN       = C.VI_SUCCESS_EVENT_EN
	VI_SUCCESS_EVENT_DIS      = C.VI_SUCCESS_EVENT_DIS
	VI_SUCCESS_QUEUE_EMPTY    = C.VI_SUCCESS_QUEUE_EMPTY
	VI_SUCCESS_TERM_CHAR      = C.VI_SUCCESS_TERM_CHAR
	VI_SUCCESS_MAX_CNT        = C.VI_SUCCESS_MAX_CNT
	VI_SUCCESS_DEV_NPRESENT   = C.VI_SUCCESS_DEV_NPRESENT
	VI_SUCCESS_TRIG_MAPPED    = C.VI_SUCCESS_TRIG_MAPPED
	VI_SUCCESS_QUEUE_NEMPTY   = C.VI_SUCCESS_QUEUE_NEMPTY
	VI_SUCCESS_NCHAIN         = C.VI_SUCCESS_NCHAIN
	VI_SUCCESS_NESTED_SHARED  = C.VI_SUCCESS_NESTED_SHARED
	VI_SUCCESS_NESTED_EXCLUSI = C.VI_SUCCESS_NESTED_EXCLUSIVE
	VI_SUCCESS_SYNC           = C.VI_SUCCESS_SYNC

	VI_WARN_QUEUE_OVERFLOW  = C.VI_WARN_QUEUE_OVERFLOW
	VI_WARN_CONFIG_NLOADED  = C.VI_WARN_CONFIG_NLOADED
	VI_WARN_NULL_OBJECT     = C.VI_WARN_NULL_OBJECT
	VI_WARN_NSUP_ATTR_STATE = C.VI_WARN_NSUP_ATTR_STATE
	VI_WARN_UNKNOWN_STATUS  = C.VI_WARN_UNKNOWN_STATUS
	VI_WARN_NSUP_BUF        = C.VI_WARN_NSUP_BUF
	VI_WARN_EXT_FUNC_NIMPL  = C.VI_WARN_EXT_FUNC_NIMPL

	VI_ERROR_SYSTEM_ERROR     = C.VI_ERROR_SYSTEM_ERROR
	VI_ERROR_INV_OBJECT       = C.VI_ERROR_INV_OBJECT
	VI_ERROR_RSRC_LOCKED      = C.VI_ERROR_RSRC_LOCKED
	VI_ERROR_INV_EXPR         = C.VI_ERROR_INV_EXPR
	VI_ERROR_RSRC_NFOUND      = C.VI_ERROR_RSRC_NFOUND
	VI_ERROR_INV_RSRC_NAME    = C.VI_ERROR_INV_RSRC_NAME
	VI_ERROR_INV_ACC_MODE     = C.VI_ERROR_INV_ACC_MODE
	VI_ERROR_TMO              = C.VI_ERROR_TMO
	VI_ERROR_CLOSING_FAILED   = C.VI_ERROR_CLOSING_FAILED
	VI_ERROR_INV_DEGREE       = C.VI_ERROR_INV_DEGREE
	VI_ERROR_INV_JOB_ID       = C.VI_ERROR_INV_JOB_ID
	VI_ERROR_NSUP_ATTR        = C.VI_ERROR_NSUP_ATTR
	VI_ERROR_NSUP_ATTR_STATE  = C.VI_ERROR_NSUP_ATTR_STATE
	VI_ERROR_ATTR_READONLY    = C.VI_ERROR_ATTR_READONLY
	VI_ERROR_INV_LOCK_TYPE    = C.VI_ERROR_INV_LOCK_TYPE
	VI_ERROR_INV_ACCESS_KEY   = C.VI_ERROR_INV_ACCESS_KEY
	VI_ERROR_INV_EVENT        = C.VI_ERROR_INV_EVENT
	VI_ERROR_INV_MECH         = C.VI_ERROR_INV_MECH
	VI_ERROR_HNDLR_NINSTALLED = C.VI_ERROR_HNDLR_NINSTALLED
	VI_ERROR_INV_HNDLR_REF    = C.VI_ERROR_INV_HNDLR_REF
	VI_ERROR_INV_CONTEXT      = C.VI_ERROR_INV_CONTEXT
	VI_ERROR_QUEUE_OVERFLOW   = C.VI_ERROR_QUEUE_OVERFLOW
	VI_ERROR_NENABLED         = C.VI_ERROR_NENABLED
	VI_ERROR_ABORT            = C.VI_ERROR_ABORT
	VI_ERROR_RAW_WR_PROT_VIOL = C.VI_ERROR_RAW_WR_PROT_VIOL
	VI_ERROR_RAW_RD_PROT_VIOL = C.VI_ERROR_RAW_RD_PROT_VIOL
	VI_ERROR_OUTP_PROT_VIOL   = C.VI_ERROR_OUTP_PROT_VIOL
	VI_ERROR_INP_PROT_VIOL    = C.VI_ERROR_INP_PROT_VIOL
	VI_ERROR_BERR             = C.VI_ERROR_BERR
	VI_ERROR_IN_PROGRESS      = C.VI_ERROR_IN_PROGRESS
	VI_ERROR_INV_SETUP        = C.VI_ERROR_INV_SETUP
	VI_ERROR_QUEUE_ERROR      = C.VI_ERROR_QUEUE_ERROR
	VI_ERROR_ALLOC            = C.VI_ERROR_ALLOC
	VI_ERROR_INV_MASK         = C.VI_ERROR_INV_MASK
	VI_ERROR_IO               = C.VI_ERROR_IO
	VI_ERROR_INV_FMT          = C.VI_ERROR_INV_FMT
	VI_ERROR_NSUP_FMT         = C.VI_ERROR_NSUP_FMT
	VI_ERROR_LINE_IN_USE      = C.VI_ERROR_LINE_IN_USE
	VI_ERROR_LINE_NRESERVED   = C.VI_ERROR_LINE_NRESERVED
	VI_ERROR_NSUP_MODE        = C.VI_ERROR_NSUP_MODE
	VI_ERROR_SRQ_NOCCURRED    = C.VI_ERROR_SRQ_NOCCURRED
	VI_ERROR_INV_SPACE        = C.VI_ERROR_INV_SPACE
	VI_ERROR_INV_OFFSET       = C.VI_ERROR_INV_OFFSET
	VI_ERROR_INV_WIDTH        = C.VI_ERROR_INV_WIDTH
	VI_ERROR_NSUP_OFFSET      = C.VI_ERROR_NSUP_OFFSET
	VI_ERROR_NSUP_VAR_WIDTH   = C.VI_ERROR_NSUP_VAR_WIDTH
	VI_ERROR_WINDOW_NMAPPED   = C.VI_ERROR_WINDOW_NMAPPED
	VI_ERROR_RESP_PENDING     = C.VI_ERROR_RESP_PENDING
	VI_ERROR_NLISTENERS       = C.VI_ERROR_NLISTENERS
	VI_ERROR_NCIC             = C.VI_ERROR_NCIC
	VI_ERROR_NSYS_CNTLR       = C.VI_ERROR_NSYS_CNTLR
	VI_ERROR_NSUP_OPER        = C.VI_ERROR_NSUP_OPER
	VI_ERROR_INTR_PENDING     = C.VI_ERROR_INTR_PENDING
	VI_ERROR_ASRL_PARITY      = C.VI_ERROR_ASRL_PARITY
	VI_ERROR_ASRL_FRAMING     = C.VI_ERROR_ASRL_FRAMING
	VI_ERROR_ASRL_OVERRUN     = C.VI_ERROR_ASRL_OVERRUN
	VI_ERROR_TRIG_NMAPPED     = C.VI_ERROR_TRIG_NMAPPED
	VI_ERROR_NSUP_ALIGN_OFFSE = C.VI_ERROR_NSUP_ALIGN_OFFSET
	VI_ERROR_USER_BUF         = C.VI_ERROR_USER_BUF
	VI_ERROR_RSRC_BUSY        = C.VI_ERROR_RSRC_BUSY
	VI_ERROR_NSUP_WIDTH       = C.VI_ERROR_NSUP_WIDTH
	VI_ERROR_INV_PARAMETER    = C.VI_ERROR_INV_PARAMETER
	VI_ERROR_INV_PROT         = C.VI_ERROR_INV_PROT
	VI_ERROR_INV_SIZE         = C.VI_ERROR_INV_SIZE
	VI_ERROR_WINDOW_MAPPED    = C.VI_ERROR_WINDOW_MAPPED
	VI_ERROR_NIMPL_OPER       = C.VI_ERROR_NIMPL_OPER
	VI_ERROR_INV_LENGTH       = C.VI_ERROR_INV_LENGTH
	VI_ERROR_INV_MODE         = C.VI_ERROR_INV_MODE
	VI_ERROR_SESN_NLOCKED     = C.VI_ERROR_SESN_NLOCKED
	VI_ERROR_MEM_NSHARED      = C.VI_ERROR_MEM_NSHARED
	VI_ERROR_LIBRARY_NFOUND   = C.VI_ERROR_LIBRARY_NFOUND
	VI_ERROR_NSUP_INTR        = C.VI_ERROR_NSUP_INTR
	VI_ERROR_INV_LINE         = C.VI_ERROR_INV_LINE
	VI_ERROR_FILE_ACCESS      = C.VI_ERROR_FILE_ACCESS
	VI_ERROR_FILE_IO          = C.VI_ERROR_FILE_IO
	VI_ERROR_NSUP_LINE        = C.VI_ERROR_NSUP_LINE
	VI_ERROR_NSUP_MECH        = C.VI_ERROR_NSUP_MECH
	VI_ERROR_INTF_NUM_NCONFIG = C.VI_ERROR_INTF_NUM_NCONFIG
	VI_ERROR_CONN_LOST        = C.VI_ERROR_CONN_LOST
	VI_ERROR_MACHINE_NAVAIL   = C.VI_ERROR_MACHINE_NAVAIL
	VI_ERROR_NPERMISSION      = C.VI_ERROR_NPERMISSION

	// Other VISA Definitions
	VI_FIND_BUFLEN = C.VI_FIND_BUFLEN

	VI_INTF_GPIB     = C.VI_INTF_GPIB
	VI_INTF_VXI      = C.VI_INTF_VXI
	VI_INTF_GPIB_VXI = C.VI_INTF_GPIB_VXI
	VI_INTF_ASRL     = C.VI_INTF_ASRL
	VI_INTF_PXI      = C.VI_INTF_PXI
	VI_INTF_TCPIP    = C.VI_INTF_TCPIP
	VI_INTF_USB      = C.VI_INTF_USB

	VI_PROT_NORMAL        = C.VI_PROT_NORMAL
	VI_PROT_FDC           = C.VI_PROT_FDC
	VI_PROT_HS488         = C.VI_PROT_HS488
	VI_PROT_4882_STRS     = C.VI_PROT_4882_STRS
	VI_PROT_USBTMC_VENDOR = C.VI_PROT_USBTMC_VENDOR

	VI_FDC_NORMAL = C.VI_FDC_NORMAL
	VI_FDC_STREAM = C.VI_FDC_STREAM

	VI_LOCAL_SPACE     = C.VI_LOCAL_SPACE
	VI_A16_SPACE       = C.VI_A16_SPACE
	VI_A24_SPACE       = C.VI_A24_SPACE
	VI_A32_SPACE       = C.VI_A32_SPACE
	VI_A64_SPACE       = C.VI_A64_SPACE
	VI_PXI_ALLOC_SPACE = C.VI_PXI_ALLOC_SPACE
	VI_PXI_CFG_SPACE   = C.VI_PXI_CFG_SPACE
	VI_PXI_BAR0_SPACE  = C.VI_PXI_BAR0_SPACE
	VI_PXI_BAR1_SPACE  = C.VI_PXI_BAR1_SPACE
	VI_PXI_BAR2_SPACE  = C.VI_PXI_BAR2_SPACE
	VI_PXI_BAR3_SPACE  = C.VI_PXI_BAR3_SPACE
	VI_PXI_BAR4_SPACE  = C.VI_PXI_BAR4_SPACE
	VI_PXI_BAR5_SPACE  = C.VI_PXI_BAR5_SPACE
	VI_OPAQUE_SPACE    = C.VI_OPAQUE_SPACE

	VI_UNKNOWN_LA      = C.VI_UNKNOWN_LA
	VI_UNKNOWN_SLOT    = C.VI_UNKNOWN_SLOT
	VI_UNKNOWN_LEVEL   = C.VI_UNKNOWN_LEVEL
	VI_UNKNOWN_CHASSIS = C.VI_UNKNOWN_CHASSIS

	VI_QUEUE         = C.VI_QUEUE
	VI_HNDLR         = C.VI_HNDLR
	VI_SUSPEND_HNDLR = C.VI_SUSPEND_HNDLR
	VI_ALL_MECH      = C.VI_ALL_MECH

	VI_ANY_HNDLR = C.VI_ANY_HNDLR

	VI_TRIG_ALL         = C.VI_TRIG_ALL
	VI_TRIG_SW          = C.VI_TRIG_SW
	VI_TRIG_TTL0        = C.VI_TRIG_TTL0
	VI_TRIG_TTL1        = C.VI_TRIG_TTL1
	VI_TRIG_TTL2        = C.VI_TRIG_TTL2
	VI_TRIG_TTL3        = C.VI_TRIG_TTL3
	VI_TRIG_TTL4        = C.VI_TRIG_TTL4
	VI_TRIG_TTL5        = C.VI_TRIG_TTL5
	VI_TRIG_TTL6        = C.VI_TRIG_TTL6
	VI_TRIG_TTL7        = C.VI_TRIG_TTL7
	VI_TRIG_ECL0        = C.VI_TRIG_ECL0
	VI_TRIG_ECL1        = C.VI_TRIG_ECL1
	VI_TRIG_ECL2        = C.VI_TRIG_ECL2
	VI_TRIG_ECL3        = C.VI_TRIG_ECL3
	VI_TRIG_ECL4        = C.VI_TRIG_ECL4
	VI_TRIG_ECL5        = C.VI_TRIG_ECL5
	VI_TRIG_STAR_SLOT1  = C.VI_TRIG_STAR_SLOT1
	VI_TRIG_STAR_SLOT2  = C.VI_TRIG_STAR_SLOT2
	VI_TRIG_STAR_SLOT3  = C.VI_TRIG_STAR_SLOT3
	VI_TRIG_STAR_SLOT4  = C.VI_TRIG_STAR_SLOT4
	VI_TRIG_STAR_SLOT5  = C.VI_TRIG_STAR_SLOT5
	VI_TRIG_STAR_SLOT6  = C.VI_TRIG_STAR_SLOT6
	VI_TRIG_STAR_SLOT7  = C.VI_TRIG_STAR_SLOT7
	VI_TRIG_STAR_SLOT8  = C.VI_TRIG_STAR_SLOT8
	VI_TRIG_STAR_SLOT9  = C.VI_TRIG_STAR_SLOT9
	VI_TRIG_STAR_SLOT10 = C.VI_TRIG_STAR_SLOT10
	VI_TRIG_STAR_SLOT11 = C.VI_TRIG_STAR_SLOT11
	VI_TRIG_STAR_SLOT12 = C.VI_TRIG_STAR_SLOT12
	VI_TRIG_STAR_INSTR  = C.VI_TRIG_STAR_INSTR
	VI_TRIG_PANEL_IN    = C.VI_TRIG_PANEL_IN
	VI_TRIG_PANEL_OUT   = C.VI_TRIG_PANEL_OUT
	VI_TRIG_STAR_VXI0   = C.VI_TRIG_STAR_VXI0
	VI_TRIG_STAR_VXI1   = C.VI_TRIG_STAR_VXI1
	VI_TRIG_STAR_VXI2   = C.VI_TRIG_STAR_VXI2

	VI_TRIG_PROT_DEFAULT   = C.VI_TRIG_PROT_DEFAULT
	VI_TRIG_PROT_ON        = C.VI_TRIG_PROT_ON
	VI_TRIG_PROT_OFF       = C.VI_TRIG_PROT_OFF
	VI_TRIG_PROT_SYNC      = C.VI_TRIG_PROT_SYNC
	VI_TRIG_PROT_RESERVE   = C.VI_TRIG_PROT_RESERVE
	VI_TRIG_PROT_UNRESERVE = C.VI_TRIG_PROT_UNRESERVE

	VI_READ_BUF           = C.VI_READ_BUF
	VI_WRITE_BUF          = C.VI_WRITE_BUF
	VI_READ_BUF_DISCARD   = C.VI_READ_BUF_DISCARD
	VI_WRITE_BUF_DISCARD  = C.VI_WRITE_BUF_DISCARD
	VI_IO_IN_BUF          = C.VI_IO_IN_BUF
	VI_IO_OUT_BUF         = C.VI_IO_OUT_BUF
	VI_IO_IN_BUF_DISCARD  = C.VI_IO_IN_BUF_DISCARD
	VI_IO_OUT_BUF_DISCARD = C.VI_IO_OUT_BUF_DISCARD

	VI_FLUSH_ON_ACCESS = C.VI_FLUSH_ON_ACCESS
	VI_FLUSH_WHEN_FULL = C.VI_FLUSH_WHEN_FULL
	VI_FLUSH_DISABLE   = C.VI_FLUSH_DISABLE

	VI_NMAPPED              = C.VI_NMAPPED
	VI_USE_OPERS            = C.VI_USE_OPERS
	VI_DEREF_ADDR           = C.VI_DEREF_ADDR
	VI_DEREF_ADDR_BYTE_SWAP = C.VI_DEREF_ADDR_BYTE_SWAP

	VI_TMO_IMMEDIATE = C.VI_TMO_IMMEDIATE
	VI_TMO_INFINITE  = C.VI_TMO_INFINITE

	VI_NO_LOCK        = C.VI_NO_LOCK
	VI_EXCLUSIVE_LOCK = C.VI_EXCLUSIVE_LOCK
	VI_SHARED_LOCK    = C.VI_SHARED_LOCK
	VI_LOAD_CONFIG    = C.VI_LOAD_CONFIG

	VI_NO_SEC_ADDR = C.VI_NO_SEC_ADDR

	VI_ASRL_PAR_NONE  = C.VI_ASRL_PAR_NONE
	VI_ASRL_PAR_ODD   = C.VI_ASRL_PAR_ODD
	VI_ASRL_PAR_EVEN  = C.VI_ASRL_PAR_EVEN
	VI_ASRL_PAR_MARK  = C.VI_ASRL_PAR_MARK
	VI_ASRL_PAR_SPACE = C.VI_ASRL_PAR_SPACE

	VI_ASRL_STOP_ONE  = C.VI_ASRL_STOP_ONE
	VI_ASRL_STOP_ONE5 = C.VI_ASRL_STOP_ONE5
	VI_ASRL_STOP_TWO  = C.VI_ASRL_STOP_TWO

	VI_ASRL_FLOW_NONE     = C.VI_ASRL_FLOW_NONE
	VI_ASRL_FLOW_XON_XOFF = C.VI_ASRL_FLOW_XON_XOFF
	VI_ASRL_FLOW_RTS_CTS  = C.VI_ASRL_FLOW_RTS_CTS
	VI_ASRL_FLOW_DTR_DSR  = C.VI_ASRL_FLOW_DTR_DSR

	VI_ASRL_END_NONE     = C.VI_ASRL_END_NONE
	VI_ASRL_END_LAST_BIT = C.VI_ASRL_END_LAST_BIT
	VI_ASRL_END_TERMCHAR = C.VI_ASRL_END_TERMCHAR
	VI_ASRL_END_BREAK    = C.VI_ASRL_END_BREAK

	VI_STATE_ASSERTED   = C.VI_STATE_ASSERTED
	VI_STATE_UNASSERTED = C.VI_STATE_UNASSERTED
	VI_STATE_UNKNOWN    = C.VI_STATE_UNKNOWN

	VI_BIG_ENDIAN    = C.VI_BIG_ENDIAN
	VI_LITTLE_ENDIAN = C.VI_LITTLE_ENDIAN

	VI_DATA_PRIV  = C.VI_DATA_PRIV
	VI_DATA_NPRIV = C.VI_DATA_NPRIV
	VI_PROG_PRIV  = C.VI_PROG_PRIV
	VI_PROG_NPRIV = C.VI_PROG_NPRIV
	VI_BLCK_PRIV  = C.VI_BLCK_PRIV
	VI_BLCK_NPRIV = C.VI_BLCK_NPRIV
	VI_D64_PRIV   = C.VI_D64_PRIV
	VI_D64_NPRIV  = C.VI_D64_NPRIV
	VI_D64_2EVME  = C.VI_D64_2EVME
	VI_D64_SST160 = C.VI_D64_SST160
	VI_D64_SST267 = C.VI_D64_SST267
	VI_D64_SST320 = C.VI_D64_SST320

	VI_WIDTH_8  = C.VI_WIDTH_8
	VI_WIDTH_16 = C.VI_WIDTH_16
	VI_WIDTH_32 = C.VI_WIDTH_32
	VI_WIDTH_64 = C.VI_WIDTH_64

	VI_GPIB_REN_DEASSERT        = C.VI_GPIB_REN_DEASSERT
	VI_GPIB_REN_ASSERT          = C.VI_GPIB_REN_ASSERT
	VI_GPIB_REN_DEASSERT_GTL    = C.VI_GPIB_REN_DEASSERT_GTL
	VI_GPIB_REN_ASSERT_ADDRESS  = C.VI_GPIB_REN_ASSERT_ADDRESS
	VI_GPIB_REN_ASSERT_LLO      = C.VI_GPIB_REN_ASSERT_LLO
	VI_GPIB_REN_ASSERT_ADDRESS_ = C.VI_GPIB_REN_ASSERT_ADDRESS_LLO
	VI_GPIB_REN_ADDRESS_GTL     = C.VI_GPIB_REN_ADDRESS_GTL

	VI_GPIB_ATN_DEASSERT        = C.VI_GPIB_ATN_DEASSERT
	VI_GPIB_ATN_ASSERT          = C.VI_GPIB_ATN_ASSERT
	VI_GPIB_ATN_DEASSERT_HANDSH = C.VI_GPIB_ATN_DEASSERT_HANDSHAKE
	VI_GPIB_ATN_ASSERT_IMMEDIAT = C.VI_GPIB_ATN_ASSERT_IMMEDIATE

	VI_GPIB_HS488_DISABLED = C.VI_GPIB_HS488_DISABLED
	VI_GPIB_HS488_NIMPL    = C.VI_GPIB_HS488_NIMPL

	VI_GPIB_UNADDRESSED = C.VI_GPIB_UNADDRESSED
	VI_GPIB_TALKER      = C.VI_GPIB_TALKER
	VI_GPIB_LISTENER    = C.VI_GPIB_LISTENER

	VI_VXI_CMD16        = C.VI_VXI_CMD16
	VI_VXI_CMD16_RESP16 = C.VI_VXI_CMD16_RESP16
	VI_VXI_RESP16       = C.VI_VXI_RESP16
	VI_VXI_CMD32        = C.VI_VXI_CMD32
	VI_VXI_CMD32_RESP16 = C.VI_VXI_CMD32_RESP16
	VI_VXI_CMD32_RESP32 = C.VI_VXI_CMD32_RESP32
	VI_VXI_RESP32       = C.VI_VXI_RESP32

	VI_ASSERT_SIGNAL       = C.VI_ASSERT_SIGNAL
	VI_ASSERT_USE_ASSIGNED = C.VI_ASSERT_USE_ASSIGNED
	VI_ASSERT_IRQ1         = C.VI_ASSERT_IRQ1
	VI_ASSERT_IRQ2         = C.VI_ASSERT_IRQ2
	VI_ASSERT_IRQ3         = C.VI_ASSERT_IRQ3
	VI_ASSERT_IRQ4         = C.VI_ASSERT_IRQ4
	VI_ASSERT_IRQ5         = C.VI_ASSERT_IRQ5
	VI_ASSERT_IRQ6         = C.VI_ASSERT_IRQ6
	VI_ASSERT_IRQ7         = C.VI_ASSERT_IRQ7

	VI_UTIL_ASSERT_SYSRESET  = C.VI_UTIL_ASSERT_SYSRESET
	VI_UTIL_ASSERT_SYSFAIL   = C.VI_UTIL_ASSERT_SYSFAIL
	VI_UTIL_DEASSERT_SYSFAIL = C.VI_UTIL_DEASSERT_SYSFAIL

	VI_VXI_CLASS_MEMORY   = C.VI_VXI_CLASS_MEMORY
	VI_VXI_CLASS_EXTENDED = C.VI_VXI_CLASS_EXTENDED
	VI_VXI_CLASS_MESSAGE  = C.VI_VXI_CLASS_MESSAGE
	VI_VXI_CLASS_REGISTER = C.VI_VXI_CLASS_REGISTER
	VI_VXI_CLASS_OTHER    = C.VI_VXI_CLASS_OTHER

	VI_PXI_ADDR_NONE = C.VI_PXI_ADDR_NONE
	VI_PXI_ADDR_MEM  = C.VI_PXI_ADDR_MEM
	VI_PXI_ADDR_IO   = C.VI_PXI_ADDR_IO
	VI_PXI_ADDR_CFG  = C.VI_PXI_ADDR_CFG

	VI_TRIG_UNKNOWN = C.VI_TRIG_UNKNOWN

	VI_PXI_LBUS_UNKNOWN         = C.VI_PXI_LBUS_UNKNOWN
	VI_PXI_LBUS_NONE            = C.VI_PXI_LBUS_NONE
	VI_PXI_LBUS_STAR_TRIG_BUS_0 = C.VI_PXI_LBUS_STAR_TRIG_BUS_0
	VI_PXI_LBUS_STAR_TRIG_BUS_1 = C.VI_PXI_LBUS_STAR_TRIG_BUS_1
	VI_PXI_LBUS_STAR_TRIG_BUS_2 = C.VI_PXI_LBUS_STAR_TRIG_BUS_2
	VI_PXI_LBUS_STAR_TRIG_BUS_3 = C.VI_PXI_LBUS_STAR_TRIG_BUS_3
	VI_PXI_LBUS_STAR_TRIG_BUS_4 = C.VI_PXI_LBUS_STAR_TRIG_BUS_4
	VI_PXI_LBUS_STAR_TRIG_BUS_5 = C.VI_PXI_LBUS_STAR_TRIG_BUS_5
	VI_PXI_LBUS_STAR_TRIG_BUS_6 = C.VI_PXI_LBUS_STAR_TRIG_BUS_6
	VI_PXI_LBUS_STAR_TRIG_BUS_7 = C.VI_PXI_LBUS_STAR_TRIG_BUS_7
	VI_PXI_LBUS_STAR_TRIG_BUS_8 = C.VI_PXI_LBUS_STAR_TRIG_BUS_8
	VI_PXI_LBUS_STAR_TRIG_BUS_9 = C.VI_PXI_LBUS_STAR_TRIG_BUS_9
	VI_PXI_STAR_TRIG_CONTROLLER = C.VI_PXI_STAR_TRIG_CONTROLLER

	// Backward Compatibility Macros
	VI_ERROR_INV_SESSION    = C.VI_ERROR_INV_SESSION
	VI_INFINITE             = C.VI_INFINITE
	VI_NORMAL               = C.VI_NORMAL
	VI_FDC                  = C.VI_FDC
	VI_HS488                = C.VI_HS488
	VI_ASRL488              = C.VI_ASRL488
	VI_ASRL_IN_BUF          = C.VI_ASRL_IN_BUF
	VI_ASRL_OUT_BUF         = C.VI_ASRL_OUT_BUF
	VI_ASRL_IN_BUF_DISCARD  = C.VI_ASRL_IN_BUF_DISCARD
	VI_ASRL_OUT_BUF_DISCARD = C.VI_ASRL_OUT_BUF_DISCARD

	// National Instruments
	VI_INTF_RIO      = C.VI_INTF_RIO
	VI_INTF_FIREWIRE = C.VI_INTF_FIREWIRE

	VI_ATTR_SYNC_MXI_ALLOW_EN = C.VI_ATTR_SYNC_MXI_ALLOW_EN

	// This is for VXI SERVANT resources
	VI_EVENT_VXI_DEV_CMD      = C.VI_EVENT_VXI_DEV_CMD
	VI_ATTR_VXI_DEV_CMD_TYPE  = C.VI_ATTR_VXI_DEV_CMD_TYPE
	VI_ATTR_VXI_DEV_CMD_VALUE = C.VI_ATTR_VXI_DEV_CMD_VALUE

	VI_VXI_DEV_CMD_TYPE_16 = C.VI_VXI_DEV_CMD_TYPE_16
	VI_VXI_DEV_CMD_TYPE_32 = C.VI_VXI_DEV_CMD_TYPE_32

	// mode values include VI_VXI_RESP16, VI_VXI_RESP32, and the next 2 values
	VI_VXI_RESP_NONE       = C.VI_VXI_RESP_NONE
	VI_VXI_RESP_PROT_ERROR = C.VI_VXI_RESP_PROT_ERROR

	// This is for VXI TTL Trigger routing
	VI_ATTR_VXI_TRIG_LINES_EN = C.VI_ATTR_VXI_TRIG_LINES_EN
	VI_ATTR_VXI_TRIG_DIR      = C.VI_ATTR_VXI_TRIG_DIR

	// This allows extended Serial support on Win32 and on NI ENET Serial products
	VI_ATTR_ASRL_DISCARD_NULL  = C.VI_ATTR_ASRL_DISCARD_NULL
	VI_ATTR_ASRL_CONNECTED     = C.VI_ATTR_ASRL_CONNECTED
	VI_ATTR_ASRL_BREAK_STATE   = C.VI_ATTR_ASRL_BREAK_STATE
	VI_ATTR_ASRL_BREAK_LEN     = C.VI_ATTR_ASRL_BREAK_LEN
	VI_ATTR_ASRL_ALLOW_TRANSMI = C.VI_ATTR_ASRL_ALLOW_TRANSMIT
	VI_ATTR_ASRL_WIRE_MODE     = C.VI_ATTR_ASRL_WIRE_MODE

	VI_ASRL_WIRE_485_4         = C.VI_ASRL_WIRE_485_4
	VI_ASRL_WIRE_485_2_DTR_ECH = C.VI_ASRL_WIRE_485_2_DTR_ECHO
	VI_ASRL_WIRE_485_2_DTR_CTR = C.VI_ASRL_WIRE_485_2_DTR_CTRL
	VI_ASRL_WIRE_485_2_AUTO    = C.VI_ASRL_WIRE_485_2_AUTO
	VI_ASRL_WIRE_232_DTE       = C.VI_ASRL_WIRE_232_DTE
	VI_ASRL_WIRE_232_DCE       = C.VI_ASRL_WIRE_232_DCE
	VI_ASRL_WIRE_232_AUTO      = C.VI_ASRL_WIRE_232_AUTO

	VI_EVENT_ASRL_BREAK    = C.VI_EVENT_ASRL_BREAK
	VI_EVENT_ASRL_CTS      = C.VI_EVENT_ASRL_CTS
	VI_EVENT_ASRL_DSR      = C.VI_EVENT_ASRL_DSR
	VI_EVENT_ASRL_DCD      = C.VI_EVENT_ASRL_DCD
	VI_EVENT_ASRL_RI       = C.VI_EVENT_ASRL_RI
	VI_EVENT_ASRL_CHAR     = C.VI_EVENT_ASRL_CHAR
	VI_EVENT_ASRL_TERMCHAR = C.VI_EVENT_ASRL_TERMCHAR

// This is for fast viPeek/viPoke macros

// #if defined(NIVISA_PEEKPOKE)
// 	 #if defined(NIVISA_PEEKPOKE_SUPP)
// 	  #undef NIVISA_PEEKPOKE_SUPP
// 	 #endif

// 	#if (defined(WIN32) || defined(_WIN32) || defined(__WIN32__) || defined(__NT__)) && !defined(_NI_mswin16_)
// 	 /* This macro is supported for all Win32 compilers, including CVI. */
// 	 #define NIVISA_PEEKPOKE_SUPP
// 	#elif (defined(_WINDOWS) || defined(_Windows)) && !defined(_CVI_) && !defined(_NI_mswin16_)
// 	 /* This macro is supported for Borland and Microsoft compilers on Win16, but not CVI. */
// 	 #define NIVISA_PEEKPOKE_SUPP
// 	#elif defined(_CVI_) && defined(_NI_sparc_)
// 	 /* This macro is supported for Solaris 1 and 2, from CVI only. */
// 	 #define NIVISA_PEEKPOKE_SUPP
// 	#else
// 	 /* This macro is not supported on other platforms. */
// 	#endif

// 	#if defined(NIVISA_PEEKPOKE_SUPP)

// 		extern ViBoolean NI_viImplVISA1;
// 		ViStatus _VI_FUNC NI_viOpenDefaultRM (ViPSession vi);
// 		#define viOpenDefaultRM(vi) NI_viOpenDefaultRM(vi)

// 		#define viPeek8(vi,addr,val)                                                \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((ViPUInt8)(val)) = *((volatile ViUInt8 _VI_PTR)(addr)));     \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPeek8)((vi),(addr),(val));                                      \
// 		      }                                                                     \
// 		   }

// 		#define viPoke8(vi,addr,val)                                                \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((volatile ViUInt8 _VI_PTR)(addr)) = ((ViUInt8)(val)));       \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPoke8)((vi),(addr),(val));                                      \
// 		      }                                                                     \
// 		   }

// 		#define viPeek16(vi,addr,val)                                               \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((ViPUInt16)(val)) = *((volatile ViUInt16 _VI_PTR)(addr)));   \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPeek16)((vi),(addr),(val));                                     \
// 		      }                                                                     \
// 		   }

// 		#define viPoke16(vi,addr,val)                                               \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((volatile ViUInt16 _VI_PTR)(addr)) = ((ViUInt16)(val)));     \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPoke16)((vi),(addr),(val));                                     \
// 		      }                                                                     \
// 		   }

// 		#define viPeek32(vi,addr,val)                                               \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((ViPUInt32)(val)) = *((volatile ViUInt32 _VI_PTR)(addr)));   \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPeek32)((vi),(addr),(val));                                     \
// 		      }                                                                     \
// 		   }

// 		#define viPoke32(vi,addr,val)                                               \
// 		   {                                                                        \
// 		      if ((NI_viImplVISA1) && (*((ViPUInt32)(vi))))                         \
// 		      {                                                                     \
// 		         do (*((volatile ViUInt32 _VI_PTR)(addr)) = ((ViUInt32)(val)));     \
// 		         while (**((volatile ViUInt8 _VI_PTR _VI_PTR)(vi)) & 0x10);         \
// 		      }                                                                     \
// 		      else                                                                  \
// 		      {                                                                     \
// 		         (viPoke32)((vi),(addr),(val));                                     \
// 		      }                                                                     \
// 		   }

// 	#endif
// #endif

// #if defined(NIVISA_PXI) || defined(PXISAVISA_PXI)
// 	VI_ATTR_PXI_USE_PREALLOC_POOL = C.VI_ATTR_PXI_USE_PREALLOC_POOL
// #endif

// #if defined(NIVISA_USB)
// 	VI_ATTR_USB_BULK_OUT_PIPE = C.VI_ATTR_USB_BULK_OUT_PIPE
// 	VI_ATTR_USB_BULK_IN_PIPE = C.VI_ATTR_USB_BULK_IN_PIPE
// 	VI_ATTR_USB_INTR_IN_PIPE = C.VI_ATTR_USB_INTR_IN_PIPE
// 	VI_ATTR_USB_CLASS = C.VI_ATTR_USB_CLASS
// 	VI_ATTR_USB_SUBCLASS = C.VI_ATTR_USB_SUBCLASS
// 	VI_ATTR_USB_ALT_SETTING = C.VI_ATTR_USB_ALT_SETTING
// 	VI_ATTR_USB_END_IN = C.VI_ATTR_USB_END_IN
// 	VI_ATTR_USB_NUM_INTFCS = C.VI_ATTR_USB_NUM_INTFCS
// 	VI_ATTR_USB_NUM_PIPES = C.VI_ATTR_USB_NUM_PIPES
// 	VI_ATTR_USB_BULK_OUT_STATUS = C.VI_ATTR_USB_BULK_OUT_STATUS
// 	VI_ATTR_USB_BULK_IN_STATUS = C.VI_ATTR_USB_BULK_IN_STATUS
// 	VI_ATTR_USB_INTR_IN_STATUS = C.VI_ATTR_USB_INTR_IN_STATUS
// 	VI_ATTR_USB_CTRL_PIPE = C.VI_ATTR_USB_CTRL_PIPE

// 	VI_USB_PIPE_STATE_UNKNOWN = C.VI_USB_PIPE_STATE_UNKNOWN
// 	VI_USB_PIPE_READY = C.VI_USB_PIPE_READY
// 	VI_USB_PIPE_STALLED = C.VI_USB_PIPE_STALLED

// 	VI_USB_END_NONE = C.VI_USB_END_NONE
// 	VI_USB_END_SHORT = C.VI_USB_END_SHORT
// 	VI_USB_END_SHORT_OR_COUNT = C.VI_USB_END_SHORT_OR_COUNT
// #endif

// 	VI_ATTR_FIREWIRE_DEST_UPPER_OFFSET = C.VI_ATTR_FIREWIRE_DEST_UPPER_OFFSET
// 	VI_ATTR_FIREWIRE_SRC_UPPER_OFFSET = C.VI_ATTR_FIREWIRE_SRC_UPPER_OFFSET
// 	VI_ATTR_FIREWIRE_WIN_UPPER_OFFSET = C.VI_ATTR_FIREWIRE_WIN_UPPER_OFFSET
// 	VI_ATTR_FIREWIRE_VENDOR_ID = C.VI_ATTR_FIREWIRE_VENDOR_ID
// 	VI_ATTR_FIREWIRE_LOWER_CHIP_ID = C.VI_ATTR_FIREWIRE_LOWER_CHIP_ID
// 	VI_ATTR_FIREWIRE_UPPER_CHIP_ID = C.VI_ATTR_FIREWIRE_UPPER_CHIP_ID

// 	VI_FIREWIRE_DFLT_SPACE = C.VI_FIREWIRE_DFLT_SPACE
)