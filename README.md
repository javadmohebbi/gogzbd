# Golang-Bitdefender GravityZone API
**gobdgz** go-module helps GravityZone API developers who are using Google GoLang, integrate their apps easier.

### Integrations
Here is the integrated features which you can use in this module
- [x] **Accounts**: You can read [this guide](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=11&zoom=100,33,112) for more information
    - [x] **getAccountsList**: Get list of all Control Center accounts [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L72)
    - [x] **createAccount**: Create a new Control Center account [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L87)
    - [x] **deleteAccount**: Delete Control Center accounts using accountId [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L249)
    - [x] **updateAccount**: Update Control Center accounts using accountId [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L124)
    - [x] **getNotificationsSettings**: Get information about notifications settings for an account [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L167)
    - [x] **configureNotificationsSettings**: Configure/Update notifications settings for an account [(example)](https://github.com/javadmohebbi/gobdgz/blob/abef6a486f9f048261ee70d33246ebad2dd1f111/example/accounts/main.go#L193)

- [x] **Network**
    - [x] **getContainers**: Returns the network containers.
    - [x] **getNetworkInventoryItems**: Returns network inventory items.
    - [x] **createScanTask**: Returns true if the task was successfully created.
    - [x] **createReconfigureClientTask**: Creates a new Reconfigure Client task.
    - [x] **getScanTasksList**: Returns the list of scan tasks.
    - [x] **getEndpointsList**: Returns the list of endpoints.
    - [x] **getManagedEndpointDetails**: Returns the details about a managed endpoints
    - [x] **createCustomGroup**: Creates a new group under an existing one or under Computers and Groups.
    - [x] **deleteCustomGroup**: Deletes a custom group.
    - [x] **moveCustomGroup**: Moves a custom group under another custom group.
    - [x] **moveEndpoints**: Moves the specified list of endpoints to a custom group.
    - [x] **deleteEndpoint**: Deletes a specified endpoint.
    - [x] **setEndpointLabel**: Sets a label to an endpoint.
    - [ ] **createScanTaskByMac**: Generates scan tasks for managed endpoints identified by MAC address.
    - [ ] **assignPolicy**: This method is used to assign a policy template on the specified endpoints or containers.

- [x] **Policies**
  - [x] **Import** policy from one GravityZone console to another one. [(HOW TO GUIDE)](example/policy-export-import/)

# Get the module
Run this command to install **gobdgz** module on your system
```
go get github.com/javadmohebbi/gobdgz
```


# Documentation
If you are interested in reading this document, you might want to integrate **Bitdefender GravityZone** with your infrastructure. This module have been developing based on **Bitdefender GravityZone On-Premises API Guide** which is available in [this link](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf)

### Quick Start Guide
1. Get API Keys
- An API Key is required to communicate with GravityZone API service. To generate an API key you can read [this document](http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=7&zoom=100,33,85)

2. Get **gobdgz** module
    - Run this command to get the module: ```go get github.com/javadmohebbi/gobdgz```

3. Use these examples to get more information about how you can use this repo for integration
    - [Accounts examples](example/accounts)
    - [Networks examples](example/network)
