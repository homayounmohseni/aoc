#include <iostream>
#include <string>


using namespace std;


int main() {
	string line;
	int most_calories = 0, calories = 0;
	while (getline(cin, line)) {
		if (line.empty()) {
			most_calories = max(calories, most_calories);
			calories = 0;
			continue;
		}
		calories += stoi(line);
	}
	cout << most_calories << "\n";
	return 0;
}
