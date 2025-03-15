package palette

import (
	"github.com/gookit/color"
)

// ColorPalette representa um conjunto de cores usadas na aplicação
type ColorPalette struct {
	ErrorLine        color.RGBColor
	ErrorText        color.RGBColor
	WarningLine      color.RGBColor
	WarningText      color.RGBColor
	UpdateLine       color.RGBColor
	InsertDeleteLine color.RGBColor
	InsertDeleteText color.RGBColor
	MsgPending       color.RGBColor
	TableHeader      color.RGBColor
	TaskDone         color.RGBColor
	TaskPending      color.RGBColor
	HeaderGreeting   color.RGBColor
	Quote            color.RGBColor
	Author           color.RGBColor
	BackgroundBar    color.RGBColor
	CompleteBar      color.RGBColor
	FinishedBar      color.RGBColor
}

// LoadColorPalette return a ColorPalette based on defined consts
func LoadColorPalette() ColorPalette {
	return ColorPalette{
		ErrorLine:        getColorFromConst(PLSErrorLineStyle),
		ErrorText:        getColorFromConst(PLSErrorTextStyle),
		WarningLine:      getColorFromConst(PLSWarningLineStyle),
		WarningText:      getColorFromConst(PLSWarningTextStyle),
		UpdateLine:       getColorFromConst(PLSUpdateLineStyle),
		InsertDeleteLine: getColorFromConst(PLSInsertDeleteLineStyle),
		InsertDeleteText: getColorFromConst(PLSInsertDeleteTextStyle),
		MsgPending:       getColorFromConst(PLSMsgPendingStyle),
		TableHeader:      getColorFromConst(PLSTableHeaderStyle),
		TaskDone:         getColorFromConst(PLSTaskDoneStyle),
		TaskPending:      getColorFromConst(PLSTaskPendingStyle),
		HeaderGreeting:   getColorFromConst(PLSHeaderGreetingsStyle),
		Quote:            getColorFromConst(PLSQuoteStyle),
		Author:           getColorFromConst(PLSAuthorStyle),
		BackgroundBar:    getColorFromConst(PLSBackgroundBarStyle),
		CompleteBar:      getColorFromConst(PLSCompleteBarStyle),
		FinishedBar:      getColorFromConst(PLSFinishedBarStyle),
	}
}

// getColorFromConst convert a hexdecimal code to color.RGBColor
func getColorFromConst(hexColor string) color.RGBColor {
	if hexColor == "" {
		return color.RGBColor{}
	}
	return color.Hex(hexColor)
}

