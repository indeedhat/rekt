package rectangled_test

import (
	"fmt"
	"testing"

	"github.com/indeedhat/rectangled"
	"github.com/stretchr/testify/require"
)

var rectangleOverlapTouchTests = []struct {
	rect1    rectangled.Rectangle[string]
	rect2    rectangled.Rectangle[string]
	overlaps bool
	touch1   []rectangled.Edge
	touch2   []rectangled.Edge
}{
	{
		rectangled.NewRectangle("above", 0, 0, 10, 10),
		rectangled.NewRectangle("below", 0, 20, 10, 30),
		false,
		nil,
		nil,
	},
	{
		rectangled.NewRectangle("left", 0, 0, 10, 10),
		rectangled.NewRectangle("right", 20, 0, 30, 10),
		false,
		nil,
		nil,
	},
	{
		rectangled.NewRectangle("touching-above", 0, 0, 10, 10),
		rectangled.NewRectangle("touching-below", 0, 10, 10, 20),
		false,
		[]rectangled.Edge{rectangled.Bottom},
		[]rectangled.Edge{rectangled.Top},
	},
	{
		rectangled.NewRectangle("touching-left", 0, 0, 10, 10),
		rectangled.NewRectangle("touching-right", 10, 0, 20, 10),
		false,
		[]rectangled.Edge{rectangled.Right},
		[]rectangled.Edge{rectangled.Left},
	},
	{
		rectangled.NewRectangle("full-overlap", 0, 0, 10, 10),
		rectangled.NewRectangle("full-overlap", 0, 0, 10, 10),
		true,
		[]rectangled.Edge{rectangled.Top, rectangled.Right, rectangled.Bottom, rectangled.Left},
		[]rectangled.Edge{rectangled.Top, rectangled.Right, rectangled.Bottom, rectangled.Left},
	},
	{
		rectangled.NewRectangle("inside-of", 5, 5, 10, 10),
		rectangled.NewRectangle("surrounding", 0, 0, 15, 15),
		true,
		nil,
		nil,
	},
	{
		rectangled.NewRectangle("top-left", 0, 0, 10, 10),
		rectangled.NewRectangle("top-left-surrounding", 0, 0, 15, 15),
		true,
		[]rectangled.Edge{rectangled.Top, rectangled.Left},
		[]rectangled.Edge{rectangled.Top, rectangled.Left},
	},
	{
		rectangled.NewRectangle("bottom-right", 5, 5, 15, 15),
		rectangled.NewRectangle("bottom-right-surrounding", 0, 0, 15, 15),
		true,
		[]rectangled.Edge{rectangled.Right, rectangled.Bottom},
		[]rectangled.Edge{rectangled.Right, rectangled.Bottom},
	},
	{
		rectangled.NewRectangle("overlap-above", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-below", 0, 9, 10, 19),
		true,
		[]rectangled.Edge{rectangled.Right, rectangled.Left},
		[]rectangled.Edge{rectangled.Right, rectangled.Left},
	},
	{
		rectangled.NewRectangle("overlap-left", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-right", 9, 0, 19, 10),
		true,
		[]rectangled.Edge{rectangled.Top, rectangled.Bottom},
		[]rectangled.Edge{rectangled.Top, rectangled.Bottom},
	},
	{
		rectangled.NewRectangle("bottom-right-corner", 0, 0, 10, 10),
		rectangled.NewRectangle("top-left corner", 9, 9, 19, 19),
		true,
		nil,
		nil,
	},
	{
		rectangled.NewRectangle("bottom-left-corner", 9, 0, 19, 10),
		rectangled.NewRectangle("top-right corner", 0, 9, 10, 19),
		true,
		nil,
		nil,
	},
}

func TestRectangleOverlap(t *testing.T) {
	for _, testCase := range rectangleOverlapTouchTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect1.Overlaps(testCase.rect2))
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.overlaps, testCase.rect2.Overlaps(testCase.rect1))
		})
	}
}

func TestRectangleTouches(t *testing.T) {
	for _, testCase := range rectangleOverlapTouchTests {
		t.Run(testCase.rect1.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch1, testCase.rect1.Touches(testCase.rect2))
		})
		t.Run(testCase.rect2.ID, func(t *testing.T) {
			require.Equal(t, testCase.touch2, testCase.rect2.Touches(testCase.rect1))
		})
	}
}

var rectangleOffsetTests = []struct {
	rect   rectangled.Rectangle[string]
	target rectangled.Rectangle[string]
}{
	{
		rectangled.NewRectangle("pos offset", 0, 0, 10, 10),
		rectangled.NewRectangle("pos offset", 10, 10, 0, 0),
	},
	{
		rectangled.NewRectangle("neg offset", 0, 0, 10, 10),
		rectangled.NewRectangle("neg offset", -10, -10, 0, 0),
	},
	{
		rectangled.NewRectangle("zero offset", 0, 0, 10, 10),
		rectangled.NewRectangle("zero offset", 0, 0, 0, 0),
	},
}

func TestRectangleOffset(t *testing.T) {
	for _, testCase := range rectangleOffsetTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			offset := testCase.rect.Offset(testCase.target)

			require.Equal(t, testCase.rect.X+testCase.target.X, offset.X)
			require.Equal(t, testCase.rect.W+testCase.target.X, offset.W)
			require.Equal(t, testCase.rect.Y+testCase.target.Y, offset.Y)
			require.Equal(t, testCase.rect.Z+testCase.target.X, offset.Z)
		})
	}
}

var rectangleAreaTests = []struct {
	rect     rectangled.Rectangle[string]
	expected int
}{
	{rectangled.NewRectangle("zero height", 0, 0, 10, 0), 0},
	{rectangled.NewRectangle("zero width", 0, 0, 0, 10), 0},
	{rectangled.NewRectangle("zero size", 0, 0, 0, 0), 0},
	{rectangled.NewRectangle("positive area", 0, 0, 10, 10), 100},
	{rectangled.NewRectangle("negative area", 0, 0, -10, -10), 100},
}

func TestRectangleArea(t *testing.T) {
	for _, testCase := range rectangleAreaTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.rect.Area())
		})
	}
}

var rectangleValidityTests = []struct {
	rect     rectangled.Rectangle[string]
	expected error
}{
	{rectangled.NewRectangle("zero height", 0, 0, 10, 0), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("zero width", 0, 0, 0, 10), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("zero size", 0, 0, 0, 0), rectangled.ErrZoroArea},
	{rectangled.NewRectangle("flipped points", 0, 0, -10, -10), rectangled.ErrBadPoints},
	{rectangled.NewRectangle("valid", 0, 0, 10, 10), nil},
	{rectangled.NewRectangle("valid in negative", -10, -10, 0, 0), nil},
}

func TestRectangleValidate(t *testing.T) {
	for _, testCase := range rectangleValidityTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			if testCase.expected == nil {
				require.Nil(t, testCase.rect.Validate())
			} else {
				require.ErrorIs(t, testCase.expected, testCase.rect.Validate())
			}
		})
	}
}

var rectangleOverlappingAreaTests = []struct {
	rect1    rectangled.Rectangle[string]
	rect2    rectangled.Rectangle[string]
	expected *rectangled.Rectangle[string]
}{
	{
		rectangled.NewRectangle("overlap-1", 0, 0, 10, 10),
		rectangled.NewRectangle("overlap-2", 5, 5, 15, 15),
		&rectangled.Rectangle[string]{"overlap-2", 5, 5, 10, 10},
	},
	{
		rectangled.NewRectangle("overlap-2", 5, 5, 15, 15),
		rectangled.NewRectangle("overlap-1", 0, 0, 10, 10),
		&rectangled.Rectangle[string]{"overlap-1", 5, 5, 10, 10},
	},
	{
		rectangled.NewRectangle("no-overlap-1", 0, 0, 10, 10),
		rectangled.NewRectangle("no-overlap-2", 10, 10, 20, 20),
		nil,
	},
}

func TestRectangleOverlappingArea(t *testing.T) {
	for _, testCase := range rectangleOverlappingAreaTests {
		var testName = fmt.Sprintf("%s -> %s", testCase.rect1.ID, testCase.rect2.ID)

		t.Run(testName, func(t *testing.T) {
			overlap := testCase.rect1.OverlappingArea(testCase.rect2)

			if testCase.expected == nil {
				require.Nil(t, overlap)
				return
			}

			require.NotNil(t, overlap)
			require.Equal(t, *overlap, *testCase.expected)
		})
	}
}

var rectangleWidthHeightTests = []struct {
	rect           rectangled.Rectangle[string]
	expectedWidth  int
	expectedHeight int
}{
	{rectangled.NewRectangle("simple", 0, 0, 10, 10), 10, 10},
	{rectangled.NewRectangle("negative", -20, -20, -10, -10), 10, 10},
	{rectangled.NewRectangle("negative->posative", -10, -10, 10, 10), 20, 20},
	// although this requires a Rectangle that would not pass the Validate checks its worth testing imo
	{rectangled.NewRectangle("zero-size", 0, 0, 0, 0), 0, 0},
}

func TestRectangleWidth(t *testing.T) {
	for _, testCase := range rectangleWidthHeightTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expectedWidth, testCase.rect.Width())
		})
	}
}

func TestRectangleHeight(t *testing.T) {
	for _, testCase := range rectangleWidthHeightTests {
		t.Run(testCase.rect.ID, func(t *testing.T) {
			require.Equal(t, testCase.expectedHeight, testCase.rect.Height())
		})
	}
}

func edge(id string, x, y, w, z int) *rectangled.EdgeCoordinates[string] {
	return &rectangled.EdgeCoordinates[string]{
		ID: id,
		X:  x,
		Y:  y,
		W:  w,
		Z:  z,
	}
}

var rectangleTouchCoordinatesTests = []struct {
	rect      rectangled.Rectangle[string]
	target    rectangled.Rectangle[string]
	coords    [4]*rectangled.EdgeCoordinates[string]
	coordsRev [4]*rectangled.EdgeCoordinates[string]
}{
	{
		rectangled.NewRectangle("below", 0, 10, 20, 20),
		rectangled.NewRectangle("above", 5, 0, 15, 10),
		[4]*rectangled.EdgeCoordinates[string]{
			edge("above", 5, 10, 15, 10), nil, nil, nil,
		},
		[4]*rectangled.EdgeCoordinates[string]{
			nil, nil, edge("below", 5, 10, 15, 10), nil,
		},
	},
	{
		rectangled.NewRectangle("left", 0, 0, 10, 20),
		rectangled.NewRectangle("right", 10, 5, 15, 15),
		[4]*rectangled.EdgeCoordinates[string]{
			nil, edge("right", 10, 5, 10, 15), nil, nil,
		},
		[4]*rectangled.EdgeCoordinates[string]{
			nil, nil, nil, edge("left", 10, 5, 10, 15),
		},
	},
}

func TestRectangleEdgeCoordinates(t *testing.T) {
	for _, testCase := range rectangleTouchCoordinatesTests {
		t.Run(fmt.Sprintf("%s -> %s", testCase.rect.ID, testCase.target.ID), func(t *testing.T) {
			require.Equal(t, testCase.coords[0], testCase.rect.TouchCoordinates(testCase.target, rectangled.Top))
			require.Equal(t, testCase.coords[1], testCase.rect.TouchCoordinates(testCase.target, rectangled.Right))
			require.Equal(t, testCase.coords[2], testCase.rect.TouchCoordinates(testCase.target, rectangled.Bottom))
			require.Equal(t, testCase.coords[3], testCase.rect.TouchCoordinates(testCase.target, rectangled.Left))
		})

		t.Run(fmt.Sprintf("%s -> %s", testCase.target.ID, testCase.rect.ID), func(t *testing.T) {
			require.Equal(t, testCase.coordsRev[0], testCase.target.TouchCoordinates(testCase.rect, rectangled.Top))
			require.Equal(t, testCase.coordsRev[1], testCase.target.TouchCoordinates(testCase.rect, rectangled.Right))
			require.Equal(t, testCase.coordsRev[2], testCase.target.TouchCoordinates(testCase.rect, rectangled.Bottom))
			require.Equal(t, testCase.coordsRev[3], testCase.target.TouchCoordinates(testCase.rect, rectangled.Left))
		})
	}
}
