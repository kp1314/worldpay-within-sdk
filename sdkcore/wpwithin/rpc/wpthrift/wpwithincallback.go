// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package wpthrift

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc/wpthrift/wpthrift_types"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = wpthrift_types.GoUnusedProtection__

type WPWithinCallback interface { //WorldpayWithin Callback Service - RPC clients implement this service to enable callbacks from WorldpayWithin Service

	// Parameters:
	//  - ServiceID
	//  - ServiceDeliveryToken
	//  - UnitsToSupply
	BeginServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsToSupply int32) (err error)
	// Parameters:
	//  - ServiceID
	//  - ServiceDeliveryToken
	//  - UnitsReceived
	EndServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsReceived int32) (err error)
}

//WorldpayWithin Callback Service - RPC clients implement this service to enable callbacks from WorldpayWithin Service
type WPWithinCallbackClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewWPWithinCallbackClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *WPWithinCallbackClient {
	return &WPWithinCallbackClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewWPWithinCallbackClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *WPWithinCallbackClient {
	return &WPWithinCallbackClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - ServiceID
//  - ServiceDeliveryToken
//  - UnitsToSupply
func (p *WPWithinCallbackClient) BeginServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsToSupply int32) (err error) {
	if err = p.sendBeginServiceDelivery(serviceID, serviceDeliveryToken, unitsToSupply); err != nil {
		return
	}
	return p.recvBeginServiceDelivery()
}

func (p *WPWithinCallbackClient) sendBeginServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsToSupply int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("beginServiceDelivery", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WPWithinCallbackBeginServiceDeliveryArgs{
		ServiceID:            serviceID,
		ServiceDeliveryToken: serviceDeliveryToken,
		UnitsToSupply:        unitsToSupply,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WPWithinCallbackClient) recvBeginServiceDelivery() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "beginServiceDelivery" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "beginServiceDelivery failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "beginServiceDelivery failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error90 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error91 error
		error91, err = error90.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error91
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "beginServiceDelivery failed: invalid message type")
		return
	}
	result := WPWithinCallbackBeginServiceDeliveryResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Err != nil {
		err = result.Err
		return
	}
	return
}

// Parameters:
//  - ServiceID
//  - ServiceDeliveryToken
//  - UnitsReceived
func (p *WPWithinCallbackClient) EndServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsReceived int32) (err error) {
	if err = p.sendEndServiceDelivery(serviceID, serviceDeliveryToken, unitsReceived); err != nil {
		return
	}
	return p.recvEndServiceDelivery()
}

func (p *WPWithinCallbackClient) sendEndServiceDelivery(serviceID int32, serviceDeliveryToken *wpthrift_types.ServiceDeliveryToken, unitsReceived int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("endServiceDelivery", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := WPWithinCallbackEndServiceDeliveryArgs{
		ServiceID:            serviceID,
		ServiceDeliveryToken: serviceDeliveryToken,
		UnitsReceived:        unitsReceived,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *WPWithinCallbackClient) recvEndServiceDelivery() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "endServiceDelivery" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "endServiceDelivery failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "endServiceDelivery failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error92 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error93 error
		error93, err = error92.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error93
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "endServiceDelivery failed: invalid message type")
		return
	}
	result := WPWithinCallbackEndServiceDeliveryResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.Err != nil {
		err = result.Err
		return
	}
	return
}

type WPWithinCallbackProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      WPWithinCallback
}

func (p *WPWithinCallbackProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *WPWithinCallbackProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *WPWithinCallbackProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewWPWithinCallbackProcessor(handler WPWithinCallback) *WPWithinCallbackProcessor {

	self94 := &WPWithinCallbackProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self94.processorMap["beginServiceDelivery"] = &wPWithinCallbackProcessorBeginServiceDelivery{handler: handler}
	self94.processorMap["endServiceDelivery"] = &wPWithinCallbackProcessorEndServiceDelivery{handler: handler}
	return self94
}

func (p *WPWithinCallbackProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x95 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x95.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x95

}

type wPWithinCallbackProcessorBeginServiceDelivery struct {
	handler WPWithinCallback
}

func (p *wPWithinCallbackProcessorBeginServiceDelivery) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WPWithinCallbackBeginServiceDeliveryArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("beginServiceDelivery", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WPWithinCallbackBeginServiceDeliveryResult{}
	var err2 error
	if err2 = p.handler.BeginServiceDelivery(args.ServiceID, args.ServiceDeliveryToken, args.UnitsToSupply); err2 != nil {
		switch v := err2.(type) {
		case *wpthrift_types.Error:
			result.Err = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing beginServiceDelivery: "+err2.Error())
			oprot.WriteMessageBegin("beginServiceDelivery", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("beginServiceDelivery", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type wPWithinCallbackProcessorEndServiceDelivery struct {
	handler WPWithinCallback
}

func (p *wPWithinCallbackProcessorEndServiceDelivery) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := WPWithinCallbackEndServiceDeliveryArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("endServiceDelivery", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := WPWithinCallbackEndServiceDeliveryResult{}
	var err2 error
	if err2 = p.handler.EndServiceDelivery(args.ServiceID, args.ServiceDeliveryToken, args.UnitsReceived); err2 != nil {
		switch v := err2.(type) {
		case *wpthrift_types.Error:
			result.Err = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing endServiceDelivery: "+err2.Error())
			oprot.WriteMessageBegin("endServiceDelivery", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("endServiceDelivery", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - ServiceID
//  - ServiceDeliveryToken
//  - UnitsToSupply
type WPWithinCallbackBeginServiceDeliveryArgs struct {
	ServiceID            int32                                `thrift:"serviceID,1" json:"serviceID"`
	ServiceDeliveryToken *wpthrift_types.ServiceDeliveryToken `thrift:"serviceDeliveryToken,2" json:"serviceDeliveryToken"`
	UnitsToSupply        int32                                `thrift:"unitsToSupply,3" json:"unitsToSupply"`
}

func NewWPWithinCallbackBeginServiceDeliveryArgs() *WPWithinCallbackBeginServiceDeliveryArgs {
	return &WPWithinCallbackBeginServiceDeliveryArgs{}
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) GetServiceID() int32 {
	return p.ServiceID
}

var WPWithinCallbackBeginServiceDeliveryArgs_ServiceDeliveryToken_DEFAULT *wpthrift_types.ServiceDeliveryToken

func (p *WPWithinCallbackBeginServiceDeliveryArgs) GetServiceDeliveryToken() *wpthrift_types.ServiceDeliveryToken {
	if !p.IsSetServiceDeliveryToken() {
		return WPWithinCallbackBeginServiceDeliveryArgs_ServiceDeliveryToken_DEFAULT
	}
	return p.ServiceDeliveryToken
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) GetUnitsToSupply() int32 {
	return p.UnitsToSupply
}
func (p *WPWithinCallbackBeginServiceDeliveryArgs) IsSetServiceDeliveryToken() bool {
	return p.ServiceDeliveryToken != nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ServiceID = v
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) readField2(iprot thrift.TProtocol) error {
	p.ServiceDeliveryToken = &wpthrift_types.ServiceDeliveryToken{}
	if err := p.ServiceDeliveryToken.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ServiceDeliveryToken), err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.UnitsToSupply = v
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("beginServiceDelivery_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceID", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceID: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ServiceID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceID: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceDeliveryToken", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:serviceDeliveryToken: ", p), err)
	}
	if err := p.ServiceDeliveryToken.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ServiceDeliveryToken), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:serviceDeliveryToken: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("unitsToSupply", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:unitsToSupply: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.UnitsToSupply)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.unitsToSupply (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:unitsToSupply: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackBeginServiceDeliveryArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WPWithinCallbackBeginServiceDeliveryArgs(%+v)", *p)
}

// Attributes:
//  - Err
type WPWithinCallbackBeginServiceDeliveryResult struct {
	Err *wpthrift_types.Error `thrift:"err,1" json:"err,omitempty"`
}

func NewWPWithinCallbackBeginServiceDeliveryResult() *WPWithinCallbackBeginServiceDeliveryResult {
	return &WPWithinCallbackBeginServiceDeliveryResult{}
}

var WPWithinCallbackBeginServiceDeliveryResult_Err_DEFAULT *wpthrift_types.Error

func (p *WPWithinCallbackBeginServiceDeliveryResult) GetErr() *wpthrift_types.Error {
	if !p.IsSetErr() {
		return WPWithinCallbackBeginServiceDeliveryResult_Err_DEFAULT
	}
	return p.Err
}
func (p *WPWithinCallbackBeginServiceDeliveryResult) IsSetErr() bool {
	return p.Err != nil
}

func (p *WPWithinCallbackBeginServiceDeliveryResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryResult) readField1(iprot thrift.TProtocol) error {
	p.Err = &wpthrift_types.Error{}
	if err := p.Err.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Err), err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("beginServiceDelivery_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WPWithinCallbackBeginServiceDeliveryResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErr() {
		if err := oprot.WriteFieldBegin("err", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:err: ", p), err)
		}
		if err := p.Err.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Err), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:err: ", p), err)
		}
	}
	return err
}

func (p *WPWithinCallbackBeginServiceDeliveryResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WPWithinCallbackBeginServiceDeliveryResult(%+v)", *p)
}

// Attributes:
//  - ServiceID
//  - ServiceDeliveryToken
//  - UnitsReceived
type WPWithinCallbackEndServiceDeliveryArgs struct {
	ServiceID            int32                                `thrift:"serviceID,1" json:"serviceID"`
	ServiceDeliveryToken *wpthrift_types.ServiceDeliveryToken `thrift:"serviceDeliveryToken,2" json:"serviceDeliveryToken"`
	UnitsReceived        int32                                `thrift:"unitsReceived,3" json:"unitsReceived"`
}

func NewWPWithinCallbackEndServiceDeliveryArgs() *WPWithinCallbackEndServiceDeliveryArgs {
	return &WPWithinCallbackEndServiceDeliveryArgs{}
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) GetServiceID() int32 {
	return p.ServiceID
}

var WPWithinCallbackEndServiceDeliveryArgs_ServiceDeliveryToken_DEFAULT *wpthrift_types.ServiceDeliveryToken

func (p *WPWithinCallbackEndServiceDeliveryArgs) GetServiceDeliveryToken() *wpthrift_types.ServiceDeliveryToken {
	if !p.IsSetServiceDeliveryToken() {
		return WPWithinCallbackEndServiceDeliveryArgs_ServiceDeliveryToken_DEFAULT
	}
	return p.ServiceDeliveryToken
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) GetUnitsReceived() int32 {
	return p.UnitsReceived
}
func (p *WPWithinCallbackEndServiceDeliveryArgs) IsSetServiceDeliveryToken() bool {
	return p.ServiceDeliveryToken != nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ServiceID = v
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) readField2(iprot thrift.TProtocol) error {
	p.ServiceDeliveryToken = &wpthrift_types.ServiceDeliveryToken{}
	if err := p.ServiceDeliveryToken.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ServiceDeliveryToken), err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.UnitsReceived = v
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("endServiceDelivery_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceID", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceID: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ServiceID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceID: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceDeliveryToken", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:serviceDeliveryToken: ", p), err)
	}
	if err := p.ServiceDeliveryToken.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ServiceDeliveryToken), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:serviceDeliveryToken: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("unitsReceived", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:unitsReceived: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.UnitsReceived)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.unitsReceived (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:unitsReceived: ", p), err)
	}
	return err
}

func (p *WPWithinCallbackEndServiceDeliveryArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WPWithinCallbackEndServiceDeliveryArgs(%+v)", *p)
}

// Attributes:
//  - Err
type WPWithinCallbackEndServiceDeliveryResult struct {
	Err *wpthrift_types.Error `thrift:"err,1" json:"err,omitempty"`
}

func NewWPWithinCallbackEndServiceDeliveryResult() *WPWithinCallbackEndServiceDeliveryResult {
	return &WPWithinCallbackEndServiceDeliveryResult{}
}

var WPWithinCallbackEndServiceDeliveryResult_Err_DEFAULT *wpthrift_types.Error

func (p *WPWithinCallbackEndServiceDeliveryResult) GetErr() *wpthrift_types.Error {
	if !p.IsSetErr() {
		return WPWithinCallbackEndServiceDeliveryResult_Err_DEFAULT
	}
	return p.Err
}
func (p *WPWithinCallbackEndServiceDeliveryResult) IsSetErr() bool {
	return p.Err != nil
}

func (p *WPWithinCallbackEndServiceDeliveryResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryResult) readField1(iprot thrift.TProtocol) error {
	p.Err = &wpthrift_types.Error{}
	if err := p.Err.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Err), err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("endServiceDelivery_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WPWithinCallbackEndServiceDeliveryResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErr() {
		if err := oprot.WriteFieldBegin("err", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:err: ", p), err)
		}
		if err := p.Err.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Err), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:err: ", p), err)
		}
	}
	return err
}

func (p *WPWithinCallbackEndServiceDeliveryResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WPWithinCallbackEndServiceDeliveryResult(%+v)", *p)
}
