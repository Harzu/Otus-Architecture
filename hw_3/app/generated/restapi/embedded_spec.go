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
    "/users": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get users list",
        "operationId": "UserList",
        "responses": {
          "200": {
            "description": "Users list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserProperties"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Update user",
        "operationId": "UpdateUser",
        "parameters": [
          {
            "description": "User params",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User updated"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Create new user",
        "operationId": "CreateUser",
        "parameters": [
          {
            "description": "New user params",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete users",
        "operationId": "DeleteUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "User emails",
              "type": "array",
              "items": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Users deleted"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
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
        "email",
        "firstName",
        "lastName",
        "passwordHash"
      ],
      "properties": {
        "email": {
          "description": "User email address",
          "type": "string",
          "minLength": 1
        },
        "firstName": {
          "description": "User first name",
          "type": "string"
        },
        "lastName": {
          "description": "User last name",
          "type": "string"
        },
        "passwordHash": {
          "description": "User password hash",
          "type": "string"
        }
      }
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
    "/users": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Get users list",
        "operationId": "UserList",
        "responses": {
          "200": {
            "description": "Users list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/UserProperties"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Update user",
        "operationId": "UpdateUser",
        "parameters": [
          {
            "description": "User params",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User updated"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Create new user",
        "operationId": "CreateUser",
        "parameters": [
          {
            "description": "New user params",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserProperties"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Delete users",
        "operationId": "DeleteUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "User emails",
              "type": "array",
              "items": {
                "type": "string",
                "minLength": 1
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Users deleted"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/FailResponse"
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
        "email",
        "firstName",
        "lastName",
        "passwordHash"
      ],
      "properties": {
        "email": {
          "description": "User email address",
          "type": "string",
          "minLength": 1
        },
        "firstName": {
          "description": "User first name",
          "type": "string"
        },
        "lastName": {
          "description": "User last name",
          "type": "string"
        },
        "passwordHash": {
          "description": "User password hash",
          "type": "string"
        }
      }
    }
  }
}`))
}