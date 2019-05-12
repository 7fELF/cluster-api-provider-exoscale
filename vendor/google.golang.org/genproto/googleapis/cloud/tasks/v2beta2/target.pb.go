// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta2/target.proto

package tasks // import "google.golang.org/genproto/googleapis/cloud/tasks/v2beta2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The HTTP method used to execute the task.
type HttpMethod int32

const (
	// HTTP method unspecified
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
)

var HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNSPECIFIED",
	1: "POST",
	2: "GET",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
}
var HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNSPECIFIED": 0,
	"POST":                    1,
	"GET":                     2,
	"HEAD":                    3,
	"PUT":                     4,
	"DELETE":                  5,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}
func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{0}
}

// Pull target.
type PullTarget struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullTarget) Reset()         { *m = PullTarget{} }
func (m *PullTarget) String() string { return proto.CompactTextString(m) }
func (*PullTarget) ProtoMessage()    {}
func (*PullTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{0}
}
func (m *PullTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullTarget.Unmarshal(m, b)
}
func (m *PullTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullTarget.Marshal(b, m, deterministic)
}
func (dst *PullTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullTarget.Merge(dst, src)
}
func (m *PullTarget) XXX_Size() int {
	return xxx_messageInfo_PullTarget.Size(m)
}
func (m *PullTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_PullTarget.DiscardUnknown(m)
}

var xxx_messageInfo_PullTarget proto.InternalMessageInfo

// The pull message contains data that can be used by the caller of
// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] to process the
// task.
//
// This proto can only be used for tasks in a queue which has
// [pull_target][google.cloud.tasks.v2beta2.Queue.pull_target] set.
type PullMessage struct {
	// A data payload consumed by the worker to execute the task.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// The task's tag.
	//
	// Tags allow similar tasks to be processed in a batch. If you label
	// tasks with a tag, your worker can
	// [lease tasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] with the
	// same tag using
	// [filter][google.cloud.tasks.v2beta2.LeaseTasksRequest.filter]. For example,
	// if you want to aggregate the events associated with a specific user once a
	// day, you could tag tasks with the user ID.
	//
	// The task's tag can only be set when the
	// [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//
	// The tag must be less than 500 characters.
	//
	// SDK compatibility: Although the SDK allows tags to be either
	// string or
	// [bytes](https://cloud.google.com/appengine/docs/standard/java/javadoc/com/google/appengine/api/taskqueue/TaskOptions.html#tag-byte:A-),
	// only UTF-8 encoded tags can be used in Cloud Tasks. If a tag isn't UTF-8
	// encoded, the tag will be empty when the task is returned by Cloud Tasks.
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullMessage) Reset()         { *m = PullMessage{} }
func (m *PullMessage) String() string { return proto.CompactTextString(m) }
func (*PullMessage) ProtoMessage()    {}
func (*PullMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{1}
}
func (m *PullMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessage.Unmarshal(m, b)
}
func (m *PullMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessage.Marshal(b, m, deterministic)
}
func (dst *PullMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessage.Merge(dst, src)
}
func (m *PullMessage) XXX_Size() int {
	return xxx_messageInfo_PullMessage.Size(m)
}
func (m *PullMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessage proto.InternalMessageInfo

func (m *PullMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PullMessage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// App Engine HTTP target.
//
// The task will be delivered to the App Engine application hostname
// specified by its
// [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget] and
// [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]. The
// documentation for
// [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]
// explains how the task's host URL is constructed.
//
// Using [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget]
// requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
type AppEngineHttpTarget struct {
	// Overrides for the
	// [task-level
	// app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	//
	// If set, `app_engine_routing_override` is used for all tasks in
	// the queue, no matter what the setting is for the
	// [task-level
	// app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,1,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *AppEngineHttpTarget) Reset()         { *m = AppEngineHttpTarget{} }
func (m *AppEngineHttpTarget) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpTarget) ProtoMessage()    {}
func (*AppEngineHttpTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{2}
}
func (m *AppEngineHttpTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpTarget.Unmarshal(m, b)
}
func (m *AppEngineHttpTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpTarget.Marshal(b, m, deterministic)
}
func (dst *AppEngineHttpTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpTarget.Merge(dst, src)
}
func (m *AppEngineHttpTarget) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpTarget.Size(m)
}
func (m *AppEngineHttpTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpTarget.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpTarget proto.InternalMessageInfo

func (m *AppEngineHttpTarget) GetAppEngineRoutingOverride() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRoutingOverride
	}
	return nil
}

// App Engine HTTP request.
//
// The message defines the HTTP request that is sent to an App Engine app when
// the task is dispatched.
//
// This proto can only be used for tasks in a queue which has
// [app_engine_http_target][google.cloud.tasks.v2beta2.Queue.app_engine_http_target]
// set.
//
// Using [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]
// requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
//
// The task will be delivered to the App Engine app which belongs to the same
// project as the queue. For more information, see
// [How Requests are
// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by
// [dispatch
// files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
//
// The [AppEngineRouting][google.cloud.tasks.v2beta2.AppEngineRouting] used to
// construct the URL that the task is delivered to can be set at the queue-level
// or task-level:
//
// * If set,
//    [app_engine_routing_override][google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override]
//    is used for all tasks in the queue, no matter what the setting
//    is for the
//    [task-level
//    app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
//
//
// The `url` that the task will be sent to is:
//
// * `url =` [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] `+`
//   [relative_url][google.cloud.tasks.v2beta2.AppEngineHttpRequest.relative_url]
//
// The task attempt has succeeded if the app's request handler returns
// an HTTP response code in the range [`200` - `299`]. `503` is
// considered an App Engine system error instead of an application
// error. Requests returning error `503` will be retried regardless of
// retry configuration and not counted against retry counts.
// Any other response code or a failure to receive a response before the
// deadline is a failed attempt.
type AppEngineHttpRequest struct {
	// The HTTP method to use for the request. The default is POST.
	//
	// The app's request handler for the task's target URL must be able to handle
	// HTTP requests with this http_method, otherwise the task attempt will fail
	// with error code 405 (Method Not Allowed). See
	// [Writing a push task request
	// handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler)
	// and the documentation for the request handlers in the language your app is
	// written in e.g.
	// [Python Request
	// Handler](https://cloud.google.com/appengine/docs/python/tools/webapp/requesthandlerclass).
	HttpMethod HttpMethod `protobuf:"varint,1,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.tasks.v2beta2.HttpMethod" json:"http_method,omitempty"`
	// Task-level setting for App Engine routing.
	//
	// If set,
	// [app_engine_routing_override][google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override]
	// is used for all tasks in the queue, no matter what the setting is for the
	// [task-level
	// app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRouting *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing,json=appEngineRouting,proto3" json:"app_engine_routing,omitempty"`
	// The relative URL.
	//
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path and query string arguments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters.
	RelativeUrl string `protobuf:"bytes,3,opt,name=relative_url,json=relativeUrl,proto3" json:"relative_url,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	// Repeated headers are not supported but a header value can contain commas.
	//
	// Cloud Tasks sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//   This header can be modified, but Cloud Tasks will append
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//   modified `User-Agent`.
	//
	// If the task has a
	// [payload][google.cloud.tasks.v2beta2.AppEngineHttpRequest.payload], Cloud
	// Tasks sets the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   `"application/octet-stream"`. The default can be overridden by explicitly
	//   setting `Content-Type` to a particular media type when the
	//   [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//   For example, `Content-Type` can be set to `"application/json"`.
	// * `Content-Length`: This is computed by Cloud Tasks. This value is
	//   output only.   It cannot be changed.
	//
	// The headers below cannot be set or overridden:
	//
	// * `Host`
	// * `X-Google-*`
	// * `X-AppEngine-*`
	//
	// In addition, Cloud Tasks sets some headers when the task is dispatched,
	// such as headers containing information about the task; see
	// [request
	// headers](https://cloud.google.com/appengine/docs/python/taskqueue/push/creating-handlers#reading_request_headers).
	// These headers are set only when the task is dispatched, so they are not
	// visible when the task is returned in a Cloud Tasks response.
	//
	// Although there is no specific limit for the maximum number of headers or
	// the size, there is a limit on the maximum size of the
	// [Task][google.cloud.tasks.v2beta2.Task]. For more information, see the
	// [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask]
	// documentation.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Payload.
	//
	// The payload will be sent as the HTTP message body. A message
	// body, and thus a payload, is allowed only if the HTTP method is
	// POST or PUT. It is an error to set a data payload on a task with
	// an incompatible [HttpMethod][google.cloud.tasks.v2beta2.HttpMethod].
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineHttpRequest) Reset()         { *m = AppEngineHttpRequest{} }
func (m *AppEngineHttpRequest) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpRequest) ProtoMessage()    {}
func (*AppEngineHttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{3}
}
func (m *AppEngineHttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpRequest.Unmarshal(m, b)
}
func (m *AppEngineHttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpRequest.Marshal(b, m, deterministic)
}
func (dst *AppEngineHttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpRequest.Merge(dst, src)
}
func (m *AppEngineHttpRequest) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpRequest.Size(m)
}
func (m *AppEngineHttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpRequest proto.InternalMessageInfo

func (m *AppEngineHttpRequest) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *AppEngineHttpRequest) GetAppEngineRouting() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRouting
	}
	return nil
}

func (m *AppEngineHttpRequest) GetRelativeUrl() string {
	if m != nil {
		return m.RelativeUrl
	}
	return ""
}

func (m *AppEngineHttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AppEngineHttpRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// App Engine Routing.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	// App service.
	//
	// By default, the task is sent to the service which is the default
	// service when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is
	// not parsable into
	// [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For
	// example, some tasks which were created using the App Engine SDK use a
	// custom domain name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable,
	// then [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the
	// empty string.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// App version.
	//
	// By default, the task is sent to the version which is the default
	// version when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is
	// not parsable into
	// [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For
	// example, some tasks which were created using the App Engine SDK use a
	// custom domain name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable,
	// then [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the
	// empty string.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// App instance.
	//
	// By default, the task is sent to an instance which is available when
	// the task is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information, see
	// [App Engine Standard request
	// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request
	// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. The host that the task is sent to.
	//
	// For more information, see
	// [How Requests are
	// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	//
	// The host is constructed as:
	//
	//
	// * `host = [application_domain_name]`</br>
	//   `| [service] + '.' + [application_domain_name]`</br>
	//   `| [version] + '.' + [application_domain_name]`</br>
	//   `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//   `| [instance] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version_dot_service] + '.' + [application_domain_name]`
	//
	// * `application_domain_name` = The domain name of the app, for
	//   example <app-id>.appspot.com, which is associated with the
	//   queue's project ID. Some tasks which were created using the App Engine
	//   SDK use a custom domain name.
	//
	// * `service =`
	// [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `version =`
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	// * `version_dot_service =`
	//   [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.' +`
	//   [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `instance =`
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]
	//
	// * `instance_dot_service =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//   +` [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `instance_dot_version =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//   +` [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	// * `instance_dot_version_dot_service =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.'
	//   +` [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.'
	//   +` [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service] is empty,
	// then the task will be sent to the service which is the default service when
	// the task is attempted.
	//
	// If [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] is empty,
	// then the task will be sent to the version which is the default version when
	// the task is attempted.
	//
	// If [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is
	// empty, then the task will be sent to an instance which is available when
	// the task is attempted.
	//
	// If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], or
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is
	// invalid, then the task will be sent to the default version of the default
	// service when the task is attempted.
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineRouting) Reset()         { *m = AppEngineRouting{} }
func (m *AppEngineRouting) String() string { return proto.CompactTextString(m) }
func (*AppEngineRouting) ProtoMessage()    {}
func (*AppEngineRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_target_94aeace9d01cd65d, []int{4}
}
func (m *AppEngineRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineRouting.Unmarshal(m, b)
}
func (m *AppEngineRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineRouting.Marshal(b, m, deterministic)
}
func (dst *AppEngineRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineRouting.Merge(dst, src)
}
func (m *AppEngineRouting) XXX_Size() int {
	return xxx_messageInfo_AppEngineRouting.Size(m)
}
func (m *AppEngineRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineRouting proto.InternalMessageInfo

func (m *AppEngineRouting) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AppEngineRouting) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppEngineRouting) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *AppEngineRouting) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*PullTarget)(nil), "google.cloud.tasks.v2beta2.PullTarget")
	proto.RegisterType((*PullMessage)(nil), "google.cloud.tasks.v2beta2.PullMessage")
	proto.RegisterType((*AppEngineHttpTarget)(nil), "google.cloud.tasks.v2beta2.AppEngineHttpTarget")
	proto.RegisterType((*AppEngineHttpRequest)(nil), "google.cloud.tasks.v2beta2.AppEngineHttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.tasks.v2beta2.AppEngineHttpRequest.HeadersEntry")
	proto.RegisterType((*AppEngineRouting)(nil), "google.cloud.tasks.v2beta2.AppEngineRouting")
	proto.RegisterEnum("google.cloud.tasks.v2beta2.HttpMethod", HttpMethod_name, HttpMethod_value)
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta2/target.proto", fileDescriptor_target_94aeace9d01cd65d)
}

var fileDescriptor_target_94aeace9d01cd65d = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0x5c, 0xa7, 0xbf, 0xc6, 0xd1, 0x27, 0x6b, 0xa9, 0x84, 0x95, 0xa2, 0xaa, 0xe4, 0x00,
	0x15, 0x42, 0xb6, 0x14, 0x2e, 0x50, 0x84, 0x50, 0x4b, 0x4c, 0x53, 0x89, 0x12, 0xcb, 0x75, 0x84,
	0x54, 0x0e, 0xd6, 0x26, 0x19, 0x1c, 0x2b, 0xee, 0xae, 0xd9, 0x5d, 0x5b, 0xca, 0x95, 0x3b, 0xff,
	0x33, 0xf2, 0xda, 0x09, 0x69, 0x80, 0x0a, 0x6e, 0xf3, 0x66, 0xde, 0xbc, 0xc9, 0xbc, 0xf1, 0x06,
	0x9e, 0x26, 0x9c, 0x27, 0x19, 0x7a, 0x93, 0x8c, 0x17, 0x53, 0x4f, 0x51, 0x39, 0x97, 0x5e, 0xd9,
	0x1b, 0xa3, 0xa2, 0x3d, 0x4f, 0x51, 0x91, 0xa0, 0x72, 0x73, 0xc1, 0x15, 0x27, 0x9d, 0x9a, 0xe8,
	0x6a, 0xa2, 0xab, 0x89, 0x6e, 0x43, 0xec, 0x3c, 0x6a, 0x44, 0x68, 0x9e, 0x7a, 0x94, 0x31, 0xae,
	0xa8, 0x4a, 0x39, 0x93, 0x75, 0x67, 0xe7, 0xa8, 0xa9, 0x6a, 0x34, 0x2e, 0xbe, 0x78, 0xd3, 0x42,
	0x68, 0x42, 0x5d, 0xef, 0xb6, 0x01, 0x82, 0x22, 0xcb, 0x22, 0x3d, 0xad, 0xfb, 0x0a, 0xac, 0x0a,
	0x5d, 0xa1, 0x94, 0x34, 0x41, 0xe2, 0xc0, 0x6e, 0x4e, 0x17, 0x19, 0xa7, 0x53, 0xc7, 0x38, 0x36,
	0x4e, 0xda, 0xe1, 0x12, 0x12, 0x1b, 0x4c, 0x45, 0x13, 0x67, 0xeb, 0xd8, 0x38, 0xd9, 0x0f, 0xab,
	0xb0, 0xfb, 0xcd, 0x80, 0x07, 0x67, 0x79, 0xee, 0xb3, 0x24, 0x65, 0x38, 0x50, 0x2a, 0xaf, 0x25,
	0xc9, 0x1c, 0x0e, 0x69, 0x9e, 0xc7, 0xa8, 0xf3, 0xb1, 0xe0, 0x85, 0x4a, 0x59, 0x12, 0xf3, 0x12,
	0x85, 0x48, 0xa7, 0xa8, 0x75, 0xad, 0xde, 0x73, 0xf7, 0xcf, 0x0b, 0xba, 0x2b, 0xd5, 0xb0, 0x6e,
	0x0e, 0x1d, 0xba, 0x91, 0x19, 0x36, 0x6a, 0xdd, 0xef, 0x26, 0x1c, 0xdc, 0xf9, 0x11, 0x21, 0x7e,
	0x2d, 0x50, 0x2a, 0x72, 0x01, 0xd6, 0x4c, 0xa9, 0x3c, 0xbe, 0x45, 0x35, 0xe3, 0xf5, 0x36, 0xff,
	0xf7, 0x9e, 0xdc, 0x37, 0xb5, 0xea, 0xbe, 0xd2, 0xec, 0x10, 0x66, 0xab, 0x98, 0xdc, 0x00, 0xf9,
	0x75, 0x1d, 0xed, 0xc3, 0xbf, 0x6e, 0x61, 0x6f, 0x6e, 0x41, 0x1e, 0x43, 0x5b, 0x60, 0x46, 0x55,
	0x5a, 0x62, 0x5c, 0x88, 0xcc, 0x31, 0xb5, 0xbb, 0xd6, 0x32, 0x37, 0x12, 0x19, 0xf9, 0x04, 0xbb,
	0x33, 0xa4, 0x53, 0x14, 0xd2, 0x69, 0x1d, 0x9b, 0x27, 0x56, 0xef, 0xcd, 0x5f, 0xcd, 0x5c, 0xb3,
	0xc2, 0x1d, 0xd4, 0xfd, 0x3e, 0x53, 0x62, 0x11, 0x2e, 0xd5, 0xd6, 0x4f, 0xbd, 0x7d, 0xe7, 0xd4,
	0x9d, 0x53, 0x68, 0xaf, 0xb7, 0x54, 0xa7, 0x9f, 0xe3, 0x42, 0x5b, 0xb8, 0x1f, 0x56, 0x21, 0x39,
	0x80, 0xed, 0x92, 0x66, 0x05, 0x36, 0x9f, 0x43, 0x0d, 0x4e, 0xb7, 0x5e, 0x1a, 0xdd, 0x12, 0xec,
	0xcd, 0xbd, 0xab, 0x49, 0x12, 0x45, 0x99, 0x4e, 0xb0, 0xd1, 0x58, 0xc2, 0xaa, 0x52, 0xa2, 0x90,
	0x29, 0x67, 0x8d, 0xd2, 0x12, 0x92, 0x0e, 0xec, 0xa5, 0x4c, 0x2a, 0xca, 0x26, 0xd8, 0xb8, 0xb2,
	0xc2, 0x84, 0x40, 0x6b, 0xc6, 0xa5, 0x72, 0x5a, 0x3a, 0xaf, 0xe3, 0x67, 0x9f, 0x01, 0x7e, 0xde,
	0x8f, 0x1c, 0xc2, 0xc3, 0x41, 0x14, 0x05, 0xf1, 0x95, 0x1f, 0x0d, 0x86, 0xfd, 0x78, 0xf4, 0xf1,
	0x3a, 0xf0, 0xdf, 0x5d, 0xbe, 0xbf, 0xf4, 0xfb, 0xf6, 0x7f, 0x64, 0x0f, 0x5a, 0xc1, 0xf0, 0x3a,
	0xb2, 0x0d, 0xb2, 0x0b, 0xe6, 0x85, 0x1f, 0xd9, 0x5b, 0x55, 0x6a, 0xe0, 0x9f, 0xf5, 0x6d, 0xb3,
	0x4a, 0x05, 0xa3, 0xc8, 0x6e, 0x11, 0x80, 0x9d, 0xbe, 0xff, 0xc1, 0x8f, 0x7c, 0x7b, 0xfb, 0x3c,
	0x87, 0xa3, 0x09, 0xbf, 0xbd, 0xc7, 0xf7, 0x73, 0xab, 0xfe, 0xf6, 0x83, 0xea, 0x85, 0x05, 0xc6,
	0xcd, 0xdb, 0x86, 0x9a, 0xf0, 0x8c, 0xb2, 0xc4, 0xe5, 0x22, 0xf1, 0x12, 0x64, 0xfa, 0xfd, 0x79,
	0x75, 0x89, 0xe6, 0xa9, 0xfc, 0xdd, 0xbf, 0xc0, 0x6b, 0x8d, 0xc6, 0x3b, 0x9a, 0xfb, 0xe2, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x5d, 0x40, 0x26, 0x30, 0x04, 0x00, 0x00,
}
