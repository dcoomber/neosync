// @generated by protoc-gen-connect-query v1.4.2 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/auth.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { CheckTokenRequest, CheckTokenResponse, GetAuthorizeUrlRequest, GetAuthorizeUrlResponse, GetAuthStatusRequest, GetAuthStatusResponse, GetCliIssuerRequest, GetCliIssuerResponse, LoginCliRequest, LoginCliResponse, RefreshCliRequest, RefreshCliResponse } from "./auth_pb.js";

/**
 * Used by the CLI to login to Neosync with OAuth.
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.LoginCli
 */
export const loginCli = {
  localName: "loginCli",
  name: "LoginCli",
  kind: MethodKind.Unary,
  I: LoginCliRequest,
  O: LoginCliResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;

/**
 * Used by the CLI to refresh an expired Neosync accesss token.
 * This should only be used if an access token was previously retrieved from the `LoginCli` or `RefreshCli` methods.
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.RefreshCli
 */
export const refreshCli = {
  localName: "refreshCli",
  name: "RefreshCli",
  kind: MethodKind.Unary,
  I: RefreshCliRequest,
  O: RefreshCliResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;

/**
 * Empty endpoint to simply check if the provided access token is valid
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.CheckToken
 */
export const checkToken = {
  localName: "checkToken",
  name: "CheckToken",
  kind: MethodKind.Unary,
  I: CheckTokenRequest,
  O: CheckTokenResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;

/**
 * Used by the CLI to retrieve Auth Issuer information
 * @deprecated
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.GetCliIssuer
 */
export const getCliIssuer = {
  localName: "getCliIssuer",
  name: "GetCliIssuer",
  kind: MethodKind.Unary,
  I: GetCliIssuerRequest,
  O: GetCliIssuerResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;

/**
 * Used by the CLI to retrieve an Authorize URL for use with OAuth login.
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.GetAuthorizeUrl
 */
export const getAuthorizeUrl = {
  localName: "getAuthorizeUrl",
  name: "GetAuthorizeUrl",
  kind: MethodKind.Unary,
  I: GetAuthorizeUrlRequest,
  O: GetAuthorizeUrlResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;

/**
 * Returns the auth status of the API server. Whether or not the backend has authentication enabled.
 * This is used by clients to make decisions on whether or not they should send access tokens to the API.
 *
 * @generated from rpc mgmt.v1alpha1.AuthService.GetAuthStatus
 */
export const getAuthStatus = {
  localName: "getAuthStatus",
  name: "GetAuthStatus",
  kind: MethodKind.Unary,
  I: GetAuthStatusRequest,
  O: GetAuthStatusResponse,
  service: {
    typeName: "mgmt.v1alpha1.AuthService"
  }
} as const;
