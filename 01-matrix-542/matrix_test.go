/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : matrix_test.go
 * CreateDate : 2019-04-24 09:00:22
 * */

package main

import (
	"testing"
)

/*
 public int[][] updateMatrix_1(int[][] matrix) {
        row = matrix.length;
        col = matrix[0].length;
        // 第一次遍历，正向遍历，根据相邻左元素和上元素得出当前元素的对应结果
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (matrix[i][j] == 1) {
                    matrix[i][j] = row + col;
                }
                if (i > 0) {
                    matrix[i][j] = Math.min(matrix[i][j], matrix[i - 1][j] + 1);
                }
                if (j > 0) {
                    matrix[i][j] = Math.min(matrix[i][j], matrix[i][j - 1] + 1);
                }
            }
        }
        // 第二次遍历，反向遍历，根据相邻右元素和下元素及当前元素的结果得出最终结果
        for (int i = row - 1; i >= 0; i--) {
            for (int j = col - 1; j >= 0; j--) {
                if (i < row - 1) {
                    matrix[i][j] = Math.min(matrix[i][j], matrix[i + 1][j] + 1);
                }
                if (j < col - 1) {
                    matrix[i][j] = Math.min(matrix[i][j], matrix[i][j + 1] + 1);
                }
            }
        }
        return matrix;
    }
*/

func matrixEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for i := 0; i < len(m1); i++ {
		if len(m1[i]) != len(m2[i]) {
			return false
		}

		for j := 0; j < len(m2); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}

	return true
}

func TestLeet(t *testing.T) {
	in := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}

	out := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}

	got := updateMatrix(in)
	if !matrixEqual(out, got) {
		t.Fatalf("ERROR got:%+v want:%+v", got, out)
	}
}

/* vim: set tabstop=4 set shiftwidth=4 */
