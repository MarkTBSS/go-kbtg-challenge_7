// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users/:id/wallets": {
            "get": {
                "description": "Get all user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_id"
                ],
                "summary": "Get all user_id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets": {
            "get": {
                "description": "Get all wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get all wallets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets/query": {
            "get": {
                "description": "Get all wallet_type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet_type"
                ],
                "summary": "Get all wallet_type",
                "parameters": [
                    {
                        "enum": [
                            "Savings",
                            "Credit Card",
                            "Crypto Wallet"
                        ],
                        "type": "string",
                        "description": "wallet type",
                        "name": "wallet_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wallet.Wallet"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wallet.Err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "wallet.Err": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "wallet.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 100
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-03-25T14:19:00.729237Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "wallet_name": {
                    "type": "string",
                    "example": "John's Wallet"
                },
                "wallet_type": {
                    "type": "string",
                    "example": "Create Card"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "Sophisticated Wallet API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
