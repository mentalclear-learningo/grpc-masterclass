# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: blog.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nblog.proto\x12\x04\x62log\"E\n\x04\x42log\x12\n\n\x02id\x18\x01 \x01(\t\x12\x11\n\tauthor_id\x18\x02 \x01(\t\x12\r\n\x05title\x18\x03 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\"-\n\x11\x43reateBlogRequest\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog\".\n\x12\x43reateBlogResponse\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog\"\"\n\x0fReadBlogRequest\x12\x0f\n\x07\x62log_id\x18\x01 \x01(\t\",\n\x10ReadBlogResponse\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog\"-\n\x11UpdateBlogRequest\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog\".\n\x12UpdateBlogResponse\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog\"$\n\x11\x44\x65leteBlogRequest\x12\x0f\n\x07\x62log_id\x18\x01 \x01(\t\"%\n\x12\x44\x65leteBlogResponse\x12\x0f\n\x07\x64\x65leted\x18\x01 \x01(\t\"\x11\n\x0fListBlogRequest\",\n\x10ListBlogResponse\x12\x18\n\x04\x62log\x18\x01 \x01(\x0b\x32\n.blog.Blog2\xd2\x02\n\x0b\x42logService\x12\x41\n\nCreateBlog\x12\x17.blog.CreateBlogRequest\x1a\x18.blog.CreateBlogResponse\"\x00\x12;\n\x08ReadBlog\x12\x15.blog.ReadBlogRequest\x1a\x16.blog.ReadBlogResponse\"\x00\x12\x41\n\nUpdateBlog\x12\x17.blog.UpdateBlogRequest\x1a\x18.blog.UpdateBlogResponse\"\x00\x12\x41\n\nDeleteBlog\x12\x17.blog.DeleteBlogRequest\x1a\x18.blog.DeleteBlogResponse\"\x00\x12=\n\x08ListBlog\x12\x15.blog.ListBlogRequest\x1a\x16.blog.ListBlogResponse\"\x00\x30\x01\x42\x0fZ\r./blog/blogpbb\x06proto3')



_BLOG = DESCRIPTOR.message_types_by_name['Blog']
_CREATEBLOGREQUEST = DESCRIPTOR.message_types_by_name['CreateBlogRequest']
_CREATEBLOGRESPONSE = DESCRIPTOR.message_types_by_name['CreateBlogResponse']
_READBLOGREQUEST = DESCRIPTOR.message_types_by_name['ReadBlogRequest']
_READBLOGRESPONSE = DESCRIPTOR.message_types_by_name['ReadBlogResponse']
_UPDATEBLOGREQUEST = DESCRIPTOR.message_types_by_name['UpdateBlogRequest']
_UPDATEBLOGRESPONSE = DESCRIPTOR.message_types_by_name['UpdateBlogResponse']
_DELETEBLOGREQUEST = DESCRIPTOR.message_types_by_name['DeleteBlogRequest']
_DELETEBLOGRESPONSE = DESCRIPTOR.message_types_by_name['DeleteBlogResponse']
_LISTBLOGREQUEST = DESCRIPTOR.message_types_by_name['ListBlogRequest']
_LISTBLOGRESPONSE = DESCRIPTOR.message_types_by_name['ListBlogResponse']
Blog = _reflection.GeneratedProtocolMessageType('Blog', (_message.Message,), {
  'DESCRIPTOR' : _BLOG,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.Blog)
  })
_sym_db.RegisterMessage(Blog)

CreateBlogRequest = _reflection.GeneratedProtocolMessageType('CreateBlogRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEBLOGREQUEST,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.CreateBlogRequest)
  })
_sym_db.RegisterMessage(CreateBlogRequest)

CreateBlogResponse = _reflection.GeneratedProtocolMessageType('CreateBlogResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEBLOGRESPONSE,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.CreateBlogResponse)
  })
_sym_db.RegisterMessage(CreateBlogResponse)

ReadBlogRequest = _reflection.GeneratedProtocolMessageType('ReadBlogRequest', (_message.Message,), {
  'DESCRIPTOR' : _READBLOGREQUEST,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.ReadBlogRequest)
  })
_sym_db.RegisterMessage(ReadBlogRequest)

ReadBlogResponse = _reflection.GeneratedProtocolMessageType('ReadBlogResponse', (_message.Message,), {
  'DESCRIPTOR' : _READBLOGRESPONSE,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.ReadBlogResponse)
  })
_sym_db.RegisterMessage(ReadBlogResponse)

UpdateBlogRequest = _reflection.GeneratedProtocolMessageType('UpdateBlogRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEBLOGREQUEST,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.UpdateBlogRequest)
  })
_sym_db.RegisterMessage(UpdateBlogRequest)

UpdateBlogResponse = _reflection.GeneratedProtocolMessageType('UpdateBlogResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEBLOGRESPONSE,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.UpdateBlogResponse)
  })
_sym_db.RegisterMessage(UpdateBlogResponse)

DeleteBlogRequest = _reflection.GeneratedProtocolMessageType('DeleteBlogRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEBLOGREQUEST,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.DeleteBlogRequest)
  })
_sym_db.RegisterMessage(DeleteBlogRequest)

DeleteBlogResponse = _reflection.GeneratedProtocolMessageType('DeleteBlogResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEBLOGRESPONSE,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.DeleteBlogResponse)
  })
_sym_db.RegisterMessage(DeleteBlogResponse)

ListBlogRequest = _reflection.GeneratedProtocolMessageType('ListBlogRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTBLOGREQUEST,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.ListBlogRequest)
  })
_sym_db.RegisterMessage(ListBlogRequest)

ListBlogResponse = _reflection.GeneratedProtocolMessageType('ListBlogResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTBLOGRESPONSE,
  '__module__' : 'blog_pb2'
  # @@protoc_insertion_point(class_scope:blog.ListBlogResponse)
  })
_sym_db.RegisterMessage(ListBlogResponse)

_BLOGSERVICE = DESCRIPTOR.services_by_name['BlogService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\r./blog/blogpb'
  _BLOG._serialized_start=20
  _BLOG._serialized_end=89
  _CREATEBLOGREQUEST._serialized_start=91
  _CREATEBLOGREQUEST._serialized_end=136
  _CREATEBLOGRESPONSE._serialized_start=138
  _CREATEBLOGRESPONSE._serialized_end=184
  _READBLOGREQUEST._serialized_start=186
  _READBLOGREQUEST._serialized_end=220
  _READBLOGRESPONSE._serialized_start=222
  _READBLOGRESPONSE._serialized_end=266
  _UPDATEBLOGREQUEST._serialized_start=268
  _UPDATEBLOGREQUEST._serialized_end=313
  _UPDATEBLOGRESPONSE._serialized_start=315
  _UPDATEBLOGRESPONSE._serialized_end=361
  _DELETEBLOGREQUEST._serialized_start=363
  _DELETEBLOGREQUEST._serialized_end=399
  _DELETEBLOGRESPONSE._serialized_start=401
  _DELETEBLOGRESPONSE._serialized_end=438
  _LISTBLOGREQUEST._serialized_start=440
  _LISTBLOGREQUEST._serialized_end=457
  _LISTBLOGRESPONSE._serialized_start=459
  _LISTBLOGRESPONSE._serialized_end=503
  _BLOGSERVICE._serialized_start=506
  _BLOGSERVICE._serialized_end=844
# @@protoc_insertion_point(module_scope)
