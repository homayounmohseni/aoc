#include <string>
#include <vector>
#include <limits>
#include <cassert>
#include <iostream>

using namespace std;
int main() {
	vector<vector<int>> field;
	string line;
	while(getline(cin, line)) {
		vector<int> row;
		row.reserve(line.length());
		for (char c : line) {
			row.push_back(c - '0');
		}
		field.push_back(row);
	}

	int rows = field.size();
	assert(rows > 0);
	int cols = field[0].size();

	vector<vector<int>> max_left, max_right, max_top, max_buttom;
	max_left = max_right = max_top = max_buttom =
		vector<vector<int>>(rows, vector<int>(cols, 0));

	for (int i = 1; i < rows - 1; i++) {
		for (int j = 1; j < cols - 1; j++) {
			max_left[i][j] = max(max_left[i][j - 1], field[i][j - 1]);
		}
	}
	for (int i = 1; i < rows - 1; i++) {
		for (int j = cols - 2; j > 0; j--) {
			max_right[i][j] = max(max_right[i][j + 1], field[i][j + 1]);
		}
	}
	for (int j = 1; j < cols - 1; j++) {
		for (int i = 1; i < rows - 1; i++) {
			max_top[i][j] = max(max_top[i - 1][j], field[i - 1][j]);
		}
	}
	for (int j = 1; j < cols; j++) {
		for (int i = rows - 2; i > 0; i--) {
			max_buttom[i][j] = max(max_buttom[i + 1][j], field[i + 1][j]);
		}
	}

		
	int visible_cnt = 2 * (rows - 1) + 2 * (cols - 1);
	for (int i = 1; i < rows - 1; i++) {
		for (int j = 1; j < cols - 1; j++) {
			if (field[i][j] > max_left[i][j] || field[i][j] > max_right[i][j] ||
					field[i][j] > max_top[i][j] || field[i][j] > max_buttom[i][j]) {
				visible_cnt++;
			}
		}
	}

	cout << visible_cnt << "\n";

	return 0;
}

