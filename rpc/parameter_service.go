// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type ParameterService interface {
	// Parameters:
	//  - Keys
	//  - Values
	Push(keys []int64, values []float64) (ex *ParameterServiceException, err error)
	// Parameters:
	//  - Keys
	Pull(keys []int64) (r []float64, ex *ParameterServiceException, err error)
}

type ParameterServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewParameterServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ParameterServiceClient {
	return &ParameterServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewParameterServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ParameterServiceClient {
	return &ParameterServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Keys
//  - Values
func (p *ParameterServiceClient) Push(keys []int64, values []float64) (ex *ParameterServiceException, err error) {
	if err = p.sendPush(keys, values); err != nil {
		return
	}
	return p.recvPush()
}

func (p *ParameterServiceClient) sendPush(keys []int64, values []float64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("push", thrift.CALL, p.SeqId)
	args0 := NewPushArgs()
	args0.Keys = keys
	args0.Values = values
	err = args0.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *ParameterServiceClient) recvPush() (ex *ParameterServiceException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result1 := NewPushResult()
	err = result1.Read(iprot)
	iprot.ReadMessageEnd()
	if result1.Ex != nil {
		ex = result1.Ex
	}
	return
}

// Parameters:
//  - Keys
func (p *ParameterServiceClient) Pull(keys []int64) (r []float64, ex *ParameterServiceException, err error) {
	if err = p.sendPull(keys); err != nil {
		return
	}
	return p.recvPull()
}

func (p *ParameterServiceClient) sendPull(keys []int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("pull", thrift.CALL, p.SeqId)
	args4 := NewPullArgs()
	args4.Keys = keys
	err = args4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *ParameterServiceClient) recvPull() (value []float64, ex *ParameterServiceException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result5 := NewPullResult()
	err = result5.Read(iprot)
	iprot.ReadMessageEnd()
	value = result5.Success
	if result5.Ex != nil {
		ex = result5.Ex
	}
	return
}

type ParameterServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ParameterService
}

func (p *ParameterServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ParameterServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ParameterServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewParameterServiceProcessor(handler ParameterService) *ParameterServiceProcessor {

	self8 := &ParameterServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["push"] = &parameterServiceProcessorPush{handler: handler}
	self8.processorMap["pull"] = &parameterServiceProcessorPull{handler: handler}
	return self8
}

func (p *ParameterServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type parameterServiceProcessorPush struct {
	handler ParameterService
}

func (p *parameterServiceProcessorPush) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewPushArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("push", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewPushResult()
	if result.Ex, err = p.handler.Push(args.Keys, args.Values); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing push: "+err.Error())
		oprot.WriteMessageBegin("push", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("push", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type parameterServiceProcessorPull struct {
	handler ParameterService
}

func (p *parameterServiceProcessorPull) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewPullArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("pull", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewPullResult()
	if result.Success, result.Ex, err = p.handler.Pull(args.Keys); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing pull: "+err.Error())
		oprot.WriteMessageBegin("pull", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("pull", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type PushArgs struct {
	Keys   []int64   `thrift:"keys,1"`
	Values []float64 `thrift:"values,2"`
}

func NewPushArgs() *PushArgs {
	return &PushArgs{}
}

func (p *PushArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PushArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Keys = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem10 int64
		if v, err := iprot.ReadI64(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem10 = v
		}
		p.Keys = append(p.Keys, _elem10)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *PushArgs) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Values = make([]float64, 0, size)
	for i := 0; i < size; i++ {
		var _elem11 float64
		if v, err := iprot.ReadDouble(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem11 = v
		}
		p.Values = append(p.Values, _elem11)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *PushArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("push_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PushArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Keys != nil {
		if err := oprot.WriteFieldBegin("keys", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:keys: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Keys)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Keys {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:keys: %s", p, err)
		}
	}
	return err
}

func (p *PushArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Values != nil {
		if err := oprot.WriteFieldBegin("values", thrift.LIST, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:values: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.DOUBLE, len(p.Values)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Values {
			if err := oprot.WriteDouble(float64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:values: %s", p, err)
		}
	}
	return err
}

func (p *PushArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PushArgs(%+v)", *p)
}

type PushResult struct {
	Ex *ParameterServiceException `thrift:"ex,1"`
}

func NewPushResult() *PushResult {
	return &PushResult{}
}

func (p *PushResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PushResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = NewParameterServiceException()
	if err := p.Ex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Ex)
	}
	return nil
}

func (p *PushResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("push_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Ex != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PushResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Ex != nil {
		if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:ex: %s", p, err)
		}
		if err := p.Ex.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Ex)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:ex: %s", p, err)
		}
	}
	return err
}

func (p *PushResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PushResult(%+v)", *p)
}

type PullArgs struct {
	Keys []int64 `thrift:"keys,1"`
}

func NewPullArgs() *PullArgs {
	return &PullArgs{}
}

func (p *PullArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PullArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Keys = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem12 int64
		if v, err := iprot.ReadI64(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem12 = v
		}
		p.Keys = append(p.Keys, _elem12)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *PullArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("pull_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PullArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Keys != nil {
		if err := oprot.WriteFieldBegin("keys", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:keys: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Keys)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Keys {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:keys: %s", p, err)
		}
	}
	return err
}

func (p *PullArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PullArgs(%+v)", *p)
}

type PullResult struct {
	Success []float64                  `thrift:"success,0"`
	Ex      *ParameterServiceException `thrift:"ex,1"`
}

func NewPullResult() *PullResult {
	return &PullResult{}
}

func (p *PullResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PullResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Success = make([]float64, 0, size)
	for i := 0; i < size; i++ {
		var _elem13 float64
		if v, err := iprot.ReadDouble(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem13 = v
		}
		p.Success = append(p.Success, _elem13)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *PullResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = NewParameterServiceException()
	if err := p.Ex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Ex)
	}
	return nil
}

func (p *PullResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("pull_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Ex != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PullResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.DOUBLE, len(p.Success)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Success {
			if err := oprot.WriteDouble(float64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *PullResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Ex != nil {
		if err := oprot.WriteFieldBegin("ex", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:ex: %s", p, err)
		}
		if err := p.Ex.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Ex)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:ex: %s", p, err)
		}
	}
	return err
}

func (p *PullResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PullResult(%+v)", *p)
}
