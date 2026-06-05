package structsTut

import "fmt"

type Student struct {
	name    string
	age     uint8
	address string
	email   string
	Course
}

type Course struct {
	CourseName   string
	duration     uint8
	moduleLeader string
}

func Computing() {
	BSC_Computing := Course{
		CourseName:   "BSc Hons in Computing",
		duration:     3,
		moduleLeader: "Anushree Paudyal",
	}
	Pranish := Student{
		name:    "Pranish",
		age:     20,
		address: "Kadaghari",
		email:   "pranish@gmail.com",
		Course:  BSC_Computing,
	}

	var moduleLeader, moduleLeaderErr = getModuleLeaderOfStudent(&Pranish)
	if moduleLeaderErr != nil {
		fmt.Println(moduleLeaderErr)
	} else {
		fmt.Printf("Module Leader of %s is %s\n", Pranish.name, moduleLeader)
	}

	var isAbove18, ageErr = isStudentAbove18(&Pranish)
	if ageErr != nil {
		fmt.Println(ageErr)
	} else if isAbove18 {
		fmt.Printf("%s is above 18\n", Pranish.name)
	} else {
		fmt.Printf("%s is below 18\n", Pranish.name)
	}

}

func getModuleLeaderOfStudent(student *Student) (string, error) {

	if student == nil || student.moduleLeader == "" {
		return "", fmt.Errorf("module leader not found")
	}
	return student.moduleLeader, nil
}

func isStudentAbove18(student *Student) (bool, error) {
	if student == nil {
		return false, fmt.Errorf("student not found")
	}
	return student.age > 18, nil
}
