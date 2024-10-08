// Code generated by thriftgo (0.3.15). DO NOT EDIT.

package hertz_demo

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// 引入其他 Thrift 文件
// include "user.thrift"
type GetUserInfoReq struct {
	UUID   string  `thrift:"uuid,1,required" form:"uuid,required" json:"uuid,required" query:"uuid,required"`
	Openid *string `thrift:"openid,2,optional" form:"openid" json:"openid,omitempty" query:"openid"`
}

func NewGetUserInfoReq() *GetUserInfoReq {
	return &GetUserInfoReq{}
}

func (p *GetUserInfoReq) InitDefault() {
}

func (p *GetUserInfoReq) GetUUID() (v string) {
	return p.UUID
}

var GetUserInfoReq_Openid_DEFAULT string

func (p *GetUserInfoReq) GetOpenid() (v string) {
	if !p.IsSetOpenid() {
		return GetUserInfoReq_Openid_DEFAULT
	}
	return *p.Openid
}

var fieldIDToName_GetUserInfoReq = map[int16]string{
	1: "uuid",
	2: "openid",
}

func (p *GetUserInfoReq) IsSetOpenid() bool {
	return p.Openid != nil
}

func (p *GetUserInfoReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUUID bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetUUID = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetUUID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetUserInfoReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetUserInfoReq[fieldId]))
}

func (p *GetUserInfoReq) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UUID = _field
	return nil
}
func (p *GetUserInfoReq) ReadField2(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Openid = _field
	return nil
}

func (p *GetUserInfoReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfoReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetUserInfoReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("uuid", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UUID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetUserInfoReq) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetOpenid() {
		if err = oprot.WriteFieldBegin("openid", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Openid); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetUserInfoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserInfoReq(%+v)", *p)

}

type GetUserInfoResp struct {
	UUID        string `thrift:"uuid,1,required" form:"uuid,required" json:"uuid,required" query:"uuid,required"`
	Openid      string `thrift:"openid,2,required" form:"openid,required" json:"openid,required" query:"openid,required"`
	Name        string `thrift:"name,3,required" form:"name,required" json:"name,required" query:"name,required"`
	ProjectName string `thrift:"projectName,4,required" form:"projectName,required" json:"projectName,required" query:"projectName,required"`
}

func NewGetUserInfoResp() *GetUserInfoResp {
	return &GetUserInfoResp{}
}

func (p *GetUserInfoResp) InitDefault() {
}

func (p *GetUserInfoResp) GetUUID() (v string) {
	return p.UUID
}

func (p *GetUserInfoResp) GetOpenid() (v string) {
	return p.Openid
}

func (p *GetUserInfoResp) GetName() (v string) {
	return p.Name
}

func (p *GetUserInfoResp) GetProjectName() (v string) {
	return p.ProjectName
}

var fieldIDToName_GetUserInfoResp = map[int16]string{
	1: "uuid",
	2: "openid",
	3: "name",
	4: "projectName",
}

func (p *GetUserInfoResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUUID bool = false
	var issetOpenid bool = false
	var issetName bool = false
	var issetProjectName bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetUUID = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetOpenid = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetName = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
				issetProjectName = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetUUID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetOpenid {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetName {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetProjectName {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetUserInfoResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetUserInfoResp[fieldId]))
}

func (p *GetUserInfoResp) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UUID = _field
	return nil
}
func (p *GetUserInfoResp) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Openid = _field
	return nil
}
func (p *GetUserInfoResp) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Name = _field
	return nil
}
func (p *GetUserInfoResp) ReadField4(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ProjectName = _field
	return nil
}

func (p *GetUserInfoResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfoResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetUserInfoResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("uuid", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UUID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetUserInfoResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("openid", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Openid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetUserInfoResp) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *GetUserInfoResp) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("projectName", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ProjectName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *GetUserInfoResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserInfoResp(%+v)", *p)

}

type User interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoReq) (r *GetUserInfoResp, err error)
}

type UserClient struct {
	c thrift.TClient
}

func NewUserClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UserClient {
	return &UserClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewUserClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UserClient {
	return &UserClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewUserClient(c thrift.TClient) *UserClient {
	return &UserClient{
		c: c,
	}
}

func (p *UserClient) Client_() thrift.TClient {
	return p.c
}

func (p *UserClient) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (r *GetUserInfoResp, err error) {
	var _args UserGetUserInfoArgs
	_args.Req = req
	var _result UserGetUserInfoResult
	if err = p.Client_().Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type UserProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      User
}

func (p *UserProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UserProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UserProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUserProcessor(handler User) *UserProcessor {
	self := &UserProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetUserInfo", &userProcessorGetUserInfo{handler: handler})
	return self
}
func (p *UserProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type userProcessorGetUserInfo struct {
	handler User
}

func (p *userProcessorGetUserInfo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := UserGetUserInfoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetUserInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := UserGetUserInfoResult{}
	var retval *GetUserInfoResp
	if retval, err2 = p.handler.GetUserInfo(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetUserInfo: "+err2.Error())
		oprot.WriteMessageBegin("GetUserInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetUserInfo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type UserGetUserInfoArgs struct {
	Req *GetUserInfoReq `thrift:"req,1"`
}

func NewUserGetUserInfoArgs() *UserGetUserInfoArgs {
	return &UserGetUserInfoArgs{}
}

func (p *UserGetUserInfoArgs) InitDefault() {
}

var UserGetUserInfoArgs_Req_DEFAULT *GetUserInfoReq

func (p *UserGetUserInfoArgs) GetReq() (v *GetUserInfoReq) {
	if !p.IsSetReq() {
		return UserGetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_UserGetUserInfoArgs = map[int16]string{
	1: "req",
}

func (p *UserGetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserGetUserInfoArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserGetUserInfoArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserGetUserInfoArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGetUserInfoReq()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *UserGetUserInfoArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfo_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserGetUserInfoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserGetUserInfoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserGetUserInfoArgs(%+v)", *p)

}

type UserGetUserInfoResult struct {
	Success *GetUserInfoResp `thrift:"success,0,optional"`
}

func NewUserGetUserInfoResult() *UserGetUserInfoResult {
	return &UserGetUserInfoResult{}
}

func (p *UserGetUserInfoResult) InitDefault() {
}

var UserGetUserInfoResult_Success_DEFAULT *GetUserInfoResp

func (p *UserGetUserInfoResult) GetSuccess() (v *GetUserInfoResp) {
	if !p.IsSetSuccess() {
		return UserGetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_UserGetUserInfoResult = map[int16]string{
	0: "success",
}

func (p *UserGetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserGetUserInfoResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserGetUserInfoResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserGetUserInfoResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewGetUserInfoResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *UserGetUserInfoResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetUserInfo_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserGetUserInfoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UserGetUserInfoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserGetUserInfoResult(%+v)", *p)

}
