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
    "description": "Otus architecture homework cache API",
    "title": "Otus-Arch-homework",
    "version": "1.0.0"
  },
  "paths": {
    "/products": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Search products by params",
        "operationId": "GetProducts",
        "parameters": [
          {
            "$ref": "#/parameters/ETag"
          },
          {
            "$ref": "#/parameters/GetProductsRequest"
          }
        ],
        "responses": {
          "200": {
            "description": "Get products success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Products"
              }
            },
            "headers": {
              "ETag": {
                "type": "string"
              }
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
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
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "Products": {
      "type": "object",
      "required": [
        "type",
        "name",
        "price"
      ],
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1
        },
        "price": {
          "type": "integer"
        },
        "type": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  },
  "parameters": {
    "ETag": {
      "type": "string",
      "name": "Etag",
      "in": "header"
    },
    "GetProductsRequest": {
      "name": "Body",
      "in": "body",
      "required": true,
      "schema": {
        "type": "object",
        "required": [
          "limit",
          "type"
        ],
        "properties": {
          "limit": {
            "type": "integer",
            "minimum": 1
          },
          "type": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Otus architecture homework cache API",
    "title": "Otus-Arch-homework",
    "version": "1.0.0"
  },
  "paths": {
    "/products": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Search products by params",
        "operationId": "GetProducts",
        "parameters": [
          {
            "type": "string",
            "name": "Etag",
            "in": "header"
          },
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "limit",
                "type"
              ],
              "properties": {
                "limit": {
                  "type": "integer",
                  "minimum": 1
                },
                "type": {
                  "type": "string",
                  "minLength": 1
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Get products success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Products"
              }
            },
            "headers": {
              "ETag": {
                "type": "string"
              }
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/FailResponse"
            }
          },
          "500": {
            "description": "Internal server error",
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
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "Products": {
      "type": "object",
      "required": [
        "type",
        "name",
        "price"
      ],
      "properties": {
        "name": {
          "type": "string",
          "minLength": 1
        },
        "price": {
          "type": "integer",
          "minimum": 0
        },
        "type": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  },
  "parameters": {
    "ETag": {
      "type": "string",
      "name": "Etag",
      "in": "header"
    },
    "GetProductsRequest": {
      "name": "Body",
      "in": "body",
      "required": true,
      "schema": {
        "type": "object",
        "required": [
          "limit",
          "type"
        ],
        "properties": {
          "limit": {
            "type": "integer",
            "minimum": 1
          },
          "type": {
            "type": "string",
            "minLength": 1
          }
        }
      }
    }
  }
}`))
}