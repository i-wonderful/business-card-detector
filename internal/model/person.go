package model

import "fmt"

type Person struct {
	Email []string `json:"Emails"`
	Site  []string `json:"Websites"`

	Phone    []string `json:"Phone Numbers"`
	Skype    []string `json:"Skype IDs"`
	Telegram []string `json:"Telegram IDs"`

	Name string `json:"Name"`

	Organization string `json:"Organization"`
	JobTitle     string `json:"Job Title"`

	Other string `json:"other"`
}

func (p *Person) Print() {
	fmt.Println("email:", p.Email)
	fmt.Println("site:", p.Site)
	fmt.Println("phone:", p.Phone)
	fmt.Println("name:", p.Name)
	fmt.Println("company:", p.Organization)
	fmt.Println("telegram:", p.Telegram)
	fmt.Println("skype:", p.Skype)
	fmt.Println("jobTitle:", p.JobTitle)
	fmt.Println("other:", p.Other)
}

func NewPerson(fields map[string]interface{}) *Person {
	p := &Person{
		Email:        []string{},
		Site:         []string{},
		Phone:        []string{},
		Name:         fields["name"].(string),
		JobTitle:     fields["jobTitle"].(string),
		Organization: fields["company"].(string),
		Telegram:     []string{},
		Skype:        []string{},
		Other:        fields["other"].(string),
	}
	if fields["email"] != "" {
		p.Email = append(p.Email, fields["email"].(string))
	}

	if fields["site"] != "" {
		p.Site = append(p.Site, fields["site"].(string))
	}

	if fields["phone"] != "" {
		p.Phone = append(p.Phone, fields["phone"].([]string)...)
	}
	if fields["telegram"] != "" {
		p.Telegram = append(p.Telegram, fields["telegram"].(string))
	}
	if fields["skype"] != "" {
		p.Skype = append(p.Skype, fields["skype"].(string))
	}
	return p
}
