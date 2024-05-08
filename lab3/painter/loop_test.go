package painter

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"
	"testing"
	"time"

	"golang.org/x/exp/shiny/screen"
)

func TestLoop_Post(t *testing.T) {
	var (
		l        Loop
		receiver testReceiver
	)

	l.Receiver = &receiver

	time.Sleep(3 * time.Second)
	l.Start(testScreen{})
	l.Post(OperationFunc(WhiteFill))
	l.Post(OperationFunc(GreenFill))
	l.Post(OperationFunc(WhiteFill))
	l.Post(UpdateOp)

	var testOperation []string

	l.Post(OperationFunc(func(t screen.Texture) {
		testOperation = append(testOperation, "op1")
		l.Post(OperationFunc(func(t screen.Texture) {
			testOperation = append(testOperation, "op3")
		}))
	}))
	l.Post(OperationFunc(func(t screen.Texture) {
		testOperation = append(testOperation, "op2")
	}))

	l.StopAndWait()

	if !reflect.DeepEqual([]string{"op1", "op2", "op3"}, testOperation) {
		t.Error("Bad operations sequence", testOperation)
	}

	if receiver.lastTesture == nil {
		t.Fatal("Texture not update")
	}

	tt := receiver.lastTesture.(*testTexture)
	if tt.lastColor != color.White {
		t.Error("Last color isn't white")
	}
}

type testReceiver struct {
	lastTesture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.lastTesture = t
}

type testScreen struct {
}

func (ts testScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	panic("not inplement")
}

func (ts testScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return &testTexture{
		size: size,
	}, nil
}

func (ts testScreen) NewWindow(opt *screen.NewWindowOptions) (screen.Window, error) {
	panic("not inplement")
}

type testTexture struct {
	size image.Point

	lastColor color.Color
}

func (tt *testTexture) Release() {}

func (tt *testTexture) Size() image.Point {
	return tt.size
}

func (tt *testTexture) Bounds() image.Rectangle {
	return image.Rect(0, 0, tt.Size().X, tt.Size().Y)
}

func (tt *testTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {

}

func (tt *testTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	tt.lastColor = src
}
