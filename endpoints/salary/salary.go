package salary

import (
	"net/http"
	"html/template"
	"fmt"
	"mntbudget/Database"
	"log"
)

type Salary struct{
	// from DB
	Profile_id int
	Description string
	Bruto_month float32
	Bruto_year float32
	Estimated_net float32
	Monthly_tax float32
	Yearly_tax float32
	Dutch_30 int

	//Calculated
	Year_taxable float32

}

func (s *Salary ) explode_salary( salary Salary ){
	var (
		tax_37 float32
		tax_40 float32
		tax_52 float32
		remainer float32
	)
	s.Bruto_year = s.Bruto_month * 12.9599
	if s.Dutch_30 > 0 {
		s.Year_taxable = s.Bruto_year * 0.7
	}else{
		s.Year_taxable = s.Bruto_year
	}

	if s.Year_taxable > 19992 {
		tax_37 = 19992 * 0.3655
		remainer = s.Year_taxable - tax_37
	}else {
		tax_37 = s.Year_taxable * 0.3655
		remainer = 0
	}
	if remainer > 0  {
		if 66421 < s.Year_taxable {
			tax_40 = ( 66421 - 19992 ) * 0.4040
			remainer = s.Year_taxable - 66421
		}else{
			tax_40 = ( s.Year_taxable - 19992 ) * 0.4040
		}
	}
	if remainer > 0 {
		tax_52 = ( s.Year_taxable - 66421 ) * 0.52
	}

	s.Yearly_tax = tax_37 + tax_40 + tax_52
	s.Monthly_tax = s.Yearly_tax / 12.9599
	s.Estimated_net = s.Bruto_year/ 12 - s.Monthly_tax


}





// end point code

func errorHandler( w http.ResponseWriter, r *http.Request, status int){
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		fmt.Fprint(w,"404 not found")
	}
	if status == http.StatusUnauthorized {
		fmt.Fprint(w, "401 not authorized")
	}
}

func salaryHandler (w http.ResponseWriter, r *http.Request){
	// login

	address := r.URL.Path[len( "/salary/"):]
	fmt.Println(address)

	if r.Method == "GET" {
		var (
			salaries []Salary
		)

		fmt.Println( "GET")

		db, err := Database.Database()

		if  err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query( "SELECT * FROM salary")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			salary := Salary{}
			err := rows.Scan(
				&salary.Profile_id,
				&salary.Description,
				&salary.Bruto_month,
				&salary.Bruto_year,
				&salary.Estimated_net,
				&salary.Monthly_tax,
				&salary.Yearly_tax,
				&salary.Dutch_30,
			)
			if  err != nil {
				fmt.Println(err)
			}
			salaries = append(salaries, salary)
		}


		defer rows.Close()

		fmt.Printf("%+v\n",salaries)

		t, err := template.ParseFiles( "./Templates/endpoints/salary.t")
		err = t.Execute(w, salaries)
		if err !=  nil{
			fmt.Println(err)
		}
	}

}

func init(){
	http.HandleFunc("/salary/", salaryHandler)
}
