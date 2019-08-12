![marcio venancio](images/neo-data.png)

<br>

# Neo data - For data integration

This Project is an example for Data Integration. 
 
The project is an API built on top of Golang programming language which serves endpoints to consume and integrate Companies data information that can be added in two ways
* Importing CSV files
* Caling an Endpoint to create a new record.

## Further explanation

The API makes use of MongoDB as NoSql database to store the information.
The data, after entered, is processed and then, if already exist, the information is merged, otherwise, a new record is added.


## Project Structure
![marcio venancio](images/api_proj_structure.png)

## How to run
1. It is necessary to have set up Golang environment with all the environment variables. Below there is 
an example of 'go env' command from my Linux environment.
<br/>

![marcio venancio](images/golang-env.png) 

2. It is necessary as well to have the NoSql database MongoDb installed. Could be a MongoDb docker image or even a local instance, but it is necessary to run under the default port 27017.

3. Now, clone the project inside the work folder. For instance, based on the Golang environment variables above, the work folder is $GOPATH/src. After the project is cloned, the path should look like this $GOPATH/src/neo-data

4. Now, it is possible to run the project. You can install manually the dependencies executing on a terminal instance 
 * *go get gopkg.in/mgo.v2/bson*,
 * *go get github.com/gorilla/mux*
 * *go get github.com/stretchr/testify/assert*
 
   After that, it is possible to run the project with the command "go run main.go", then if everything is ok, the API will be listening in the http port 3000.
   
   In addition to that, you may want to compile and generate an executable, inside the neo-data folder, execute the command "**make all**". This will generate an exe inside $GOPATH/bin

## Project Structure
There are a couple of automated tests that invoke each end point in order to check if it is up and running.
<br/>

![marcio venancio](images/api_main_test.png)

<br/>

## Endpoints

* [POST] *Create* - **"/api/company"**  

* Sent Json *
```json
{    
    "name": "TESTCREATE_01",
    "zip": "12345",
    "website": "website"
}
```
> ![marcio venancio](images/api-create-001.png)
* [PUT] - *Update* - **"/api/company/{id}"**

*Sent json*
```json
{
    "id": "5d502bf8cad9901c5d085c04",
    "name": "TESTECREATE_01",
    "zip": "99999",
    "website": "website"
}
``` 
> ![marcio venancio](images/api-update-001.png)

* [DELETE] - *Delete* - **"/api/company/{id}"**
> ![marcio venancio](images/api-delete-001.png)
<br/><br/>

* [GET] - *Get company by ID* - **"/api/company/{id}"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-getbyid-001.png)
<br/><br/>

* [GET] - *Get company by name* - **"/api/company/name/{name}"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-getbyname-001.png)
<br/><br/>

* [GET] - *Get compay by name and zip code* - **"/api/company/name/{name}/zip/{zip}"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-getbyname-zip-001.png)
<br/><br/>

* [GET] - *Get all companies data* - **"/api/company"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-getall-001.png)
<br/><br/>

* [POST] - *Import data (batch)* - **"/api/company/batch"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-batch-001.png)
> 