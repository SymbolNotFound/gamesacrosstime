package categories

type Category string

const (
	CATEGORY_ABSTRACT    CategoryID = "abstract-strategy"
	CATEGORY_CARDS       CategoryID = "card-game"
	CATEGORY_DICE        CategoryID = "dice-game"
	CATEGORY_ROLL_WRITE  CategoryID = "roll-n-write"
	CATEGORY_THEMATIC    CategoryID = "thematic"
	CATEGORY_COLLECTIBLE CategoryID = "collectible"
	CATEGORY_PARTY       CategoryID = "party"
	CATEGORY_FAMILY      CategoryID = "family"
	CATEGORY_CHILDRENS   CategoryID = "childrens"
	CATEGORY_EDUCATIONAL CategoryID = "educational"
	CATEGORY_EXPLORATION CategoryID = "exploration"
	CATEGORY_STRATEGY    CategoryID = "strategy"
	CATEGORY_WARGAME     CategoryID = "wargame"
	CATEGORY_MINIATURES  CategoryID = "miniatures"
	CATEGORY_TERRITORY   CategoryID = "territory"
	CATEGORY_PUZZLE      CategoryID = "puzzle"
	CATEGORY_WORD_GAME   CategoryID = "word-game"
	CATEGORY_PRINTABLE   CategoryID = "print-n-play"
	CATEGORY_ELECTRONIC  CategoryID = "electronic"
)

const EXPECTED_SUBCATEGORIES_COUNT int = 11

type CategoryID string

var Categories map[CategoryID]Category

func parseSubcategories(inCategories []string) []CategoryID {
	categories := make([]CategoryID, 0)
	for i, s := range inCategories {
		if s == "0" {
			continue
		}
		switch i {
		case 0: // BGGId (ignore)
		case 1: // Exploration
		case 2: // Miniatures
		case 3: // Territory Building
		case 4: // Card Game
		case 5: // Educational
		case 6: // Puzzle
		case 7: // Collectible Components
		case 8: // Word Game
		case 9: // Print & Play
		case 10: // Electronic
		}
	}

	return categories
}

func init() {
	// Create this map only once, so it can be called from within parseCategories.
	Categories = map[CategoryID]Category{
		CATEGORY_ABSTRACT:    Category("Abstract Strategy"),
		CATEGORY_CARDS:       Category("Card Game"),
		CATEGORY_DICE:        Category("Dice Game"),
		CATEGORY_ROLL_WRITE:  Category("Roll 'N' Write"),
		CATEGORY_THEMATIC:    Category("Thematic"),
		CATEGORY_COLLECTIBLE: Category("Collectible"),
		CATEGORY_PARTY:       Category("Party"),
		CATEGORY_FAMILY:      Category("Family"),
		CATEGORY_CHILDRENS:   Category("Childrens"),
		CATEGORY_EDUCATIONAL: Category("Educational"),
		CATEGORY_EXPLORATION: Category("Exploration"),
		CATEGORY_STRATEGY:    Category("Strategy"),
		CATEGORY_WARGAME:     Category("Wargame"),
		CATEGORY_MINIATURES:  Category("Miniatures"),
		CATEGORY_TERRITORY:   Category("Territory"),
		CATEGORY_PUZZLE:      Category("Puzzle"),
		CATEGORY_WORD_GAME:   Category("Word Game"),
		CATEGORY_PRINTABLE:   Category("Print 'N' Play"),
		CATEGORY_ELECTRONIC:  Category("Electronic"),
	}
}
