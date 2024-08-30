package commands

type Command uint8

const (
	Noop Command = iota

	// Views
	SwitchToEditorView
	SwitchToConnectionsView
	HelpPopup

	// Movement: Basic
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
	// Movement: Jumps
	GotoNext
	GotoPrev
	GotoStart
	GotoEnd
	GotoTop
	GotoBottom

	// Movement: Page
	PageNext
	PagePrev

	// Tabs
	TabNext
	TabPrev
	TabFirst
	TabLast
	TabClose

	// Operations
	Copy
	Edit
	Save
	Delete
	Search
	Quit
	Execute
	OpenInExternalEditor
	AppendNewRow
	SortAsc
	SortDesc
)

func (c Command) String() string {
	switch c {
	case Noop:
		return "Noop"
	// Views
	case SwitchToEditorView:
		return "SwitchToEditorView"
	case SwitchToConnectionsView:
		return "SwitchToConnectionsView"
	case HelpPopup:
		return "HelpPopup"

	// Movement: Basic
	case MoveUp:
		return "MoveUp"
	case MoveDown:
		return "MoveDown"
	case MoveLeft:
		return "MoveRight"
	case MoveRight:
		return "MoveRight"
	// Movement: Jumps
	case GotoNext:
		return "GotoNext"
	case GotoPrev:
		return "GotoPrev"
	case GotoStart:
		return "GotoStart"
	case GotoEnd:
		return "GotoEnd"
	case GotoTop:
		return "GotoTop"
	case GotoBottom:
		return "GotoBottom"

	// Movement: Page
	case PageNext:
		return "PageNext"
	case PagePrev:
		return "PagePrev"

	// Tabs
	case TabNext:
		return "TabNext"
	case TabPrev:
		return "TabPrev"
	case TabFirst:
		return "TabFirst"
	case TabLast:
		return "TabLast"
	case TabClose:
		return "TabClose"

	// Operations
	case Copy:
		return "Copy"
	case Edit:
		return "Edit"
	case Save:
		return "Save"
	case Delete:
		return "Delete"
	case Search:
		return "Search"
	case Quit:
		return "Quit"
	case Execute:
		return "Execute"
	case OpenInExternalEditor:
		return "OpenInExternalEditor"
	case AppendNewRow:
		return "AppendNewRow"
	case SortAsc:
		return "SortAsc"
	case SortDesc:
		return "SortDesc"
	}
	return "Unknown"
}
