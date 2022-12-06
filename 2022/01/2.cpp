#include <iostream>
#include <string>


using namespace std;


int main() {
	string line;
	int  calories = 0;
	int top1 = 0, top2 = 0, top3 = 0;
	while (getline(cin, line)) {
		if (line.empty()) {
			if (calories > top1) {
				top3 = top2;
				top2 = top1;
				top1 = calories;
			}
			else if (calories > top2) {
				top3 = top2;
				top2 = calories;
			}
			else if (calories > top3) {
				top3 = calories;
			}

			calories = 0;
			continue;
		}
		calories += stoi(line);
	}
	cout << top1 + top2 + top3 << "\n";
	return 0;
}
