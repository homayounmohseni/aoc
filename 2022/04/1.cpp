#include <iostream>
#include <string>

using namespace std;

int main() {
	string line;

	int cnt = 0;
	while (getline(cin, line)) {
		pair<int, int> first, second;
		int i;
		for (i = 0; line[i] != '-' && i < line.length(); i++);
		first.first = stoi(line.substr(0, i));
		
		int start = ++i;
		for (; line[i] != ',' && i < line.length(); i++);
		first.second = stoi(line.substr(start, i - start));
		start = ++i;
		for (; line[i] != '-' && i < line.length(); i++);
		second.first = stoi(line.substr(start, i - start));
		start = i + 1;
		second.second = stoi(line.substr(start, line.length() - start));



		if ((first.first <= second.first && first.second >= second.second) ||
				(first.first >= second.first && first.second <= second.second)) {
			cnt++;
		}
	}

	cout << cnt << "\n";
	return 0;
}

				

