package main
import "fmt"
import "bufio"
import "os"

// 対話形式で引数を受ける方法
// 競プロでは対話形式でない問題の場合でも、これで設問が用意する引数を読み込むことができる
func main () {
	var sc = bufio.NewScanner(os.Stdin)	
	var input string

	// sc.Scan で入力状態になる
	// そこで入力された値を、 sc.Textで取得可能
	// 入力状態になるとき、何もターミナル側にはメッセージが表示されない点に注意
	// 入力状態であることを明確に伝えるためには fmt.Print を使用すると良い
	// 【ここから先は 競プロの場合の説明】
	// sc.Scan 次の引数にポインタを進め、sc.text でその位置にある入力を読み込む
	// つまり、 sc.Scan と sc.Text はほとんどの場合同時に利用する
	// sc.Scan を実行するまで、 sc.Text は何度実行しても同じ値を返す
	fmt.Print("何か入力してください:")
	sc.Scan()
	input = sc.Text()

	fmt.Println("あなたが入力した値は[" + input + "]です")
	
}