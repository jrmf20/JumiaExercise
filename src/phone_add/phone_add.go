package phone_add

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

//Country dic structure binding the country name and regex
type country_dic struct {
	name  string
	regex *regexp.Regexp
}

//map containing code as key and the name/format registry the country names
type clist map[int]country_dic

//public variable to the countries
var PhoneCDic clist
var regcode *regexp.Regexp

func init() {
	PhoneCDic = clist{
		237: country_dic{"Cameroon", regexp.MustCompile(`\((237)\)\ ?([2368]\d{7,8}$)`)},
		251: country_dic{"Ethiopia", regexp.MustCompile(`\((251)\)\ ?([1-59]\d{8}$)`)},
		212: country_dic{"Moroco", regexp.MustCompile(`\((212)\)\ ?([5-9]\d{8}$)`)},
		258: country_dic{"Mozambique", regexp.MustCompile(`\((258)\)\ ?([28]\d{7,8}$)`)},
		256: country_dic{"Uganda", regexp.MustCompile(`\((256)\)\ ?(\d{9}$)`)},
	}
	regcode = regexp.MustCompile(`\((\d+)\).+`)
}

//Phone Number "class"
type PhoneAdd struct {
	id, number, code int
	Name, country    string
	Valid            bool
}

//Populate/Create PhoneAdd "class" given db info
func CreatePhoneAdd(id int, name string, number string) (p PhoneAdd) {
	p.id = id
	p.Name = name
	p.country, p.code, p.number, p.Valid = numberToInfo(number)
	return
}

//Retrives additional info from db number
func numberToInfo(n_string string) (country string, code int, number int, valid bool) {
	valid = false
	code_string := regcode.FindStringSubmatch(n_string)
	if code_string == nil {
		//fmt.Println(fmt.Errorf("Invalid Country Code").Error())
		return
	}
	code_int, err := strconv.Atoi(code_string[1])
	//should always work
	if err != nil {
		//fmt.Println(fmt.Errorf("Invalid Country Code").Error())
		return
	}
	code = code_int

	if count_det, ok := PhoneCDic[code_int]; !ok {
		//fmt.Println(fmt.Errorf("Unmatched Country Code").Error())
		return
	} else {
		country = count_det.name
		if matches := count_det.regex.FindStringSubmatch(n_string); matches != nil {
			number_int, err := strconv.Atoi(matches[2])
			//should always work
			if err != nil {
				//fmt.Println(fmt.Errorf("Unmatched Number Sequence").Error())
				return
			}
			number = number_int
			valid = true
		}
	}
	return
}

func (p *PhoneAdd) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"id\" : %d,", p.id))
	buffer.WriteString(fmt.Sprintf("\"name\" : \"%s\",", p.Name))
	buffer.WriteString(fmt.Sprintf("\"number\" : %d,", p.number))
	buffer.WriteString(fmt.Sprintf("\"country\" : \"%s\",", p.country))
	buffer.WriteString(fmt.Sprintf("\"state\" : %t", p.Valid))
	buffer.WriteString(fmt.Sprintf("}"))
	return buffer.Bytes(), nil
}
