## Go Project Cli

It helps to create structured service for go project. It creates a module with the following structure.
`
├── service
│ └── moduleOne.model.go
│ └── moduleOne.repo.go
│ └── moduleTwo.model.go
│ └── moduleTwo.repo.go
│ └── service
│ └── moduleOne.go
│ └── moduleTwo.go

`

## Installation

`  
go build
go install

`

## Usage

**_Create a service_**
`go-project service <service-name> <module-name> <module-name> <module-name> ...`

**_Create a module_**
`go-project module <module-name> model repo service etc ...`
