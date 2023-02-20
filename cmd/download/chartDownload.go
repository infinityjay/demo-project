package main

type CsvFilter interface {
	Filter()
}

type Add struct{}

func (f *Add) Filter() {
	//
}

type Sum struct{}

func (f *Sum) Filter() {
	//
}

type Decrease struct{}

func (f *Decrease) Filter() {
	//
}

func getFuncObj(funName string) CsvFilter {
	switch funName {
	case "add":
		return &Add{}
	}
	return nil
}

func main() {
	funcList := []string{
		"add",
		"sum",
	}
	for _, funName := range funcList {
		getFuncObj(funName).Filter()
	}

}
