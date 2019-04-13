# Requirement

* Install the protobuf latest [protobuf compiler](https://github.com/protocolbuffers/protobuf/releases) and add the bin folder to your path.
* You should also have a working go installation, and install the go protobuf extension by running `go get github.com/golang/protobuf/protoc-gen-go`. Confirmed to be working is go version 1.12 +.

# Updating protobuf

Whenever you update the protobuf definition, please run `update.sh` for your own sake. This will generate the source files based on the proto file for each project that depends on it. 

# Setting up protobuf

To use the definition in a project, you'll likely need to generate the source files for your project using `protoc` and output the generated files to wherever you have them. The call for this should be added to `update.sh` and don't forget to update the `.gitignore`. 
