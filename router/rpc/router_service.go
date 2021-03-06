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

type RouterService interface {
	// Parameters:
	//  - Nodes
	SetNodes(nodes []string) (ex *RouterException, err error)
	// Parameters:
	//  - Key
	GetNodesByFeature(key []int64) (r []string, ex *RouterException, err error)
}

type RouterServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewRouterServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RouterServiceClient {
	return &RouterServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewRouterServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RouterServiceClient {
	return &RouterServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Nodes
func (p *RouterServiceClient) SetNodes(nodes []string) (ex *RouterException, err error) {
	if err = p.sendSetNodes(nodes); err != nil {
		return
	}
	return p.recvSetNodes()
}

func (p *RouterServiceClient) sendSetNodes(nodes []string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("setNodes", thrift.CALL, p.SeqId)
	args0 := NewSetNodesArgs()
	args0.Nodes = nodes
	err = args0.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *RouterServiceClient) recvSetNodes() (ex *RouterException, err error) {
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
	result1 := NewSetNodesResult()
	err = result1.Read(iprot)
	iprot.ReadMessageEnd()
	if result1.Ex != nil {
		ex = result1.Ex
	}
	return
}

// Parameters:
//  - Key
func (p *RouterServiceClient) GetNodesByFeature(key []int64) (r []string, ex *RouterException, err error) {
	if err = p.sendGetNodesByFeature(key); err != nil {
		return
	}
	return p.recvGetNodesByFeature()
}

func (p *RouterServiceClient) sendGetNodesByFeature(key []int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("getNodesByFeature", thrift.CALL, p.SeqId)
	args4 := NewGetNodesByFeatureArgs()
	args4.Key = key
	err = args4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *RouterServiceClient) recvGetNodesByFeature() (value []string, ex *RouterException, err error) {
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
	result5 := NewGetNodesByFeatureResult()
	err = result5.Read(iprot)
	iprot.ReadMessageEnd()
	value = result5.Success
	if result5.Ex != nil {
		ex = result5.Ex
	}
	return
}

type RouterServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      RouterService
}

func (p *RouterServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RouterServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RouterServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRouterServiceProcessor(handler RouterService) *RouterServiceProcessor {

	self8 := &RouterServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["setNodes"] = &routerServiceProcessorSetNodes{handler: handler}
	self8.processorMap["getNodesByFeature"] = &routerServiceProcessorGetNodesByFeature{handler: handler}
	return self8
}

func (p *RouterServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type routerServiceProcessorSetNodes struct {
	handler RouterService
}

func (p *routerServiceProcessorSetNodes) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSetNodesArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("setNodes", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSetNodesResult()
	if result.Ex, err = p.handler.SetNodes(args.Nodes); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing setNodes: "+err.Error())
		oprot.WriteMessageBegin("setNodes", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("setNodes", thrift.REPLY, seqId); err2 != nil {
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

type routerServiceProcessorGetNodesByFeature struct {
	handler RouterService
}

func (p *routerServiceProcessorGetNodesByFeature) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetNodesByFeatureArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getNodesByFeature", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetNodesByFeatureResult()
	if result.Success, result.Ex, err = p.handler.GetNodesByFeature(args.Key); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getNodesByFeature: "+err.Error())
		oprot.WriteMessageBegin("getNodesByFeature", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("getNodesByFeature", thrift.REPLY, seqId); err2 != nil {
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

type SetNodesArgs struct {
	Nodes []string `thrift:"nodes,1"`
}

func NewSetNodesArgs() *SetNodesArgs {
	return &SetNodesArgs{}
}

func (p *SetNodesArgs) Read(iprot thrift.TProtocol) error {
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

func (p *SetNodesArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Nodes = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem10 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem10 = v
		}
		p.Nodes = append(p.Nodes, _elem10)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *SetNodesArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("setNodes_args"); err != nil {
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

func (p *SetNodesArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Nodes != nil {
		if err := oprot.WriteFieldBegin("nodes", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:nodes: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Nodes)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Nodes {
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:nodes: %s", p, err)
		}
	}
	return err
}

func (p *SetNodesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetNodesArgs(%+v)", *p)
}

type SetNodesResult struct {
	Ex *RouterException `thrift:"ex,1"`
}

func NewSetNodesResult() *SetNodesResult {
	return &SetNodesResult{}
}

func (p *SetNodesResult) Read(iprot thrift.TProtocol) error {
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

func (p *SetNodesResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = NewRouterException()
	if err := p.Ex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Ex)
	}
	return nil
}

func (p *SetNodesResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("setNodes_result"); err != nil {
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

func (p *SetNodesResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *SetNodesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetNodesResult(%+v)", *p)
}

type GetNodesByFeatureArgs struct {
	Key []int64 `thrift:"key,1"`
}

func NewGetNodesByFeatureArgs() *GetNodesByFeatureArgs {
	return &GetNodesByFeatureArgs{}
}

func (p *GetNodesByFeatureArgs) Read(iprot thrift.TProtocol) error {
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

func (p *GetNodesByFeatureArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Key = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem11 int64
		if v, err := iprot.ReadI64(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem11 = v
		}
		p.Key = append(p.Key, _elem11)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *GetNodesByFeatureArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getNodesByFeature_args"); err != nil {
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

func (p *GetNodesByFeatureArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Key != nil {
		if err := oprot.WriteFieldBegin("key", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Key)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Key {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:key: %s", p, err)
		}
	}
	return err
}

func (p *GetNodesByFeatureArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetNodesByFeatureArgs(%+v)", *p)
}

type GetNodesByFeatureResult struct {
	Success []string         `thrift:"success,0"`
	Ex      *RouterException `thrift:"ex,1"`
}

func NewGetNodesByFeatureResult() *GetNodesByFeatureResult {
	return &GetNodesByFeatureResult{}
}

func (p *GetNodesByFeatureResult) Read(iprot thrift.TProtocol) error {
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

func (p *GetNodesByFeatureResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Success = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var _elem12 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem12 = v
		}
		p.Success = append(p.Success, _elem12)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *GetNodesByFeatureResult) readField1(iprot thrift.TProtocol) error {
	p.Ex = NewRouterException()
	if err := p.Ex.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Ex)
	}
	return nil
}

func (p *GetNodesByFeatureResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getNodesByFeature_result"); err != nil {
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

func (p *GetNodesByFeatureResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Success)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
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

func (p *GetNodesByFeatureResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GetNodesByFeatureResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetNodesByFeatureResult(%+v)", *p)
}
