package palette

import (
	"github.com/charmbracelet/lipgloss"
)

type ColorStyles struct {
	HeaderStyle           lipgloss.Style
	DividerStyle          lipgloss.Style
	TaskCompletedStyle    lipgloss.Style
	TaskPendingStyle      lipgloss.Style
	StrikeStyle           lipgloss.Style
	StatusCharStyle       lipgloss.Style
	HeaderGreetingStyle   lipgloss.Style
	QuoteStyle            lipgloss.Style
	CompletedsBar         lipgloss.Style
	RemaingTasksBarStyle  lipgloss.Style
	TaskCompletedBarStyle lipgloss.Style
	TaskRatioStyles       lipgloss.Style
	CenteredStyle         lipgloss.Style
	ErrorStyle            lipgloss.Style
	InfoStyle             lipgloss.Style
}

func NewPalette() ColorStyles {
	headerStyle := lipgloss.NewStyle().
		Bold(false).
		Foreground(PinkLilac).
		Align(lipgloss.Center)

	dividerStyle := lipgloss.NewStyle().
		Foreground(PurpleOrchid)

	taskCompletedStyle := lipgloss.NewStyle().
		Foreground(GrayAsh)

	taskPendingStyle := lipgloss.NewStyle().
		Foreground(PurpleLavender)

	strikeStyle := lipgloss.NewStyle().
		Strikethrough(true).
		Align(lipgloss.Center)

	statusCharStyle := lipgloss.NewStyle().
		Foreground(GreenMint)

	headerGreetingStyle := lipgloss.NewStyle().
		Foreground(YellowGolden)

	quoteStyle := lipgloss.NewStyle().
		Foreground(GraySilver).
		Align(lipgloss.Center)

	completedsBarStryle := lipgloss.NewStyle().
		Foreground(PinkNeon)

	remaingTasksBarStyle := lipgloss.NewStyle().
		Foreground(GrayCharcoal).
		Bold(true)

	taskCompletedBarStyle := lipgloss.NewStyle().
		Foreground(PinkNeon).
		Bold(true)

	tasksRatioStyles := lipgloss.NewStyle().
		Foreground(GreenOlive)

	centeredStyle := lipgloss.NewStyle().
		Align(lipgloss.Center)

	errorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(PinkCarmine))

	infoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(GraySilver)).
		Faint(true)

	return ColorStyles{
		HeaderStyle:           headerStyle,
		DividerStyle:          dividerStyle,
		TaskCompletedStyle:    taskCompletedStyle,
		TaskPendingStyle:      taskPendingStyle,
		StrikeStyle:           strikeStyle,
		StatusCharStyle:       statusCharStyle,
		HeaderGreetingStyle:   headerGreetingStyle,
		QuoteStyle:            quoteStyle,
		CompletedsBar:         completedsBarStryle,
		RemaingTasksBarStyle:  remaingTasksBarStyle,
		TaskCompletedBarStyle: taskCompletedBarStyle,
		TaskRatioStyles:       tasksRatioStyles,
		CenteredStyle:         centeredStyle,
		ErrorStyle:            errorStyle,
		InfoStyle:             infoStyle,
	}
}
