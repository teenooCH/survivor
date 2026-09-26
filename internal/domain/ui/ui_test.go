package ui_test

import (
	"image/color"
	"testing"

	"survivor/internal/domain/graph"
	"survivor/internal/domain/ui"
)

// fakeRenderer is a minimal graph.Image + graph.TextRenderer used as the
// draw target in tests, recording the last DrawText call it received.
type fakeRenderer struct {
	drawTextCalls int
	lastOpt       graph.TextOpt
}

func (f *fakeRenderer) DrawImage(graph.Image, graph.DrawOpt) {}
func (f *fakeRenderer) Dimensions() (width, height float64)  { return 0, 0 }
func (f *fakeRenderer) SetDimensions(width, height float64)  {}

func (f *fakeRenderer) DrawText(opts graph.TextOpt) {
	f.drawTextCalls++
	f.lastOpt = opts
}
func (f *fakeRenderer) FillRect(x, y, width, height float64, clr color.Color)                {}
func (f *fakeRenderer) StrokeRect(x, y, width, height, strokeWidth float64, clr color.Color) {}

// plainImage is a graph.Image that does NOT implement graph.TextRenderer.
type plainImage struct{}

func (plainImage) DrawImage(graph.Image, graph.DrawOpt) {}
func (plainImage) Dimensions() (width, height float64)  { return 0, 0 }
func (plainImage) SetDimensions(width, height float64)  {}

func TestText_Draw(t *testing.T) {
	txt := ui.NewText("hello", 10, 20)
	txt.SetColor(color.RGBA{R: 255, A: 255})
	txt.SetSize(24)
	txt.SetBackground(ui.NewBackground(color.Black))
	txt.SetBorder(ui.NewBorder(color.White, 2))

	target := &fakeRenderer{}
	txt.Draw(target)

	if target.drawTextCalls != 1 {
		t.Fatalf("expected DrawText to be called once, got %d", target.drawTextCalls)
	}

	got := target.lastOpt
	if got.Text != "hello" || got.X != 10 || got.Y != 20 || got.Size != 24 {
		t.Errorf("unexpected TextOpt: %+v", got)
	}
	if got.Background.Transparent || got.Background.Color != color.Black {
		t.Errorf("unexpected background: %+v", got.Background)
	}
	if !got.Border.Visible || got.Border.Width != 2 {
		t.Errorf("unexpected border: %+v", got.Border)
	}
}

func TestText_Draw_NotVisible(t *testing.T) {
	txt := ui.NewText("hello", 0, 0)
	txt.SetVisible(false)

	target := &fakeRenderer{}
	txt.Draw(target)

	if target.drawTextCalls != 0 {
		t.Errorf("expected DrawText not to be called when invisible, got %d calls", target.drawTextCalls)
	}
}

func TestText_Draw_UnsupportedTarget(t *testing.T) {
	txt := ui.NewText("hello", 0, 0)

	// Should not panic when the target doesn't support text rendering.
	txt.Draw(plainImage{})
}

func TestHUD_Draw(t *testing.T) {
	hud := ui.NewHUD()
	a := ui.NewText("a", 0, 0)
	b := ui.NewText("b", 0, 0)
	hud.AddWidget(a)
	hud.AddWidget(b)

	target := &fakeRenderer{}
	hud.Draw(target)

	if target.drawTextCalls != 2 {
		t.Fatalf("expected 2 DrawText calls, got %d", target.drawTextCalls)
	}

	if !hud.RemoveWidget(a) {
		t.Fatalf("expected RemoveWidget to find widget a")
	}

	target = &fakeRenderer{}
	hud.Draw(target)

	if target.drawTextCalls != 1 {
		t.Fatalf("expected 1 DrawText call after removal, got %d", target.drawTextCalls)
	}
}
