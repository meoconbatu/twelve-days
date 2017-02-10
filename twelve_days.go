package twelve

const testVersion = 1

type givenThing struct {
	quantity int
	thing    string
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

var number2Word = []string{"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
var ordinal = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

func Ordinal(number int) string {
	return ordinal[number-1]
}
func SpellOut(number int) string {
	return number2Word[number-1]
}
func Gift(index int) string {
	result := ""
	for i := index; i > 0; i-- {
                if i == 1 && index != 1 {
                        result += ", and "
                } else {
                        result += ", "
                }
                result += SpellOut(givenThings[i-1].quantity) + " " + givenThings[i-1].thing
        }
	return result
}
func Verse(index int) string {
	return "On the " + Ordinal(index) + " day of Christmas my true love gave to me" + Gift(index) + " in a Pear Tree."
}
func Song() string {
	var result string
	numberOfVerse := 12
	for i := 1; i <= numberOfVerse; i++ {
		result += Verse(i) + "\n"
	}
	return result
}
