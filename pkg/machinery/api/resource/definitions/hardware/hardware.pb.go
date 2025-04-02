// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: resource/definitions/hardware/hardware.proto

package hardware

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MemoryModuleSpec represents a single Memory.
type MemoryModuleSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Size          uint32                 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	DeviceLocator string                 `protobuf:"bytes,2,opt,name=device_locator,json=deviceLocator,proto3" json:"device_locator,omitempty"`
	BankLocator   string                 `protobuf:"bytes,3,opt,name=bank_locator,json=bankLocator,proto3" json:"bank_locator,omitempty"`
	Speed         uint32                 `protobuf:"varint,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Manufacturer  string                 `protobuf:"bytes,5,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,6,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	AssetTag      string                 `protobuf:"bytes,7,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	ProductName   string                 `protobuf:"bytes,8,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemoryModuleSpec) Reset() {
	*x = MemoryModuleSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoryModuleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryModuleSpec) ProtoMessage() {}

func (x *MemoryModuleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryModuleSpec.ProtoReflect.Descriptor instead.
func (*MemoryModuleSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{0}
}

func (x *MemoryModuleSpec) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MemoryModuleSpec) GetDeviceLocator() string {
	if x != nil {
		return x.DeviceLocator
	}
	return ""
}

func (x *MemoryModuleSpec) GetBankLocator() string {
	if x != nil {
		return x.BankLocator
	}
	return ""
}

func (x *MemoryModuleSpec) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *MemoryModuleSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *MemoryModuleSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *MemoryModuleSpec) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *MemoryModuleSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

// PCIDeviceSpec represents a single processor.
type PCIDeviceSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Class         string                 `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Subclass      string                 `protobuf:"bytes,2,opt,name=subclass,proto3" json:"subclass,omitempty"`
	Vendor        string                 `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Product       string                 `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	ClassId       string                 `protobuf:"bytes,5,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	SubclassId    string                 `protobuf:"bytes,6,opt,name=subclass_id,json=subclassId,proto3" json:"subclass_id,omitempty"`
	VendorId      string                 `protobuf:"bytes,7,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,8,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PCIDeviceSpec) Reset() {
	*x = PCIDeviceSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PCIDeviceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCIDeviceSpec) ProtoMessage() {}

func (x *PCIDeviceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCIDeviceSpec.ProtoReflect.Descriptor instead.
func (*PCIDeviceSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{1}
}

func (x *PCIDeviceSpec) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *PCIDeviceSpec) GetSubclass() string {
	if x != nil {
		return x.Subclass
	}
	return ""
}

func (x *PCIDeviceSpec) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *PCIDeviceSpec) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *PCIDeviceSpec) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *PCIDeviceSpec) GetSubclassId() string {
	if x != nil {
		return x.SubclassId
	}
	return ""
}

func (x *PCIDeviceSpec) GetVendorId() string {
	if x != nil {
		return x.VendorId
	}
	return ""
}

func (x *PCIDeviceSpec) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// PCIDriverRebindConfigSpec describes PCI rebind configuration.
type PCIDriverRebindConfigSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pciid         string                 `protobuf:"bytes,1,opt,name=pciid,proto3" json:"pciid,omitempty"`
	TargetDriver  string                 `protobuf:"bytes,2,opt,name=target_driver,json=targetDriver,proto3" json:"target_driver,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PCIDriverRebindConfigSpec) Reset() {
	*x = PCIDriverRebindConfigSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PCIDriverRebindConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCIDriverRebindConfigSpec) ProtoMessage() {}

func (x *PCIDriverRebindConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCIDriverRebindConfigSpec.ProtoReflect.Descriptor instead.
func (*PCIDriverRebindConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{2}
}

func (x *PCIDriverRebindConfigSpec) GetPciid() string {
	if x != nil {
		return x.Pciid
	}
	return ""
}

func (x *PCIDriverRebindConfigSpec) GetTargetDriver() string {
	if x != nil {
		return x.TargetDriver
	}
	return ""
}

// PCIDriverRebindStatusSpec describes status of rebinded drivers.
type PCIDriverRebindStatusSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pciid         string                 `protobuf:"bytes,1,opt,name=pciid,proto3" json:"pciid,omitempty"`
	TargetDriver  string                 `protobuf:"bytes,2,opt,name=target_driver,json=targetDriver,proto3" json:"target_driver,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PCIDriverRebindStatusSpec) Reset() {
	*x = PCIDriverRebindStatusSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PCIDriverRebindStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCIDriverRebindStatusSpec) ProtoMessage() {}

func (x *PCIDriverRebindStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCIDriverRebindStatusSpec.ProtoReflect.Descriptor instead.
func (*PCIDriverRebindStatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{3}
}

func (x *PCIDriverRebindStatusSpec) GetPciid() string {
	if x != nil {
		return x.Pciid
	}
	return ""
}

func (x *PCIDriverRebindStatusSpec) GetTargetDriver() string {
	if x != nil {
		return x.TargetDriver
	}
	return ""
}

// ProcessorSpec represents a single processor.
type ProcessorSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Socket        string                 `protobuf:"bytes,1,opt,name=socket,proto3" json:"socket,omitempty"`
	Manufacturer  string                 `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName   string                 `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	MaxSpeed      uint32                 `protobuf:"varint,4,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty"`
	BootSpeed     uint32                 `protobuf:"varint,5,opt,name=boot_speed,json=bootSpeed,proto3" json:"boot_speed,omitempty"`
	Status        uint32                 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,7,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	AssetTag      string                 `protobuf:"bytes,8,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	PartNumber    string                 `protobuf:"bytes,9,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	CoreCount     uint32                 `protobuf:"varint,10,opt,name=core_count,json=coreCount,proto3" json:"core_count,omitempty"`
	CoreEnabled   uint32                 `protobuf:"varint,11,opt,name=core_enabled,json=coreEnabled,proto3" json:"core_enabled,omitempty"`
	ThreadCount   uint32                 `protobuf:"varint,12,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessorSpec) Reset() {
	*x = ProcessorSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessorSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorSpec) ProtoMessage() {}

func (x *ProcessorSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorSpec.ProtoReflect.Descriptor instead.
func (*ProcessorSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessorSpec) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *ProcessorSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *ProcessorSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProcessorSpec) GetMaxSpeed() uint32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *ProcessorSpec) GetBootSpeed() uint32 {
	if x != nil {
		return x.BootSpeed
	}
	return 0
}

func (x *ProcessorSpec) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProcessorSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *ProcessorSpec) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *ProcessorSpec) GetPartNumber() string {
	if x != nil {
		return x.PartNumber
	}
	return ""
}

func (x *ProcessorSpec) GetCoreCount() uint32 {
	if x != nil {
		return x.CoreCount
	}
	return 0
}

func (x *ProcessorSpec) GetCoreEnabled() uint32 {
	if x != nil {
		return x.CoreEnabled
	}
	return 0
}

func (x *ProcessorSpec) GetThreadCount() uint32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

// SystemInformationSpec represents the system information obtained from smbios.
type SystemInformationSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Manufacturer  string                 `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Version       string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	SerialNumber  string                 `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Uuid          string                 `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	WakeUpType    string                 `protobuf:"bytes,6,opt,name=wake_up_type,json=wakeUpType,proto3" json:"wake_up_type,omitempty"`
	SkuNumber     string                 `protobuf:"bytes,7,opt,name=sku_number,json=skuNumber,proto3" json:"sku_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SystemInformationSpec) Reset() {
	*x = SystemInformationSpec{}
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemInformationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInformationSpec) ProtoMessage() {}

func (x *SystemInformationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_hardware_hardware_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInformationSpec.ProtoReflect.Descriptor instead.
func (*SystemInformationSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_hardware_hardware_proto_rawDescGZIP(), []int{5}
}

func (x *SystemInformationSpec) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *SystemInformationSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SystemInformationSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SystemInformationSpec) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *SystemInformationSpec) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SystemInformationSpec) GetWakeUpType() string {
	if x != nil {
		return x.WakeUpType
	}
	return ""
}

func (x *SystemInformationSpec) GetSkuNumber() string {
	if x != nil {
		return x.SkuNumber
	}
	return ""
}

var File_resource_definitions_hardware_hardware_proto protoreflect.FileDescriptor

const file_resource_definitions_hardware_hardware_proto_rawDesc = "" +
	"\n" +
	",resource/definitions/hardware/hardware.proto\x12#talos.resource.definitions.hardware\"\x8f\x02\n" +
	"\x10MemoryModuleSpec\x12\x12\n" +
	"\x04size\x18\x01 \x01(\rR\x04size\x12%\n" +
	"\x0edevice_locator\x18\x02 \x01(\tR\rdeviceLocator\x12!\n" +
	"\fbank_locator\x18\x03 \x01(\tR\vbankLocator\x12\x14\n" +
	"\x05speed\x18\x04 \x01(\rR\x05speed\x12\"\n" +
	"\fmanufacturer\x18\x05 \x01(\tR\fmanufacturer\x12#\n" +
	"\rserial_number\x18\x06 \x01(\tR\fserialNumber\x12\x1b\n" +
	"\tasset_tag\x18\a \x01(\tR\bassetTag\x12!\n" +
	"\fproduct_name\x18\b \x01(\tR\vproductName\"\xeb\x01\n" +
	"\rPCIDeviceSpec\x12\x14\n" +
	"\x05class\x18\x01 \x01(\tR\x05class\x12\x1a\n" +
	"\bsubclass\x18\x02 \x01(\tR\bsubclass\x12\x16\n" +
	"\x06vendor\x18\x03 \x01(\tR\x06vendor\x12\x18\n" +
	"\aproduct\x18\x04 \x01(\tR\aproduct\x12\x19\n" +
	"\bclass_id\x18\x05 \x01(\tR\aclassId\x12\x1f\n" +
	"\vsubclass_id\x18\x06 \x01(\tR\n" +
	"subclassId\x12\x1b\n" +
	"\tvendor_id\x18\a \x01(\tR\bvendorId\x12\x1d\n" +
	"\n" +
	"product_id\x18\b \x01(\tR\tproductId\"V\n" +
	"\x19PCIDriverRebindConfigSpec\x12\x14\n" +
	"\x05pciid\x18\x01 \x01(\tR\x05pciid\x12#\n" +
	"\rtarget_driver\x18\x02 \x01(\tR\ftargetDriver\"V\n" +
	"\x19PCIDriverRebindStatusSpec\x12\x14\n" +
	"\x05pciid\x18\x01 \x01(\tR\x05pciid\x12#\n" +
	"\rtarget_driver\x18\x02 \x01(\tR\ftargetDriver\"\x8a\x03\n" +
	"\rProcessorSpec\x12\x16\n" +
	"\x06socket\x18\x01 \x01(\tR\x06socket\x12\"\n" +
	"\fmanufacturer\x18\x02 \x01(\tR\fmanufacturer\x12!\n" +
	"\fproduct_name\x18\x03 \x01(\tR\vproductName\x12\x1b\n" +
	"\tmax_speed\x18\x04 \x01(\rR\bmaxSpeed\x12\x1d\n" +
	"\n" +
	"boot_speed\x18\x05 \x01(\rR\tbootSpeed\x12\x16\n" +
	"\x06status\x18\x06 \x01(\rR\x06status\x12#\n" +
	"\rserial_number\x18\a \x01(\tR\fserialNumber\x12\x1b\n" +
	"\tasset_tag\x18\b \x01(\tR\bassetTag\x12\x1f\n" +
	"\vpart_number\x18\t \x01(\tR\n" +
	"partNumber\x12\x1d\n" +
	"\n" +
	"core_count\x18\n" +
	" \x01(\rR\tcoreCount\x12!\n" +
	"\fcore_enabled\x18\v \x01(\rR\vcoreEnabled\x12!\n" +
	"\fthread_count\x18\f \x01(\rR\vthreadCount\"\xf2\x01\n" +
	"\x15SystemInformationSpec\x12\"\n" +
	"\fmanufacturer\x18\x01 \x01(\tR\fmanufacturer\x12!\n" +
	"\fproduct_name\x18\x02 \x01(\tR\vproductName\x12\x18\n" +
	"\aversion\x18\x03 \x01(\tR\aversion\x12#\n" +
	"\rserial_number\x18\x04 \x01(\tR\fserialNumber\x12\x12\n" +
	"\x04uuid\x18\x05 \x01(\tR\x04uuid\x12 \n" +
	"\fwake_up_type\x18\x06 \x01(\tR\n" +
	"wakeUpType\x12\x1d\n" +
	"\n" +
	"sku_number\x18\a \x01(\tR\tskuNumberBz\n" +
	"+dev.talos.api.resource.definitions.hardwareZKgithub.com/siderolabs/talos/pkg/machinery/api/resource/definitions/hardwareb\x06proto3"

var (
	file_resource_definitions_hardware_hardware_proto_rawDescOnce sync.Once
	file_resource_definitions_hardware_hardware_proto_rawDescData []byte
)

func file_resource_definitions_hardware_hardware_proto_rawDescGZIP() []byte {
	file_resource_definitions_hardware_hardware_proto_rawDescOnce.Do(func() {
		file_resource_definitions_hardware_hardware_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resource_definitions_hardware_hardware_proto_rawDesc), len(file_resource_definitions_hardware_hardware_proto_rawDesc)))
	})
	return file_resource_definitions_hardware_hardware_proto_rawDescData
}

var file_resource_definitions_hardware_hardware_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resource_definitions_hardware_hardware_proto_goTypes = []any{
	(*MemoryModuleSpec)(nil),          // 0: talos.resource.definitions.hardware.MemoryModuleSpec
	(*PCIDeviceSpec)(nil),             // 1: talos.resource.definitions.hardware.PCIDeviceSpec
	(*PCIDriverRebindConfigSpec)(nil), // 2: talos.resource.definitions.hardware.PCIDriverRebindConfigSpec
	(*PCIDriverRebindStatusSpec)(nil), // 3: talos.resource.definitions.hardware.PCIDriverRebindStatusSpec
	(*ProcessorSpec)(nil),             // 4: talos.resource.definitions.hardware.ProcessorSpec
	(*SystemInformationSpec)(nil),     // 5: talos.resource.definitions.hardware.SystemInformationSpec
}
var file_resource_definitions_hardware_hardware_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resource_definitions_hardware_hardware_proto_init() }
func file_resource_definitions_hardware_hardware_proto_init() {
	if File_resource_definitions_hardware_hardware_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resource_definitions_hardware_hardware_proto_rawDesc), len(file_resource_definitions_hardware_hardware_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_hardware_hardware_proto_goTypes,
		DependencyIndexes: file_resource_definitions_hardware_hardware_proto_depIdxs,
		MessageInfos:      file_resource_definitions_hardware_hardware_proto_msgTypes,
	}.Build()
	File_resource_definitions_hardware_hardware_proto = out.File
	file_resource_definitions_hardware_hardware_proto_goTypes = nil
	file_resource_definitions_hardware_hardware_proto_depIdxs = nil
}
