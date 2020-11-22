package itlssp

import (
	"fmt"
)

const (
	xSTX = 0x7F
)

type Channel struct {
	Value     int
	Level     int
	Channel   byte
	Recycling bool
	Currency  []byte
}

func (this *Channel) String() string {
	return fmt.Sprintf(`{"Value":%d,"Level":%d,"Channel":%d,"Recycling":%v,"Currency":%s}`,
		this.Value, this.Level, this.Channel, this.Recycling, string(this.Currency))
}

type SSPResponse int

const (
	SspResponseOk              SSPResponse = 0xF0
	SspResponseCommandUnknown  SSPResponse = 0xF2
	SspResponseWrongParams     SSPResponse = 0xF3
	SspResponseParamOutOfRange SSPResponse = 0xF4
	SspResponseCannotProcess   SSPResponse = 0xF5
	SspResponseSoftwareError   SSPResponse = 0xF6
	SspResponseFail            SSPResponse = 0xF8
	SspResponseKeyNotSet       SSPResponse = 0xFA
)

func (code SSPResponse) String() string {
	switch code {
	case SspResponseOk:
		return "Success"
	case SspResponseCannotProcess:
		return "CANNOT PROCESS COMMAND"
	case SspResponseFail:
		return "Command response is FAIL"
	case SspResponseKeyNotSet:
		return "Command response is KEY NOT SET, renegotiate keys"
	case SspResponseParamOutOfRange:
		return "Command response is PARAM OUT OF RANGE"
	case SspResponseSoftwareError:
		return "Command response is SOFTWARE ERROR"
	case SspResponseCommandUnknown:
		return "Command response is UNKNOWN COMMAND"
	case SspResponseWrongParams:
		return "Command response is WRONG PARAMETERS"
	default:
		return "UNKNOWN command"
	}
}

type SspCommand byte

const (
	// generics command
	SspCmdReset               SspCommand = 0x01
	SspCmdSetupRequest        SspCommand = 0x05
	SspCmdHostProtocolVersion SspCommand = 0x06
	SspCmdPoll                SspCommand = 0x07
	SspCmdDisable             SspCommand = 0x09
	SspCmdEnable              SspCommand = 0x0A
	SspCmdGetSerialNumber     SspCommand = 0x0C
	SspCmdSync                SspCommand = 0x11
	SspCmdFirmwareVersion     SspCommand = 0x20
	SspCmdGetDatasetVersion   SspCommand = 0x21
	SspCmdGetBuildRevision    SspCommand = 0x4F
	SspCmdPollWithAck         SspCommand = 0x56
	SspCmdEventAck            SspCommand = 0x57
	SspCmdGetCounter          SspCommand = 0x58
	SspCmdResetCounter        SspCommand = 0x59
	SspCmdSetGenerator        SspCommand = 0x4A
	SspCmdSetModulus          SspCommand = 0x4B
	SspCmdRequestKeyExchange  SspCommand = 0x4C
	// validator command
	SspCmdSetInhibits           SspCommand = 0x02
	SspCmdDisplayOn             SspCommand = 0x03
	SspCmdDisplayOff            SspCommand = 0x04
	SspCmdRejectNote            SspCommand = 0x08
	SspCmdUnitData              SspCommand = 0x0D
	SspCmdChannelValueRequest   SspCommand = 0x0E
	SspCmdChannelSecurityData   SspCommand = 0x14
	SspCmdLastRejectCode        SspCommand = 0x17
	SspCmdHold                  SspCommand = 0x18
	SspCmdGetNotePositions      SspCommand = 0x41
	SspCmdPayoutLastNote        SspCommand = 0x42
	SspCmdStackLastNote         SspCommand = 0x43
	SspCmdSetValueReportingType SspCommand = 0x45
	// payout devices
	SspCmdHaltPayout           SspCommand = 0x38
	SspCmdSetDenominationRoute SspCommand = 0x3B
	SspCmdGetDenominationRoute SspCommand = 0x3C
	SspCmdEmptyAll             SspCommand = 0x3F
	SspCmdSmartEmpty           SspCommand = 0x52
	SspCmdEnablePayout         SspCommand = 0x5C
	SspCmdDisablePayout        SspCommand = 0x5B
)

func (code SspCommand) String() string {
	switch code {
	case SspCmdReset:
		return "RESET COMMAND"
	case SspCmdHostProtocolVersion:
		return "HOST PROTOCOL VERSION"
	case SspCmdPoll:
		return "POLL COMMAND"
	case SspCmdGetSerialNumber:
		return "GET SERIAL NUMBER"
	case SspCmdSync:
		return "SYNC COMMAND"
	case SspCmdEnable:
		return "ENABLE COMMAND"
	case SspCmdDisable:
		return "DISABLE COMMAND"
	case SspCmdFirmwareVersion:
		return "GET FIRMWARE VERSION"
	case SspCmdGetDatasetVersion:
		return "GET DATASET VERSION"
	case SspCmdGetBuildRevision:
		return "GET BUILD REVISION"
	case SspCmdPollWithAck:
		return "POLL WITH ACK"
	case SspCmdEventAck:
		return "EVENT ACK"
	case SspCmdResetCounter:
		return "RESET COUNTER"
	case SspCmdGetCounter:
		return "GET COUNTER"
	case SspCmdSetInhibits:
		return "SET INHIBITS COMMAND"
	case SspCmdSetupRequest:
		return "SETUP REQUEST COMMAND"
	case SspCmdDisplayOn:
		return "DISPLAY ON COMMAND"
	case SspCmdDisplayOff:
		return "DISPLAY OFF COMMAND"
	case SspCmdRejectNote:
		return "REJECT NOTE"
	case SspCmdUnitData:
		return "UNIT DATA"
	case SspCmdChannelValueRequest:
		return "CHANNEL VALUE REQUEST"
	case SspCmdChannelSecurityData:
		return "CHANNEL SECURITY"
	case SspCmdLastRejectCode:
		return "LAST REJECT CODE"
	case SspCmdHold:
		return "HOLD"
	case SspCmdEnablePayout:
		return "ENABLE PAYOUT COMMAND"
	case SspCmdDisablePayout:
		return "DISABLE PAYOUT COMMAND"
	case SspCmdSetValueReportingType:
		return "SET VALUE REPORTING TYPE COMMAND"
	case SspCmdHaltPayout:
		return "HALT PAYOUT"
	case SspCmdSetDenominationRoute:
		return "SET DENOMINATION ROUTE"
	case SspCmdGetDenominationRoute:
		return "GET DENOMINATION ROUTE"
	case SspCmdEmptyAll:
		return "EMPTY ALL"
	case SspCmdSmartEmpty:
		return "SMART EMPTY"
	case SspCmdPayoutLastNote:
		return "PAYOUT LAST NOTE COMMAND"
	case SspCmdGetNotePositions:
		return "GET NOTE POSITIONS COMMAND"
	case SspCmdStackLastNote:
		return "STACK LAST NOTE COMMAND"
	case SspCmdSetGenerator:
		return "SET GENERATOR COMMAND"
	case SspCmdSetModulus:
		return "SET MODULUS COMMAND"
	case SspCmdRequestKeyExchange:
		return "KEY EXCHANGE COMMAND"
	case 0xF1:
		return "RESET RESPONSE"
	case 0xEF:
		return "NOTE READ RESPONSE"
	case 0xEE:
		return "CREDIT RESPONSE"
	case 0xED:
		return "REJECTING RESPONSE"
	case 0xEC:
		return "REJECTED RESPONSE"
	case 0xCC:
		return "STACKING RESPONSE"
	case 0xEB:
		return "STACKED RESPONSE"
	case 0xEA:
		return "SAFE JAM RESPONSE"
	case 0xE9:
		return "UNSAFE JAM RESPONSE"
	case 0xE8:
		return "DISABLED RESPONSE"
	case 0xE6:
		return "FRAUD ATTEMPT RESPONSE"
	case 0xE7:
		return "STACKER FULL RESPONSE"
	case 0xE1:
		return "NOTE CLEARED FROM FRONT RESPONSE"
	case 0xE2:
		return "NOTE CLEARED TO CASHBOX RESPONSE"
	case 0xE3:
		return "CASHBOX REMOVED RESPONSE"
	case 0xE4:
		return "CASHBOX REPLACED RESPONSE"
	case 0xDB:
		return "NOTE STORED RESPONSE"
	case 0xDA:
		return "NOTE DISPENSING RESPONSE"
	case 0xD2:
		return "NOTE DISPENSED RESPONSE"
	case 0xC9:
		return "NOTE TRANSFERRED TO STACKER RESPONSE"
	case 0xF0:
		return "OK RESPONSE"
	case 0xF2:
		return "UNKNOWN RESPONSE"
	case 0xF3:
		return "WRONG PARAMS RESPONSE"
	case 0xF4:
		return "PARAM OUT OF RANGE RESPONSE"
	case 0xF5:
		return "CANNOT PROCESS RESPONSE"
	case 0xF6:
		return "SOFTWARE ERROR RESPONSE"
	case 0xF8:
		return "FAIL RESPONSE"
	case 0xFA:
		return "KEY NOT SET RESPONSE"
	default:
		return "Byte command name unsupported"
	}
}
