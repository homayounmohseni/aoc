#include <iostream>
#include <string>
#include <vector>


int get_index(char c);

using namespace std;
int main() {
	int prsum = 0;

	vector<string> rs_str(3, "");
	while (getline(cin, rs_str[0])) {
		for (int i = 1; i < 3; i++) {
			getline(cin, rs_str[i]);
		}

		vector<vector<int>> rs(3, vector<int>(52, 0));
		for (int i = 0; i < 3; i++) {
			for (int j = 0; j < rs_str[i].length(); j++) {
				rs[i][get_index(rs_str[i][j])]++;
			}
		}


		for (int i = 0; i < 52; i++) {
			bool cond = true;
			for (int j = 0; j < 3; j++) {
				cond &= (rs[j][i] > 0);
			}
			if (cond) {
				prsum += i + 1;
				break;
			}
		}
	}
	cout << prsum << "\n";
	return 0;
}

int get_index(char c) {
	if (c >= 'a' && c <= 'z') {
		return c - 'a';
	}
	else if (c >= 'A' && c <= 'Z') {
		return c - 'A' + 26;
	}
	return -1;
}
