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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger Walletsuro",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "bevzenko.sergey@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localshot",
  "basePath": "/v1",
  "paths": {
    "/create": {
      "post": {
        "tags": [
          "wallet"
        ],
        "summary": "Create new wallet account",
        "operationId": "createWallet",
        "parameters": [
          {
            "type": "integer",
            "name": "user_id",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/deposit": {
      "post": {
        "description": "Deposit money to the wallet",
        "tags": [
          "wallet"
        ],
        "summary": "Deposit money to the wallet",
        "operationId": "deposit",
        "parameters": [
          {
            "type": "integer",
            "name": "user_id",
            "in": "header",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "amount": {
                  "type": "integer"
                },
                "wallet_id": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/info": {
      "get": {
        "tags": [
          "wallet"
        ],
        "summary": "info",
        "operationId": "info",
        "responses": {
          "200": {
            "description": "ok"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Wallets",
      "name": "wallet",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger Walletsuro",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "bevzenko.sergey@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localshot",
  "basePath": "/v1",
  "paths": {
    "/create": {
      "post": {
        "tags": [
          "wallet"
        ],
        "summary": "Create new wallet account",
        "operationId": "createWallet",
        "parameters": [
          {
            "type": "integer",
            "name": "user_id",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/deposit": {
      "post": {
        "description": "Deposit money to the wallet",
        "tags": [
          "wallet"
        ],
        "summary": "Deposit money to the wallet",
        "operationId": "deposit",
        "parameters": [
          {
            "type": "integer",
            "name": "user_id",
            "in": "header",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "amount": {
                  "type": "integer"
                },
                "wallet_id": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/info": {
      "get": {
        "tags": [
          "wallet"
        ],
        "summary": "info",
        "operationId": "info",
        "responses": {
          "200": {
            "description": "ok"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Wallets",
      "name": "wallet",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    }
  ]
}`))
}