// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var quote_pb = require('./quote_pb.js');

function serialize_sale_CustomerId(arg) {
  if (!(arg instanceof quote_pb.CustomerId)) {
    throw new Error('Expected argument of type sale.CustomerId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sale_CustomerId(buffer_arg) {
  return quote_pb.CustomerId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sale_ProductRequest(arg) {
  if (!(arg instanceof quote_pb.ProductRequest)) {
    throw new Error('Expected argument of type sale.ProductRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sale_ProductRequest(buffer_arg) {
  return quote_pb.ProductRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sale_Quote(arg) {
  if (!(arg instanceof quote_pb.Quote)) {
    throw new Error('Expected argument of type sale.Quote');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sale_Quote(buffer_arg) {
  return quote_pb.Quote.deserializeBinary(new Uint8Array(buffer_arg));
}


var QuoteServiceService = exports.QuoteServiceService = {
  getQuote: {
    path: '/sale.QuoteService/GetQuote',
    requestStream: false,
    responseStream: false,
    requestType: quote_pb.CustomerId,
    responseType: quote_pb.Quote,
    requestSerialize: serialize_sale_CustomerId,
    requestDeserialize: deserialize_sale_CustomerId,
    responseSerialize: serialize_sale_Quote,
    responseDeserialize: deserialize_sale_Quote,
  },
  addProduct: {
    path: '/sale.QuoteService/AddProduct',
    requestStream: false,
    responseStream: false,
    requestType: quote_pb.ProductRequest,
    responseType: quote_pb.Quote,
    requestSerialize: serialize_sale_ProductRequest,
    requestDeserialize: deserialize_sale_ProductRequest,
    responseSerialize: serialize_sale_Quote,
    responseDeserialize: deserialize_sale_Quote,
  },
  removeProduct: {
    path: '/sale.QuoteService/RemoveProduct',
    requestStream: false,
    responseStream: false,
    requestType: quote_pb.ProductRequest,
    responseType: quote_pb.Quote,
    requestSerialize: serialize_sale_ProductRequest,
    requestDeserialize: deserialize_sale_ProductRequest,
    responseSerialize: serialize_sale_Quote,
    responseDeserialize: deserialize_sale_Quote,
  },
  updateQuantity: {
    path: '/sale.QuoteService/UpdateQuantity',
    requestStream: false,
    responseStream: false,
    requestType: quote_pb.ProductRequest,
    responseType: quote_pb.Quote,
    requestSerialize: serialize_sale_ProductRequest,
    requestDeserialize: deserialize_sale_ProductRequest,
    responseSerialize: serialize_sale_Quote,
    responseDeserialize: deserialize_sale_Quote,
  },
};

exports.QuoteServiceClient = grpc.makeGenericClientConstructor(QuoteServiceService);
