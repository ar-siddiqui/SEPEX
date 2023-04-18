// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Seth Lawler",
            "email": "slawler@dewberry.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "[LandingPage Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_landing_page)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Landing Page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/conformance": {
            "get": {
                "description": "[Conformance Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_conformance_classes)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "API Conformance List",
                "responses": {
                    "200": {
                        "description": "hello:[\"dolly\"]",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/jobs": {
            "get": {
                "description": "[Job List Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_retrieve_job_results)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Summary of all (cached) Jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/jobs/{jobID}": {
            "get": {
                "description": "[Job Results Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_retrieve_job_results)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Job Results",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "[Dismss Job Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#ats_dismiss)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Dismiss Job",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/jobs/{jobID}/logs": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Job Logs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jobs.JobLog"
                        }
                    }
                }
            }
        },
        "/processes": {
            "get": {
                "description": "[Process List Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_process_list)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "processes"
                ],
                "summary": "List Available Processes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/processes/{processID}": {
            "get": {
                "description": "[Process Description Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_process_description)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "processes"
                ],
                "summary": "Describe Process Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "processID",
                        "name": "processID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/processes/{processID}/execution": {
            "post": {
                "description": "[Execute Process Specification](https://docs.ogc.org/is/18-062r2/18-062r2.html#sc_create_job)",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "processes"
                ],
                "summary": "Execute Process",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "jobs.JobLog": {
            "type": "object",
            "properties": {
                "api_log": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "container_log": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    },
    "externalDocs": {
        "description": "Schemas",
        "url": "http://schemas.opengis.net/ogcapi/processes/part1/1.0/openapi/schemas/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "dev-3.5.23",
	Host:             "localhost:5050",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Process-API Server",
	Description:      "An OGC compliant(ish) process server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
