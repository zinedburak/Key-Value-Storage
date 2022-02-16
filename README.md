# Key-Value-Storage
[![Go](https://github.com/zinedburak/Key_Value_Storage/actions/workflows/go.yml/badge.svg)](https://github.com/zinedburak/Key_Value_Storage/actions/workflows/go.yml)

This is a simple In memory key value storage application that is written in Golang


# How To Run The Application
- Make sure that you have docker installed
- Make sure the port 8000 is unoccupied in docker
- Open a Terminal in the project Folder
- Simply run `docker-compose up` command

# General Flow of The Application
The Application consists of three main parts which are Controllers, Routers, Models.
The models part consist of the two different models which are Storage and KeyValue
The Storage model is the main model for database it has two attributes which are keyValue Map and the path that the database is stored
This model is created with the CreateStore functionality which takes two parameters DbName and saveInterval. Dbname is used to create a file
in the current folder that the application is. The Save interval is used to create an internal clock that each n'th passes since the application has started the application 
will save the in memory key value pairs to the created file. Since the keyValue attribute of the Store is not exported to other files it has getter and setter functions to use it on the endpoints.
It also has two functions for internal use which are loadDataFromFile and saveDataToFile. The loadDataFromFile function loads data from the database file if the file exists.
The saveDataToFile is used to read the current memory and store the key value pairs in the memory to the file. There are also 4 different endpoints that is explained below section



     
# Different Endpoints
 - ## Get Key Value Pair 
   - A Function for getting the specific key value pair
   - The function will take a single parameter which is Key to
   - An example curl can be found below
   - ```bash
     curl --location --request GET 'http://127.0.0.1:8000/api/get-key-value' \--header 'Content-Type: application/json' \--data-raw '{"key": "ThirdKey"}' 
     ```
 - ## Set Key Value Pair 
   - A Function for setting a new key value pair if key does not exist and, if the key exist the value will be updated
   - This function will take a key value pair as a parameter
   - An example curl can be found below
   - ```bash
     curl --location --request POST 'http://127.0.0.1:8000/api/set-key-value' \--header 'Content-Type: application/json' \--data-raw '{"key": "First Key","value": "First Value"}'
     ```
   ## Get All Key Value Pair 
   - A Function for getting all  the key value pairs in the storage
   - This function will not get any parameters 
   - An example curl can be found below
   - ```bash
     curl --location --request GET 'http://127.0.0.1:8000/api/get-all-key-value' \--data-raw ''
     ```
   
    ## Flush All Data 
    - A function for saving all the stored data that is in the memory to the file 
    - This function will not get any parameters
    - An example curl can be found below
    - ```bash
      curl --location --request GET 'http://127.0.0.1:8000/api/save-all-key-value' \--header 'Content-Type: application/json' \--data-raw ''
      ```
#### 


