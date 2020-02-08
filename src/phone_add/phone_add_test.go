package phone_add

import (
    "testing"
    "reflect"
)


func TestPhoneTranslation(t *testing.T) {
	no_code := "() 11231356" 
	unmatched_country := "(99999) 25113516" 
	invalid_number := "(251) 134d15"	
	valid_numbers := map[string]PhoneNumber{
			"(237) 699209115":{0,699209115,237,"","Cameroon",true},
			"(251) 911168450":{0,911168450,251,"","Ethiopia",true},
			"(212) 609892534":{0,609892534,212,"","Moroco",true},
			"(258) 847651504":{0,847651504,258,"","Mozambique",true},
			"(256) 750306263":{0,750306263,256,"","Uganda",true}}
	test_class := PhoneNumber{}
	var expected_result PhoneNumber
	
	
	test_class.CreatePhoneNumber(0, "", no_code)
	expected_result = PhoneNumber{0,0,0,"","",false}
	if reflect.DeepEqual(test_class, expected_result) {
		t.Log("No Code: Passed")
	}else{
		t.Log("No Code: Failed")
		t.Fail()
	}
	
	
	test_class.CreatePhoneNumber(0,"", unmatched_country)
	expected_result = PhoneNumber{0,0,99999,"","",false}
	if reflect.DeepEqual(test_class, expected_result) {
		t.Log("Unmatched Country: Passed")
	}else{
		t.Log("Unmatched Country: Failed")
		t.Fail()
	}
	
	
	test_class.CreatePhoneNumber(0,"", invalid_number)
	expected_result = PhoneNumber{0,0,251,"","Ethiopia",false}
	if reflect.DeepEqual(test_class, expected_result) {
		t.Log("Invalid Number: Passed")
	}else{
		t.Log("Invalid Number: Failed")
		t.Fail()
	}
	
	for number,result := range valid_numbers{
		test_class.CreatePhoneNumber(0,"", number)
		expected_result = result
		t.Log(test_class)
		t.Log(expected_result)
		if !reflect.DeepEqual(test_class, expected_result) {
			t.Logf("Valid Numbers: Failed at number %s",number)
			t.FailNow()
		}
	}
}

