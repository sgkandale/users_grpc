package user

var AllUsers = []User{
	{
		Id:      1,
		Fname:   "Ram",
		City:    "Rampur",
		Phone:   1234567890,
		Height:  1.50,
		Married: true,
	},
	{
		Id:      2,
		Fname:   "Shyam",
		City:    "Shyampur",
		Phone:   2345678901,
		Height:  1.54,
		Married: false,
	},
	{
		Id:      3,
		Fname:   "Raju",
		City:    "Raju Nagar",
		Phone:   3456789012,
		Height:  1.55,
		Married: true,
	},
	{
		Id:      4,
		Fname:   "Shantabai",
		City:    "Shantapur",
		Phone:   4567890123,
		Height:  1.52,
		Married: true,
	},
	{
		Id:      5,
		Fname:   "Kamla",
		City:    "Kamla Nagar",
		Phone:   5678901234,
		Height:  1.61,
		Married: false,
	},
	{
		Id:      6,
		Fname:   "Bharat",
		City:    "Bharatpur",
		Phone:   6789012345,
		Height:  1.58,
		Married: true,
	},
	{
		Id:      7,
		Fname:   "Ganga",
		City:    "Gangapur",
		Phone:   7890123456,
		Height:  1.55,
		Married: true,
	},
	{
		Id:      8,
		Fname:   "Bala",
		City:    "Balapur",
		Phone:   8901234567,
		Height:  1.58,
		Married: false,
	},
	{
		Id:      9,
		Fname:   "Rajesh",
		City:    "Rajesh Nagar",
		Phone:   9012345678,
		Height:  1.60,
		Married: true,
	},
	{
		Id:      10,
		Fname:   "Suresh",
		City:    "Suresh Nagar",
		Phone:   1234567890,
		Height:  1.50,
		Married: true,
	},
}

func GetAllUsers() []User {
	return AllUsers
}
