package wakaf

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetPDFContent() core.Maroto {
	cfg := config.NewBuilder().WithDebug(false).Build()
	m := maroto.New(cfg)

	m.AddAutoRow(image.NewFromFileCol(12, "examples/wakaf/assets/logo.png", props.Rect{
		Percent:            40,
		JustReferenceWidth: true,
	}))
	m.AddAutoRow(text.NewCol(12, "SERTIFIKAT WAKAF", props.Text{
		Style: fontstyle.Bold,
		Size:  32,
		Top:   5,
	}))
	m.AddAutoRow(text.NewCol(12, "Waqf Certificate", props.Text{
		Top:   3,
		Style: fontstyle.Italic,
		Size:  8,
		Color: &props.Color{Red: 169, Green: 169, Blue: 169},
	}))
	m.AddAutoRow(text.NewCol(12, "Sebuah doa dan apresiasi", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Size:  20,
	}))
	m.AddAutoRow(text.NewCol(12, "Untuk Anda,", props.Text{
		Top:  2,
		Size: 20,
	}))
	m.AddAutoRow(text.NewCol(12, "With Gratitude and Sincere Prayers, For You", props.Text{
		Top:   1,
		Style: fontstyle.Italic,
		Size:  8,
		Color: &props.Color{Red: 169, Green: 169, Blue: 169},
	}))

	return m
}
