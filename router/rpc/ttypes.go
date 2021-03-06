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

var GoUnusedProtection__ int

type RouterException struct {
	Status  int32  `thrift:"status,1"`
	Message string `thrift:"message,2"`
}

func NewRouterException() *RouterException {
	return &RouterException{}
}

func (p *RouterException) Read(iprot thrift.TProtocol) error {
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

func (p *RouterException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Status = v
	}
	return nil
}

func (p *RouterException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Message = v
	}
	return nil
}

func (p *RouterException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RouterException"); err != nil {
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

func (p *RouterException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:status: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return fmt.Errorf("%T.status (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:status: %s", p, err)
	}
	return err
}

func (p *RouterException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.message (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:message: %s", p, err)
	}
	return err
}

func (p *RouterException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RouterException(%+v)", *p)
}
