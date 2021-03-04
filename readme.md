## Import from NASA APOD API
API documentation can be found [here](https://api.nasa.gov)

## Purpose
The main purpose of this application is to try Go with different routine tasks like:
1. REST API's
2. Databases
3. Google Cloud
4. Work with files
5. Bulk import

## Notes
This is only the first step of more general idea: import, translation, building frontend etc.  
As for now there is only importer logic with db connector.  
All db related logic has been placed in db package, just for simplicity.  
All modules have been placed into the same repository for simplicity.  

## Requirements
* Golang
* MongoDB

## Executing
To execute import process use ./dailyimport/dailyimport
