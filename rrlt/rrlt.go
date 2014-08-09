package rrlt

type Record struct {
	Code int
	Msg  string
	Data Resume
}

type Resume struct {
	Resume_Id          string
	User_Id            string
	Owner_User_Id      string
	Name               string
	Gender             string
	Email              string
	BirthYear          string
	BirthMonth         string
	Mobile             string
	Mobile_Verified    string
	CityName           string
	Marital            string
	WorkYear           string
	WorkMonth          string
	EduLevelName       string
	CurrentStatusName  string
	CurrentCompany     string
	CurrentPosition    string
	CurrentIndustry    string
	CurrentSalaryName  string
	AvatarUrl          string
	Labels             string
	Resume_Text        string
	Resume_Source_Type string
	Works              []Work
	Projects           []Project
}

type Work struct {
	Work_Id       string
	Resume_Id     string
	Company_Name  string
	Industry_Name string
	Position_Name string
	SYear         string
	SMonth        string
	EYear         string
	EMonth        string
	Department    string
	JD            string
	UntilNow      string
}

type Project struct {
	Project_Id   string
	Resume_Id    string
	Project_Name string
	Project_Desc string
	SYear        string
	SMonth       string
	EYear        string
	EMonth       string
	UntilNow     string
}
