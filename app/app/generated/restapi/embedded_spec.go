// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Otus architecture homework 3",
    "title": "Otus-Arch-homework",
    "version": "1.0.0"
  },
  "paths": {
    "/users/me": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get user info",
        "operationId": "GetUserInfo",
        "parameters": [
          {
            "$ref": "#/parameters/XUserEmail"
          },
          {
            "$ref": "#/parameters/XUserFirstName"
          },
          {
            "$ref": "#/parameters/XUserLastName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FailResponse": {
      "description": "Error response",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "description": "Error message",
          "type": "string",
          "minLength": 1
        }
      }
    },
    "UserProperties": {
      "description": "User info",
      "type": "object",
      "required": [
        "Email",
        "FirstName",
        "LastName"
      ],
      "properties": {
        "Email": {
          "description": "User email address",
          "type": "string",
          "minLength": 1
        },
        "FirstName": {
          "description": "User first name",
          "type": "string"
        },
        "LastName": {
          "description": "User last name",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "XUserEmail": {
      "type": "string",
      "name": "X-User-Email",
      "in": "header",
      "required": true
    },
    "XUserFirstName": {
      "type": "string",
      "name": "X-User-First-Name",
      "in": "header",
      "required": true
    },
    "XUserLastName": {
      "type": "string",
      "name": "X-User-Last-Name",
      "in": "header",
      "required": true
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Otus architecture homework 3",
    "title": "Otus-Arch-homework",
    "version": "1.0.0"
  },
  "paths": {
    "/users/me": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get user info",
        "operationId": "GetUserInfo",
        "parameters": [
          {
            "type": "string",
            "name": "X-User-Email",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "X-User-First-Name",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "X-User-Last-Name",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FailResponse": {
      "description": "Error response",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "description": "Error message",
          "type": "string",
          "minLength": 1
        }
      }
    },
    "UserProperties": {
      "description": "User info",
      "type": "object",
      "required": [
        "Email",
        "FirstName",
        "LastName"
      ],
      "properties": {
        "Email": {
          "description": "User email address",
          "type": "string",
          "minLength": 1
        },
        "FirstName": {
          "description": "User first name",
          "type": "string"
        },
        "LastName": {
          "description": "User last name",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "XUserEmail": {
      "type": "string",
      "name": "X-User-Email",
      "in": "header",
      "required": true
    },
    "XUserFirstName": {
      "type": "string",
      "name": "X-User-First-Name",
      "in": "header",
      "required": true
    },
    "XUserLastName": {
      "type": "string",
      "name": "X-User-Last-Name",
      "in": "header",
      "required": true
    }
  }
}`))
}
