/**
 * @fileoverview gRPC-Web generated client stub for ui
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var flow_flow_pb = require('../flow/flow_pb.js')

var ui_notifications_pb = require('../ui/notifications_pb.js')

var ui_status_pb = require('../ui/status_pb.js')
const proto = {};
proto.ui = require('./ui_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ui.UIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ui.UIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ui.GetEventsRequest,
 *   !proto.ui.GetEventsResponse>}
 */
const methodDescriptor_UI_GetEvents = new grpc.web.MethodDescriptor(
  '/ui.UI/GetEvents',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.ui.GetEventsRequest,
  proto.ui.GetEventsResponse,
  /**
   * @param {!proto.ui.GetEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.GetEventsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ui.GetEventsRequest,
 *   !proto.ui.GetEventsResponse>}
 */
const methodInfo_UI_GetEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ui.GetEventsResponse,
  /**
   * @param {!proto.ui.GetEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.GetEventsResponse.deserializeBinary
);


/**
 * @param {!proto.ui.GetEventsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.ui.GetEventsResponse>}
 *     The XHR Node Readable Stream
 */
proto.ui.UIClient.prototype.getEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ui.UI/GetEvents',
      request,
      metadata || {},
      methodDescriptor_UI_GetEvents);
};


/**
 * @param {!proto.ui.GetEventsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.ui.GetEventsResponse>}
 *     The XHR Node Readable Stream
 */
proto.ui.UIPromiseClient.prototype.getEvents =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/ui.UI/GetEvents',
      request,
      metadata || {},
      methodDescriptor_UI_GetEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ui.GetStatusRequest,
 *   !proto.ui.GetStatusResponse>}
 */
const methodDescriptor_UI_GetStatus = new grpc.web.MethodDescriptor(
  '/ui.UI/GetStatus',
  grpc.web.MethodType.UNARY,
  ui_status_pb.GetStatusRequest,
  ui_status_pb.GetStatusResponse,
  /**
   * @param {!proto.ui.GetStatusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  ui_status_pb.GetStatusResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ui.GetStatusRequest,
 *   !proto.ui.GetStatusResponse>}
 */
const methodInfo_UI_GetStatus = new grpc.web.AbstractClientBase.MethodInfo(
  ui_status_pb.GetStatusResponse,
  /**
   * @param {!proto.ui.GetStatusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  ui_status_pb.GetStatusResponse.deserializeBinary
);


/**
 * @param {!proto.ui.GetStatusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ui.GetStatusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ui.GetStatusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ui.UIClient.prototype.getStatus =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ui.UI/GetStatus',
      request,
      metadata || {},
      methodDescriptor_UI_GetStatus,
      callback);
};


/**
 * @param {!proto.ui.GetStatusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ui.GetStatusResponse>}
 *     A native promise that resolves to the response
 */
proto.ui.UIPromiseClient.prototype.getStatus =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ui.UI/GetStatus',
      request,
      metadata || {},
      methodDescriptor_UI_GetStatus);
};


module.exports = proto.ui;

