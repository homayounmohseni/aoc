#include <iostream>
#include <string>

using namespace std;

int find_outcome_score(char, char);
int find_move_score(char);

int main() {
	char opponents, yours;

	int score = 0;

	string line;
	while (getline(cin, line)) {
		if (line.length() != 3) {
			cout << "WTF\n";
			continue;
		}
		char first = line[0];
		char second = line[2];
		char move;

		if (second == 'X') {
			move = (first - 'A' - 1 + 3) % 3 + 'X';
		}
		else if (second == 'Y') {
			move = first - 'A' + 'X';
		}
		else if (second == 'Z') {
			move = (first - 'A' + 1) % 3 + 'X';
		}
		score += find_outcome_score(first, move) +
			find_move_score(move);
	}
	cout << score << "\n";
	return 0;
}

int find_outcome_score(char first, char second) {
	char diff = second - 'X' - first - 'A';
	diff = (diff + 3) % 3;
	
	int result;
	if (diff == 1) {
		result = 6;
	}
	else if (diff == 0) {
		result = 3;
	}
	else {
		result = 0;
	}
	return result;
}
		
int find_move_score(char move) {
	char move_z = move - 'X';
	return move_z + 1;
}


