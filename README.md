![marcio venancio](images/neo-data.png)
#Neo data - For data integration

This Project is an example for Data Ingration. 
 
The project is an API built on top of Golang which serves endpoints to consume and integrate Companies data information that can be added in two ways
* Importing CSV files
* Caling an Endpoint to create a new record.

## Further explanation

The API make use of MongoDB as NoSql database to store the information.
The data, after entered, is processed and then, if already exist, the information is merged, otherwise, a new record is saved.


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
> **Request**
> ![marcio venancio](images/api-getbyid-001.png)
<br/><br/>

* [GET] - *Get company by name* - **"/api/company/name/{name}"**
<br/><br/>
> **Request**
> ![marcio venancio](images/api-getbyname-001.png)
<br/><br/>

* [GET] - *Get compay by name and zip code* - **"/api/company/name/{name}/zip/{zip}"**
<br/><br/>
> **Request**
> ![marcio venancio](images/api-getbyname-zip-001.png)
<br/><br/>

* [GET] - *Get all companies data* - **"/api/company"**
<br/><br/>
> **Request / Response**
> ![marcio venancio](images/api-getall-001.png)
<br/><br/>

* [POST] - *Import data (batch)* - **"/api/company/batch"**
<br/>
> **Request / Response**
> ![marcio venancio](images/api-batch-001.png)
<br/><br/>






*italico*
_italico_

**negrito**
__negrito__


1. aaaa
2. asdas

# Teste
- 1
- 2

* 2
* 2

(Google)(http://www.google.com)
[Build](/build/classes)


> teste de texto 
>

## Tasks
- [x] Done
- [ ] undone


##Tables

| cabecalho | cabecalho |
|-----------|-----------|
| aaaaaa    |  asdasdasd|
|1|2|





