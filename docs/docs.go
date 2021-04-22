// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://www.abtasty.com/solutions-product-teams/",
            "email": "support@flagship.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/activate": {
            "post": {
                "description": "Activate a campaign for a visitor ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaigns"
                ],
                "summary": "Activate a campaign",
                "operationId": "activate",
                "parameters": [
                    {
                        "description": "Campaign activation request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.activateBody"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/campaigns": {
            "post": {
                "description": "Get all campaigns value and metadata for a visitor ID and context",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaigns"
                ],
                "summary": "Get all campaigns for the visitor",
                "operationId": "get-campaigns",
                "parameters": [
                    {
                        "description": "Campaigns request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.campaignsBodySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CampaignsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/campaigns/{id}": {
            "post": {
                "description": "Get a single campaign value and metadata for a visitor ID and context",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaigns"
                ],
                "summary": "Get a single campaigns for the visitor",
                "operationId": "get-campaign",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Campaign ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Campaign request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.campaignsBodySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Campaign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/flags": {
            "post": {
                "description": "Get all flags value and metadata for a visitor ID and context",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flags"
                ],
                "summary": "Get all flags",
                "operationId": "get-flags",
                "parameters": [
                    {
                        "description": "Flag request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.campaignsBodySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/handlers.FlagInfos"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/flags/{key}": {
            "post": {
                "description": "Get a single flag value and metadata for a visitor ID and context",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flags"
                ],
                "summary": "Get flag by name",
                "operationId": "get-flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Flag key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Flag request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.campaignsBodySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.FlagInfos"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/flags/{key}/activate": {
            "post": {
                "description": "Activate a flag by its key for a visitor ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flags"
                ],
                "summary": "Activate a flag key",
                "operationId": "activate-flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Flag key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Flag activation request body",
                        "name": "flagActivation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.flagActivateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.flagActivated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/flags/{key}/value": {
            "post": {
                "description": "Get a single flag value for a visitor ID and context",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flags"
                ],
                "summary": "Get flag value by name",
                "operationId": "get-flag-value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Flag key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Flag request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.campaignsBodySwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.HTTPError"
                        }
                    }
                }
            }
        },
        "/health": {
            "post": {
                "description": "Get current health status for the API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Get health status",
                "operationId": "health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.FlagInfos"
                        }
                    }
                }
            }
        },
        "/hits": {
            "post": {
                "description": "Send a hit to Flagship datacollect",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/gif"
                ],
                "tags": [
                    "Hits"
                ],
                "summary": "Send a hit",
                "operationId": "send-hit",
                "parameters": [
                    {
                        "description": "Hit request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httputils.hit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CampaignsResponse": {
            "type": "object",
            "properties": {
                "campaigns": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Campaign"
                    }
                },
                "panic": {
                    "type": "boolean"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        },
        "handlers.FlagInfos": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/handlers.FlagMetadata"
                },
                "value": {
                    "type": "object"
                }
            }
        },
        "handlers.FlagMetadata": {
            "type": "object",
            "properties": {
                "campaignId": {
                    "type": "string"
                },
                "variationGroupID": {
                    "type": "string"
                },
                "variationID": {
                    "type": "string"
                }
            }
        },
        "handlers.activateBody": {
            "type": "object",
            "required": [
                "caid",
                "cid",
                "vaid",
                "vid"
            ],
            "properties": {
                "caid": {
                    "type": "string"
                },
                "cid": {
                    "type": "string"
                },
                "vaid": {
                    "type": "string"
                },
                "vid": {
                    "type": "string"
                }
            }
        },
        "handlers.campaignsBodyContextSwagger": {
            "type": "object",
            "properties": {
                "key_bool": {
                    "type": "boolean"
                },
                "key_number": {
                    "type": "number"
                },
                "key_string": {
                    "type": "string"
                }
            }
        },
        "handlers.campaignsBodySwagger": {
            "type": "object",
            "required": [
                "context",
                "visitor_id"
            ],
            "properties": {
                "context": {
                    "$ref": "#/definitions/handlers.campaignsBodyContextSwagger"
                },
                "trigger_hit": {
                    "type": "boolean"
                },
                "visitor_id": {
                    "type": "string"
                }
            }
        },
        "handlers.flagActivateBody": {
            "type": "object",
            "required": [
                "visitorId"
            ],
            "properties": {
                "visitorId": {
                    "type": "string"
                }
            }
        },
        "handlers.flagActivated": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "httputils.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "httputils.hit": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "string"
                },
                "t": {
                    "type": "string"
                },
                "vid": {
                    "type": "string"
                }
            }
        },
        "model.Campaign": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "variation": {
                    "$ref": "#/definitions/model.ClientVariation"
                },
                "variationGroupId": {
                    "type": "string"
                }
            }
        },
        "model.ClientVariation": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "modifications": {
                    "$ref": "#/definitions/model.Modification"
                },
                "reference": {
                    "type": "boolean"
                }
            }
        },
        "model.Modification": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Flagship Decision Host",
	Description: "This is the Flagship Decision Host API documentation",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
