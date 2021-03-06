// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/v1/job_service.proto

package dataproc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetJobRequest struct {
	// ID of the cluster to request a job from.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the job to return.
	//
	// To get a job ID make a [JobService.List] request.
	JobId                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobRequest) Reset()         { *m = GetJobRequest{} }
func (m *GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobRequest) ProtoMessage()    {}
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17500a3fdd75bbf9, []int{0}
}

func (m *GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobRequest.Unmarshal(m, b)
}
func (m *GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobRequest.Marshal(b, m, deterministic)
}
func (m *GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobRequest.Merge(m, src)
}
func (m *GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobRequest.Size(m)
}
func (m *GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobRequest proto.InternalMessageInfo

func (m *GetJobRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *GetJobRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ListJobsRequest struct {
	// ID of the cluster to list jobs for.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListJobsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListJobsResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters jobs listed in the response.
	//
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Job.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. The value. Must be 3-63 characters long and match the regular expression `^[a-z][-a-z0-9]{1,61}[a-z0-9].
	// Example of a filter: `name=my-job`.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17500a3fdd75bbf9, []int{1}
}

func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (m *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(m, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

func (m *ListJobsRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListJobsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListJobsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListJobsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListJobsResponse struct {
	// List of jobs for the specified cluster.
	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListJobsRequest.page_size], use `next_page_token` as the value
	// for the [ListJobsRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17500a3fdd75bbf9, []int{2}
}

func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (m *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(m, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

func (m *ListJobsResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *ListJobsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateJobRequest struct {
	// ID of the cluster to create a job for.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the job.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Specification for the job.
	//
	// Types that are valid to be assigned to JobSpec:
	//	*CreateJobRequest_MapreduceJob
	//	*CreateJobRequest_SparkJob
	//	*CreateJobRequest_PysparkJob
	//	*CreateJobRequest_HiveJob
	JobSpec              isCreateJobRequest_JobSpec `protobuf_oneof:"job_spec"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17500a3fdd75bbf9, []int{3}
}

func (m *CreateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobRequest.Unmarshal(m, b)
}
func (m *CreateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobRequest.Merge(m, src)
}
func (m *CreateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateJobRequest.Size(m)
}
func (m *CreateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobRequest proto.InternalMessageInfo

func (m *CreateJobRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isCreateJobRequest_JobSpec interface {
	isCreateJobRequest_JobSpec()
}

type CreateJobRequest_MapreduceJob struct {
	MapreduceJob *MapreduceJob `protobuf:"bytes,3,opt,name=mapreduce_job,json=mapreduceJob,proto3,oneof"`
}

type CreateJobRequest_SparkJob struct {
	SparkJob *SparkJob `protobuf:"bytes,4,opt,name=spark_job,json=sparkJob,proto3,oneof"`
}

type CreateJobRequest_PysparkJob struct {
	PysparkJob *PysparkJob `protobuf:"bytes,5,opt,name=pyspark_job,json=pysparkJob,proto3,oneof"`
}

type CreateJobRequest_HiveJob struct {
	HiveJob *HiveJob `protobuf:"bytes,6,opt,name=hive_job,json=hiveJob,proto3,oneof"`
}

func (*CreateJobRequest_MapreduceJob) isCreateJobRequest_JobSpec() {}

func (*CreateJobRequest_SparkJob) isCreateJobRequest_JobSpec() {}

func (*CreateJobRequest_PysparkJob) isCreateJobRequest_JobSpec() {}

func (*CreateJobRequest_HiveJob) isCreateJobRequest_JobSpec() {}

func (m *CreateJobRequest) GetJobSpec() isCreateJobRequest_JobSpec {
	if m != nil {
		return m.JobSpec
	}
	return nil
}

func (m *CreateJobRequest) GetMapreduceJob() *MapreduceJob {
	if x, ok := m.GetJobSpec().(*CreateJobRequest_MapreduceJob); ok {
		return x.MapreduceJob
	}
	return nil
}

func (m *CreateJobRequest) GetSparkJob() *SparkJob {
	if x, ok := m.GetJobSpec().(*CreateJobRequest_SparkJob); ok {
		return x.SparkJob
	}
	return nil
}

func (m *CreateJobRequest) GetPysparkJob() *PysparkJob {
	if x, ok := m.GetJobSpec().(*CreateJobRequest_PysparkJob); ok {
		return x.PysparkJob
	}
	return nil
}

func (m *CreateJobRequest) GetHiveJob() *HiveJob {
	if x, ok := m.GetJobSpec().(*CreateJobRequest_HiveJob); ok {
		return x.HiveJob
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateJobRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateJobRequest_MapreduceJob)(nil),
		(*CreateJobRequest_SparkJob)(nil),
		(*CreateJobRequest_PysparkJob)(nil),
		(*CreateJobRequest_HiveJob)(nil),
	}
}

type CreateJobMetadata struct {
	// ID of the cluster that the job is being created for.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the job being created.
	JobId                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateJobMetadata) Reset()         { *m = CreateJobMetadata{} }
func (m *CreateJobMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateJobMetadata) ProtoMessage()    {}
func (*CreateJobMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_17500a3fdd75bbf9, []int{4}
}

func (m *CreateJobMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateJobMetadata.Unmarshal(m, b)
}
func (m *CreateJobMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateJobMetadata.Marshal(b, m, deterministic)
}
func (m *CreateJobMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateJobMetadata.Merge(m, src)
}
func (m *CreateJobMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateJobMetadata.Size(m)
}
func (m *CreateJobMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateJobMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateJobMetadata proto.InternalMessageInfo

func (m *CreateJobMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateJobMetadata) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobRequest)(nil), "yandex.cloud.dataproc.v1.GetJobRequest")
	proto.RegisterType((*ListJobsRequest)(nil), "yandex.cloud.dataproc.v1.ListJobsRequest")
	proto.RegisterType((*ListJobsResponse)(nil), "yandex.cloud.dataproc.v1.ListJobsResponse")
	proto.RegisterType((*CreateJobRequest)(nil), "yandex.cloud.dataproc.v1.CreateJobRequest")
	proto.RegisterType((*CreateJobMetadata)(nil), "yandex.cloud.dataproc.v1.CreateJobMetadata")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/v1/job_service.proto", fileDescriptor_17500a3fdd75bbf9)
}

var fileDescriptor_17500a3fdd75bbf9 = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x4f, 0x13, 0x5d,
	0x14, 0x7e, 0x87, 0x96, 0xbe, 0xed, 0x01, 0x02, 0xef, 0x5d, 0x35, 0x0d, 0xe4, 0x85, 0xd1, 0xf0,
	0x51, 0x69, 0xa7, 0x53, 0x83, 0xf1, 0x03, 0x4c, 0xac, 0x31, 0x40, 0x23, 0x4a, 0xa6, 0xae, 0x24,
	0xa4, 0xb9, 0x33, 0x73, 0x2d, 0xb7, 0xb4, 0x73, 0xc7, 0xb9, 0xb7, 0x0d, 0x14, 0xd9, 0xb8, 0x31,
	0x61, 0xe1, 0xc6, 0xb5, 0x6b, 0xf7, 0x26, 0xfe, 0x06, 0xd8, 0xaa, 0x7f, 0xc1, 0x85, 0x6b, 0x97,
	0xae, 0xcc, 0xdc, 0x99, 0xd2, 0x8f, 0x30, 0x04, 0xd9, 0x9d, 0xdb, 0xf3, 0xcc, 0x73, 0x9e, 0xf3,
	0x59, 0xc8, 0x1e, 0x62, 0xc7, 0x26, 0x07, 0x9a, 0xd5, 0x60, 0x2d, 0x5b, 0xb3, 0xb1, 0xc0, 0xae,
	0xc7, 0x2c, 0xad, 0xad, 0x6b, 0x75, 0x66, 0x56, 0x39, 0xf1, 0xda, 0xd4, 0x22, 0x79, 0xd7, 0x63,
	0x82, 0xa1, 0x74, 0x80, 0xcd, 0x4b, 0x6c, 0xbe, 0x8b, 0xcd, 0xb7, 0xf5, 0xcc, 0x74, 0x8d, 0xb1,
	0x5a, 0x83, 0x68, 0xd8, 0xa5, 0x1a, 0x76, 0x1c, 0x26, 0xb0, 0xa0, 0xcc, 0xe1, 0xc1, 0x77, 0x19,
	0xf5, 0xb2, 0x18, 0x21, 0x66, 0x7e, 0x00, 0xc3, 0x5c, 0xe2, 0x49, 0x8a, 0x9e, 0x15, 0xe2, 0x66,
	0x06, 0x70, 0x6d, 0xdc, 0xa0, 0x76, 0xbf, 0x7b, 0x76, 0xc0, 0xed, 0xcb, 0x19, 0x22, 0x50, 0x31,
	0x4c, 0xac, 0x13, 0x51, 0x66, 0xa6, 0x41, 0x5e, 0xb7, 0x08, 0x17, 0xe8, 0x16, 0x80, 0xd5, 0x68,
	0x71, 0x41, 0xbc, 0x2a, 0xb5, 0xd3, 0xca, 0xac, 0xb2, 0x98, 0x2a, 0x8d, 0xff, 0x3c, 0xd5, 0x95,
	0x93, 0x33, 0x3d, 0xbe, 0xba, 0xb6, 0x52, 0x30, 0x52, 0xa1, 0x7f, 0xd3, 0x46, 0x37, 0x20, 0xe1,
	0xd7, 0x85, 0xda, 0xe9, 0x91, 0x0b, 0x80, 0xa3, 0x75, 0x66, 0x6e, 0xda, 0xea, 0x17, 0x05, 0x26,
	0x9f, 0x52, 0xee, 0x07, 0xe1, 0xd7, 0x8a, 0xb2, 0x00, 0x29, 0x17, 0xd7, 0x48, 0x95, 0xd3, 0x0e,
	0x91, 0x81, 0x62, 0x25, 0xf8, 0x7d, 0xaa, 0x27, 0x56, 0xd7, 0xf4, 0x42, 0xa1, 0x60, 0x24, 0x7d,
	0x67, 0x85, 0x76, 0x08, 0x5a, 0x04, 0x90, 0x40, 0xc1, 0xf6, 0x89, 0x93, 0x8e, 0x49, 0xd6, 0xd4,
	0xc9, 0x99, 0x3e, 0x2a, 0x91, 0x86, 0x64, 0x79, 0xe1, 0xfb, 0x90, 0x0a, 0x89, 0x57, 0xb4, 0x21,
	0x88, 0x97, 0x8e, 0x4b, 0x14, 0x9c, 0x9c, 0x9d, 0xf3, 0x85, 0x1e, 0xb5, 0x09, 0x53, 0x3d, 0xd9,
	0xdc, 0x65, 0x0e, 0x27, 0x48, 0x87, 0x78, 0x9d, 0x99, 0x3c, 0xad, 0xcc, 0xc6, 0x16, 0xc7, 0x8a,
	0x33, 0xf9, 0xa8, 0x11, 0xc8, 0xfb, 0x15, 0x95, 0x50, 0x34, 0x0f, 0x93, 0x0e, 0x39, 0x10, 0xd5,
	0x3e, 0x65, 0xb2, 0x58, 0xc6, 0x84, 0xff, 0xf3, 0x76, 0x57, 0x92, 0xfa, 0x31, 0x06, 0x53, 0x8f,
	0x3d, 0x82, 0x05, 0xb9, 0x6e, 0x37, 0x56, 0x20, 0xee, 0xe0, 0x26, 0x09, 0x7b, 0x31, 0xf7, 0xeb,
	0x54, 0x9f, 0x79, 0xb3, 0x83, 0x73, 0x9d, 0xdd, 0x9d, 0x1c, 0xce, 0x75, 0x0a, 0xb9, 0x7b, 0xbb,
	0x47, 0xfa, 0xf2, 0x1d, 0xfd, 0x78, 0x27, 0x7c, 0x19, 0x12, 0x8e, 0xb6, 0x60, 0xa2, 0x89, 0x5d,
	0x8f, 0xd8, 0x2d, 0x8b, 0x54, 0xeb, 0xcc, 0x94, 0x85, 0x1b, 0x2b, 0xce, 0x47, 0x27, 0xb7, 0xd5,
	0x85, 0x97, 0x99, 0xb9, 0xf1, 0x8f, 0x31, 0xde, 0xec, 0x7b, 0xa3, 0x47, 0x90, 0xe2, 0x2e, 0xf6,
	0xf6, 0x25, 0x55, 0x5c, 0x52, 0xa9, 0xd1, 0x54, 0x15, 0x1f, 0x1a, 0xd0, 0x24, 0x79, 0x68, 0xa3,
	0x75, 0x18, 0x73, 0x0f, 0x7b, 0x24, 0xa3, 0x92, 0xe4, 0x66, 0x34, 0xc9, 0xf6, 0x21, 0xef, 0xd1,
	0x80, 0x7b, 0xfe, 0x42, 0x0f, 0x21, 0xb9, 0x47, 0xdb, 0x41, 0x56, 0x09, 0xc9, 0x32, 0x17, 0xcd,
	0xb2, 0x41, 0xdb, 0x61, 0x42, 0xff, 0xee, 0x05, 0x66, 0x09, 0x20, 0x29, 0xf7, 0xde, 0x25, 0x96,
	0x8a, 0xe1, 0xbf, 0xf3, 0xf6, 0x6c, 0x11, 0x81, 0xfd, 0x4f, 0xff, 0xae, 0x3f, 0xff, 0x0f, 0x6d,
	0x4b, 0x72, 0x68, 0x53, 0x8a, 0x5f, 0x63, 0x00, 0x65, 0x66, 0x56, 0x82, 0x33, 0x83, 0xde, 0x2b,
	0x10, 0xf7, 0x27, 0x10, 0x2d, 0x45, 0x8b, 0x1e, 0x5a, 0xac, 0x4c, 0xf6, 0x2a, 0xd0, 0x60, 0x98,
	0x55, 0xed, 0xed, 0xf7, 0x1f, 0x1f, 0x46, 0x96, 0xd0, 0xc2, 0xc0, 0x11, 0x0a, 0xf5, 0x72, 0xed,
	0xa8, 0x97, 0xd9, 0xb1, 0x26, 0x47, 0xf9, 0x93, 0x02, 0x89, 0xa0, 0x06, 0xe8, 0x92, 0x38, 0xc3,
	0x43, 0x9c, 0x19, 0xaa, 0x79, 0xef, 0x04, 0x3d, 0xef, 0x5a, 0xea, 0xb3, 0xcf, 0xdf, 0xb2, 0xe9,
	0x8b, 0x0a, 0x1c, 0x2b, 0x33, 0x53, 0xca, 0x5c, 0x56, 0xaf, 0x2a, 0xf3, 0xbe, 0x92, 0x45, 0xef,
	0x14, 0x88, 0xad, 0x13, 0x81, 0x16, 0xa2, 0x65, 0x0e, 0x9c, 0xbd, 0xcc, 0xe5, 0xab, 0xac, 0xde,
	0x95, 0x1a, 0x8a, 0xa8, 0x70, 0x45, 0x0d, 0xda, 0x51, 0xd0, 0xe9, 0xe3, 0xd2, 0x3e, 0x4c, 0x0f,
	0x30, 0x63, 0x97, 0xf6, 0xb3, 0x97, 0xe2, 0xdb, 0x1b, 0xe5, 0xca, 0xcb, 0x27, 0x35, 0x2a, 0xf6,
	0x5a, 0x66, 0xde, 0x62, 0x4d, 0x2d, 0x80, 0xe7, 0x82, 0x9b, 0x5d, 0x63, 0xb9, 0x1a, 0x71, 0xe4,
	0xad, 0xd6, 0xa2, 0xfe, 0x37, 0x1e, 0x74, 0x6d, 0x33, 0x21, 0x81, 0xb7, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x06, 0x98, 0x3a, 0x65, 0xc7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	// Retrieves a list of jobs for a cluster.
	List(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	// Creates a job for a cluster.
	Create(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified job.
	Get(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*Job, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) List(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.JobService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) Create(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.JobService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) Get(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.v1.JobService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	// Retrieves a list of jobs for a cluster.
	List(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	// Creates a job for a cluster.
	Create(context.Context, *CreateJobRequest) (*operation.Operation, error)
	// Returns the specified job.
	Get(context.Context, *GetJobRequest) (*Job, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) List(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedJobServiceServer) Create(ctx context.Context, req *CreateJobRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedJobServiceServer) Get(ctx context.Context, req *GetJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.JobService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).List(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.JobService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Create(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.v1.JobService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Get(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.v1.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _JobService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _JobService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _JobService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/v1/job_service.proto",
}
