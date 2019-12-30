/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : matrix.go
 * CreateDate : 2019-04-24 08:46:51
 * */

package main

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func updateMatrix(matrix [][]int) [][]int {
	row := len(matrix)
	if row <= 0 {
		return nil
	}

	col := len(matrix[0])
	if col <= 0 {
		return nil
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 先把它设置为最大值
			if matrix[i][j] == 1 {
				matrix[i][j] = row + col
			}

			// 左边
			if i > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
			}

			// 上方
			if j > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i < row-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
			}
			if j < col-1 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}

	return matrix
}

func main() {

}

/* vim: set tabstop=4 set shiftwidth=4 */
