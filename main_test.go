package simplesoapserver

import (
	"testing"
)

var input string = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
        <SendArray xmlns="http://tempuri.org/">
            <WSID>12</WSID>
            <UserName>userName</UserName>
            <Password>passW0rd</Password>
            <Domain>dom.org</Domain>
            <RecipientNumber>
                <string>+989168030360</string>
                <string>+989111111111</string>
            </RecipientNumber>
            <MessageBody>this is body of message</MessageBody>
            <SpecialNumber>+9890000611</SpecialNumber>
            <IsFlashMessage>1</IsFlashMessage>
            <CheckingMessageID>
                <long>95959511</long>
                <long>45454545</long>
            </CheckingMessageID>
        </SendArray>
    </soap:Body>
</soap:Envelope>`

func TestFindSoapMethod(t *testing.T) {
	var AllMethodes []string
	AllMethodes = append(AllMethodes, "SendArray")
	_, err := FindSoapMethod([]byte(input), AllMethodes)
	if err != nil {
		t.Error("Methode not found. FindSoapMethod function err:", err)
	}
}

func TestSoapRequestFieldParse(t *testing.T) {
	var AddRequest []string
	AddRequest = append(AddRequest, "UserName", "Password", "MessageBody", "string")
	_, err := SoapRequestFieldParse(AddRequest, []byte(input))
	if err != nil {
		t.Error("Fields not found. SoapRequestFieldParse function err:", err)
	}
}
func TestSoapResponse(t *testing.T) {
	endReturn := `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <SendArrayResponse xmlns="http://tempuri.org/">
      <SendArrayResult>
        <long>20202</long>
      </SendArrayResult>
    </SendArrayResponse>
  </soap:Body>
</soap:Envelope>`

	responseCode := [][]string{
		{"long", "20202"},
	}
	if SoapResponse("SendArray", responseCode) != endReturn {
		t.Error("SoapResponse function is not work.")
	}
}
