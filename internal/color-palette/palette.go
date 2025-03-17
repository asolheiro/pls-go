package palette

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gookit/color"
)

type ColorStyles struct {
	HeaderStyle lipgloss.Style
	DividerStyle lipgloss.Style
	TaskCompletedStyle lipgloss.Style
	TaskPendingStyle lipgloss.Style
	StrikeStyle lipgloss.Style
	StatusCharStyle lipgloss.Style
}

func NewPalette() ColorStyles {
	headerStyle := lipgloss.NewStyle().
		Bold(false).
		Foreground(lipgloss.Color("#D47BD5")).
		Align(lipgloss.Center)

	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#D77DD8"))

	taskCompletedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888C84"))

	taskPendingStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#B38EE9"))

	strikeStyle := lipgloss.NewStyle().
		Strikethrough(true).
		Align(lipgloss.Center)

	statusCharStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A5D9A3"))

	return ColorStyles{
		HeaderStyle: headerStyle,
		DividerStyle: dividerStyle,
		TaskCompletedStyle: taskCompletedStyle,
		TaskPendingStyle: taskPendingStyle,
		StrikeStyle: strikeStyle,
		StatusCharStyle: statusCharStyle,
	}
}

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

