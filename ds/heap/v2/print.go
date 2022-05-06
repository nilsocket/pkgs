package heap

import (
	"fmt"
	"strings"
)

func (h *Heap[T]) PrintTree() {
	height := h.Height()
	b := &strings.Builder{}
	fillGap(height + 1)
	b.WriteString(spacePrefix(height+1) + h.Elems[0].String() + "\n")
	h.printTree(b, height+1, height)
	fmt.Println(b.String())
}

const (
	baseElemWidth = 7
)

func (h *Heap[T]) printTree(b *strings.Builder, maxHeight, curHeight int) {
	if curHeight > 0 {
		bars(b, maxHeight, curHeight)
		h.printValues(b, maxHeight, curHeight)
		h.printTree(b, maxHeight, curHeight-1)
	}
}

func (h *Heap[T]) printValues(b *strings.Builder, mh, ch int) {
	if ch > 0 {
		sp := spacePrefix(ch)
		ne := 1 << (mh - ch)

		b.WriteString(sp)

		f, t := ne-1, (ne*2)-1
		for ; f < t && f < h.Size; f += 2 {
			b.WriteString(h.Elems[f].String())
			if f+1 < h.Size {
				b.WriteString(valGap(ch))
				b.WriteString(h.Elems[f+1].String())
			}
			b.WriteString(barGap(ch))
		}
		b.WriteString("\n")
	}
}

func valGap(ch int) string {
	return strings.Repeat(" ", (gaps[ch].mbl * 2))
}

func bars(b *strings.Builder, mh, ch int) {
	if ch > 0 {
		sp := spacePrefix(ch)
		ne := 1 << (mh - ch)

		b.WriteString(sp)
		for i := ne; i > 0; i -= 2 {
			b.WriteString(bar(ch) + barGap(ch))
		}
		b.WriteString("\n")
	}
}

func barGap(ch int) string {
	return strings.Repeat(" ", gaps[ch].gap)
}

func bar(ch int) string {
	mb := strings.Repeat("─", gaps[ch].mbl)
	return "┌" + mb + "┴" + mb + "┐"
}

type gapData struct {
	gap, mbl int
}

var gaps = make(map[int]gapData)

func spacePrefix(ch int) string {
	return strings.Repeat(" ", gaps[ch].gap/2)
}

func fillGap(ch int) {
	if ch == 0 {
		gaps[ch] = gapData{gap: 1, mbl: 2}
	} else {
		fillGap(ch - 1)
		t := gaps[ch-1]
		gap := (t.mbl * 2) + 2 + t.gap
		mbl := gap / 2
		gaps[ch] = gapData{gap: gap, mbl: mbl}
	}
}

func (e Elem[T]) String() string {
	if e.Data == nil {
		return fmt.Sprintf("%v", e.Key)
	}
	return fmt.Sprintf("%v-%v", e.Key, e.Data)
}
