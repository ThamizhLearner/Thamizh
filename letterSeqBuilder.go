package script

// Helps construct the immutable letter-seq efficiently
type LetterSeqBuilder struct {
	letters []Letter // Letter accumulator
}

// New letter-seq builder, with specified initial capacity.
func NewLetterSeqBuilder(capacity int) LetterSeqBuilder {
	return LetterSeqBuilder{letters: make([]Letter, 0, capacity)}
}

func (this *LetterSeqBuilder) Append(seq LetterSeq) {
	this.letters = append(this.letters, seq.letters...)
}

func (this *LetterSeqBuilder) AppendRangeN(seq LetterSeq, startLetterIdx int, letterCount int) {
	this.letters = append(this.letters, seq.TakeRangeN(startLetterIdx, letterCount).letters...)
}

func (this *LetterSeqBuilder) AppendFirstN(seq LetterSeq, letterCount int) {
	this.letters = append(this.letters, seq.TakeFirstN(letterCount).letters...)
}

func (this *LetterSeqBuilder) AppendLastN(seq LetterSeq, letterCount int) {
	this.letters = append(this.letters, seq.TakeLastN(letterCount).letters...)
}

func (this *LetterSeqBuilder) AppendLetter(letter Letter) {
	this.letters = append(this.letters, letter)
}

// Constructs and returns the composed String
func (this *LetterSeqBuilder) Build() LetterSeq {
	letters2 := this.letters
	this.letters = nil // Ensure this does not get used again! [Effectively, resets this letter-seq builder]
	return UnsafeNewLetterSeq(letters2)
}
