package tx

type Input struct {
	HashSrcTx string // 输入来源交易的 hash
	IndexSrcOutput string // 输入来源交易输出的索引
}
