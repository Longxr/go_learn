# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: arith.proto

import sys

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

_b = sys.version_info[0] < 3 and (
    lambda x: x) or (lambda x: x.encode('latin1'))
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor.FileDescriptor(
    name='arith.proto',
    package='pb',
    syntax='proto3',
    serialized_options=None,
    serialized_pb=_b('\n\x0b\x61rith.proto\x12\x02pb\"$\n\x0c\x41rithRequest\x12\t\n\x01\x61\x18\x01 \x01(\x05\x12\t\n\x01\x62\x18\x02 \x01(\x05\"6\n\rArithResponse\x12\x0b\n\x03pro\x18\x01 \x01(\x05\x12\x0b\n\x03quo\x18\x02 \x01(\x05\x12\x0b\n\x03rem\x18\x03 \x01(\x05\x32n\n\x0c\x41rithService\x12/\n\x08Multiply\x12\x10.pb.ArithRequest\x1a\x11.pb.ArithResponse\x12-\n\x06\x44ivide\x12\x10.pb.ArithRequest\x1a\x11.pb.ArithResponseb\x06proto3')
)


_ARITHREQUEST = _descriptor.Descriptor(
    name='ArithRequest',
    full_name='pb.ArithRequest',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='a', full_name='pb.ArithRequest.a', index=0,
            number=1, type=5, cpp_type=1, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='b', full_name='pb.ArithRequest.b', index=1,
            number=2, type=5, cpp_type=1, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=19,
    serialized_end=55,
)


_ARITHRESPONSE = _descriptor.Descriptor(
    name='ArithResponse',
    full_name='pb.ArithResponse',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='pro', full_name='pb.ArithResponse.pro', index=0,
            number=1, type=5, cpp_type=1, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='quo', full_name='pb.ArithResponse.quo', index=1,
            number=2, type=5, cpp_type=1, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='rem', full_name='pb.ArithResponse.rem', index=2,
            number=3, type=5, cpp_type=1, label=1,
            has_default_value=False, default_value=0,
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=57,
    serialized_end=111,
)

DESCRIPTOR.message_types_by_name['ArithRequest'] = _ARITHREQUEST
DESCRIPTOR.message_types_by_name['ArithResponse'] = _ARITHRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ArithRequest = _reflection.GeneratedProtocolMessageType('ArithRequest', (_message.Message,), dict(
    DESCRIPTOR=_ARITHREQUEST,
    __module__='arith_pb2'
    # @@protoc_insertion_point(class_scope:pb.ArithRequest)
))
_sym_db.RegisterMessage(ArithRequest)

ArithResponse = _reflection.GeneratedProtocolMessageType('ArithResponse', (_message.Message,), dict(
    DESCRIPTOR=_ARITHRESPONSE,
    __module__='arith_pb2'
    # @@protoc_insertion_point(class_scope:pb.ArithResponse)
))
_sym_db.RegisterMessage(ArithResponse)


_ARITHSERVICE = _descriptor.ServiceDescriptor(
    name='ArithService',
    full_name='pb.ArithService',
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    serialized_start=113,
    serialized_end=223,
    methods=[
        _descriptor.MethodDescriptor(
            name='Multiply',
            full_name='pb.ArithService.Multiply',
            index=0,
            containing_service=None,
            input_type=_ARITHREQUEST,
            output_type=_ARITHRESPONSE,
            serialized_options=None,
        ),
        _descriptor.MethodDescriptor(
            name='Divide',
            full_name='pb.ArithService.Divide',
            index=1,
            containing_service=None,
            input_type=_ARITHREQUEST,
            output_type=_ARITHRESPONSE,
            serialized_options=None,
        ),
    ])
_sym_db.RegisterServiceDescriptor(_ARITHSERVICE)

DESCRIPTOR.services_by_name['ArithService'] = _ARITHSERVICE

# @@protoc_insertion_point(module_scope)
