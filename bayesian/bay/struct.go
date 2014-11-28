package bay

type TD struct {
	// consider every string in lower case
	// if not, convert it to lower case
	// class = positive, negative like sentiment
	// for the purpose of multiple classes
	// we use integer format
	// for example, the preference degree as class
	// 	will span from 1 to 10; 10 is the most preferred
	// weight values are 10 * class
	Class  int
	Weight int
	Text   string
}
