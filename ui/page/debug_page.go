package page

import (
	"gioui.org/layout"
	"github.com/planetdecred/godcr/ui/decredmaterial"
	"github.com/planetdecred/godcr/ui/load"
	"github.com/planetdecred/godcr/ui/page/components"
	"github.com/planetdecred/godcr/ui/values"
)

const DebugPageID = "Debug"

type debugItem struct {
	text   string
	page   string
	action func()
}

type DebugPage struct {
	*load.Load
	debugItems []debugItem
	list       *decredmaterial.ClickableList

	backButton decredmaterial.IconButton
}

func NewDebugPage(l *load.Load) *DebugPage {
	debugItems := []debugItem{
		{
			text: "Check wallet logs",
			page: LogPageID,
			action: func() {
				l.ChangeFragment(NewLogPage(l))
			},
		},
		{
			text: "Check statistics",
			page: StatisticsPageID,
			action: func() {
				l.ChangeFragment(NewStatPage(l))
			},
		},
	}

	pg := &DebugPage{
		Load:       l,
		debugItems: debugItems,
		list:       l.Theme.NewClickableList(layout.Vertical),
	}
	pg.list.Radius = decredmaterial.Radius(14)

	pg.backButton, _ = components.SubpageHeaderButtons(l)

	return pg
}

func (pg *DebugPage) ID() string {
	return DebugPageID
}

func (pg *DebugPage) OnResume() {

}

func (pg *DebugPage) Handle() {
	if clicked, item := pg.list.ItemClicked(); clicked {
		pg.debugItems[item].action()
	}
}

func (pg *DebugPage) OnClose() {}

func (pg *DebugPage) debugItem(gtx C, i int) D {
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.UniformInset(values.MarginPadding15).Layout(gtx, pg.Theme.Body1(pg.debugItems[i].text).Layout)
		}),
		layout.Flexed(1, func(gtx C) D {
			return layout.E.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(values.MarginPadding15).Layout(gtx, func(gtx C) D {
					ic := decredmaterial.NewIcon(pg.Icons.ChevronRight)
					return ic.Layout(gtx, values.MarginPadding22)
				})
			})
		}),
	)
}

func (pg *DebugPage) layoutDebugItems(gtx C) {
	background := pg.Theme.Color.Surface
	card := pg.Theme.Card()
	card.Color = background
	card.Layout(gtx, func(gtx C) D {
		return pg.list.Layout(gtx, len(pg.debugItems), func(gtx C, i int) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return pg.debugItem(gtx, i)
				}),
				layout.Rigid(func(gtx C) D {
					if i == len(pg.debugItems)-1 {
						return layout.Dimensions{}
					}
					return layout.Inset{
						Left: values.MarginPadding16,
					}.Layout(gtx, pg.Theme.Separator().Layout)
				}),
			)
		})
	})
}

func (pg *DebugPage) Layout(gtx C) D {
	container := func(gtx C) D {
		sp := components.SubPage{
			Load:       pg.Load,
			Title:      "Debug",
			BackButton: pg.backButton,
			Back: func() {
				pg.PopFragment()
			},
			Body: func(gtx C) D {
				pg.layoutDebugItems(gtx)
				return layout.Dimensions{Size: gtx.Constraints.Max}
			},
		}
		return sp.Layout(gtx)

	}
	return components.UniformPadding(gtx, container)
}
