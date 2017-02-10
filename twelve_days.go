package twelve

const testVersion = 1

type givenThing struct {
	quantity int
	thing	string
}

var givenThings = []givenThing{
	{1, "Partridge"},
	{2, "Turtle Doves"},
	{3, "French Hens"},
	{4, "Calling Birds"},
	{5, "Gold Rings"},
	{6, "Geese-a-Laying"},
	{7, "Swans-a-Swimming"},
	{8, "Maids-a-Milking"},
	{9, "Ladies Dancing"},
	{10, "Lords-a-Leaping"},
	{11, "Pipers Piping"},
	{12, "Drummers Drumming"},
}
var  number2Word = [12]string{"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"} 
func Ordinal (number int) string {
	var result string
	switch number {
		case 1:
			result = "first"
		case 2:
			result = "second"
		case 3:
                        result = "third"
		case 5:
                        result = "fifth"
                case 8:
                        result = "eighth"
                case 9:
                        result = "ninth"
                case 12:
                        result = "twelfth"
		default:
			result = SpellOut(number) + "th"
	}
	return result
}
func SpellOut (number int) string {
	return number2Word[number-1]
}
func Verse(index int) string {
	result := "On the"
	result += " " + Ordinal(index)
	result += " day of Christmas my true love gave to me"
	for i := index; i > 0; i-- {
		if i == 1 && index != 1 {
			result += ", and "
		} else {
			result += ", "
			}
		
		result += SpellOut(givenThings[i-1].quantity) + " " + givenThings[i-1].thing
	} 
	result += " in a Pear Tree."
	return result
}
func Song() string {
	var result string
	for i:= 1; i<=12; i++ {
		result += Verse(i) + "\n"			
	}
	return result		
}
