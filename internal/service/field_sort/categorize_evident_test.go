package field_sort

type testCase struct {
	name                string
	input               []string
	expected            map[string]interface{}
	expectedNotDetected []string
}

var testCases = []testCase{
	{
		name:  "Phone numbers detection",
		input: []string{"+12345678901", "+98765432109"},
		expected: map[string]interface{}{
			FIELD_PHONE: []string{"+12345678901", "+98765432109"},
		},
		expectedNotDetected: []string{},
	},
	{
		name:  "Email addresses detection",
		input: []string{"example@example.com", "another.email@gmail.com"},
		expected: map[string]interface{}{
			FIELD_EMAIL: []string{"example@example.com", "another.email@gmail.com"},
		},
		expectedNotDetected: []string{},
	},
	//{
	//	name:  "Telegram usernames detection",
	//	input: []string{"@username", "@another_username"},
	//	expected: map[string]interface{}{
	//		FIELD_TELEGRAM: []string{"@username", "@another_username"},
	//	},
	//},
	//{
	//	name:  "Skype usernames detection",
	//	input: []string{"skype:username", "skype:another_username"},
	//	expected: map[string]interface{}{
	//		FIELD_SKYPE: []string{"skype:username", "skype:another_username"},
	//	},
	//},
	{
		name:                "Non-detectable items",
		input:               []string{"random string", "not a number", "nope"},
		expected:            map[string]interface{}{},
		expectedNotDetected: []string{"random string", "not a number", "nope"},
	},
	{
		name: "Basic test",
		input: []string{
			"+1 (123) 456-7890",
			"user@example.com",
			"@telegram_user",
			"live:skype_user",
			"Some unrecognized text",
		},
		expected: map[string]interface{}{
			FIELD_PHONE:    []string{"+1 (123) 456-7890"},
			FIELD_EMAIL:    []string{"user@example.com"},
			FIELD_TELEGRAM: []string{"@telegram_user"},
			FIELD_SKYPE:    "live:skype_user",
		},
		expectedNotDetected: []string{"Some unrecognized text"},
	},
	{
		name: "Multiple phones and emails",
		input: []string{
			"+1 (123) 456-7890",
			"+7 987 654 32 10",
			"user1@example.com",
			"user2@example.com",
		},
		expected: map[string]interface{}{
			FIELD_PHONE: []string{"+1 (123) 456-7890", "+7 987 654 32 10"},
			FIELD_EMAIL: []string{"user1@example.com", "user2@example.com"},
		},
		expectedNotDetected: []string{},
	},
	{
		name:  "Valid Phone",
		input: []string{"call me at 123-456-7890"},
		expected: map[string]interface{}{
			FIELD_PHONE: []string{" 123-456-7890"},
		},
		expectedNotDetected: []string{},
	},
	{
		"Phone with slash",
		[]string{"+55(71)3206-8267/807"},
		map[string]interface{}{
			FIELD_PHONE: []string{"+55(71)3206-8267/807"},
		},
		[]string{},
	},
	{
		name:  "Valid Email",
		input: []string{"send email to example@example.com"},
		expected: map[string]interface{}{
			FIELD_EMAIL: []string{"example@example.com"},
		},
		expectedNotDetected: []string{},
	},
	{
		name:  "Valid Telegram",
		input: []string{"contact me @telegram_handle"},
		expected: map[string]interface{}{
			FIELD_TELEGRAM: []string{"@telegram_handle"},
		},
		expectedNotDetected: []string{},
	},
}

//func TestCategorizeEvident(t *testing.T) {
//	service := &Service{}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			recognized, _ ,_ := service.categorizeEvidentFields(tc.input)
//			filteredRecognized := removeEmptyFields(recognized)
//			assert.Equal(t, tc.expected, filteredRecognized, "Recognized should be equal")
//			//assert.Equal(t, tc.expectedNotDetected, notDetected, "notDetected should be equal")
//		})
//	}
//}

func removeEmptyFields(recognized map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	for key, value := range recognized {
		switch v := value.(type) {
		case string:
			if v != "" {
				filtered[key] = value
			}
		case []string:
			if len(v) > 0 {
				filtered[key] = value
			}
		}
	}
	return filtered
}
