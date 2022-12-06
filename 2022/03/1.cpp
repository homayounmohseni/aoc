#include <iostream>
#include <string>
#include <vector>


int get_index(char c);

using namespace std;
int main() {
	string line;
	int prsum = 0;
	while (getline(cin, line)) {
		vector<char> left(52, 0);
		vector<char> right(52, 0);
		for (int i = 0; i < line.length() / 2; i++) {
			left[get_index(line[i])]++;
		}
		for (int i = line.length() / 2; i < line.length(); i++) {
			right[get_index(line[i])]++;
		}

		for (int i = 0; i < 52; i++) {
			if (left[i] && right[i]) {
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
