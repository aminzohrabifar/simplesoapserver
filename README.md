
# Simple Soap Server for Go Language

Using SOAP, you can set up a web server and process all XML values that are sent via POST.
#
There are indeed many SOAP packages available for the Go language, but these packages cannot handle cases where the user sends an array request to you.


## Features

- Very Simple and Usefull
- <b>Parse Array Inputs</b>
- Parse Methode
- Simple Response


## Example

### Example SOAP Request.
As you can see this package can parse this Soap Request.

```xml
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
        <SendArray xmlns="http://tempuri.org/">
            <WSID>int</WSID>
            <UserName>string</UserName>
            <Password>string</Password>
            <Domain>string</Domain>
            <RecipientNumber>
                <string>string</string>
                <string>string</string>
            </RecipientNumber>
            <MessageBody>string</MessageBody>
            <SpecialNumber>string</SpecialNumber>
            <IsFlashMessage>string</IsFlashMessage>
            <CheckingMessageID>
                <long>long</long>
                <long>long</long>
            </CheckingMessageID>
        </SendArray>
    </soap:Body>
</soap:Envelope>
```
    
## Installation

```sh
go get github.com/aminzohrabifar/simplesoapserver
```

### Example with Echo Framework

```go
package main

import (
	"fmt"
	"github.com/aminzohrabifar/simplesoapserver"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

func main() {

	e := echo.New()
	e.POST("/wsdl", func(c echo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			log.Println(string(err.Error()))
		}
		
		// Get Varable From Body
		var AddRequest []string
		AddRequest = append(AddRequest, "UserName", "Password", "MessageBody", "RecipientNumber")
		match ,err := simplesoapserver.SoapRequestFieldParse(AddRequest, b)
		if err != nil {
			log.Println(err)
		}
		
		// Get Method Used From Body
		var AllMethodes []string
		AllMethodes = append(AllMethodes, "GetCredit", "SendArray", "GetMessageID", "GetMessageStatus")
		Methode, err := simplesoapserver.FindSoapMethod(b, AllMethodes)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%+v", match)
		
		// Response to Request
		m := [][]string{
			{"long", "20202"},
			{"long", "23330202"},
		}
		return c.XMLBlob(http.StatusOK, []byte(simplesoapserver.SoapResponse(Methode, m)))
	})
	e.Start("127.0.0.1:8080")
}


```
## Function Reference

#### To process XML, you need to use the following function:


```go
  match,err := simplesoapserver.SoapRequestFieldParse(AddRequest, b)
```

| Parameter | Type     | Description                                     |
| :-------- | :------- |:------------------------------------------------|
| `match` | `map[string][]string` | a slice of passed values                        |
| `err`      | `error` | returns an error if passed values are not found |
| `AddRequest` | `[]string` | a slice of variables needed for processing      |
| `b` | `[]byte` | the sent body text                              |

#### To determine the sent methods, use the following function:

```go
  Method, err := simplesoapserver.FindSoapMethod(b, AllMethods)
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Method`      | `string` | the name of the used method as a string |
| `err`      | `error` | returns an error if the sent methods are not found |
| `b` | `[]byte` | the sent body text |


#### To respond to a request, you can use the following command:

```go
  return c.XMLBlob(http.StatusOK, []byte(simplesoapserver.SoapResponse(Method, m)))

```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Method`      | `string` | **Required**. the method name |
| `m`      | `[][]string` | **Required**. a slice of variables that should be sent |


## Authors

- [@Amin Zohrabi Far](https://www.github.com/aminzohrabifar)

