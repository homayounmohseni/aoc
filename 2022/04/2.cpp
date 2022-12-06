#include <iostream>
#include <string>

using namespace std;

bool does_contain(pair<int, int>, pair<int, int>);
bool does_overlap(pair<int, int>, pair<int, int>);

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



		if (does_overlap(first, second)) {
			cnt++;
		}
	}

	cout << cnt << "\n";
	return 0;
}


bool does_contain(pair<int, int> p1, pair<int, int> p2) {
	if ((p1.first <= p2.first && p1.second >= p2.second) ||
			(p1.first >= p2.first && p1.second <= p2.second)) {
		return true;
	}
	return false;
}

bool does_overlap(pair<int, int> p1, pair<int, int> p2) {
	if ((p1.first <= p2.first && p1.second >= p2.first) ||
			(p1.first >= p2.first && p1.first <= p2.second)) {
		return true;
	}
	return false;
}
