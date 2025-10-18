package ethcal

// Month names in Amharic
var MonthNames = []string{
	"", // 0-index placeholder
	"መስከረም",
	"ጥቅምት",
	"ኅዳር",
	"ታኅሣሥ",
	"ጥር",
	"የካቲት",
	"መጋቢት",
	"ሚያዝያ",
	"ግንቦት",
	"ሰኔ",
	"ሐምሌ",
	"ነሐሴ",
	"ጳጉሜን",
}

// Day names in Amharic
var DayNames = []string{
	"",       // 0-index placeholder
	"ሰኞ",      // Monday
	"ማክሰኞ",    // Tuesday
	"ረቡዕ",     // Wednesday
	"ሐሙስ",     // Thursday
	"ዓርብ",     // Friday
	"ቅዳሜ",     // Saturday
	"እሑድ",     // Sunday
}

// Day names in long form
var DayNamesLong = []string{
	"",
	"ሰኞ",
	"ማክሰኞ",
	"ረቡዕ",
	"ሐሙስ",
	"ዓርብ",
	"ቅዳሜ",
	"እሑድ",
}

// Orthodox day names (Saints' days)
var OrthodoxDayNames = []string{
	"",
	"ልደታ",
	"ታደዎስ",
	"ቅዱስ",
	"ሩፋኤል",
	"አቡነ",
	"ቅዱሳን",
	"ሥላሴ",
	"እግዝእትነ",
	"ቅዱስ",
	"ገብርኤል",
	"ዮሐንስ",
	"ሚካኤል",
	"እየሱስ",
	"አማኑኤል",
	"ቂርቆስ",
	"እግዚአብሔር",
	"እስጢፋኖስ",
	"መድኃኔዓለም",
	"ደብረዘይት",
	"ሕንፅተ",
	"መስቀል",
	"ሩፋኤል",
	"ድንግል",
	"ሚካኤል",
	"አቡነ",
	"እግዚአብሔር",
	"መድኃኔዓለም",
	"ጊዮርጊስ",
	"በዓለ",
	"ምህረት",
}

// Orthodox year names (Evangelists)
var OrthodoxYearNames = []string{
	"ዮሐንስ",  // John
	"ማቴዎስ",  // Matthew
	"ማርቆስ",  // Mark
	"ሉቃስ",   // Luke
}

// Geez numbers
var GeezNumbers = []string{
	"",
	"፩", "፪", "፫", "፬", "፭", "፮", "፯", "፰", "፱",
	"፲", "፲፩", "፲፪", "፲፫", "፲፬", "፲፭", "፲፮", "፲፯", "፲፰", "፲፱",
	"፳", "፳፩", "፳፪", "፳፫", "፳፬", "፳፭", "፳፮", "፳፯", "፳፰", "፳፱",
	"፴",
}

// Time of day names
var TimeOfDayNames = map[string]string{
	"midnight":  "እኩለ፡ሌሊት",
	"morning":   "ጡዋት",
	"noon":      "ቀትር",
	"afternoon": "ከሰዓት",
	"evening":   "ምሽት",
	"night":     "ሌሊት",
}

// Era names
const (
	EraAM     = "ዓ/ም" // ዓመተ ምሕረት (Anno Mundi - Year of Mercy)
	EraBC     = "ዓ/ዓ" // ዓመተ ዓለም (Before Christ)
)

// Date format constants
const (
	// Example: ዓርብ፣ ግንቦት 04 ቀን 02:35:45 ውደቀት EAT 2009 ዓ/ም
	DateEthiopian = "l፣ F d ቀን H:i:s A T Y E"
	
	// Example: ዓርብ፣ ግንቦት 04 ቀን (ዮሐንስ) 10:35:45 ረፋድ EAT 2009 (ማርቆስ) ዓ/ም
	DateEthiopianOrthodox = "l፣ F d ቀን (x) H:i:s A T Y (X) E"
	
	// Example: ቅዳሜ፣ ግንቦት ፬ ቀን 10:35:45 ረፋድ EAT ፳፻፲ ዓ/ም
	DateGeez = "l፣ F V ቀን H:i:s A T K E"
	
	// Example: ረቡዕ፣ ግንቦት ፭ ቀን (አቦ) 10:35:45 ረፋድ EAT ፳፻፯ (ዮሐንስ) ዓ/ም
	DateGeezOrthodox = "l፣ F V ቀን (x) H:i:s A T K (X) E"
)
